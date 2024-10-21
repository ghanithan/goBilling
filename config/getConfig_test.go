package config

import (
	"reflect"
	"testing"
)

func TestGetConfig(t *testing.T) {
	t.Run(
		"Testing getConfig",
		func(t *testing.T) {
			want := &Config{
				Server: ServerConfig{
					Host: "http://localhost",
					Port: "3000",
				},
			}
			got, err := GetConfig()
			if err != nil {
				t.Fatalf("Error in fetching config: %s", err)
			}
			if !reflect.DeepEqual(want, got) {
				t.Fatalf("expected %q got %q", want, got)

			}
		},
	)
}