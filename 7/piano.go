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

//steps to get up and running

//phase 0.5 get the android app running with an image

//phase 1, on android app, when you hit the image, a sound plays

//phase 1.5 two images are on the screen, a different sound will play when each image is hit

//phase 2 multiple images, multiple sounds

