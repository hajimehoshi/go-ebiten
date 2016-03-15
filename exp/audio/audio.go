// Copyright 2015 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package audio

import (
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type mixedPlayersStream struct {
	context      *Context
	writtenBytes int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const (
	channelNum     = 2
	bytesPerSample = 2
	bitsPerSample  = bytesPerSample * 8
)

func (s *mixedPlayersStream) Read(b []byte) (int, error) {
	s.context.Lock()
	defer s.context.Unlock()

	bytesPerFrame := s.context.sampleRate * bytesPerSample * channelNum / ebiten.FPS
	x := s.context.frames*bytesPerFrame + len(b)
	if x <= s.writtenBytes {
		return 0, nil
	}

	if len(s.context.players) == 0 {
		l := min(len(b), x-s.writtenBytes)
		l &= ^3
		copy(b, make([]byte, l))
		s.writtenBytes += l
		return l, nil
	}
	closed := []*Player{}
	l := len(b)
	for p := range s.context.players {
		_, err := p.readToBuffer(l)
		if err == io.EOF {
			closed = append(closed, p)
		} else if err != nil {
			return 0, err
		}
		l = min(p.bufferLength()/4*4, l)
	}
	l &= ^3
	b16s := [][]int16{}
	for p := range s.context.players {
		b16s = append(b16s, p.bufferToInt16(l))
	}
	for i := 0; i < l/2; i++ {
		x := 0
		for _, b16 := range b16s {
			x += int(b16[i])
		}
		if x > (1<<15)-1 {
			x = (1 << 15) - 1
		}
		if x < -(1 << 15) {
			x = -(1 << 15)
		}
		b[2*i] = byte(x)
		b[2*i+1] = byte(x >> 8)
	}
	for p := range s.context.players {
		p.proceed(l)
	}
	for _, p := range closed {
		delete(s.context.players, p)
	}
	s.writtenBytes += l
	return l, nil
}

// TODO: Enable to specify the format like Mono8?

type Context struct {
	sampleRate  int
	stream      *mixedPlayersStream
	players     map[*Player]struct{}
	innerPlayer *player
	frames      int
	sync.Mutex
}

func NewContext(sampleRate int) *Context {
	// TODO: Panic if one context exists.
	c := &Context{
		sampleRate: sampleRate,
		players:    map[*Player]struct{}{},
	}
	c.stream = &mixedPlayersStream{
		context: c,
	}
	p, err := startPlaying(c.stream, c.sampleRate)
	if err != nil {
		panic(fmt.Sprintf("audio: NewContext error: %v", err))
	}
	c.innerPlayer = p
	return c
}

// Update proceeds the inner (logical) time of the context by 1/60 second.
// This is expected to be called in the game's updating function (sync mode)
// or an independent goroutine with timers (unsync mode).
// In sync mode, the game logical time syncs the audio logical time and
// you will find audio stops when the game stops e.g. when the window is deactivated.
// In unsync mode, the audio never stops even when the game stops.
func (c *Context) Update() {
	c.Lock()
	defer c.Unlock()
	c.frames++
}

// SampleRate returns the sample rate.
func (c *Context) SampleRate() int {
	c.Lock()
	defer c.Unlock()
	return c.sampleRate
}

type Player struct {
	context *Context
	src     io.ReadSeeker
	buf     []byte
	pos     int64
}

// NewPlayer creates a new player with the given data to the given channel.
// The given data is queued to the end of the buffer.
// This may not be played immediately when data already exists in the buffer.
//
// src's format must be linear PCM (16bits, 2 channel stereo, little endian)
// without a header (e.g. RIFF header).
func (c *Context) NewPlayer(src io.ReadSeeker) (*Player, error) {
	c.Lock()
	defer c.Unlock()
	p := &Player{
		context: c,
		src:     src,
		buf:     []byte{},
	}
	// Get the current position of the source.
	pos, err := p.src.Seek(0, 1)
	if err != nil {
		return nil, err
	}
	p.pos = pos
	return p, nil
}

func (p *Player) readToBuffer(length int) (int, error) {
	bb := make([]byte, length)
	n, err := p.src.Read(bb)
	if 0 < n {
		p.buf = append(p.buf, bb[:n]...)
	}
	return n, err
}

func (p *Player) bufferToInt16(lengthInBytes int) []int16 {
	r := make([]int16, lengthInBytes/2)
	for i := 0; i < lengthInBytes/2; i++ {
		r[i] = int16(p.buf[2*i]) | (int16(p.buf[2*i+1]) << 8)
	}
	return r
}

func (p *Player) proceed(length int) {
	p.buf = p.buf[length:]
	p.pos += int64(length)
}

func (p *Player) bufferLength() int {
	return len(p.buf)
}

func (p *Player) Play() error {
	p.context.Lock()
	defer p.context.Unlock()

	p.context.players[p] = struct{}{}
	return nil
}

func (p *Player) IsPlaying() bool {
	_, ok := p.context.players[p]
	return ok
}

func (p *Player) Rewind() error {
	return p.Seek(0)
}

func (p *Player) Seek(offset time.Duration) error {
	p.buf = []byte{}
	o := int64(offset) * int64(p.context.sampleRate) / int64(time.Second)
	pos, err := p.src.Seek(o, 0)
	if err != nil {
		return err
	}
	p.pos = pos
	return nil
}

func (p *Player) Pause() error {
	p.context.Lock()
	defer p.context.Unlock()

	delete(p.context.players, p)
	return nil
}

func (p *Player) Current() time.Duration {
	return time.Duration(p.pos) * time.Second / time.Duration(p.context.sampleRate)
}

// TODO: Volume / SetVolume?
// TODO: Panning
