package go_http_pool

import (
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	app := New()
	cli := app.GetClient()
	for i := 0; i < 10; i++ {
		res, err := cli.Get("https://google.com")
		if err != nil {
			t.Errorf(err.Error())
			return
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("HTTP_ERR")

			return

		}

	}

}
