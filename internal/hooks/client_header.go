package hooks

import (
	"net/http"
	"strings"

	"github.com/sfcompute/sfc-go/internal/config"
)

// sfcClientHeader is the canonical header used by SFC services to identify
// SDK clients. The server-side parser relies on this header rather than the
// generated User-Agent so that our parsing does not break if Speakeasy
// changes the User-Agent format.
const sfcClientHeader = "x-sfc-client"

// ClientHeaderHook sets the x-sfc-client header on every outgoing request.
// The version is captured from the generated User-Agent at SDK init time.
type ClientHeaderHook struct {
	headerValue string
}

func (h *ClientHeaderHook) SDKInit(c config.SDKConfiguration) config.SDKConfiguration {
	h.headerValue = "speakeasy/go-" + parseSDKVersion(c.UserAgent)
	return c
}

func (h *ClientHeaderHook) BeforeRequest(_ BeforeRequestContext, req *http.Request) (*http.Request, error) {
	req.Header.Set(sfcClientHeader, h.headerValue)
	return req, nil
}

// parseSDKVersion extracts the SDK version (second token) from the
// Speakeasy-generated User-Agent string of the form:
//
//	"speakeasy-sdk/go <sdkVersion> <generatorVersion> <docVersion> <module>"
//
// Returns "unknown" if the User-Agent is empty or does not have at least two
// space-separated tokens.
func parseSDKVersion(userAgent string) string {
	parts := strings.Fields(userAgent)
	if len(parts) < 2 {
		return "unknown"
	}
	return parts[1]
}
