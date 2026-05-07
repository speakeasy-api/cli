package main

import (
	"context"
	"os"

	"github.com/speakeasy-api/gf/go/cli/internal/app"
)

func main() {
	app.Execute(context.Background(), os.Args)
}
