package brreg

import (
	"encoding/json"
	"fmt"
	oupdate "github.com/tjololo/nororg-streamer/pkg/brreg/structs"
	"net/http"
	"time"
)

type OrganizationUpdateService struct {
	BaseUrl string
}

func (u *OrganizationUpdateService) FetchUpdatesAfterDate(from time.Time) (updates []oupdate.OrganizationUpdate, err error) {
	getUrl := fmt.Sprintf("%s/enhetsregisteret/api/oppdateringer/enheter?dato=%s", u.BaseUrl, from.Format("2006-01-02T15:04:05.000Z"))
	r, err := http.Get(getUrl)
	if err != nil {
		return
	}
	if r.StatusCode != 200 {
		return updates, fmt.Errorf("non 200 response code returned %d, %s", r.StatusCode, err)
	}
	var response oupdate.OrganizationUpdatePage
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		return
	}
	return response.Embedded.OrganizationUpdates, nil
}
