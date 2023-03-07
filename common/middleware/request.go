package middleware

import "time"

type Userinfo struct {
	Data struct {
		SignedUp   time.Time    `json:"signedUp"`
		Identities []Identities `json:"identities"`
		Phone      string       `json:"phone"`
		Nickname   string       `json:"nickname"`
		Photo      string       `json:"photo"`
		Company    string       `json:"company"`
		Email      string       `json:"email"`
		Username   string       `json:"username"`
	} `json:"data"`
}

type Identities struct {
	LoginName   string `json:"login_name"`
	UserIdInIdp string `json:"userIdInIdp"`
	Identity    string `json:"identity"`
	UserName    string `json:"user_name"`
	AccessToken string `json:"accessToken"`
}
