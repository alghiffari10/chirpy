package main

import (
	"reflect"
	"testing"
)

func TestBadWordFilter(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"No bad words": {
			input: "I had something interesting for breakfast",
			want:  "I had something interesting for breakfast",
		},
		"Lowercase bad word": {
			input: "I hear Mastodon is better than Chirpy. sharbert I need to migrate",
			want:  "I hear Mastodon is better than Chirpy. **** I need to migrate",
		},
		"Mixed case bad words": {
			input: "I really need a kerfuffle to go to bed sooner, Fornax !",
			want:  "I really need a **** to go to bed sooner, **** !",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := badWordFilter(tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
			}
		})
	}
}
