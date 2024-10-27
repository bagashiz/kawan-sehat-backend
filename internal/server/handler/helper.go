package handler

import (
	"net/http"
	"strconv"
)

// getLimitOffset extracts the limit and offset values from the request query.
func getLimitOffset(r *http.Request) (int32, int32) {
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")
	limit32, err := stringToInt32(limit, 0)
	if err != nil {
		limit32 = 0
	}
	offset32, err := stringToInt32(offset, 0)
	if err != nil {
		offset32 = 0
	}
	return limit32, offset32
}

// stringToInt32 converts a string to an int32,
// if the string is empty, it returns the default value.
func stringToInt32(s string, def int32) (int32, error) {
	if s == "" {
		return def, nil
	}
	i64, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(i64), nil
}
