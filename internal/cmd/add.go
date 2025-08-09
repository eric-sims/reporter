// Package cmd - definitions for the cobra-cli
package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/eric-sims/reporter/internal/db"
	rmodel "github.com/eric-sims/reporter/internal/model"
	"github.com/eric-sims/reporter/internal/util"
	"github.com/spf13/cobra"
)

var (
	addDate string
	addText string
	addFile string
	addEdit bool
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a daily summary",
	RunE: func(_ *cobra.Command, _ []string) error {
		if addText == "" && addFile == "" && !addEdit {
			// Read from stdin if piped, else open editor
			fi, _ := os.Stdin.Stat()
			if (fi.Mode() & os.ModeCharDevice) == 0 {
				// piped
				s := bufio.NewScanner(os.Stdin)
				var b strings.Builder
				for s.Scan() {
					b.WriteString(s.Text())
					b.WriteString("\n")
				}
				addText = strings.TrimSpace(b.String())
			} else {
				addEdit = true
			}
		}

		if addEdit {
			edited, err := openInEditor()
			if err != nil {
				return err
			}
			addText = edited
		}

		if addFile != "" {
			b, err := os.ReadFile(addFile)
			if err != nil {
				return err
			}
			addText = string(b)
		}

		if strings.TrimSpace(addText) == "" {
			return errors.New("no summary text provided")
		}

		when, err := util.ParseDateOrToday(addDate)
		if err != nil {
			return err
		}

		database, err := db.Open()
		if err != nil {
			return err
		}
		defer database.Close()

		s := rmodel.Summary{Date: when, Text: addText}
		if err := database.UpsertSummary(&s); err != nil {
			return err
		}
		fmt.Printf("Saved summary for %s (id=%d)\n", when.Format("2006-01-02"), s.ID)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVar(&addDate, "date", time.Now().Format("2006-01-02"), "Date for the summary (YYYY-MM-DD)")
	addCmd.Flags().StringVar(&addText, "text", "", "Summary text (if empty, reads stdin or opens $EDITOR)")
	addCmd.Flags().StringVar(&addFile, "file", "", "Read summary text from file")
	addCmd.Flags().BoolVar(&addEdit, "edit", false, "Open $EDITOR to write the summary")
}

func openInEditor() (string, error) {
	ed := os.Getenv("EDITOR")
	if ed == "" {
		ed = "vi"
	}
	tf, err := os.CreateTemp("", "reporter-*.txt")
	if err != nil {
		return "", err
	}
	defer os.Remove(tf.Name())

	c := exec.Command(ed, tf.Name())
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if rErr := c.Run(); rErr != nil {
		return "", rErr
	}

	b, reErr := os.ReadFile(tf.Name())
	if reErr != nil {
		return "", reErr
	}
	return string(b), nil
}
