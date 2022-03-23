package message

import (
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

// StructRecord wraps a struct and implements the Record interface
// based on the given tag lookup as the "Keys". For example: many
// db drivers use a 'db' tag on struct fields to know how to translate
// to the database column. the GetKeys() call of this returns the 'db' tag values.
// You should always use the NewStructRecord constructor to create this.
type StructRecord struct {
	tagName    string
	record     interface{} // record holds the struct to do the tag lookup on
	tags       []string
	tagsToName map[string]string
}

// ErrNotAStruct is for when the provided arg is not a struct
var ErrNotAStruct = errors.New("not a struct")

// NewStructRecord createa a new StructRecord. The tagName arg
// is optional and will be used instead of the default field name.
// While the tagName arg is a slice, only the [0] value is used.
// The value of the tag for a given field is ignore if it is "" or "-".
// and will be skipped. If the tag value has a "," in it, the part
// before comma is used as the tag value. This allows for values like "id,omitempty"
func NewStructRecord(strct interface{}, tagName ...string) (StructRecord, error) {
	tag := ""
	if len(tagName) > 0 {
		tag = tagName[0]
	}

	// ensure that it is a struct we are working with
	t := reflect.TypeOf(strct)
	if t.Kind() == reflect.Ptr {
		t = t.Elem() // get the type the pointer points to
	}
	if t.Kind() != reflect.Struct {
		return StructRecord{}, ErrNotAStruct
	}

	// extract the tags
	tags := []string{}
	tagsToName := map[string]string{}

	for _, sf := range reflect.VisibleFields(t) {
		if sf.IsExported() && !sf.Anonymous {
			tags = append(tags, extract(sf, tag, tagsToName)...)
		}
	}

	rec := StructRecord{tagName: tag, record: strct, tags: tags, tagsToName: tagsToName}
	return rec, nil
}

func extract(f reflect.StructField, tag string, tagsToName map[string]string) []string {
	tags := []string{}
	tagVal := f.Name
	if tag != "" {
		tagVal = strings.Split(f.Tag.Get(tag), ",")[0]
		if tagVal == "" || tagVal == "-" {
			// don't add if there was not tag value for a specified tag
			return tags
		}
	}
	tags = append(tags, tagVal)
	tagsToName[tagVal] = f.Name

	return tags
}

// In implements the Inner interface
// and returns the original struct that this wraps.
func (sr StructRecord) In() interface{} {
	return sr.record
}

// Get implements the Record interface
func (sr StructRecord) Get(key string) (interface{}, bool) {
	if name, ok := sr.tagsToName[key]; ok {
		r := reflect.ValueOf(sr.record)
		return reflect.Indirect(r).FieldByName(name).Interface(), true
	}
	return nil, false
}

// GetKeys implements the Record interface
func (sr StructRecord) GetKeys() []string { return sr.tags }

// GetVals implements the Record interface
func (sr StructRecord) GetVals() []interface{} {
	vals := []interface{}{}
	for _, key := range sr.tags {
		if val, ok := sr.Get(key); ok {
			vals = append(vals, val)
		} else {
			// not ok so for some reason
			// this should never happen but if it does, return nil
			return nil
		}
	}
	return vals
}
