package types

// Output Output struct used for outputting right JSON object
type Output struct {
	Quotes  []string `json:"quotes"`
	Authors []string `json:"authors"`
}