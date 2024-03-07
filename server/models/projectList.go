package models

type ProjectList struct {
	ProjectName   string `json:"projectName"`
	ProjectUrl    string `json:"projectUrl"`
	ProjectBranch string `json:"projectBranch"`
	Username      string `json:"username"`
	Password      string `json:"password"`
}
type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
type ProjectItem struct {
	Id int `json:"id"`
	*ProjectList
}
type ReportDetails struct {
	ProjectName   string   `json:"projectName"`
	ReportContent []string `json:"reportContent"`
}
