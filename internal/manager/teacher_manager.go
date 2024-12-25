package manager

import (
	"Quotium/external/untis"
	"Quotium/internal/server/db/entity"
	"Quotium/internal/server/db/repo"
	"Quotium/util"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func UpdateTeachersInDB() bool {
	logrus.Info("Update teachers in db")
	client := untis.NewUntisClient(os.Getenv("UNTIS_SCHOOL"), os.Getenv("UNTIS_USERNAME"), os.Getenv("UNTIS_PASSWORD"), os.Getenv("UNTIS_URL"))
	err := client.Login()
	if err != nil {
		logrus.Errorf("could not login to untis. Probably wrong credentials or school. %v", err)
		return false
	}
	logrus.Info("Logged in to untis")
	untisTeachers, err := client.GetTeachers()
	if err != nil {
		logrus.Errorf("could not get teachers. You probably dont have access to this resource or your credentials are wrong. %v", err)
		return false
	}
	logrus.Infof("Retrieved %d teachers from untis", len(untisTeachers))

	dbTeachers := repo.GetAllTeachers()
	if dbTeachers == nil {
		logrus.Errorf("could not get all teachers")
		return false
	}
	logrus.Infof("Retrieved %d teachers from db", len(dbTeachers))

	var teachersToAdd []entity.Teacher
	for _, untisTeacher := range untisTeachers {
		found := false
		for _, dbTeacher := range dbTeachers {
			if untisTeacher.Id == dbTeacher.ID {
				found = true
				break
			}
		}
		if !found {
			teachersToAdd = append(teachersToAdd, entity.Teacher{
				ID:        untisTeacher.Id,
				ShortName: untisTeacher.Name,
				Name:      untisTeacher.Forename + " " + untisTeacher.LongName,
				Title:     untisTeacher.Title,
			})
		}
	}
	logrus.Info("Start downloading teacher images from untis.")
	logrus.Infof("This operation will likely take %d seconds", len(teachersToAdd))
	err = os.MkdirAll("data/teacher", 664)
	if err != nil {
		logrus.Errorf("could not create directory for downloading teachers. %v", err)
	} else {
		for _, teacher := range teachersToAdd {
			downloadPath := fmt.Sprintf("data/teacher/%d.png", teacher.ID)
			if util.FileExists(downloadPath) {
				continue
			}
			if err = client.DownloadTeacherImage(downloadPath, teacher.ID); err != nil {
				logrus.Errorf("could not download teacher image. %v", err)
			} else {
				logrus.Debugf("Downloaded image for teacher %d", teacher.ID)
			}
			time.Sleep(1 * time.Second)
		}
	}
	logrus.Info("finished downloading teacher images from untis")
	if len(teachersToAdd) != 0 {
		logrus.Infof("Found %d new teachers and add them to db", len(teachersToAdd))
		repo.AddTeachers(teachersToAdd)
		logrus.Infof("Added teachers to db now have %d teachers", len(teachersToAdd)+len(dbTeachers))
	} else {
		logrus.Info("No new teachers found")
	}
	return true
}
