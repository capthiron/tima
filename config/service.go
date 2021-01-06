package config

import "log"

type Service interface {
	SaveConfig(config *Config) *Config
	GetConfig(id uint) *Config
}

func NewService() Service {
	return defaultService{newRepository()}
}

type defaultService struct {
	repo repository
}

func (s defaultService) SaveConfig(config *Config) *Config {
	config, err := s.repo.saveConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

func (s defaultService) GetConfig(id uint) *Config {
	config, err := s.repo.getConfig(id)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
