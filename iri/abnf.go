// DO NOT EDIT: This file is autogenerated by https://github.com/0x51-dev/upeg.
package iri

import (
	"github.com/0x51-dev/rids/uri"
	. "github.com/0x51-dev/upeg/abnf/core"
	"github.com/0x51-dev/upeg/parser/op"
)

var (
	IRI           = op.Capture{Name: "IRI", Value: op.And{uri.Scheme, ':', IhierPart, op.Optional{Value: op.And{'?', Iquery}}, op.Optional{Value: op.And{'#', Ifragment}}}}
	IhierPart     = op.Capture{Name: "IhierPart", Value: op.Or{op.And{"//", Iauthority, IpathAbempty}, IpathAbsolute, IpathRootless, IpathEmpty}}
	IRIReference  = op.Capture{Name: "IRIReference", Value: op.Or{IRI, IrelativeRef}}
	AbsoluteIRI   = op.Capture{Name: "AbsoluteIRI", Value: op.And{uri.Scheme, ':', IhierPart, op.Optional{Value: op.And{'?', Iquery}}}}
	IrelativeRef  = op.Capture{Name: "IrelativeRef", Value: op.And{IrelativePart, op.Optional{Value: op.And{'?', Iquery}}, op.Optional{Value: op.And{'#', Ifragment}}}}
	IrelativePart = op.Capture{Name: "IrelativePart", Value: op.Or{op.And{"//", Iauthority, IpathAbempty}, IpathAbsolute, IpathNoscheme, IpathEmpty}}
	Iauthority    = op.Capture{Name: "Iauthority", Value: op.And{op.Optional{Value: op.And{Iuserinfo, '@'}}, Ihost, op.Optional{Value: op.And{':', uri.Port}}}}
	Iuserinfo     = op.Capture{Name: "Iuserinfo", Value: op.ZeroOrMore{Value: op.Or{Iunreserved, uri.PctEncoded, uri.SubDelims, ':'}}}
	Ihost         = op.Capture{Name: "Ihost", Value: op.Or{uri.IPLiteral, uri.IPv4address, IregName}}
	IregName      = op.Capture{Name: "IregName", Value: op.ZeroOrMore{Value: op.Or{Iunreserved, uri.PctEncoded, uri.SubDelims}}}
	Ipath         = op.Capture{Name: "Ipath", Value: op.Or{IpathAbempty, IpathAbsolute, IpathNoscheme, IpathRootless, IpathEmpty}}
	IpathAbempty  = op.Capture{Name: "IpathAbempty", Value: op.ZeroOrMore{Value: op.And{'/', Isegment}}}
	IpathAbsolute = op.Capture{Name: "IpathAbsolute", Value: op.And{'/', op.Optional{Value: op.And{IsegmentNz, op.ZeroOrMore{Value: op.And{'/', Isegment}}}}}}
	IpathNoscheme = op.Capture{Name: "IpathNoscheme", Value: op.And{IsegmentNzNc, op.ZeroOrMore{Value: op.And{'/', Isegment}}}}
	IpathRootless = op.Capture{Name: "IpathRootless", Value: op.And{IsegmentNz, op.ZeroOrMore{Value: op.And{'/', Isegment}}}}
	IpathEmpty    = op.Capture{Name: "IpathEmpty", Value: op.Repeat{Min: 0, Max: 0, Value: Ipchar}}
	Isegment      = op.Capture{Name: "Isegment", Value: op.ZeroOrMore{Value: Ipchar}}
	IsegmentNz    = op.Capture{Name: "IsegmentNz", Value: op.OneOrMore{Value: Ipchar}}
	IsegmentNzNc  = op.Capture{Name: "IsegmentNzNc", Value: op.OneOrMore{Value: op.Or{Iunreserved, uri.PctEncoded, uri.SubDelims, '@'}}}
	Ipchar        = op.Capture{Name: "Ipchar", Value: op.Or{Iunreserved, uri.PctEncoded, uri.SubDelims, ':', '@'}}
	Iquery        = op.Capture{Name: "Iquery", Value: op.ZeroOrMore{Value: op.Or{Ipchar, Iprivate, '/', '?'}}}
	Ifragment     = op.Capture{Name: "Ifragment", Value: op.ZeroOrMore{Value: op.Or{Ipchar, '/', '?'}}}
	Iunreserved   = op.Capture{Name: "Iunreserved", Value: op.Or{ALPHA, DIGIT, '-', '.', '_', '~', Ucschar}}
	Ucschar       = op.Capture{Name: "Ucschar", Value: op.Or{op.RuneRange{Min: 0xA0, Max: 0xD7FF}, op.RuneRange{Min: 0xF900, Max: 0xFDCF}, op.RuneRange{Min: 0xFDF0, Max: 0xFFEF}, op.RuneRange{Min: 0x10000, Max: 0x1FFFD}, op.RuneRange{Min: 0x20000, Max: 0x2FFFD}, op.RuneRange{Min: 0x30000, Max: 0x3FFFD}, op.RuneRange{Min: 0x40000, Max: 0x4FFFD}, op.RuneRange{Min: 0x50000, Max: 0x5FFFD}, op.RuneRange{Min: 0x60000, Max: 0x6FFFD}, op.RuneRange{Min: 0x70000, Max: 0x7FFFD}, op.RuneRange{Min: 0x80000, Max: 0x8FFFD}, op.RuneRange{Min: 0x90000, Max: 0x9FFFD}, op.RuneRange{Min: 0xA0000, Max: 0xAFFFD}, op.RuneRange{Min: 0xB0000, Max: 0xBFFFD}, op.RuneRange{Min: 0xC0000, Max: 0xCFFFD}, op.RuneRange{Min: 0xD0000, Max: 0xDFFFD}, op.RuneRange{Min: 0xE1000, Max: 0xEFFFD}}}
	Iprivate      = op.Capture{Name: "Iprivate", Value: op.Or{op.RuneRange{Min: 0xE000, Max: 0xF8FF}, op.RuneRange{Min: 0xF0000, Max: 0xFFFFD}, op.RuneRange{Min: 0x100000, Max: 0x10FFFD}}}
)
