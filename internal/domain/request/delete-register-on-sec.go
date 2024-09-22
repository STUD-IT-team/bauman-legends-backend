package request

import (
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type DeleteRegisterOnSecFilter struct {
	MasterClassId int `json:"master_class_id"`
}

func NewDeleteRegisterOnSecFilter() *DeleteRegisterOnSecFilter {
	return &DeleteRegisterOnSecFilter{}
}

func (d *DeleteRegisterOnSecFilter) Bind(r *http.Request) (err error) {
	d.MasterClassId, err = strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	return nil
}
