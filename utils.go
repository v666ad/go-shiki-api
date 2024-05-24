package shikimori

import (
	"errors"
	"io"
	"net/http"
)

func reverse[T any](list []T) []T {
	for i, j := 0, len(list)-1; i < j; {
		list[i], list[j] = list[j], list[i]
		i++
		j--
	}
	return list
}

func getErrorFromBadResponse(resp *http.Response) error {
	switch resp.StatusCode {
	case 400:
		return ErrBadRequest
	case 401:
		return ErrUnauthorized
	case 403:
		return ErrForbidden
	case 404:
		return ErrNotFound
	case 422:
		return ErrUnprocessableEntity
	case 429:
		return ErrTooManyRequests
	case 500:
		return ErrInternalServer
	default:
		data, err := io.ReadAll(resp.Body)
		if err != nil && len(data) != 0 {
			resp.Body.Close()
			return errors.New("bad status " + resp.Request.Method + " " + resp.Request.URL.String() + " -> " + resp.Status + "\n" + string(data))
		} else {
			return errors.New("bad status " + resp.Request.Method + " " + resp.Request.URL.String() + " -> " + resp.Status)
		}
	}
}
