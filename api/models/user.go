package models

import (
    "github.com/google/uuid"
)


type User struct {
    UUID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"uuid"`
    Username    string `gorm:"uniqueIndex" json:"username"`
    Company     string `json:"company"`
    PhoneNumber string `json:"phoneNumber"`
}