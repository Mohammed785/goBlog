package main

import (
	"github.com/Mohammed785/goBlog/cmd"

)

func main() {
	app:=cmd.CreateApp()
	app.Listen("127.0.0.1:8080")
}
