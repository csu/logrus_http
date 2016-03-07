package logrus_http

import (
	"github.com/Sirupsen/logrus"
	"testing"
	"net/http"
	"io"
)

func hello(w http.ResponseWriter, r *http.Request) {
	// TODO: check for secret and extra log keys
	io.WriteString(w, "Log received.")
}

func TestHookCreateAndPost(t *testing.T) {
	// start http server
	mux := http.NewServeMux()
	mux.HandleFunc("/log", hello)
	go http.ListenAndServe(":3000", mux)

	log := logrus.New()
	log.Formatter = new(logrus.JSONFormatter)

	m := make(map[string]string)
	m["secret"] = "example-secret-here"

	extras := make(map[string]interface{})
	
	hook, err := NewHttpHook("http://localhost:3000/log", "logContent", m, extras)
	if err != nil {
		t.Errorf("Unable to create hook.")
	}

	log.Hooks.Add(hook)

	log.Info("It worked!")
}