package main

import (
	"golang.org/x/mobile/gl"
	"golang.org/x/mobile/exp/audio/al"
)

func main()  {
	var x = gl.ACTIVE_ATTRIBUTE_MAX_LENGTH
	al.DeviceError()
	print(x)
}

