package unsplash

import (
	"os"
	"testing"
)

func TestValidEnvKey(t *testing.T) {
	var (
		err error
		api Api
	)
	const fakeEnvKey = "validEnvKey"
	err = os.Setenv("UNSPLASH_API_ACCESS_KEY", fakeEnvKey)
	if err != nil {
		t.Fatalf("Setenv returned an error: %v\n", err)
	}
	api, err = New()
	if err != nil {
		t.Fatalf("err is not nil: %v\n", err)
	}
	if api.apiKey != fakeEnvKey {
		t.Fatalf("Api key should be %v but got '%v' instead\n", fakeEnvKey, api.apiKey)
	}
}

func TestMissingEnvKey(t *testing.T) {
	_, err := New()
	if err != nil {
		t.Fatal("err is nil and should not be\n")
	}
}
