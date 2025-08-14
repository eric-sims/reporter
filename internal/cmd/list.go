// Package cmd - lists entries for the day
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/eric-sims/reporter/internal/db"
	"github.com/jinzhu/now"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list reports for the day",
	RunE: func(_ *cobra.Command, _ []string) error {
		start := now.BeginningOfDay()
		end := now.EndOfDay()
		database, err := db.Open()
		if err != nil {
			return err
		}
		defer database.Close()

		entries, err := database.ListSummaries(start, end)
		if err != nil {
			return err
		}
		if len(entries) == 0 {
			fmt.Println("No entries in that range.")
			return nil
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.Header([]string{"date", "text"})
		var rows [][]string
		for _, entry := range entries {
			rows = append(rows, []string{entry.Date.Format(time.DateOnly), entry.Text})
		}
		bulkErr := table.Bulk(rows)
		if bulkErr != nil {
			return bulkErr
		}
		renderErr := table.Render()
		if renderErr != nil {
			return renderErr
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
