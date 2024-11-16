// File to listen and react to the hotkey pressed
package main

import (
	"log"
	"sync"

	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
)

func main() { mainthread.Init(fn) }

func fn() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)
	go func() {
		defer waitGroup.Done()

		err := listenHotkey(hotkey.ModCtrl, hotkey.Shift, hotkey.KeyC)
		if err != nil {
			log.Println(err)
		}
	}()
	waitGroup.Wait()
}

func listenHotkey(key hotkey.Key, mods ...hotkey.Modifier) (err error) {
	ms := []hotkey.Modifier{}
	ms = append(ms, mods...)
	hk := hotkey.New(ms, key)

	err = hk.Register()
	if err != nil {
		return
	}

	// Doesn't work until hotkey is pressed
	<-hk.Keydown()
	log.Printf("hotkey: %v is down\n", hk)

	<-hk.KeyUp()
	log.Printf("hotkey: %v is up\n", hk)
	hk.Unregister()
	return

}
