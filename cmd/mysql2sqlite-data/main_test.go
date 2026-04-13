package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCoerceSQLiteValue(t *testing.T) {
	t.Parallel()

	refTime := time.Date(2024, 3, 15, 10, 30, 0, 123456789, time.FixedZone("EST", -5*3600))
	wantTime := refTime.UTC().Format(time.RFC3339Nano)

	tests := []struct {
		name     string
		v        any
		declType string
		want     any
	}{
		// nil passes through
		{name: "nil", v: nil, declType: "TEXT", want: nil},

		// []byte with BLOB type → copied []byte
		{name: "bytes_blob", v: []byte{0xCA, 0xFE}, declType: "BLOB", want: []byte{0xCA, 0xFE}},

		// []byte with non-BLOB type → string
		{name: "bytes_text", v: []byte("hello"), declType: "TEXT", want: "hello"},

		// []byte with BLOB substring in type
		{name: "bytes_mediumblob", v: []byte{0x01}, declType: "MEDIUMBLOB", want: []byte{0x01}},

		// time.Time → UTC RFC3339Nano string
		{name: "time", v: refTime, declType: "TEXT", want: wantTime},

		// bool true → 1
		{name: "bool_true", v: true, declType: "INTEGER", want: 1},

		// bool false → 0
		{name: "bool_false", v: false, declType: "INTEGER", want: 0},

		// int64 passes through
		{name: "int64", v: int64(42), declType: "INTEGER", want: int64(42)},

		// float64 passes through
		{name: "float64", v: float64(3.14), declType: "REAL", want: float64(3.14)},

		// string passes through
		{name: "string", v: "forum post", declType: "TEXT", want: "forum post"},

		// fallback: unknown type → fmt.Sprint
		{name: "fallback_int32", v: int32(7), declType: "INTEGER", want: fmt.Sprint(int32(7))},
	}
	for _, tt := range tests {
		got := coerceSQLiteValue(tt.v, tt.declType)
		switch w := tt.want.(type) {
		case nil:
			if got != nil {
				t.Errorf("coerceSQLiteValue[%s]: got %v (%T), want nil", tt.name, got, got)
			}
		case []byte:
			g, ok := got.([]byte)
			if !ok {
				t.Errorf("coerceSQLiteValue[%s]: got type %T, want []byte", tt.name, got)
			} else if string(g) != string(w) {
				t.Errorf("coerceSQLiteValue[%s]: got %v, want %v", tt.name, g, w)
			}
		default:
			if got != tt.want {
				t.Errorf("coerceSQLiteValue[%s]: got %v (%T), want %v (%T)", tt.name, got, got, tt.want, tt.want)
			}
		}
	}
}

func TestCoerceSQLiteValue_BlobCopiesBuffer(t *testing.T) {
	t.Parallel()
	orig := []byte{0x01, 0x02, 0x03}
	result := coerceSQLiteValue(orig, "BLOB").([]byte)
	// Mutating the original must not affect the returned copy.
	orig[0] = 0xFF
	if result[0] == 0xFF {
		t.Error("coerceSQLiteValue BLOB did not copy the input buffer")
	}
}
