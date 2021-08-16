package jinx

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func ResponseJSONPayload(w http.ResponseWriter, r *http.Request, code int, payload interface{}) error {
	null := make(map[string]interface{})
	resp := &Response{
		Version: Version{
			Number: "0.1.0",
		},
		Meta:       Meta{Code: http.StatusText(code)},
		Data:       payload,
		Pagination: null,
	}

	if ver, ok := r.Context().Value(CtxVersion).(Version); ok {
		resp.Version = ver
	}
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)

	w.Header().Set(HeaderContentType, MIMEApplicationJSONCharsetUTF8)

	if err := enc.Encode(resp); err != nil {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	w.WriteHeader(code)
	_, err := w.Write(buf.Bytes())
	if err != nil {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	return nil
}
