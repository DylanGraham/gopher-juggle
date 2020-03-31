package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Gopher Juggle")

	gopherIcon1 := canvas.NewImageFromFile("buff.png")

	container := fyne.NewContainerWithLayout(layout.NewGridLayout(1),
		gopherIcon1)

	// myWidget := widget.NewButtonWithIcon("", gopherIcon1, func() {
	// 	log.Println("Clicked!")
	// })

	// go changeContent(myCanvas)

	myWindow.Resize(fyne.NewSize(500, 500))
	myWindow.SetContent(container)
	myWindow.ShowAndRun()
}

func changeContent(c fyne.Canvas) {

}
