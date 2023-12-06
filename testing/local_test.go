package testing

import (
	"net/http"
	"testing"
)

const base = "http://127.0.0.1:80"

func TestAPI(t *testing.T) {
	c := http.Client{}

	resp, err := c.Get(base + "/v1/health")
	if err != nil {
		t.Logf("skipping: %s", err.Error())
		t.Skip()
		return
	}
	if resp.StatusCode != http.StatusOK {
		t.Logf("skipping: HTTP Status %s", resp.Status)
		t.Skip()
		return
	}

	// TODO: implement testing make api calls
}
