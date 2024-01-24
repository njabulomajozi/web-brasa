package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func createWindow(title string) (*gtk.Window, error) {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return nil, err
	}
	win.SetTitle(title)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	return win, err
}

func renderWindow(win *gtk.Window) {
	win.SetDefaultSize(800, 600)
	win.ShowAll()
	gtk.Main()
}

func main() {
	win, err := createWindow("Web Bra")
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}

	l, err := gtk.LabelNew("Hello, gotk3!")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	win.Add(l)

	renderWindow(win)
}
