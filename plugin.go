package concat

import (
	//"fmt"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/mholt/caddy"
	"github.com/mholt/caddy/caddyhttp/httpserver"
)

type Concat struct {
	Next httpserver.Handler
}

func (ctx Concat) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
	if strings.Contains(r.RequestURI, "??") {
		chunk := strings.Split(r.RequestURI, "??")
		files := strings.Split(chunk[1], ",")
		for _, file := range files {
			u := url.URL{
				Path:   path.Join(chunk[0], file),
				Host:   r.Host,
				Scheme: "http",
			}
			resp, err := http.Get(u.String())
			if err != nil {
				fmt.Println(err)
				return 500, err
			}
			io.Copy(w, resp.Body)
			w.Write([]byte("\n"))
		}
		return 200, nil
	}
	return ctx.Next.ServeHTTP(w, r)
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
