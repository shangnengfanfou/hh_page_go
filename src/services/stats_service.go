package services

import (
	"hh_page_go/src/repositories"
)

type StatsService struct {
	VisitsRecordRepo *repositories.VisitsRecordRepo
}

func NewStatsService() *StatsService {
	return &StatsService{
		VisitsRecordRepo: repositories.NewVisitsRecordRepo(),
	}
}

func (h *StatsService) CountAndSumByTime(sTime int, eTime int) (repositories.VisitsRecordStats, error) {
	data, err := h.VisitsRecordRepo.CountAndTotalTime(sTime, eTime)
	if err != nil {
		return repositories.VisitsRecordStats{}, err
	}
	return data, nil
}
