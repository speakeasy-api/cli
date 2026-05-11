package main

import (
	"context"
	"os"

	"github.com/speakeasy-api/cli/go/cli/internal/app"
)

func main() {
	app.Execute(context.Background(), os.Args)
}
