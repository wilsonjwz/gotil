package xhttp

const (
	POST           = "POST"
	GET            = "GET"
	TypeJSON       = "json"
	TypeXML        = "xml"
	TypeUrlencoded = "urlencoded"
	TypeForm       = "form"
	TypeFormData   = "form-data"
)

var types = map[string]string{
	TypeJSON:       "application/json",
	TypeXML:        "application/xml",
	TypeForm:       "application/x-www-form-urlencoded",
	TypeFormData:   "application/x-www-form-urlencoded",
	TypeUrlencoded: "application/x-www-form-urlencoded",
}
