// Code generated by Gojay. DO NOT EDIT.

package annotated_struct

import (
	"database/sql"
	"github.com/orbs-network/gojay"
	"time"
)

type Ints []int

// UnmarshalJSONArray decodes JSON array elements into slice
func (a *Ints) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value int
	if err := dec.Int(&value); err != nil {
		return err
	}
	*a = append(*a, value)
	return nil
}

// MarshalJSONArray encodes arrays into JSON
func (a Ints) MarshalJSONArray(enc *gojay.Encoder) {
	for _, item := range a {
		enc.Int(item)
	}
}

// IsNil checks if array is nil
func (a Ints) IsNil() bool {
	return len(a) == 0
}

type Float32s []float32

// UnmarshalJSONArray decodes JSON array elements into slice
func (a *Float32s) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value float32
	if err := dec.Float32(&value); err != nil {
		return err
	}
	*a = append(*a, value)
	return nil
}

// MarshalJSONArray encodes arrays into JSON
func (a Float32s) MarshalJSONArray(enc *gojay.Encoder) {
	for _, item := range a {
		enc.Float32(item)
	}
}

// IsNil checks if array is nil
func (a Float32s) IsNil() bool {
	return len(a) == 0
}

type SubMessagesPtr []*SubMessage

func (s *SubMessagesPtr) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value = &SubMessage{}
	if err := dec.Object(value); err != nil {
		return err
	}
	*s = append(*s, value)
	return nil
}

func (s SubMessagesPtr) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range s {
		enc.Object(s[i])
	}
}

func (s SubMessagesPtr) IsNil() bool {
	return len(s) == 0
}

type SubMessages []SubMessage

func (s *SubMessages) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value = SubMessage{}
	if err := dec.Object(&value); err != nil {
		return err
	}
	*s = append(*s, value)
	return nil
}

func (s SubMessages) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range s {
		enc.Object(&s[i])
	}
}

func (s SubMessages) IsNil() bool {
	return len(s) == 0
}

// MarshalJSONObject implements MarshalerJSONObject
func (p *Payload) MarshalJSONObject(enc *gojay.Encoder) {

}

// IsNil checks if instance is nil
func (p *Payload) IsNil() bool {
	return p == nil
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (p *Payload) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {

	switch k {

	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (p *Payload) NKeys() int { return 0 }

// MarshalJSONObject implements MarshalerJSONObject
func (m *Message) MarshalJSONObject(enc *gojay.Encoder) {
	enc.IntKey("id", m.Id)
	enc.StringKey("name", m.Name)
	enc.Float64Key("price", m.Price)
	var intsSlice = Ints(m.Ints)
	enc.ArrayKey("ints", intsSlice)
	var floatsSlice = Float32s(m.Floats)
	enc.ArrayKey("floats", floatsSlice)
	enc.ObjectKey("subMessageX", m.SubMessageX)
	var messagesXSlice = SubMessagesPtr(m.MessagesX)
	enc.ArrayKey("messagesX", messagesXSlice)
	enc.ObjectKey("SubMessageY", &m.SubMessageY)
	var messagesYSlice = SubMessages(m.MessagesY)
	enc.ArrayKey("MessagesY", messagesYSlice)
	enc.BoolKey("enabled", *m.IsTrue)
	var payloadSlice = gojay.EmbeddedJSON(m.Payload)
	enc.AddEmbeddedJSONKey("data", &payloadSlice)
	if m.SQLNullString != nil {
		enc.SQLNullStringKey("sqlNullString", m.SQLNullString)
	}
}

// IsNil checks if instance is nil
func (m *Message) IsNil() bool {
	return m == nil
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (m *Message) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {

	switch k {
	case "id":
		return dec.Int(&m.Id)

	case "name":
		return dec.String(&m.Name)

	case "price":
		return dec.Float64(&m.Price)

	case "ints":
		var aSlice = Ints{}
		err := dec.Array(&aSlice)
		if err == nil && len(aSlice) > 0 {
			m.Ints = []int(aSlice)
		}
		return err

	case "floats":
		var aSlice = Float32s{}
		err := dec.Array(&aSlice)
		if err == nil && len(aSlice) > 0 {
			m.Floats = []float32(aSlice)
		}
		return err

	case "subMessageX":
		var value = &SubMessage{}
		err := dec.Object(value)
		if err == nil {
			m.SubMessageX = value
		}

		return err

	case "messagesX":
		var aSlice = SubMessagesPtr{}
		err := dec.Array(&aSlice)
		if err == nil && len(aSlice) > 0 {
			m.MessagesX = []*SubMessage(aSlice)
		}
		return err

	case "SubMessageY":
		err := dec.Object(&m.SubMessageY)

		return err

	case "MessagesY":
		var aSlice = SubMessages{}
		err := dec.Array(&aSlice)
		if err == nil && len(aSlice) > 0 {
			m.MessagesY = []SubMessage(aSlice)
		}
		return err

	case "enabled":
		var value bool
		err := dec.Bool(&value)
		if err == nil {
			m.IsTrue = &value
		}
		return err

	case "data":
		var value = gojay.EmbeddedJSON{}
		err := dec.AddEmbeddedJSON(&value)
		if err == nil && len(value) > 0 {
			m.Payload = Payload(value)
		}
		return err

	case "sqlNullString":
		var value = &sql.NullString{}
		err := dec.SQLNullString(value)
		if err == nil {
			m.SQLNullString = value
		}
		return err

	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (m *Message) NKeys() int { return 12 }

// MarshalJSONObject implements MarshalerJSONObject
func (m *SubMessage) MarshalJSONObject(enc *gojay.Encoder) {
	enc.IntKey("id", m.Id)
	enc.StringKey("description", m.Description)
	enc.TimeKey("startDate", &m.StartTime, "2006-01-02 15:04:05")
	if m.EndTime != nil {
		enc.TimeKey("endDate", m.EndTime, "2006-01-02 15:04:05")
	}
}

// IsNil checks if instance is nil
func (m *SubMessage) IsNil() bool {
	return m == nil
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (m *SubMessage) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {

	switch k {
	case "id":
		return dec.Int(&m.Id)

	case "description":
		return dec.String(&m.Description)

	case "startDate":
		var format = "2006-01-02 15:04:05"
		var value = time.Time{}
		err := dec.Time(&value, format)
		if err == nil {
			m.StartTime = value
		}
		return err

	case "endDate":
		var format = "2006-01-02 15:04:05"
		var value = &time.Time{}
		err := dec.Time(value, format)
		if err == nil {
			m.EndTime = value
		}
		return err

	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (m *SubMessage) NKeys() int { return 4 }
