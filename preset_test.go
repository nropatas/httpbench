package httpbench

import (
	"net/http"
	"testing"

	"github.com/nropatas/httpbench/syncedtrace"
)

func generateRequestFunc(url string) func(string) (*http.Request, error) {
	return func(uniqueId string) (*http.Request, error) {
		return http.NewRequest(http.MethodGet, url, nil)
	}
}

func TestPreset_NewCorrectUrl(t *testing.T) {
	waitHook := syncedtrace.TLSHandshakeDone
	url := "http://Sheker.com"
	p := New(generateRequestFunc(url), waitHook, nil)

	req, _ := p.NewRequest("test")
	if req.URL.String() != url {
		t.Error("Preset new request URL is different than the one we set")
	}
}

func TestPreset_ResultChCreated(t *testing.T) {
	waitHook := syncedtrace.TLSHandshakeDone
	p := New(generateRequestFunc("http://Sheker.com"), waitHook, nil)

	if p.ResultCh == nil {
		t.Error("Preset result channel wasn't created")
	}
}
