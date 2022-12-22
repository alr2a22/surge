package cmd

import (
	"net/http"
	"surge/api/account"
	"surge/api/ride"
	"surge/internal/config"
	"surge/internal/db"
	"surge/internal/logger"
	"surge/internal/middlewares"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	runServerCmd = &cobra.Command{
		Use:   "run-server",
		Short: "Start http server",
		Long:  `Start http server for surge`,
		Run: func(cmd *cobra.Command, args []string) {
			runServer(cmd)
		},
	}
)

func init() {
	rootCmd.AddCommand(runServerCmd)
}

func runServer(cmd *cobra.Command) {
	config.GetConfig()
	logger.Setup()
	db.GetDBConn()
	db.GetRedisBackend()

	r := mux.NewRouter()
	r.Use(middlewares.DefaultHeaders)
	account.AddRoutes(r)
	ride.AddRoutes(r)

	logrus.Infoln("Server start at: http://127.0.0.1:3000")
	logrus.Fatal(http.ListenAndServe(":3000",
		handlers.CORS(
			handlers.AllowedHeaders([]string{
				"X-Requested-With",
				"Content-Type",
				"Authorization"}),
			handlers.AllowedMethods([]string{
				"GET",
				"POST",
				"PUT",
				"DELETE",
				"HEAD",
				"OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(r)))
}
