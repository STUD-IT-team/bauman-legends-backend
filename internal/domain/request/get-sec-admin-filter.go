package request

import (
	"net/http"
	"strconv"
)

type GetSecAdminByFilter struct {
	Limit  int
	Offset int
}

func NewGetSecAdminByFilter() *GetSecAdminByFilter {
	return &GetSecAdminByFilter{}
}

func (g *GetSecAdminByFilter) Bind(r *http.Request) (err error) {
	values := r.URL.Query()

	if values.Has("limit") {
		g.Limit, err = strconv.Atoi(values.Get("limit"))
		if err != nil {
			return err
		}
	} else {
		g.Limit = 5
	}
	if values.Has("offset") {
		g.Offset, err = strconv.Atoi(values.Get("offset"))
		if err != nil {
			return err
		}
	} else {
		g.Offset = 0
	}

	return nil
}
