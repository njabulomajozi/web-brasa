package main

import (
	"log"
	"web-bra/components"
)

func main() {
	_, err := components.Init("Web Bra")
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}

	components.AddTab("Tab 1")
	components.AddTab("Tab 2")

	components.RenderWindow()
}
