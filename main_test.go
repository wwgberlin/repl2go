package main_test

import (	
	"net/http"	
	"net/http/httptest"	
	"testing"
	"github.com/wwgberlin/repl2go/handler"
)

func TestRunMainNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/doesnotexist", nil)
	if err != nil {	
		t.Fatal(err)	
	}	
			
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.RunHandler)	

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method	
	// directly and pass in our Request and ResponseRecorder.	
	handler.ServeHTTP(rr, req)	

	// Check the status code is what we expect.	
	if status := rr.Code; status != http.StatusNotFound {		
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound);
	}
}