package conf

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config the application's configuration
type Config struct {
	Port      int64
	Host      string
	Env       string
	Config    string
	LogConfig LoggingConfig
}

// LoadConfig loads the config from a file if specified, otherwise from the environment
func LoadConfig(cmd *cobra.Command) (*Config, error) {
	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		return nil, err
	}

	viper.SetEnvPrefix("MYSERVICE")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if configFile, _ := cmd.Flags().GetString("config"); configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.AddConfigPath("./")
		viper.SetConfigType("env")
		viper.SetConfigName(".env")

	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	// Watch for changes to the configuration file and recompile
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})

	config := new(Config)
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
