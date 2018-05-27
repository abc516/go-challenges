// +build darwin linux windows

package main

import (
	"golang.org/x/mobile/app"
	"log"
	//"golang.org/x/mobile/event/lifecycle"
	//"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/event/touch"
)

func main()  {
	log.Println("app has started")
	app.Main(func(a app.App) {
		//var glctx gl.Context
		//var sz size.Event
		log.Println("app has started")
		for e := range a.Events() {
			log.Println(e)
			switch e := a.Filter(e).(type) {
			//case lifecycle.Event:
			//case size.Event:
			//	sz = e
			//
			//}
			case touch.Event:
				x := e.X
				y := e.Y
				log.Println("Our current coordiantes x: " + string(int(x)) + "and y: " + string(int(y)))
			}
		}
	})
}

//steps to get up and running

//phase 0.5 get the android app running with an image

//phase 1, on android app, when you hit the image, a sound plays

//phase 1.5 two images are on the screen, a different sound will play when each image is hit

//phase 2 multiple images, multiple sounds

