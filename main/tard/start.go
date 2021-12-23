package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"tard/mods"

	"dev.azure.com/carrotins/hdm/hdm-go.git/logging"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startCmd = func() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start server",
		Long:  "Start serer",
		Args:  startArgs,
		Run:   startRun,
	}

	return cmd
}()

var loggingConfig *logging.Config

func startArgs(cmd *cobra.Command, args []string) error {
	verbosePrint("verbose: %t", verbose)
	verbosePrint("configPath: %s", configPath)
	verbosePrint("pname: %s", viper.GetString("pname"))
	verbosePrint("server.bind: %s", viper.GetString("server.bind"))
	verbosePrint("server.http-prefix: %s", viper.GetString("server.http-prefix"))

	loggingConfig = &logging.Config{
		Name:               viper.GetString("logger.name"),
		Console:            viper.GetBool("logger.console"),
		Filename:           viper.GetString("logger.filename"),
		Append:             viper.GetBool("logger.append"),
		MaxSize:            viper.GetInt("logger.maxsize"),
		MaxBackups:         viper.GetInt("logger.maxbackups"),
		MaxAge:             viper.GetInt("logger.maxage"),
		Compress:           viper.GetBool("logger.compress"),
		DefaultPrefixWidth: viper.GetInt("logger.prefix-width"),
		DefaultLevel:       logging.ParseLogLevel(viper.GetString("log-level")),
	}

	for i := range viper.Get("logger.levels").([]interface{}) {
		prefix := fmt.Sprintf("logger.levels.%d", i)
		pattern := viper.GetString(prefix + ".pattern")
		levelName := viper.GetString(prefix + ".level")

		level := logging.ParseLogLevel(levelName)

		lcfg := logging.LevelConfig{Pattern: pattern, Level: level}
		loggingConfig.Levels = append(loggingConfig.Levels, lcfg)
	}

	// if viper.GetBool("metrics.enabled") {
	// 	metricConfig = &metrics.MetricsReporterConfig{
	// 		Url:      viper.GetString("metrics.url"),
	// 		Database: viper.GetString("metrics.database"),
	// 		Username: viper.GetString("metrics.username"),
	// 		Password: viper.GetString("metrics.password"),
	// 		Align:    viper.GetBool("metrics.align"),
	// 		Interval: viper.GetDuration("metrics.interval"),
	// 	}
	// }

	return nil
}

func startRun(cmd *cobra.Command, args []string) {
	// Set Logging
	logging.Configure(loggingConfig)

	log := logging.GetLog("tard-server")

	// Banner
	log.Info(BootBanner(viper.GetString("pname"), mods.VersionDescription()))

	// Wait Signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func handleVersion(c *gin.Context) {
	c.String(http.StatusOK, mods.VersionDescription())
}
