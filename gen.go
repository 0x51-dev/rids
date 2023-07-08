package ri

// DEPS: go install github.com/0x51-dev/upeg/cmd/abnf
//go:generate abnf --in=abnf/uri.abnf --out=u/abnf.go --package=u --ignore=pct-encoded,unreserved,reserved,gen-delims,sub-delims --custom=IPv6address --importCore
//go:generate abnf --in=abnf/iri.abnf --out=i/abnf.go --package=i --dependencies=github.com/0x51-dev/ri/u,scheme,port,pct-encoded,sub-delims,IP-literal,IPv4address --importCore
