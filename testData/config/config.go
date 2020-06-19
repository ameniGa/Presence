package config

// TTMakeConf represents table test structure of LoadConfig test
type TTLoadConf struct {
	Name   string
	EnvVar string
	IsProd bool
	IsTest bool
}

// CreateTTLoadConf creates table test for LoadConfig test
func CreateTTLoadConf() []TTLoadConf {
	tt := []TTLoadConf{
		{
			Name:   "test config",
			EnvVar: "test",
			IsTest: true,
		},
		{
			Name:   "default config",
			EnvVar: "production",
			IsProd: true,
		},
	}
	return tt
}
