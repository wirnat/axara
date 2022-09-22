package model

import "time"

type Branch struct {
	BaseModel
	CompanyID      int64      `json:"company_id"`
	Name           string     `json:"name"` //@meta validate_store:true
	Description    *string    `json:"description"`
	Email          string     `json:"email"` //@meta validate_store:true
	Phone          string     `json:"phone"` //@meta validate_store:true
	PicName        string     `json:"pic_name"`
	PicPhone       string     `json:"pic_phone"`
	PicEmail       string     `json:"pic_email"`
	Address        string     `json:"address"`
	Status         string     `json:"status"`
	VerifiedStatus string     `json:"verified_status"`
	OpenStatus     string     `json:"open_status"`
	ProfileImageID *int64     `json:"profile_image_id"` //MEDIA
	OpenedAt       *time.Time `json:"opened_at"`
	ClosedAt       *time.Time `json:"closed_at"`
	Latitude       float64    `json:"latitude"`
	Longitude      float64    `json:"longitude"`
}

/*
	@model Branch
	@service branch
*/
