package ri

// DEPS: go install github.com/0x51-dev/upeg/cmd/abnf
//go:generate abnf --in=abnf/uri.abnf --out=uri/abnf.go --package=uri --ignore=pct-encoded,unreserved,reserved,gen-delims,sub-delims --custom=IPv6address --importCore
//go:generate abnf --in=abnf/iri.abnf --out=iri/abnf.go --package=iri --dependencies=github.com/0x51-dev/rids/uri,scheme,port,pct-encoded,sub-delims,IP-literal,IPv4address --importCore
