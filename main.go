package main

import (
	_ "embed"

	"github.com/webview/webview"
)

func main() {
	Url := ("https://lishogi.org/")
	debug := false
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Lishogi")
	w.SetSize(800, 600, webview.HintNone)

	w.Navigate(Url)

	w.Run()
}
