package request

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type GetTeamById struct {
	Id int `json:"id"`
}

func (t *GetTeamById) Bind(req *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		return fmt.Errorf("can't Atoi id on DeleteFeed.Bind: %w", err)
	}

	t.Id = id

	return t.validate()
}

func (t *GetTeamById) validate() error {
	if t.Id == 0 {
		return fmt.Errorf("team id is required")
	}

	return nil
}
