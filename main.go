package main

import (
	app "./application"
)

func main() {
	a := app.GoService{}
	a.Initialize()

	a.Run()

	a.WaitForShutdown()
}
