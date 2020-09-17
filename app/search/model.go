package search

type Results struct {
	Definition string `json:"definition"`
	PartOfSpeech string `json:"partOfSpeech"`
	Synonyms []string `json:"synonyms"`
	TypeOf []string `json:"typeOf"`
	HasTypes []string `json:"hasTypes"`
	Examples []string `json:"examples"`
}

type Pronunciation struct {
	All string `json:"all"`
}

type WholeResult struct {
	Word string `json:"word"`
	Results []Results `json:"results"`
	Pronunciation `json:"pronunciation"`
	Frequency float64 `json:"frequency"`
}
