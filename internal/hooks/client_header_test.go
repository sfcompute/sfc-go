package hooks

import (
	"net/http"
	"testing"

	"github.com/sfcompute/sfc-go/internal/config"
)

func TestClientHeaderHook(t *testing.T) {
	tests := []struct {
		name      string
		userAgent string
		want      string
	}{
		{
			name:      "speakeasy generated user agent",
			userAgent: "speakeasy-sdk/go 0.0.1 2.881.17 0.1.0 github.com/sfcompute/sfc-go",
			want:      "speakeasy/go-0.0.1",
		},
		{
			name:      "empty user agent",
			userAgent: "",
			want:      "speakeasy/go-unknown",
		},
		{
			name:      "single-token user agent",
			userAgent: "speakeasy-sdk/go",
			want:      "speakeasy/go-unknown",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := &ClientHeaderHook{}
			h.SDKInit(config.SDKConfiguration{UserAgent: tc.userAgent})

			req, err := http.NewRequest(http.MethodGet, "https://api.sfcompute.com/v1/foo", nil)
			if err != nil {
				t.Fatalf("NewRequest: %v", err)
			}

			if _, err := h.BeforeRequest(BeforeRequestContext{}, req); err != nil {
				t.Fatalf("BeforeRequest: %v", err)
			}

			if got := req.Header.Get(sfcClientHeader); got != tc.want {
				t.Errorf("x-sfc-client = %q, want %q", got, tc.want)
			}
		})
	}
}
