// Package cmd - lists entries for the day [default: today}
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

var (
	listDate string
	listWeek bool
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List reports for a given day",
	Long:  "List reports for a given day (local time). Use --date YYYY-MM-DD to choose a day; defaults to today.",
	RunE: func(_ *cobra.Command, _ []string) error {
		var start, end time.Time
		if listDate == "" {
			// Today in local time
			start = now.BeginningOfDay()
			end = now.EndOfDay()
		} else {
			d, err := time.Parse(time.DateOnly, listDate)
			if err != nil {
				return fmt.Errorf("invalid --date %q: %w", listDate, err)
			}
			// Treat the parsed date in local time
			start = time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.Local)
			end = start.Add(24*time.Hour - time.Nanosecond)
		}

		if listWeek {
			start = now.BeginningOfWeek()
			end = now.EndOfWeek()
		}

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
	listCmd.Flags().StringVarP(&listDate, "date", "d", "", "Date to list (YYYY-MM-DD); defaults to today")
	listCmd.Flags().BoolVarP(&listWeek, "week", "w", false, "List all week reports")
}
