package config

import (
	cfgTD "github.com/ameniGa/timeTracker/testData/config"
	"os"
	"testing"
)

func Test_LoadConfig(t *testing.T) {
	for _, testCase := range cfgTD.CreateTTLoadConf() {
		t.Run(testCase.Name, func(t *testing.T) {
			os.Setenv(ConfEnvVar, testCase.EnvVar)
			conf, err := LoadConfig()
			t.Logf("%v",conf)
			if err != nil ||
				(testCase.IsProd && os.Getenv(ConfEnvVar) != "production" && conf.Tag == "dev") ||
				(testCase.IsTest && os.Getenv(ConfEnvVar) != "test" && conf.Tag == "prod") {
				t.Errorf("expected loading config for %v environment, got error: %v", testCase.EnvVar, err)
			}
		})
	}
}
