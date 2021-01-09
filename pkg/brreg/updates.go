package brreg

import (
	"encoding/json"
	"fmt"
	oupdate "github.com/tjololo/nororg-streamer/pkg/brreg/structs"
	"net/http"
	"time"
)

//OrganizationUpdateService fetch organization updates
type OrganizationUpdateService struct {
	BaseURL string
}

//FetchUpdatesAfterDate fetch updates after date
func (u *OrganizationUpdateService) FetchUpdatesAfterDate(from time.Time) (updates []oupdate.OrganizationUpdate, err error) {
	getURL := fmt.Sprintf("%s/enhetsregisteret/api/oppdateringer/enheter?dato=%s", u.BaseURL, from.Format("2006-01-02T15:04:05.000Z"))
	r, err := http.Get(getURL)
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

//FetchUpdatesAfterID fetch updates after id
func (u *OrganizationUpdateService) FetchUpdatesAfterID(id int64) (updates []oupdate.OrganizationUpdate, err error) {
	getURL := fmt.Sprintf("%s/enhetsregisteret/api/oppdateringer/enheter?oppdateringsid=%d", u.BaseURL, id)
	r, err := http.Get(getURL)
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
