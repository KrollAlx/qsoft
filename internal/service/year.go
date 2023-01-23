package service

import (
	"fmt"
	"time"
)

type Year interface {
	WhereYear(year int) string
}

type YearService struct {
}

func NewYearService() *YearService {
	return &YearService{}
}

func (s *YearService) WhereYear(year int) string {
	target := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	if year < time.Now().Year() {
		daysGone := int64(time.Since(target).Hours()) / 24
		return fmt.Sprintf("Days gone: %d", daysGone)
	} else {
		daysLeft := int64(time.Until(target).Hours()) / 24
		return fmt.Sprintf("Days left: %d", daysLeft)
	}
}
