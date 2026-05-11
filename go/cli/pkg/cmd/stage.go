package cmd

import "github.com/speakeasy-api/gf/go/cli/internal/app"

type StageFunctionOptions = app.StageFunctionOptions
type StageOpenAPIOptions = app.StageOpenAPIOptions

var StageFunction = app.DoStageFunction
var StageOpenAPI = app.DoStageOpenAPI
