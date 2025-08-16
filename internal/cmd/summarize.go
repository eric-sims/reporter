package cmd

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/eric-sims/reporter/internal/db"
	"github.com/eric-sims/reporter/internal/ollama"
	"github.com/eric-sims/reporter/internal/openai"
	"github.com/eric-sims/reporter/internal/summarize"
	"github.com/eric-sims/reporter/internal/util"
	"github.com/spf13/cobra"
)

var (
	weekISO  string
	fromDate string
	toDate   string
)

var weekCmd = &cobra.Command{
	Use:   "summarize",
	Short: "Summarize a week's entries with Ollama",
	Long:  "Collect summaries in a week range and generate an LLM-powered recap using a local Ollama server. Default's to current week.",
	RunE: func(cmd *cobra.Command, _ []string) error {
		var start, end time.Time
		var err error
		switch {
		case weekISO != "":
			start, end, _ = util.ISOWeekBounds(weekISO)
		case fromDate != "" || toDate != "":
			if fromDate == "" || toDate == "" {
				return errors.New("both --from and --to are required together")
			}
			start, err = util.ParseDate(fromDate)
			if err != nil {
				return err
			}
			end, err = util.ParseDate(toDate)
			if err != nil {
				return err
			}
			// inclusive range: end of day
			end = end.Add(23*time.Hour + 59*time.Minute + 59*time.Second)
		default:
			start, end = util.ThisWeek()
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

		prompt := summarize.WeeklyPrompt(entries, start, end)

		var resp string
		var errAPI error
		if useOpenAI {
			client := openai.NewClient(openAIKey)
			resp, errAPI = client.Generate(cmd.Context(), prompt)
		} else {
			client := ollama.NewClient(ollamaHost)
			resp, errAPI = client.Generate(cmd.Context(), model, prompt)
		}
		if errAPI != nil {
			return errAPI
		}
		fmt.Println(strings.TrimSpace(resp))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(weekCmd)
	weekCmd.Flags().StringVar(&weekISO, "week", "", "ISO week like 2025-W32")
	weekCmd.Flags().StringVar(&fromDate, "from", "", "Start date YYYY-MM-DD")
	weekCmd.Flags().StringVar(&toDate, "to", "", "End date YYYY-MM-DD")
}
