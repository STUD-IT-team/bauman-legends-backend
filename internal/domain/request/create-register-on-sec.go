package request

import (
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type CreateRegisterOnSecFilter struct {
	MasterClassId int `json:"master_class_id"`
}

func NewCreateRegisterOnSecFilter() *CreateRegisterOnSecFilter {
	return &CreateRegisterOnSecFilter{}
}

func (d *CreateRegisterOnSecFilter) Bind(r *http.Request) (err error) {
	d.MasterClassId, err = strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	return nil
}
