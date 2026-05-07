package api

import (
	"fmt"

	sdk "github.com/speakeasy-api/gf/go/sdk"
)

func newSDK(scheme string, host string) *sdk.Gram {
	return sdk.New(
		sdk.WithServerURL(fmt.Sprintf("%s://%s", scheme, host)),
		sdk.WithClient(goaSharedHTTPClient),
	)
}
