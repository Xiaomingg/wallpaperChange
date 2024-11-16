// A program to launch on startup and change wallpaper using a keyboard command

// Step 1: Choose Wallpaper to choose between
// Step 2: Bind it to a key
// Step 3: Make launch on startup

package main
package keyListener
package changeImage

import (
	"log"
	"os"
)

func main() {
	img1, err := os.ReadFile("wallpaper1.")
	img2, err := os.ReadFile("wallpaper2")
	if err != nil {
		log.Fatal(err)
	}
}
