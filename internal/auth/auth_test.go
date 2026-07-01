package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestValidateJWT(t *testing.T) {
	userID := uuid.New()

	tests := []struct {
		name      string
		secret    string
		expiresIn time.Duration
		validator string
		wantErr   bool
	}{
		{
			name:      "valid token",
			secret:    "secret",
			expiresIn: time.Hour,
			validator: "secret",
			wantErr:   false,
		},
		{
			name:      "expired token",
			secret:    "secret",
			expiresIn: -time.Hour,
			validator: "secret",
			wantErr:   true,
		},
		{
			name:      "wrong secret",
			secret:    "secret",
			expiresIn: time.Hour,
			validator: "wrong-secret",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := MakeJWT(userID, tt.secret, tt.expiresIn)
			if err != nil {
				t.Fatalf("MakeJWT() error = %v", err)
			}

			gotID, err := ValidateJWT(token, tt.validator)

			if (err != nil) != tt.wantErr {
				t.Fatalf("ValidateJWT() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && gotID != userID {
				t.Fatalf("ValidateJWT() = %v, want %v", gotID, userID)
			}
		})
	}
}

func TestGetBearerToken(t *testing.T) {
	tests := []struct {
		name    string
		header  string
		want    string
		wantErr bool
	}{
		{
			name:    "valid bearer token",
			header:  "Bearer abc123",
			want:    "abc123",
			wantErr: false,
		},
		{
			name:    "missing authorization header",
			header:  "",
			wantErr: true,
		},
		{
			name:    "missing bearer prefix",
			header:  "abc123",
			wantErr: true,
		},
		{
			name:    "extra whitespace",
			header:  "Bearer   abc123   ",
			want:    "abc123",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}

			if tt.header != "" {
				headers.Set("Authorization", tt.header)
			}

			got, err := GetBearerToken(headers)

			if (err != nil) != tt.wantErr {
				t.Fatalf("GetBearerToken() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Fatalf("GetBearerToken() = %q, want %q", got, tt.want)
			}
		})
	}
}
