package main

import (
	_ "github.com/caddy-dns/cloudflare"
	_ "github.com/caddy-dns/route53"
	caddycmd "github.com/caddyserver/caddy/v2/cmd"
	_ "github.com/caddyserver/caddy/v2/modules/standard"
	_ "github.com/sbc64/caddy-ext"
)

func main() {
	caddycmd.Main()
}
