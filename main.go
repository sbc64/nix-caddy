package main

import (
	_ "github.com/caddy-dns/cloudflare"
    _ "github.com/sbc64/caddy-ext"
	caddycmd "github.com/caddyserver/caddy/v2/cmd"
	_ "github.com/caddyserver/caddy/v2/modules/standard"
)

func main() {
	caddycmd.Main()
}
