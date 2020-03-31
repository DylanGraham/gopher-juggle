package main

import (
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Gopher Juggle")

	gopherIcon1 := canvas.NewImageFromFile("buff.png")
	content := widget.NewButton("click me", func() {
		log.Println("tapped")
	})

	container := fyne.NewContainerWithLayout(layout.NewMaxLayout(), content, gopherIcon1)

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
