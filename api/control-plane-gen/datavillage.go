package api

// if v is empty string, returns empty OptString
func (o OptString) From(v string) OptString {
	o.Value = v
	o.Set = v != ""
	return o
}
