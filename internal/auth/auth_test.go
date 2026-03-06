package auth

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAPIKey(t *testing.T) {
	validHeader := http.Header{}
	invalidHeader := http.Header{}

	validHeader.Set("Authorization", "ApiKey 010")
	invalidHeader.Set("Authorization", "ApiKey")

	tests := map[string]struct {
		header  http.Header
		want    string
		wantErr bool
	}{
		"valid header": {
			header:  validHeader,
			want:    "010",
			wantErr: false,
		},
		"invalid header": {
			header:  invalidHeader,
			wantErr: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.header)
			if (err != nil) != tc.wantErr {
				t.Fatalf("unexprected error: %v", err)
			}
			if tc.wantErr {
				return
			}

			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Error(diff)
			}
		})
	}
}
