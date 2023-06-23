package models

type Tracking struct {
	TrackNumber        string `json:"trackNumber" gorm:"primaryKey;auto_increment;not_null"`
	Sensor             string `json:"sensor" binding:"required" `
	Category           string `json:"category" binding:"required" `
	Time               string `json:"time" gorm:"not null" binding:"required"`
	Latitude           string `json:"latitude" `
	Longitude          string `json:"longitude"`
	Course             string `json:"course"`
	Speed              string `json:"speed"`
	Heading            string `json:"heading"`
	Altitude           string `json:"altitude" `
	Pitch              string `json:"pitch"`
	Roll               string `json:"roll"`
	Yaw                string `json:"yaw"`
	AccelarationX      string `json:"accelerationX"`
	AccelarationY      string `json:"accelerationY" `
	AccelarationZ      string `json:"accelerationZ"`
	VelocityX          string `json:"velocityX"`
	VelocityY          string `json:"velocityY"`
	VelocityZ          string `json:"velocityZ"`
	Humidity           string `json:"humidity" `
	WindSpeed          string `json:"windSpeed"`
	WindDirection      string `json:"windDirection"`
	AirTemperature     string `json:"airTemperature"`
	BarometricPressure string `json:"barometricPressure"`
}

// type Tracking struct {
// 	ID        uint64    `json:"trackNumber" gorm:"primaryKey;auto_increment;not_null"`
// 	Username  string    `json:"sensor" binding:"required" `
// 	Email     string    `json:"category" binding:"required" `
// 	Fullname  string    `json:"time" gorm:"not null" binding:"required"`
// 	CreatedBy uint64    `json:"latitude" `
// 	CreatedAt time.Time `json:"longitude"`
// 	UpdatedBy uint64    `json:"course"`
// 	UpdatedAt time.Time `json:"speed"`
// 	IsDeleted bool      `json:"heading"`
// 	CreatedBy uint64    `json:"altitudetime
// 	CreatedAt time.Time `json:"pitch"`
// 	UpdatedBy uint64    `json:"roll"`
// 	UpdatedAt time.Time `json:"yaw"`
// 	IsDeleted bool      `json:"accelerationX"`
// 	CreatedBy uint64    `json:"accelerationY" `
// 	CreatedAt time.Time `json:"accelerationZ"`
// 	UpdatedBy uint64    `json:"velocityX"`
// 	UpdatedAt time.Time `json:"velocityY"`
// 	IsDeleted bool      `json:"velocityZ"`
// 	CreatedBy uint64    `json:"humidity" `
// 	CreatedAt time.Time `json:"windSpeed"`
// 	UpdatedBy uint64    `json:"windDirection"`
// 	UpdatedAt time.Time `json:"airTemperature"`
// 	IsDeleted bool      `json:"barometricPressure"`
// }
