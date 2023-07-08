package u

import (
	"fmt"
	"github.com/0x51-dev/upeg/parser"
	"github.com/0x51-dev/upeg/parser/op"
)

var ipv6addressOperator = op.Capture{
	Name: "IPv6address",
	Value: op.Or{
		op.And{op.Repeat{Min: 6, Max: 6, Value: op.And{H16, ':'}}, Ls32},
		op.And{"::", op.Repeat{Min: 5, Max: 5, Value: op.And{H16, ':'}}, Ls32},
		op.And{op.Optional{Value: H16}, "::", op.Repeat{Min: 4, Max: 4, Value: op.And{H16, ':'}}, Ls32},
		op.And{op.Optional{Value: op.And{op.Repeat{Max: 1, Value: op.And{H16, ':', op.Not{Value: ':'}}}, H16}}, "::", op.Repeat{Min: 3, Max: 3, Value: op.And{H16, ':'}}, Ls32},
		op.And{op.Optional{Value: op.And{op.Repeat{Max: 2, Value: op.And{H16, ':', op.Not{Value: ':'}}}, H16}}, "::", op.Repeat{Min: 2, Max: 2, Value: op.And{H16, ':'}}, Ls32},
		op.And{op.Optional{Value: op.And{op.Repeat{Max: 3, Value: op.And{H16, ':', op.Not{Value: ':'}}}, H16}}, "::", H16, ':', Ls32},
		op.And{op.Optional{Value: op.And{op.Repeat{Max: 4, Value: op.And{H16, ':', op.Not{Value: ':'}}}, H16}}, "::", Ls32},
		op.And{op.Optional{Value: op.And{op.Repeat{Max: 5, Value: op.And{H16, ':', op.Not{Value: ':'}}}, H16}}, "::", H16},
		op.And{op.Optional{Value: op.And{op.Repeat{Max: 6, Value: op.And{H16, ':', op.Not{Value: ':'}}}, H16}}, "::"}},
}

type IPv6addressOperator struct{}

func (I IPv6addressOperator) Match(_ parser.Cursor, p *parser.Parser) (end parser.Cursor, err error) {
	return p.Match(ipv6addressOperator)
}

func (I IPv6addressOperator) Parse(p *parser.Parser) (end *parser.Node, err error) {
	return p.Parse(ipv6addressOperator)
}

func (I IPv6addressOperator) String() string {
	return fmt.Sprintf("{%s}", op.StringAny(ipv6addressOperator))
}
