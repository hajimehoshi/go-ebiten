// Code generated by file2byteslice. DO NOT EDIT.


package main

var ntsc_go = []byte("// Copyright 2020 The Ebiten Authors\r\n//\r\n// Licensed under the Apache License, Version 2.0 (the \"License\");\r\n// you may not use this file except in compliance with the License.\r\n// You may obtain a copy of the License at\r\n//\r\n//     http://www.apache.org/licenses/LICENSE-2.0\r\n//\r\n// Unless required by applicable law or agreed to in writing, software\r\n// distributed under the License is distributed on an \"AS IS\" BASIS,\r\n// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\r\n// See the License for the specific language governing permissions and\r\n// limitations under the License.\r\n\r\n// +build ignore\r\n\r\npackage main\r\n\r\nvar Time float\r\n\r\nfunc distort(p vec2) vec2 {\r\n\twarpX := 0.031\r\n\twarpY := 0.041\r\n\tp = p*2.0 - 1.0\r\n\tp *= vec2(1.0+(p.y*p.y)*warpX, 1.0+(p.x*p.x)*warpY)\r\n\tp.y += 0.5\r\n\tp.x += 0.51\r\n\treturn p*0.5 + 0.25\r\n}\r\n\r\nfunc colorBleeding(current_color vec4, color_left vec4) (vec4, vec4) {\r\n\tcolor_bleeding := 1.2\r\n\tcurrent_color = current_color * vec4(color_bleeding, 0.5, 1.0-color_bleeding, 1)\r\n\tcolor_left = color_left * vec4(1.0-color_bleeding, 0.5, color_bleeding, 1)\r\n\treturn current_color, color_left\r\n}\r\n\r\nfunc colorScanline(uv vec2, c vec4, time float) vec4 {\r\n\tscreen_height := 480.0\r\n\tscan_size := 2.0\r\n\tlines_velocity := 30.0\r\n\tscanline_alpha := 0.9\r\n\tlines_distance := 4.0\r\n\tline_row := floor((uv.y * screen_height / scan_size) + mod(time*lines_velocity, lines_distance))\r\n\tn := 1.0 - ceil((mod(line_row, lines_distance) / lines_distance))\r\n\tc = c - n*c*(1.0-scanline_alpha)\r\n\tc.a = 1.0\r\n\treturn c\r\n}\r\n\r\nfunc Fragment(position vec4, texCoord vec2, color vec4) vec4 {\r\n\txy := texCoord\r\n\txy = distort(xy)\r\n\r\n\tbleeding_range_x := 2.0\r\n\tbleeding_range_y := 2.0\r\n\tscreen_height := 480.0\r\n\tscreen_width := 640.0\r\n\r\n\tpixel_size_x := 1.0 / screen_width * bleeding_range_x\r\n\tpixel_size_y := 1.0 / screen_height * bleeding_range_y\r\n\tcolor_left := imageSrc0At(xy - vec2(pixel_size_x, pixel_size_y))\r\n\tcurrent_color := imageSrc0At(xy)\r\n\tcolor_left, current_color = colorBleeding(current_color, color_left)\r\n\tc := current_color + color_left\r\n\treturn colorScanline(xy, c, Time)\r\n}\r\n")
