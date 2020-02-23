package main

import (
	"github.com/NYTimes/gizmo/server/kit"
	"github.com/justindfuller/financial/internal/service"
)

func main() {
	kit.Run(service.New())
}
