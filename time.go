package xmltv

import (
	"encoding/xml"
	"fmt"
	"time"
)

type Time struct{ time.Time }

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

// parseTimeValue parses value against the accepted XMLTV layouts.
func parseTimeValue(value string) (time.Time, error) {
	for _, layout := range timeLayouts {
		if t, err := time.Parse(layout, value); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("xmltv: unable to parse time %q", value)
}

func (t *Time) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if t == nil {
		return xml.Attr{}, nil
	}

	if t.IsZero() {
		return xml.Attr{}, nil
	}

	return xml.Attr{Name: name, Value: t.Format("20060102150405 -0700")}, nil
}

func (t *Time) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value == "" {
		*t = Time{}

		return nil
	}

	tt, err := parseTimeValue(attr.Value)
	if err != nil {
		return err
	}

	*t = Time{tt}

	return nil
}

func (t *Time) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t == nil {
		return e.EncodeElement(nil, start)
	}

	if t.IsZero() {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(t.Format("20060102"), start)
}

func (t *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}

	if v == "" {
		*t = Time{}

		return nil
	}

	tt, err := parseTimeValue(v)
	if err != nil {
		return err
	}

	*t = Time{tt}

	return nil
}
