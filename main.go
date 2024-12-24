package main

import (
	"PassGet/modules/browser"
	"log"
)

func main() {
	err := browser.Get()
	if err != nil {
		log.Println("Get Browser Data Failed")
		return
	}
}
