package models

type User struct {
	UUID      string `json:"id"`
	Shortname string `json:"shortname"`
	Mail      string `json:"mail"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
	Fullname  string `json:"fullname"`
	Type      string `json:"type"`
	About     string `json:"about"`
	// * Если type = user -  вход через сайт
	// * Если type = github - вход через github
	// * Если type = google - вход через google
}

type Project struct {
	UUID         string `json:"uui"`
	UserUUID     string `json:"user_uuid"`
	CategoryUUID string `json:"category_uuid"`
	Name         string `json:"name"`
	Prewiew      string `json:"prewiew"`
	State        int    `jsin:"state"`
	// * Если state = 0 - публичный (Default)
	// * Если state = 1 - приватный для всех
	// * Если state = 2 - доступный только по api
	// * Если state = 3 - доступный по api и публично
}

type Photo struct {
	UUID        string `json:"photo_uuid"`
	ProjectUUID string `json:"project_uuid"`
	Src         string `json:"src"`
	Type        int    `json:"type"`
	// * Если type = 0: Десктопное фото
	// * Если type = 1: Мобильное фото ( делать в три колонки )
}

type Tags struct {
	UUID        int    `json:"uuid"`
	ProjectUUID int    `json:"project_uuid"`
	Name        string `json:"name"`
}

type Description struct {
	UUID        int    `json:"uuid"`
	ProjectUUID int    `json:"project_uuid"`
	Key         string `json:"key"`
	Value       string `json:"value"`
	Language    string `json:"language"`
}

type Category struct {
	UUID     string `json:"uuid"`
	UserUUID string `json:"user_uuid"`
	Name     string `json:"name"`
}
