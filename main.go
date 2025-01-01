package main

import (
	"fmt"

	hook "github.com/robotn/gohook"
)

func main() {
	fmt.Println("Listening for key presses...")
	events := hook.Start()
	defer hook.End()

	for event := range events {
		fmt.Printf("Event: %v\n", event)
		if event.Kind == hook.KeyDown && event.Rawcode == 81 { // 81 is the rawcode for 'q'
			fmt.Println("Key 'q' pressed!")
			break
		}
	}
}
