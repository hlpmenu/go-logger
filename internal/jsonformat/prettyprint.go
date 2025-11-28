package jsonformat

import (
	"bytes"
	"encoding/json"
	"io"

	"gopkg.hlmpn.dev/pkg/xprint"
)

// PrettyPrintJSONIface takes an interface and returns pretty-printed JSON bytes.
func PrettyPrintJSONIface(j any) io.Reader {
	var buf = &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetIndent("", "  ")
	if err := enc.Encode(j); err != nil {
		buf.Reset()
		xprint.Fprintf(buf, "failed to marshal to indented json: %v", err)
	}
	return buf
}

// PrettyPrintJSON takes raw JSON bytes and returns pretty-printed JSON bytes.
func PrettyPrintJSON(j []byte) io.Reader {
	var buf = &bytes.Buffer{}
	err := json.Indent(buf, j, "", "  ")
	if err != nil {
		buf.Reset()
		xprint.Fprintf(buf, "failed to indent json: %v", err)
	}
	return buf
}

// PrettyPrintJSONFromReader reads raw JSON from a reader and returns pretty-printed JSON bytes.
func PrettyPrintJSONFromReader(r io.Reader) io.Reader {
	var v any
	if err := json.NewDecoder(r).Decode(&v); err != nil {
		var errBuf = &bytes.Buffer{}
		xprint.Fprintf(errBuf, "failed to decode json from reader: %v", err)
		return errBuf
	}
	return PrettyPrintJSONIface(v)
}
