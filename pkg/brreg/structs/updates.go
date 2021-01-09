package structs

//OrganizationUpdatePage root object in update response
type OrganizationUpdatePage struct {
	Embedded EmbeddedUpdates `json:"_embedded"`
	Links    Links           `json:"_links"`
	Page     Page            `json:"page"`
}

//EmbeddedUpdates _embedded
type EmbeddedUpdates struct {
	OrganizationUpdates []OrganizationUpdate `json:"oppdaterteEnheter"`
}

//OrganizationUpdate oppdaterteEnheter
type OrganizationUpdate struct {
	UpdateID           int        `json:"oppdateringsid"`
	Date               string     `json:"dato"`
	Organizationnumber string     `json:"organisasjonsnummer"`
	ChangeType         ChangeType `json:"endringstype"`
	Links              Links      `json:"_links"`
}

//ChangeType endringstype
type ChangeType string

const (
	//Unknown ukjent
	Unknown ChangeType = "Ukjent"
	//New ny
	New = "Ny"
	//Change endring
	Change = "Endring"
	//Deletion sletting
	Deletion = "Sletting"
	//Removed fjernet
	Removed = "Fjernet"
)

//Links _links
type Links struct {
	First        Href `json:"first,omitempty"`
	Self         Href `json:"self,omitempty"`
	Next         Href `json:"next,omitempty"`
	Last         Href `json:"last,omitempty"`
	Organization Href `json:"enhet,omitempty"`
}

//Href links href
type Href struct {
	Href string `json:"href"`
}

//Page page
type Page struct {
	Size          int `json:"size"`
	TotalElements int `json:"totalElements"`
	TotalPages    int `json:"totalPages"`
	Number        int `json:"number"`
}
