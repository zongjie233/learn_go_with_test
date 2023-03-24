package v11Select

import "testing"

func TestWebsiteRacer(t *testing.T) {
	slowURL := "www.baidu.com"
	fastURL := "www.bing.com"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got '%s',want '%s'", got, want)
	}

}
