package handler

import (
	"net/http"
	"strconv"
)

// getLimitPage extracts the limit and page values from the request query,
// if the values are not present or invalid, it returns the default values,
// default limit is 0 and default page is 1.
func getLimitPage(r *http.Request) (int32, int32) {
	limit := stringToInt32(r.URL.Query().Get("limit"), 0)
	page := stringToInt32(r.URL.Query().Get("page"), 1)
	if page < 1 {
		page = 1
	}
	return limit, page
}

// stringToInt32 converts a string  to an int32,
// if the string is empty or invalid, it returns the default value.
func stringToInt32(s string, def int32) int32 {
	if s == "" {
		return def
	}
	value, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return def
	}
	return int32(value)
}
