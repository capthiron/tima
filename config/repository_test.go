package config

import (
	"fmt"
	"os"
	"testing"
)

var repo repository

func beforeRepoTest() {
	err := os.Setenv("PROFILE", "test")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	repo = newRepository()
}

func TestDefaultRepository_SaveNote(t *testing.T) {
	beforeRepoTest()

	savedConfig, err := repo.saveConfig(&Config{})
	if err != nil {
		t.Error("Failed to save the config with error:", err)
	}

	if savedConfig == nil {
		t.Error("Saved config should not be nil")
	}

	_, err = repo.getConfig(savedConfig.ID)
	if err != nil {
		t.Error("Failed to get added note with error:", err)
	}
}

func TestDefaultRepository_GetNote(t *testing.T) {
	beforeRepoTest()

	dayLength := 8.5
	breakDuration := 50

	config, _ := repo.saveConfig(&Config{
		DayLength:            dayLength,
		DefaultBreakDuration: breakDuration,
	})

	receivedConfig, err := repo.getConfig(config.ID)
	if err != nil {
		t.Error("Repository returned an error:", err)
	}

	if receivedConfig.DayLength != dayLength {
		t.Errorf("Expected config.DayLength to be %v but got %v", dayLength, config.DayLength)
	}

	if receivedConfig.DefaultBreakDuration != breakDuration {
		t.Errorf("Expected config.DefaultBreakDuration to be %v but got %v", breakDuration, config.DefaultBreakDuration)
	}
}
