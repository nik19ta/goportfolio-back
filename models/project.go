package models

// * Структура которая возвращяеться по api
type ApiProject struct {
	UUID         string `json:"uuid"`
	CategoryUUID string `json:"category_uuid"`
	Name         string `json:"name"`
	Prewiew      string `json:"prewiew"`
	State        int    `json:"state"`
	// * Если state = 0 - публичный для всех (Default)
	// * Если state = 1 - приватный для всех
	// * Если state = 2 - доступный только по api
}

type SetStateInp struct {
	ProjectUUID string `json:"project_uuid"`
	State       int    `json:"state"`
}

type DelProjectInp struct {
	UUID string `json:"uuid"`
}

type NewProjectInp struct {
	CategoryUUID string `json:"category_uuid"`
}

type RenameProjectInp struct {
	UUID  string `json:"uuid"`
	Title string `json:"title"`
}

type AddDescription struct {
	Test string `json:"text"`
	UUID string `json:"uuid"`
}

type InfoProjects struct {
	Main         Project       `json:"project"`
	Photos       []Photo       `json:"photos"`
	Descriptions []Description `json:"descriptions"`
}
