package components

import (
	"errors"

	"github.com/gotk3/gotk3/gtk"
)

var win *gtk.Window
var tabManager *gtk.Notebook

func createWindow(winTitle string) (*gtk.Window, error) {
	gtk.Init(nil)

	winInstance, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return nil, err
	}
	winInstance.SetTitle(winTitle)
	winInstance.Connect("destroy", func() {
		gtk.MainQuit()
	})

	win = winInstance

	tabManagerInstance, err := gtk.NotebookNew()
	if err != nil {
		return nil, err
	}

	tabManager = tabManagerInstance

	win.Add(tabManager)

	return win, err
}

func AddTab(tabLabel string) (int, error) {
	label, err := gtk.LabelNew(tabLabel)
	if err != nil {
		return -1, err
	}

	var index = tabManager.AppendPage(label, label)

	if index == -1 {
		return -1, errors.New("Failed to append page.")
	}

	return index, nil
}

func Init(winTitle string) (*gtk.Window, error) {
	return createWindow("Web Bra")
}

func RenderWindow() {
	win.SetDefaultSize(800, 600)
	win.ShowAll()
	gtk.Main()
}
