package request

import (
	"errors"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type CreateRegisterOnSecFilter struct {
	Id   int    `json:"id"`
	Time string `json:"time"`
}

func NewCreateRegisterOnSecFilter() *CreateRegisterOnSecFilter {
	return &CreateRegisterOnSecFilter{}
}

func (d *CreateRegisterOnSecFilter) Bind(r *http.Request) (err error) {
	values := r.URL.Query()

	if values.Has("time") {
		d.Time = values.Get("time")
	} else {
		return errors.New("time param is required")
	}

	d.Id, err = strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	return nil
}
