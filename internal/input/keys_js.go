// Copyright 2013 The Ebiten Authors
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

// Code generated by genkeys.go using 'go generate'. DO NOT EDIT.

// +build js

package input

var keyToCodes = map[Key][]string{
	Key0: {
		"Digit0",
	},
	Key1: {
		"Digit1",
	},
	Key2: {
		"Digit2",
	},
	Key3: {
		"Digit3",
	},
	Key4: {
		"Digit4",
	},
	Key5: {
		"Digit5",
	},
	Key6: {
		"Digit6",
	},
	Key7: {
		"Digit7",
	},
	Key8: {
		"Digit8",
	},
	Key9: {
		"Digit9",
	},
	KeyA: {
		"KeyA",
	},
	KeyAlt: {
		"AltLeft", "AltRight",
	},
	KeyApostrophe: {
		"Quote",
	},
	KeyB: {
		"KeyB",
	},
	KeyBackslash: {
		"Backslash",
	},
	KeyBackspace: {
		"Backspace",
	},
	KeyC: {
		"KeyC",
	},
	KeyCapsLock: {
		"CapsLock",
	},
	KeyComma: {
		"Comma",
	},
	KeyControl: {
		"ControlLeft", "ControlRight",
	},
	KeyD: {
		"KeyD",
	},
	KeyDelete: {
		"Delete",
	},
	KeyDown: {
		"ArrowDown",
	},
	KeyE: {
		"KeyE",
	},
	KeyEnd: {
		"End",
	},
	KeyEnter: {
		"Enter",
	},
	KeyEqual: {
		"Equal",
	},
	KeyEscape: {
		"Escape",
	},
	KeyF: {
		"KeyF",
	},
	KeyF1: {
		"F1",
	},
	KeyF10: {
		"F10",
	},
	KeyF11: {
		"F11",
	},
	KeyF12: {
		"F12",
	},
	KeyF2: {
		"F2",
	},
	KeyF3: {
		"F3",
	},
	KeyF4: {
		"F4",
	},
	KeyF5: {
		"F5",
	},
	KeyF6: {
		"F6",
	},
	KeyF7: {
		"F7",
	},
	KeyF8: {
		"F8",
	},
	KeyF9: {
		"F9",
	},
	KeyG: {
		"KeyG",
	},
	KeyGraveAccent: {
		"Backquote",
	},
	KeyH: {
		"KeyH",
	},
	KeyHome: {
		"Home",
	},
	KeyI: {
		"KeyI",
	},
	KeyInsert: {
		"Insert",
	},
	KeyJ: {
		"KeyJ",
	},
	KeyK: {
		"KeyK",
	},
	KeyKP0: {
		"Numpad0",
	},
	KeyKP1: {
		"Numpad1",
	},
	KeyKP2: {
		"Numpad2",
	},
	KeyKP3: {
		"Numpad3",
	},
	KeyKP4: {
		"Numpad4",
	},
	KeyKP5: {
		"Numpad5",
	},
	KeyKP6: {
		"Numpad6",
	},
	KeyKP7: {
		"Numpad7",
	},
	KeyKP8: {
		"Numpad8",
	},
	KeyKP9: {
		"Numpad9",
	},
	KeyKPAdd: {
		"NumpadAdd",
	},
	KeyKPDecimal: {
		"NumpadDecimal",
	},
	KeyKPDivide: {
		"NumpadDivide",
	},
	KeyKPEnter: {
		"NumpadEnter",
	},
	KeyKPEqual: {
		"NumpadEqual",
	},
	KeyKPMultiply: {
		"NumpadMultiply",
	},
	KeyKPSubtract: {
		"NumpadSubtract",
	},
	KeyL: {
		"KeyL",
	},
	KeyLeft: {
		"ArrowLeft",
	},
	KeyLeftBracket: {
		"BracketLeft",
	},
	KeyM: {
		"KeyM",
	},
	KeyMinus: {
		"Minus",
	},
	KeyN: {
		"KeyN",
	},
	KeyNumLock: {
		"NumLock",
	},
	KeyO: {
		"KeyO",
	},
	KeyP: {
		"KeyP",
	},
	KeyPageDown: {
		"PageDown",
	},
	KeyPageUp: {
		"PageUp",
	},
	KeyPause: {
		"Pause",
	},
	KeyPeriod: {
		"Period",
	},
	KeyPrintScreen: {
		"PrintScreen",
	},
	KeyQ: {
		"KeyQ",
	},
	KeyR: {
		"KeyR",
	},
	KeyRight: {
		"ArrowRight",
	},
	KeyRightBracket: {
		"BracketRight",
	},
	KeyS: {
		"KeyS",
	},
	KeySemicolon: {
		"Semicolon",
	},
	KeyShift: {
		"ShiftLeft", "ShiftRight",
	},
	KeySlash: {
		"Slash",
	},
	KeySpace: {
		"Space",
	},
	KeyT: {
		"KeyT",
	},
	KeyTab: {
		"Tab",
	},
	KeyU: {
		"KeyU",
	},
	KeyUp: {
		"ArrowUp",
	},
	KeyV: {
		"KeyV",
	},
	KeyW: {
		"KeyW",
	},
	KeyX: {
		"KeyX",
	},
	KeyY: {
		"KeyY",
	},
	KeyZ: {
		"KeyZ",
	},
}

