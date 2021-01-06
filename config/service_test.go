package config

import (
	"testing"
)

type mockRepo struct{}

var saveConfigMock func(config *Config) (*Config, error)
var getConfigMock func(id uint) (*Config, error)

func (m mockRepo) saveConfig(config *Config) (*Config, error) {
	return saveConfigMock(config)
}

func (m mockRepo) getConfig(id uint) (*Config, error) {
	return getConfigMock(id)
}

var service Service
var mockedRepo mockRepo

func beforeServiceTest() {
	mockedRepo = mockRepo{}
	service = defaultService{mockedRepo}
}

func TestDefaultService_SaveConfig(t *testing.T) {
	beforeServiceTest()

	configToSave := Config{
		DateFormat: "TestDateFormat",
	}

	if configToSave.ID != 0 {
		t.Errorf("configToSave.ID should be 0 but got %v", configToSave.ID)
	}

	configSaved := configToSave
	configSaved.ID = 1

	saveConfigCalled := 0
	saveConfigMock = func(config *Config) (*Config, error) {
		saveConfigCalled++
		return &configSaved, nil
	}

	if saveConfigCalled != 0 {
		t.Errorf("saveConfig() should not have been called yet but instead got %v", saveConfigCalled)
	}

	configToSave = *service.SaveConfig(&configToSave)

	if configToSave.ID != 1 {
		t.Errorf("id of configToSave should now be 1 but got %v", configToSave.ID)
	}

	if saveConfigCalled != 1 {
		t.Errorf("saveConfig() should have been called but instead got %v", saveConfigCalled)
	}
}

func TestDefaultService_GetConfig(t *testing.T) {
	beforeServiceTest()

	getConfigCalled := 0
	getConfigMock = func(id uint) (*Config, error) {
		getConfigCalled++
		return &Config{}, nil
	}

	if getConfigCalled != 0 {
		t.Errorf("getConfig() should not have been called yet but instead got %v", getConfigCalled)
	}

	service.GetConfig(1)

	if getConfigCalled != 1 {
		t.Errorf("getConfig() should have been called but instead got %v", getConfigCalled)
	}
}
