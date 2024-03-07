package service

import (
	"dailyReport/db"
	"dailyReport/models"
	"dailyReport/utils"
	"github.com/go-git/go-git/v5/plumbing/object"
	"log"
	"sync"
)

func CreateProject(params *models.ProjectList) error {
	return db.CreateProjectApi(params)
}
func GetProjectList() ([]models.ProjectItem, error) {
	return db.GetProjectListApi()
}
func DeleteProjectItem(deleteId int) error {
	return db.DeleteProjectItemApi(deleteId)
}
func GetReportDetails() ([]models.ReportDetails, error) {
	projectList, err := db.GetProjectListApi()
	if err != nil {
		return nil, err
	}

	projectReportChan := make(chan models.ReportDetails, len(projectList)) // 带缓冲的 channel

	var wg sync.WaitGroup
	for _, projectItem := range projectList {
		wg.Add(1)
		go func(item models.ProjectItem) {
			defer wg.Done()

			var content []string
			cIter, err := utils.GetHistory(item)
			if err != nil {
				log.Println(err) // Log the error instead of terminating
			} else {
				cIterErr := cIter.ForEach(func(c *object.Commit) error {
					msg := c.Message
					content = append(content, msg)
					return nil
				})
				if cIterErr != nil {
					log.Println(cIterErr) // Log the error instead of terminating
				}
			}

			report := models.ReportDetails{
				ProjectName:   item.ProjectName,
				ReportContent: content,
			}
			projectReportChan <- report // 发送报告详情到 channel
		}(projectItem)
	}

	go func() {
		wg.Wait()                // 等待所有 goroutine 完成
		close(projectReportChan) // 关闭 channel
	}()

	var reportDetailsList []models.ReportDetails
	for report := range projectReportChan {
		reportDetailsList = append(reportDetailsList, report)
	}

	return reportDetailsList, nil
}
