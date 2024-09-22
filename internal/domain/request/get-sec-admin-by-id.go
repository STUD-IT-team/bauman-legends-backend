package request

import (
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type GetSecAdminByIdFilter struct {
	Id int `json:"id"`
}

func NewGetSecAdminByIdFilter() *GetSecAdminByIdFilter {
	return &GetSecAdminByIdFilter{}
}

func (g *GetSecAdminByIdFilter) Bind(r *http.Request) (err error) {
	g.Id, err = strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	return nil
}
