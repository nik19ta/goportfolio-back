package models

type SignInp struct {
	Shortname string `json:"shortname"`
	Mail      string `json:"mail"`
	Password  string `json:"password"`
	Fullname  string `json:"fullname"`
}

type OAuthAccessResponse struct {
	AccessToken string `json:"access_token"`
}

type OAuthEmail struct {
	Email      string `json:"email"`
	Primary    bool   `json:"primary"`
	Verified   bool   `json:"verified"`
	Visibility string `json:"visibility"`
}

type OAuthUser struct {
	AvatarUrl           string `json:"avatar_url"`
	Bio                 string `json:"bio"`
	Blog                string `json:"blog"`
	Company             string `json:"company"`
	Created_at          string `json:"created_at"`
	Email               string `json:"email"`
	Events_url          string `json:"events_url"`
	Followers           int    `json:"followers"`
	Followers_url       string `json:"followers_url"`
	Following           int    `json:"following"`
	Following_url       string `json:"following_url"`
	Gists_url           string `json:"gists_url"`
	Gravatar_id         string `json:"gravatar_id"`
	Hireable            bool   `json:"hireable"`
	Html_url            string `json:"html_url"`
	Id                  int64  `json:"id"`
	Location            string `json:"location"`
	Login               string `json:"login"`
	Name                string `json:"name"`
	Node_id             string `json:"node_id"`
	Organizations_url   string `json:"organizations_url"`
	Public_gists        int    `json:"public_gists"`
	Public_repos        int    `json:"public_repos"`
	Received_events_url string `json:"received_events_url"`
	Repos_url           string `json:"repos_url"`
	Site_admin          bool   `json:"site_admin"`
	Starred_url         string `json:"starred_url"`
	Subscriptions_url   string `json:"subscriptions_url"`
	Twitter_username    string `json:"twitter_username"`
	Type                string `json:"type"`
	Ipdated_at          string `json:"updated_at"`
	Url                 string `json:"url"`
}
