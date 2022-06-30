package models

type VisitsRecord struct {
	Id         int `json:"id" gorm:"column:id;primary_key"`
	AccessTime int `json:"accessTime" gorm:"column:access_time"`
	ClosedTime int `json:"closedTime" gorm:"column:closed_time"`
}

func (VisitsRecord) TableName() string {
	return "visits_record"
}
