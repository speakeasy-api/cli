package cmd

import "github.com/speakeasy-api/cli/go/cli/internal/profile"

type Profile = profile.Profile

var LoadProfile = profile.Load
var DefaultProfilePath = profile.DefaultProfilePath
