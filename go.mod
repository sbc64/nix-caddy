module caddy

go 1.16

require (
	github.com/caddy-dns/cloudflare v0.0.0-20210607183747-91cf700356a1
	github.com/caddy-dns/route53 v1.1.2
	github.com/caddyserver/caddy/v2 v2.4.3
	github.com/sbc64/caddy-ext v0.0.0-20210407142344-e66a3e677e1f
)

replace github.com/antlr/antlr4/runtime/Go/antlr => github.com/antlr/antlr4 v0.0.0-20210521184019-c5ad59b459ec
