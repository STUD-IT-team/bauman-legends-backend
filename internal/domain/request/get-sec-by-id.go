package request

import (
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type GetSecByIdFilter struct {
	Id int `json:"id"`
}

func NewGetSecByIdFilter() *GetSecByIdFilter {
	return &GetSecByIdFilter{}
}

func (g *GetSecByIdFilter) Bind(r *http.Request) (err error) {
	g.Id, err = strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	return nil
}
