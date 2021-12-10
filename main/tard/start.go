package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"tard/mods"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
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

func startArgs(cmd *cobra.Command, args []string) error {

	return nil
}

func startRun(cmd *cobra.Command, args []string) {
	// Wait Signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func handleVersion(c *gin.Context) {
	c.String(http.StatusOK, mods.VersionDescription())
}
