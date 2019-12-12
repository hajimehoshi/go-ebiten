// Copyright 209 The Ebiten Authors
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

package ebiten

import "github.com/hajimehoshi/ebiten/internal/driver"

// An InputCursorMode represents
// a render and coordinate mode of a mouse cursor.
type InputCursorMode int

// Cursor Modes
const (
	CursorModeVisible  = InputCursorMode(driver.CursorModeVisible)
	CursorModeHidden   = InputCursorMode(driver.CursorModeHidden)
	CursorModeCaptured = InputCursorMode(driver.CursorModeCaptured)
)