var keyCodeToKeyEdge = map[int]Key{
	8:   KeyBackspace,
	9:   KeyTab,
	13:  KeyEnter,
	16:  KeyShift,
	17:  KeyControl,
	18:  KeyAlt,
	19:  KeyPause,
	20:  KeyCapsLock,
	27:  KeyEscape,
	32:  KeySpace,
	33:  KeyPageUp,
	34:  KeyPageDown,
	35:  KeyEnd,
	36:  KeyHome,
	37:  KeyLeft,
	38:  KeyUp,
	39:  KeyRight,
	40:  KeyDown,
	44:  KeyPrintScreen,
	45:  KeyInsert,
	46:  KeyDelete,
	48:  Key0,
	49:  Key1,
	50:  Key2,
	51:  Key3,
	52:  Key4,
	53:  Key5,
	54:  Key6,
	55:  Key7,
	56:  Key8,
	57:  Key9,
	65:  KeyA,
	66:  KeyB,
	67:  KeyC,
	68:  KeyD,
	69:  KeyE,
	70:  KeyF,
	71:  KeyG,
	72:  KeyH,
	73:  KeyI,
	74:  KeyJ,
	75:  KeyK,
	76:  KeyL,
	77:  KeyM,
	78:  KeyN,
	79:  KeyO,
	80:  KeyP,
	81:  KeyQ,
	82:  KeyR,
	83:  KeyS,
	84:  KeyT,
	85:  KeyU,
	86:  KeyV,
	87:  KeyW,
	88:  KeyX,
	89:  KeyY,
	90:  KeyZ,
	96:  KeyKP0,
	97:  KeyKP1,
	98:  KeyKP2,
	99:  KeyKP3,
	100: KeyKP4,
	101: KeyKP5,
	102: KeyKP6,
	103: KeyKP7,
	104: KeyKP8,
	105: KeyKP9,
	106: KeyKPMultiply,
	107: KeyKPAdd,
	109: KeyKPSubtract,
	110: KeyKPDecimal,
	111: KeyKPDivide,
	112: KeyF1,
	113: KeyF2,
	114: KeyF3,
	115: KeyF4,
	116: KeyF5,
	117: KeyF6,
	118: KeyF7,
	119: KeyF8,
	120: KeyF9,
	121: KeyF10,
	122: KeyF11,
	123: KeyF12,
	144: KeyNumLock,
	186: KeySemicolon,
	187: KeyEqual,
	188: KeyComma,
	189: KeyMinus,
	190: KeyPeriod,
	191: KeySlash,
	192: KeyGraveAccent,
	219: KeyLeftBracket,
	220: KeyBackslash,
	221: KeyRightBracket,
	222: KeyApostrophe,
}
