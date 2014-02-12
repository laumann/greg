package greg

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"bytes"
)

// TODO Construct request and response objects, and feed them to compile
func TestCompileOK(t *testing.T) {
	reqBody := bytes.NewReader([]byte("regex=a*b&regex-input=aaab"))
	r, err := http.NewRequest("POST", "http://localhost:9000/compile", reqBody)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	compile(w, r)

	t.Log("Response:", w.Body.String())

}
