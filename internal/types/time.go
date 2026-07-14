package types

import (
	"encoding/xml"
	"fmt"
	"time"
)

type XMLTVTime struct{ time.Time }

// timeLayouts lists the accepted XMLTV date/time layouts, ordered from most to
// least specific. Per the XMLTV DTD, dates are 'YYYYMMDDhhmmss' or any initial
// substring (e.g. 'YYYYMM'), optionally followed by a numeric timezone offset.
// If no explicit timezone is given, UTC is assumed.
var timeLayouts = []string{
	"20060102150405 -0700",
	"20060102150405",
	"200601021504 -0700",
	"200601021504",
	"2006010215 -0700",
	"2006010215",
	"20060102 -0700",
	"20060102",
	"200601 -0700",
	"200601",
	"2006 -0700",
	"2006",
}

// parseXMLTVTime parses value against the accepted XMLTV layouts.
func parseXMLTVTime(value string) (time.Time, error) {
	for _, layout := range timeLayouts {
		if t, err := time.Parse(layout, value); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("xmltv: unable to parse time %q", value)
}

func (t *XMLTVTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if t == nil {
		return xml.Attr{}, nil
	}

	if t.IsZero() {
		return xml.Attr{}, nil
	}

	return xml.Attr{Name: name, Value: t.Format("20060102150405 -0700")}, nil
}

func (t *XMLTVTime) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value == "" {
		*t = XMLTVTime{}

		return nil
	}

	tt, err := parseXMLTVTime(attr.Value)
	if err != nil {
		return err
	}

	*t = XMLTVTime{tt}

	return nil
}

func (t *XMLTVTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t == nil {
		return e.EncodeElement(nil, start)
	}

	if t.IsZero() {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(t.Format("20060102"), start)
}

func (t *XMLTVTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}

	if v == "" {
		*t = XMLTVTime{}

		return nil
	}

	tt, err := parseXMLTVTime(v)
	if err != nil {
		return err
	}

	*t = XMLTVTime{tt}

	return nil
}
