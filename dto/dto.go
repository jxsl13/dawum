package dto

type Data struct {
	Database    Database              `json:"Database"`
	Parliaments map[string]Parliament `json:"Parliaments"`
	Institutes  map[string]Institute  `json:"Institutes"`
	Taskers     map[string]Institute  `json:"Taskers"`
	Methods     map[string]Institute  `json:"Methods"`
	Parties     map[string]Party      `json:"Parties"`
	Surveys     map[string]Survey     `json:"Surveys"`
}
