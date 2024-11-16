// File to change the background image
package main

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/joho/godotenv"
	"pkg.go.dev/golang.org/x/sys/windows"
)

var user32DLL = windows.NewLazyDLL("user32.dll")
var procSystemParameterInfo = user32DLL.NewProc("SystemParametersInfoW")

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	imgPath, _ := windows.UTF16PtrFromString(`wallpaper1Path`)
	img2Path, _ := windows.UTF16PtrFromString(`wallpaper2Path`)
	fmt.Println("Changing...")
	procSystemParameterInfo.Call(20, 0, uintptr(unsafe.Pointer(imgPath)), 0x001A)
}
