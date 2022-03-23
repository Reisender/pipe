package message_test

import (
	"reflect"
	"testing"

	"github.com/Reisender/pipe/message"
)

type TestStructWithTags struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type TestStructWithoutTags struct {
	ID         int
	Name       string
	unexported bool
}

type TestStructWithSomeTags struct {
	ID    int    `db:"id"`
	Name  string `db:"name,omitempty"`
	Phone string
}

type TestEmbeddedStructWithTags struct {
	TestStructWithTags
}

type TestEmbeddedStructPtrWithTags struct {
	*TestStructWithTags
}

func TestStructRecord(t *testing.T) {
	t.Run("struct pointer", func(t *testing.T) {
		foo := TestStructWithTags{1, "foo"}
		_, err := message.NewStructRecord(&foo, "db")
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("with tags", func(t *testing.T) {
		foo := TestStructWithTags{1, "foo"}
		r, err := message.NewStructRecord(foo, "db")
		if err != nil {
			t.Error(err)
		}

		wantKeys := []string{"id", "name"}
		gotKeys := r.GetKeys()

		wantVals := []interface{}{1, "foo"}
		gotVals := r.GetVals()

		for i, key := range wantKeys {
			val, ok := r.Get(key)
			if !ok {
				t.Error("couldn't find key", key)
			}
			if !reflect.DeepEqual(val, wantVals[i]) {
				t.Errorf("want %v got %v", wantVals[i], val)
			}
		}

		if !reflect.DeepEqual(wantKeys, gotKeys) {
			t.Errorf("want %v got %v", wantKeys, gotKeys)
		}

		if !reflect.DeepEqual(wantVals, gotVals) {
			t.Errorf("want %v got %v", wantVals, gotVals)
		}
	})

	t.Run("without tags", func(t *testing.T) {
		foo := TestStructWithoutTags{1, "foo", false}
		r, err := message.NewStructRecord(foo)
		if err != nil {
			t.Error(err)
		}

		wantKeys := []string{"ID", "Name"}
		gotKeys := r.GetKeys()

		wantVals := []interface{}{1, "foo"}
		gotVals := r.GetVals()

		for i, key := range wantKeys {
			val, ok := r.Get(key)
			if !ok {
				t.Error("couldn't find key", key)
			}
			if !reflect.DeepEqual(val, wantVals[i]) {
				t.Errorf("want %v got %v", wantVals[i], val)
			}
		}

		if !reflect.DeepEqual(wantKeys, gotKeys) {
			t.Errorf("want %v got %v", wantKeys, gotKeys)
		}

		if !reflect.DeepEqual(wantVals, gotVals) {
			t.Errorf("want %v got %v", wantVals, gotVals)
		}
	})

	t.Run("with some tags", func(t *testing.T) {
		foo := TestStructWithSomeTags{1, "foo", "555-5555"}
		r, err := message.NewStructRecord(foo, "db")
		if err != nil {
			t.Error(err)
		}

		wantKeys := []string{"id", "name"}
		gotKeys := r.GetKeys()

		wantVals := []interface{}{1, "foo"}
		gotVals := r.GetVals()

		for i, key := range wantKeys {
			val, ok := r.Get(key)
			if !ok {
				t.Error("couldn't find key", key)
			}
			if !reflect.DeepEqual(val, wantVals[i]) {
				t.Errorf("want %v got %v", wantVals[i], val)
			}
		}

		if !reflect.DeepEqual(wantKeys, gotKeys) {
			t.Errorf("want %v got %v", wantKeys, gotKeys)
		}

		if !reflect.DeepEqual(wantVals, gotVals) {
			t.Errorf("want %v got %v", wantVals, gotVals)
		}
	})

	t.Run("with embedded struct tags", func(t *testing.T) {
		foo := TestEmbeddedStructWithTags{TestStructWithTags{1, "foo"}}
		r, err := message.NewStructRecord(foo)
		if err != nil {
			t.Error(err)
		}

		wantKeys := []string{"ID", "Name"}
		gotKeys := r.GetKeys()

		wantVals := []interface{}{1, "foo"}
		gotVals := r.GetVals()

		for i, key := range wantKeys {
			val, ok := r.Get(key)
			if !ok {
				t.Error("couldn't find key", key)
			}
			if !reflect.DeepEqual(val, wantVals[i]) {
				t.Errorf("want %v got %v", wantVals[i], val)
			}
		}

		if !reflect.DeepEqual(wantKeys, gotKeys) {
			t.Errorf("want %v got %v", wantKeys, gotKeys)
		}

		if !reflect.DeepEqual(wantVals, gotVals) {
			t.Errorf("want %v got %v", wantVals, gotVals)
		}
	})

	t.Run("with embedded struct ptr tags", func(t *testing.T) {
		foo := TestEmbeddedStructPtrWithTags{&TestStructWithTags{1, "foo"}}
		r, err := message.NewStructRecord(foo)
		if err != nil {
			t.Error(err)
		}

		wantKeys := []string{"ID", "Name"}
		gotKeys := r.GetKeys()

		wantVals := []interface{}{1, "foo"}
		gotVals := r.GetVals()

		for i, key := range wantKeys {
			val, ok := r.Get(key)
			if !ok {
				t.Error("couldn't find key", key)
			}
			if !reflect.DeepEqual(val, wantVals[i]) {
				t.Errorf("want %v got %v", wantVals[i], val)
			}
		}

		if !reflect.DeepEqual(wantKeys, gotKeys) {
			t.Errorf("want %v got %v", wantKeys, gotKeys)
		}

		if !reflect.DeepEqual(wantVals, gotVals) {
			t.Errorf("want %v got %v", wantVals, gotVals)
		}
	})
}
