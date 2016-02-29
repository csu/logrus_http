package logrus_http

import (
  "fmt"
  "github.com/Sirupsen/logrus"
  "net/http"
  "net/url"
  "os"
)

type HttpHook struct {
  RequestEndpoint string
  RequestFormKey string
  RequestExtraFields map[string]string
}

// Creates a hook to be added to an instance of logger. This is called with
// `hook, err := NewHttpHook("http://log-server/post_new_log", "logBody")`
// `if err == nil { log.Hooks.Add(hook) }`
func NewHttpHook(endpoint string, formKey string, extraFields map[string]string) (*HttpHook, error) {
  return &HttpHook{endpoint, formKey, extraFields}, nil
}

func (hook *HttpHook) Fire(entry *logrus.Entry) error {
  line, err := entry.String()
  if err != nil {
    fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
    return err
  }

  reqForm := url.Values{}

  // add in extra fields, if any
  for k, v := range hook.RequestExtraFields {
    reqForm.Set(k, v)
  }

  // add log line
  reqForm.Set(hook.RequestFormKey, line)

  resp, err = http.PostForm(hook.RequestEndpoint, reqForm)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Unable to post line to server, %v", err)
    return err
  }
  resp.Body.Close()
  return nil
}

func (hook *HttpHook) Levels() []logrus.Level {
  return logrus.AllLevels
}
