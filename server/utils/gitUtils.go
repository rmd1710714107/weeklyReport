package utils

import (
	"dailyReport/models"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
	"log"
	"os"
	"time"
)

func GetHistory(projectInfo models.ProjectItem) (object.CommitIter, error) {
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: projectInfo.ProjectUrl,
		Auth: &http.BasicAuth{
			Username: projectInfo.Username, // anything except an empty string
			Password: projectInfo.Password,
		},
		ReferenceName: plumbing.NewBranchReferenceName(projectInfo.ProjectBranch),
		SingleBranch:  true,
		Progress:      os.Stdout,
	})
	if err != nil {

		return nil, err
	}
	ref, err1 := r.Head()
	if err1 != nil {
		log.Fatalf("ref error,error is %s", err1.Error())
		return nil, err1
	}

	cIter, err2 := r.Log(&git.LogOptions{
		From:  ref.Hash(),
		Since: getFirstDateOfWeek(),
	})
	if err2 != nil {
		log.Fatalf("log error, error is %s", err2.Error())
		return nil, err2
	}
	return cIter, nil
	//index := 0
	// iterates over the commits and print each
	/*err3 := cIter.ForEach(func(c *object.Commit) error {
		if index >= 10 {
			return nil
		}
		//msg := c.Message
		if err != nil {
			return err
		}
		index++
		fmt.Println(c)
		return nil
	})
	if err3 != nil {
		log.Fatalf("foreach error, error is %s", err3.Error())
	}*/
}

/*
*
获取本周周一的日期
*/
func getFirstDateOfWeek() *time.Time {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	until := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	return &until
	//fmt.Println(weekStartDate.UnixMilli())
	//weekMonday = weekStartDate.Format("2006-01-02 00:00:00")
	//return
}
