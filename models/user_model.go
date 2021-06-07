package models

type User struct {
	Email      string `json:"email,omitempty"`
	Name       string `json:"name,omitempty"`
	TwitterId  string `json:"twitter_id,omitempty"`
	FacebookId string `json:"facebook_id,omitempty"`
	GithubId   string `json:"github_id,omitempty"`
}
