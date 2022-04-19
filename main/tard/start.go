package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"tard/mods"
	"tard/mods/tlog"

	"dev.azure.com/carrotins/hdm/hdm-go.git/httpsvr"
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

	return nil
}

func startRun(cmd *cobra.Command, args []string) {
	// Set Logging
	logging.Configure(loggingConfig)

	log := logging.GetLog("tard-server")

	// Banner
	log.Info(BootBanner(viper.GetString("pname"), mods.VersionDescription()))

	// GOMAXPROCS
	if viper.IsSet("GOMAXPROCS") {
		n := viper.GetInt("GOMAXPROCS")
		runtime.GOMAXPROCS(n)
	}
	log.Infof("GOMAXPROCS: %d", runtime.GOMAXPROCS(0))

	// Service for tlogs
	service := httpsvr.NewServer(&httpsvr.HttpServerConfig{
		DisableConsoleColor: false,
		DebugMode:           false,
		LoggingConfig: &logging.Config{
			Name:     "http-tard",
			Console:  false,
			Filename: ".", // disable http access logging
		},
		LoggerName: "tard-http-server",
	})
	httpPrefix := viper.GetString("server.http-prefix")
	serverBind := viper.GetString("server.bind")

	service.GET(httpPrefix+"/version", handleVersion)
	service.GET(httpPrefix+"/trips/:tripId/logs", tlog.HandleGetTLog())
	service.POST(httpPrefix+"/trip/log", tlog.HandlePostTLog())
	lsnr, err := net.Listen("tcp", serverBind)
	if err != nil {
		log.Errorf("listener failed, %s", err)
	}

	service.Start(lsnr)

	// Wait Signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func handleVersion(c *gin.Context) {
	c.String(http.StatusOK, mods.VersionDescription())
}
