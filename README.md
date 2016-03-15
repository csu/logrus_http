# logrus_http
A Logrus hook to post new entries to an HTTP endpoint.

## Usage
```go
hook, err := NewHttpHook("http://localhost:3000/log", "logContent", m, extras)
log.Hooks.Add(hook)
```