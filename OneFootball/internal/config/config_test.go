package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/NajeebTyson/onefootball/pkg/utils"
)

func TestLoadConfig(t *testing.T) {
	testConfigFile := "test_config.json"
	config := Config{
		TeamApi: "test api",
		Teams:   []string{"one", "two", "three"},
	}

	data, err := json.Marshal(config)
	if err != nil {
		t.Fatalf("unable to marshal json data, err: %v", err)
	}
	err = ioutil.WriteFile(testConfigFile, data, 0644)
	if err != nil {
		t.Fatal("unable to write json to data, err: ", err)
	}
	defer func() {
		os.Remove(testConfigFile)
	}()

	loadConfig, err := LoadConfig(testConfigFile)
	if err != nil {
		t.Fatal("unable to load config, err: ", err)
	}

	if loadConfig.TeamApi != config.TeamApi {
		t.Fatalf("expected team api: %s, got %s", config.TeamApi, loadConfig.TeamApi)
	}
	if !utils.IsSliceEqual(loadConfig.Teams, config.Teams) {
		t.Fatalf("expected teams: %v, got teams: %v", config.Teams, loadConfig.Teams)
	}
}
