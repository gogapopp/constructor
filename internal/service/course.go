package service

import (
	"log"
	"time"
)

func (c *courseService) StartCourseAccessExpirationCheck() {
	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		for range ticker.C {
			err := c.courseStore.CheckAndExpireCourseAccess()
			log.Print(err)
		}
	}()
}
