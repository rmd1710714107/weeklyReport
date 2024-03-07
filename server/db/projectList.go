package db

import (
	"dailyReport/models"
	"log"
)

func CreateProjectApi(params *models.ProjectList) error {
	sqlStr := `insert into projectList (projectName,projectUrl,projectBranch,username,password) values (?,?,?,?,?)`
	_, err := sqlite3.Exec(sqlStr, params.ProjectName, params.ProjectUrl, params.ProjectBranch, params.Username, params.Password)
	if err != nil {
		return err
	}
	return nil
}
func GetProjectListApi() ([]models.ProjectItem, error) {
	sqlStr := "select * from projectList"
	row, err := sqlite3.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var projectList []models.ProjectItem
	for row.Next() {
		projectItem := models.ProjectItem{
			Id: 0,
			ProjectList: &models.ProjectList{
				ProjectName:   "",
				ProjectUrl:    "",
				ProjectBranch: "",
				Username:      "",
				Password:      "",
			},
		}
		err := row.Scan(&projectItem.Id, &projectItem.ProjectName, &projectItem.ProjectList.ProjectUrl, &projectItem.ProjectBranch, &projectItem.Username, &projectItem.Password)
		if err != nil {
			log.Fatalln(err)
			return nil, err
		}
		projectList = append(projectList, projectItem)
	}
	if projectList == nil {
		return []models.ProjectItem{}, nil
	}
	return projectList, err
}
func DeleteProjectItemApi(deleteId int) error {
	sqlStr := "delete from projectList where id = ?"
	_, err := sqlite3.Exec(sqlStr, deleteId)
	if err != nil {
		return err
	}
	return nil
}
