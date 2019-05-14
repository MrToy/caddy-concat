package main

import (
	"github.com/mholt/caddy/caddy/caddymain"
	// plug in plugins here, for example:
	_ "github.com/captncraig/cors/caddy"
	_ "github.com/mrtoy/caddy-concat"
)

func main() {
	// optional: disable telemetry
	// caddymain.EnableTelemetry = false
	caddymain.Run()
}
