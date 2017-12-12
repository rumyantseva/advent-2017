package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rumyantseva/advent-2017/07-version/version"
)

func TestHome(t *testing.T) {
	w := httptest.NewRecorder()
	version.Release = "test version"
	home(w, nil)

	resp := w.Result()
	if have, want := resp.StatusCode, http.StatusOK; have != want {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", have, want)
	}

	greeting, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	info := struct {
		BuildTime string `json:"buildTime"`
		Commit    string `json:"commit"`
		Release   string `json:"release"`
	}{}
	err = json.Unmarshal(greeting, &info)
	if err != nil {
		t.Fatal(err)
	}
	if info.Release != version.Release {
		t.Errorf("Release version is wrong. Have: %s, want: %s", info.Release, version.Release)
	}
}
