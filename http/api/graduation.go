package api

import (
	"encoding/json"
	"net/http"

	"github.com/thoriqadillah/gown/http/model"
)

type GraduationAPI struct {
	client *http.Client
	url    string
	data   model.Data
}

func New(url string) GraduationAPI {
	return GraduationAPI{
		client: &http.Client{},
		url:    url,
		data:   model.Data{},
	}
}

func (g *GraduationAPI) GetGraduees(year string) *model.Data {
	res, err := g.client.Get(g.url + "&q=" + year)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&g.data); err != nil {
		panic(err)
	}

	return &g.data
}
