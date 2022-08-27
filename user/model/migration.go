package model

func migration() error {
	return DB.AutoMigrate(&User{})
}
