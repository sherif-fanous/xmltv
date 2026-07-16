package xmltv

import "encoding/xml"

type Bool bool

func (b *Bool) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if b == nil {
		return xml.Attr{}, nil
	}

	attributeValue := "no"
	if *b {
		attributeValue = "yes"
	}

	return xml.Attr{Name: name, Value: attributeValue}, nil
}

func (b *Bool) UnmarshalXMLAttr(attr xml.Attr) error {
	*b = Bool(attr.Value == "yes")

	return nil
}

func (b *Bool) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if b == nil {
		return e.EncodeElement(nil, start)
	}

	var elementText string
	switch {
	case start.Name.Local == "new":
		elementText = ""
	case bool(*b):
		elementText = "yes"
	default:
		elementText = "no"
	}

	return e.EncodeElement(elementText, start)
}

func (b *Bool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var elementText string
	if err := d.DecodeElement(&elementText, &start); err != nil {
		return err
	}

	*b = Bool(start.Name.Local == "new" || elementText == "yes")

	return nil
}
