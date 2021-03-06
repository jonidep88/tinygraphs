package isogrids

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/taironas/route"
	tgTesting "github.com/taironas/tinygraphs/testing"
)

func TestHexa16(t *testing.T) {
	t.Parallel()
	r := new(route.Router)
	r.HandleFunc("/isogrids/labs/hexa16", Hexa)
	r.HandleFunc("/isogrids/labs/hexa16/:key", Hexa)

	test := tgTesting.GenerateHandlerFunc(t, Hexa)
	for _, p := range tgTesting.GoodParams {
		recorder := test("/isogrids/labs/hexa16", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
		recorder = test("/isogrids/labs/hexa16/somekey", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}

	}

	for _, p := range tgTesting.BadParams {
		recorder := test("/isogrids/labs/hexa16", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
		recorder = test("/isogrids/labs/hexa16/somekey", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}

	}
}

func TestHexa16Cache(t *testing.T) {
	t.Parallel()
	r := new(route.Router)
	r.HandleFunc("/isogrids/labs/hexa16/:key", Hexa16)

	var etag string

	test := tgTesting.GenerateHandlerFunc(t, Hexa16)

	if recorder := test("/isogrids/labs/hexa16/somekey", "GET", nil, r); recorder != nil {
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
		etag = recorder.Header().Get("Etag")
	}

	// test caching
	if req, err := http.NewRequest("GET", "/isogrids/labs/hexa16/somekey", nil); err != nil {
		t.Errorf("%v", err)
	} else {
		req.Header.Set("If-None-Match", etag)
		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusNotModified {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusNotModified)
		}
	}
}
