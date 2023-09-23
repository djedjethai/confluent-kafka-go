// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
package recordname

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type BasicPerson struct {
	// friend age
	Number *UnionLongNull `json:"number"`
	// friend name
	Name UnionString `json:"name"`
}

const BasicPersonAvroCRC64Fingerprint = "Y0\xfc0\xae\x13kW"

func NewBasicPerson() BasicPerson {
	r := BasicPerson{}
	r.Number = NewUnionLongNull()

	r.Name = NewUnionString()

	return r
}

func DeserializeBasicPerson(r io.Reader) (BasicPerson, error) {
	t := NewBasicPerson()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeBasicPersonFromSchema(r io.Reader, schema string) (BasicPerson, error) {
	t := NewBasicPerson()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeBasicPerson(r BasicPerson, w io.Writer) error {
	var err error
	err = writeUnionLongNull(r.Number, w)
	if err != nil {
		return err
	}
	err = writeUnionString(r.Name, w)
	if err != nil {
		return err
	}
	return err
}

func (r BasicPerson) Serialize(w io.Writer) error {
	return writeBasicPerson(r, w)
}

func (r BasicPerson) Schema() string {
	return "{\"fields\":[{\"doc\":\"friend age\",\"name\":\"number\",\"type\":[\"long\",\"null\"]},{\"doc\":\"friend name\",\"name\":\"name\",\"type\":[\"string\"]}],\"name\":\"python.test.advanced.basicPerson\",\"type\":\"record\"}"
}

func (r BasicPerson) SchemaName() string {
	return "python.test.advanced.basicPerson"
}

func (_ BasicPerson) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ BasicPerson) SetInt(v int32)       { panic("Unsupported operation") }
func (_ BasicPerson) SetLong(v int64)      { panic("Unsupported operation") }
func (_ BasicPerson) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ BasicPerson) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ BasicPerson) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ BasicPerson) SetString(v string)   { panic("Unsupported operation") }
func (_ BasicPerson) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *BasicPerson) Get(i int) types.Field {
	switch i {
	case 0:
		r.Number = NewUnionLongNull()

		return r.Number
	case 1:
		r.Name = NewUnionString()

		w := types.Record{Target: &r.Name}

		return w

	}
	panic("Unknown field index")
}

func (r *BasicPerson) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *BasicPerson) NullField(i int) {
	switch i {
	case 0:
		r.Number = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ BasicPerson) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ BasicPerson) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ BasicPerson) HintSize(int)                     { panic("Unsupported operation") }
func (_ BasicPerson) Finalize()                        {}

func (_ BasicPerson) AvroCRC64Fingerprint() []byte {
	return []byte(BasicPersonAvroCRC64Fingerprint)
}

func (r BasicPerson) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["number"], err = json.Marshal(r.Number)
	if err != nil {
		return nil, err
	}
	output["name"], err = json.Marshal(r.Name)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *BasicPerson) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["number"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Number); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for number")
	}
	val = func() json.RawMessage {
		if v, ok := fields["name"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Name); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for name")
	}
	return nil
}