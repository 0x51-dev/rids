package uri_test

import (
	"github.com/0x51-dev/rids/uri"
	"github.com/0x51-dev/upeg/parser"
	"github.com/0x51-dev/upeg/parser/op"
	"testing"
)

func TestParseURI(t *testing.T) {
	for _, test := range []string{
		"ftp://ftp.is.co.za/rfc/rfc1808.txt",
		"http://www.ietf.org/rfc/rfc2396.txt",
		"ldap://[2001:db8::7]/c=GB?objectClass?one",
		"mailto:John.Doe@example.com",
		"news:comp.infosystems.www.servers.unix",
		"tel:+1-816-555-1212",
		"telnet://192.0.2.16:80/",
		"urn:oasis:names:specification:docbook:dtd:xml:4.1.2",
	} {
		p, err := parser.New([]rune(test))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := p.Parse(op.And{uri.URI, op.EOF{}}); err != nil {
			t.Fatal(err)
		}
	}
}
