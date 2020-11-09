package telraam

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Telraam"
)

type Settings struct {
	common *cfg.Common

	numberOfCameras int           `help:"Defines number of cameras to be displayed. Default is 5" optional:"true"`
	cameraIds       []interface{} `help:"The integer IDs of the Telraam mac ids you wish to report on."`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		numberOfCameras: ymlConfig.UInt("numberOfCameras", 5),
		cameraIds:       ymlConfig.UList("cameraIds"),
	}

	return &settings
}
