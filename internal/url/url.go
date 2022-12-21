package url

import "net/url"

func Encode(in string) (string, error) {
	return url.QueryEscape(in), nil
}

func Decode(in string) (string, error) {
	return url.QueryUnescape(in)
}
