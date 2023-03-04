package main

import (
	"github.com/namsral/flag"
	"go.smantic.dev/access-viz/internal/webServer"
)

func main() {
	flag.Parse()

	webServer.Start()
}
