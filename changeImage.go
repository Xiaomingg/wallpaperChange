// File to change the background image
package main

import (
	"fmt"
	"unsafe"

	"pkg.go.dev/golang.org/x/sys/windows"
)

var user32DLL = windows.NewLazyDLL("user32.dll")
var procSystemParameterInfo = user32DLL.NewProc("SystemParametersInfoW")

func main() {
	imgPath, _ := windows.UTF16PtrFromString(`C:\Users\James\Documents\coding\Projects\wallpaperChange\wallpaper1.jpg`)
	img2Path, _ := windows.UTF16PtrFromString(`C:\Users\James\Documents\coding\Projects\wallpaperChange\wallpaper2.jpg`)
	fmt.Println("Changing...")
	procSystemParameterInfo.Call(20, 0, uintptr(unsafe.Pointer(imgPath)), 0x001A)
}
