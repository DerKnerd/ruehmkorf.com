package main

import (
	"embed"
	"fmt"
	"log"
	"os"

	"ruehmkorf/cmd"
	"ruehmkorf/database"

	"github.com/spf13/cobra"

	_ "github.com/joho/godotenv/autoload"
)

var (
	//go:embed static
	static embed.FS
)

var rootCmd = &cobra.Command{
	Use:   "ruehmkorf",
	Short: "Command line tool for ruehmkorf.com",
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Run: func(_ *cobra.Command, args []string) {
		cmd.WebUi(static)
	},
}

var createUserCmd = &cobra.Command{
	Use:   "create-user",
	Short: "Create a new user",
	Run: func(_ *cobra.Command, args []string) {
		cmd.CreateUser(args[0])
	},
}

func init() {
	createUserCmd.Args = cobra.ExactArgs(1)

	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(createUserCmd)
}

func main() {
	log.Println("Preparing the database")
	database.SetupDatabase()

	defer database.GetDbMap().Db.Close()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
