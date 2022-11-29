package cmd

import (
	"github.com/ishtiaqhimel/go-api-server/api"
	"github.com/ishtiaqhimel/go-api-server/utils"
	"github.com/spf13/cobra"
	"log"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the api server",
	Long:  `Start the api server on port 3000 by default`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("start called! start the server from  point..")
		api.CallRoutes(utils.PORT)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
