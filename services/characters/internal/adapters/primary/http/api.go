package http

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/AranGarcia/droop/characters/internal/ports/api"
)

func getStringParam(q url.Values, paramName, defaultValue string) string {
	value := q.Get(paramName)
	if value == "" {
		return defaultValue
	}
	return value
}

func getIntParam(q url.Values, paramName string, defaultValue int) (int, error) {
	value := q.Get(paramName)
	if value == "" {
		return defaultValue, nil
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, InvalidInputError{
			Field:  paramName,
			Reason: "not an integer"}
	}
	return intValue, nil
}

func listCharactersRequestFromHTTP(r *http.Request) (*api.ListCharactersRequest, error) {
	apiRequest := &api.ListCharactersRequest{}
	if r == nil {
		return apiRequest, nil
	}
	var err error
	q := r.URL.Query()
	apiRequest.Sort = api.ListSortType(getStringParam(q, "sort", string(api.CreatedAt)))
	apiRequest.Size, err = getIntParam(q, "size", 10)
	if err != nil {
		return nil, err
	}
	apiRequest.Offset, err = getIntParam(q, "offset", 0)
	if err != nil {
		return nil, err
	}
	return apiRequest, nil
}
