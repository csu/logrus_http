package logrus_http

import (
  "github.com/Sirupsen/logrus"
  "testing"
)

func TestPrint(t *testing.T) {
  log := logrus.New()
  log.Formatter = new(logrus.JSONFormatter)

  m := make(map[string]string)
  m["secret"] = "example-secret-here"

  extras := make(map[string]interface{})
  
  hook, err := NewHttpHook("http://logserver.christopher.su/log", "logContent", m, extras)
  if err != nil {
    t.Errorf("Unable to create hook.")
  }

  log.Hooks.Add(hook)

  log.Info("It worked!")
}