[![Build Status](https://travis-ci.com/ikawaha/goahttpcheck.svg?branch=master)](https://travis-ci.com/ikawaha/goahttpcheck)

goahttpcheck
---


A test helper for testing endpoints of APIs generated by Goa v3.
This makes it possible to test endpoints using ivpusic/httpcheck.

# Usage

1. New checker.
1. Set the method handler constructor, mounter, and method endpoint in the checker.
1. Register the middleware with the checker's `Use` method (If any).
1. Call checker's `Test` method and test by ivpusic/httpcheck way.

# Example

see. https://github.com/ikawaha/goahttpcheck/blob/master/testdata/calc_test.go

**Design**

```go
var _ = Service("calc", func() {
	Description("The calc service performs operations on numbers.")
	Method("add", func() {
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})
		Result(Int)
		HTTP(func() {
			GET("/add/{a}/{b}")
			Response(StatusOK)
		})
	})
```
**Tests**
```go
import (
...
	"calc/gen/calc"
	"calc/gen/http/calc/server"
)

func TestCalcsrvc_Add(t *testing.T) {
	checker := goahttpcheck.New()
	var logger log.Logger
	checker.Mount(server.NewAddHandler, server.MountAddHandler, calc.NewAddEndpoint(NewCalc(&logger)))

	// see. https://github.com/ikawaha/httpcheck
	checker.Test(t, http.MethodGet, "/add/1/2").
		Check().
		HasStatus(http.StatusOK).
		Cb(func(r *http.Response) {
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("unexpected error, %v", err)
			}
			r.Body.Close()
			if got, expected := strings.TrimSpace(string(b)), "3"; got != expected {
				t.Errorf("got %+v, expected %v", b, expected)
			}
		})
}
```


**Blog**: http://ikawaha.hateblo.jp/entry/2019/12/03/154521

---

License MIT
