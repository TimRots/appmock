package mock

import (
	"os"
	"os/user"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

type Mock struct {
	Duration   time.Duration `yaml:"duration"`
	Exitstatus int           `yaml:"exit-status"`
	Stderr     string        `yaml:"stderr"`
	Stdout     string        `yaml:"stdout"`
}

// MockApp takes an application name as string argument
// and returns a Mock or an error
func MockApp(app string) (*Mock, error) {
	// strip optional dotslash prefix from appname
	appname := strings.TrimPrefix(app, "./")

	// find current user to determine homdir
	user, err := user.Current()
	if err != nil {
		return nil, err
	}

	// read config from ~/.appmock/appname.yml
	appconfig := user.HomeDir + "/.appmock/" + appname + ".yml"
	config, err := os.Open(appconfig)
	if err != nil {
		return nil, err
	}
	defer config.Close()

	// marshal config file to struct
	mock := &Mock{}
	d := yaml.NewDecoder(config)
	if err := d.Decode(&mock); err != nil {
		return nil, err
	}

	// return Mock stuct
	return mock, nil
}
