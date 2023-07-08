package jsonptr

// DEPS: go install github.com/0x51-dev/upeg/cmd/abnf
//go:generate abnf --in=abnf/uri.abnf --out=abnf/uri.go --ignore=pct-encoded,unreserved,reserved,gen-delims,sub-delims --custom=IPv6address --importCore
