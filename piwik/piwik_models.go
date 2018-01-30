package piwik

//ActionDetail map the piwik's action details from a specific visit
type ActionDetail struct {
	URL string `json:"url"`
}

//Visit map the piwik's visits array entry
type Visit struct {
	ActionDetails []ActionDetail `json:"actionDetails"`
}
