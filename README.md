# errlist

Group a few errors in one error.

## Example

```go
type query struct {
	url.Values
	errs []error
}

func (q *query) Error() error { return errlist.Error(q.errs) }

func (q *query) Int(name string) int {
	if value := q.Get(name); value != "" {
		i, err := strconv.Atoi(value)
		q.errs = errlist.Append(q.errs, err)
		return i
	}
	return 0
}

func handler(r *http.Request, w http.ResponseWriter) {
	q := query{Values: r.URL.Query()}
	a := q.Int("a")
	b := q.Int("b")
	c := q.Int("c")

	if err := q.Err(); err != nil {
		// Here we return a error combined of the errors of parsing a, b, c params.
		http.Error(w, err, http.StatusBadRequest)
		return
	}

	// ...
}
```
