package config

import (
	"errors"
	"fmt"
	"github.com/capthiron/tima/db"
	"gorm.io/gorm"
)

type repository interface {
	saveConfig(config *Config) (*Config, error)
	getConfig(id uint) (*Config, error)
}

func newRepository() repository {
	dbConn := db.Connect()

	err := dbConn.AutoMigrate(&Config{})
	if err != nil {
		fmt.Println(err)
	}

	return &defaultRepository{dbConn}
}

type defaultRepository struct {
	db *gorm.DB
}

func (r defaultRepository) saveConfig(config *Config) (*Config, error) {
	var configToSave Config
	err := r.db.First(&configToSave, config.ID).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		configToSave = *config
		err = r.db.Create(&configToSave).Error
		if err != nil {
			return &configToSave, err
		}
	}

	err = r.db.Save(&configToSave).Error
	if err != nil {
		return &configToSave, err
	}

	return &configToSave, nil
}

func (r defaultRepository) getConfig(id uint) (*Config, error) {
	var config Config
	err := r.db.First(&config, id).Error
	if err != nil {
		err = fmt.Errorf("getConfig(%v): %w", id, err)
	}
	return &config, err
}
