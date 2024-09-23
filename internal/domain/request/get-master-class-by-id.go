package request

import (
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type GetMasterClassByID struct {
	MasterClassId int `json:"master_class_id"`
}

func NewGetMasterClassByID() *GetMasterClassByID {
	return &GetMasterClassByID{}
}

func (d *GetMasterClassByID) Bind(r *http.Request) (err error) {
	d.MasterClassId, err = strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	return nil
}
