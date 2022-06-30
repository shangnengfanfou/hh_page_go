package repositories

import (
	"hh_page_go/src/models"
	"hh_page_go/src/utils"
)

type VisitsRecordRepo struct{}

func NewVisitsRecordRepo() *VisitsRecordRepo {
	return &VisitsRecordRepo{}
}

func (h *VisitsRecordRepo) InsertOne(record models.VisitsRecord) error {
	result := utils.DB.Create(&record)
	return result.Error
}

type VisitsRecordStats struct {
	Count     int `json:"count"`
	TotalTime int `json:"totalTime" gorm:"column:total_time"`
}

func (h *VisitsRecordRepo) CountAndTotalTime(sTime int, eTime int) (VisitsRecordStats, error) {
	var stats VisitsRecordStats
	result := utils.DB.Table("visits_record").Select("count(*) as count, sum(closed_time - access_time) as total_time").Where("access_time > ? AND access_time < ? AND closed_time is not null", sTime, eTime).First(&stats)
	return stats, result.Error
}
