package domain

import (
	"time"
)

type BaseModel struct {
	ID        int64      `json:"id"`
	UUID      string     `json:"uuid"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Province struct {
	BaseModel
	ParentID    int64        `json:"parent_id"`
	Code        string       `json:"code"`
	Name        string       `json:"name"`
	Nationality *Nationality `json:"nationality,omitempty" gorm:"foreignKey:ParentID"`
}

type Nationality struct {
	BaseModel
	Code        string `json:"code"`
	Country     string `json:"country"`
	ISO         string `json:"iso"`
	Nationality string `json:"nationality"`
}

type City struct {
	BaseModel
	Code     string    `json:"code"`
	ParentID int64     `json:"parent_id"`
	Name     string    `json:"name"`
	Type     string    `json:"type"`
	Group    string    `json:"group"`
	Province *Province `json:"province,omitempty" gorm:"foreignKey:ParentID"`
}

type District struct {
	BaseModel
	Code     string `json:"code"`
	ParentID int64  `json:"parent_id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Group    string `json:"group"`
	City     *City  `json:"city,omitempty" gorm:"foreignKey:ParentID"`
}

type Village struct {
	BaseModel
	Code     string    `json:"code"`
	ParentID int64     `json:"parent_id"`
	Name     string    `json:"name"`
	Type     string    `json:"type"`
	Group    string    `json:"group"`
	ZipCode  string    `json:"zip_code"`
	District *District `json:"district,omitempty" gorm:"foreignKey:ParentID"`
}

type Region struct {
	FullAddress  string `json:"full_address"`
	NationalID   int64  `json:"national_id"`
	NationalCode string `json:"national_code"`
	NationalName string `json:"national_name"`
	ProvinceID   int64  `json:"province_id"`
	ProvinceName string `json:"province_name"`
	CityID       int64  `json:"city_id"`
	CityName     string `json:"city_name"`
	DistrictID   int64  `json:"district_id"`
	DistrictName string `json:"district_name"`
	VillageID    int64  `json:"village_id"`
	VillageName  string `json:"village_name"`
	ZipCode      string `json:"zip_code"`
}
