package concat

import (
	//"fmt"
	"github.com/mholt/caddy"
	"github.com/mholt/caddy/caddyhttp/httpserver"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type Concat struct {
	Next httpserver.Handler
}

func (this Concat) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
	if strings.Contains(r.URL.String(), "??") {
		files := strings.Split(strings.Split(r.URL.RawQuery, "?")[1], ",")
		for _, file := range files {
			u := url.URL{
				Path:   path.Join(r.URL.Path, file),
				Host:   r.Host,
				Scheme: "http",
			}
			resp, err := http.Get(u.String())
			if err != nil {
				return 500, err
			}
			io.Copy(w, resp.Body)
			w.Write([]byte("\n"))
		}
		return 200, nil
	}
	return this.Next.ServeHTTP(w, r)
}

func init() {
	caddy.RegisterPlugin("concat", caddy.Plugin{
		ServerType: "http",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {
	cfg := httpserver.GetConfig(c)
	cfg.AddMiddleware(func(next httpserver.Handler) httpserver.Handler {
		return Concat{Next: next}
	})
	return nil
}
