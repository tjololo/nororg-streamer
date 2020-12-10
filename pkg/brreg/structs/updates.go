package structs

type OrganizationUpdatePage struct {
	Embedded EmbeddedUpdates `json:"_embedded"`
	Links    Links           `json:"_links"`
	Page     Page            `json:"page"`
}

type EmbeddedUpdates struct {
	OrganizationUpdates []OrganizationUpdate `json:"oppdaterteEnheter"`
}

type OrganizationUpdate struct {
	UpdateID           int        `json:"oppdateringsid"`
	Date               string     `json:"dato"`
	Organizationnumber string     `json:"organisasjonsnummer"`
	ChangeType         ChangeType `json:"endringstype"`
	Links              Links      `json:"_links"`
}

type ChangeType string

const (
	Unknown  ChangeType = "Ukjent"
	New                 = "Ny"
	Change              = "Endring"
	Deletion            = "Sletting"
	Removed             = "Fjernet"
)

type Links struct {
	First        Href `json:"first,omitempty"`
	Self         Href `json:"self,omitempty"`
	Next         Href `json:"next,omitempty"`
	Last         Href `json:"last,omitempty"`
	Organization Href `json:"enhet,omitempty"`
}

type Href struct {
	Href string `json:"href"`
}

type Page struct {
	Size          int `json:"size"`
	TotalElements int `json:"totalElements"`
	TotalPages    int `json:"totalPages"`
	Number        int `json:"number"`
}
