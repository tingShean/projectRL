package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"io/ioutil"
	"time"
)

func TestHello_world(t *testing.T) {
	// Create a request to pass to our handler.

	for i := 0; i < 61; i++ {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		h_test := httptest.NewRecorder()
		handler := http.HandlerFunc(Hello_world)

		handler.ServeHTTP(h_test, req)

		res := h_test.Result()
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
		if status := h_test.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	}

	time.Sleep(time.Second * 61)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	h_test := httptest.NewRecorder()
	handler := http.HandlerFunc(Hello_world)

	handler.ServeHTTP(h_test, req)

	res := h_test.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	if status := h_test.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
