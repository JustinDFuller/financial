package main

import (
	"github.com/NYTimes/gizmo/server/kit"
	"github.com/justindfuller/financial/service"
)

func main() {
	kit.Run(service.New())
}
