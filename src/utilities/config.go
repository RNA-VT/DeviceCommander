package utilities

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func ConfigureEnvironment() {
	// viper.SetConfigName("config")
	// viper.SetConfigType("yaml")
	// viper.AddConfigPath(".")

	// err := viper.ReadInConfig()
	// if err != nil {
	// 	panic(fmt.Errorf("fatal error config file: %s ", err))
	// }

	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	viper.AutomaticEnv()

	viper.SetDefault("ENV", "local") // local or production only
	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", 8001)
	viper.SetDefault("CLUSTER_NAME", "Flaming Hot Fleet Directory")
	viper.SetDefault("IP_ADDRESS_ROOT", "192.16.1.")
	viper.SetDefault("ARP_SCAN_DURATION", 30)
	viper.SetDefault("DISCOVERY_PERIOD", 30)
	viper.SetDefault("HEALTH_CHECK_PERIOD", 60)

	viper.SetDefault("POSTGRES_NAME", "postgres")
	viper.SetDefault("POSTGRES_HOST", "0.0.0.0")
	viper.SetDefault("POSTGRES_PORT", 5432)
	viper.SetDefault("POSTGRES_USER", "postgres")
	viper.SetDefault("POSTGRES_PASSWORD", "changeme")
}
