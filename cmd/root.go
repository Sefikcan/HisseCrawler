// Package cmd /*
package cmd

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/sefikcan/hisse-crawler/internal/asset"
	"github.com/sefikcan/hisse-crawler/pkg/parser"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "assets",
	Short: "This cli command help to get asset information by symbol",
	Long: `
Assets: 
1. thyo
2. garan
3. kchol
4. sahol

Example usage
----------------------
"thyo"
"thyo -s 1"
"garan -s 2"
"kchol -s 3"
"sahol -s 4"

For now, the above models are supported in searching on getmidas.com.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		symbol := asset.Types.Thyo
		if t, err := cmd.Flags().GetInt("type"); err == nil && t >= 1 && t <= 4 {
			switch t {
			case 1:
				symbol = asset.Types.Thyo
			case 2:
				symbol = asset.Types.Garan
			case 3:
				symbol = asset.Types.Kchol
			case 4:
				symbol = asset.Types.Sahol
			}
		}

		getAssetDetail(symbol)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hisse-crawler.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().IntP("type", "t", 1, "Asset type (1: thyo, 2: garan, 3: kchol, 4: sahol)")
}

func getAssetDetail(symbol string) {
	result, err := parser.ParseMidas(symbol)
	if err != nil {
		fmt.Println("Error fetching asset details:", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Price", "Daily Volume", "Daily Change"})
	table.SetCaption(true, "Asset information from Midas.com")
	table.Append([]string{
		result.Name.Slice(),
		result.Price.Slice(),
		result.DailyVolume.Slice(),
		result.DailyChange.Slice(),
	})

	table.SetHeaderColor(
		tablewriter.Color(tablewriter.Bold, tablewriter.BgMagentaColor),
		tablewriter.Color(tablewriter.Bold, tablewriter.BgGreenColor),
		tablewriter.Color(tablewriter.Bold, tablewriter.BgYellowColor),
		tablewriter.Color(tablewriter.Bold, tablewriter.BgBlueColor),
	)

	table.Render()
}
