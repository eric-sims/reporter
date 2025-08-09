package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	ollamaHost string
	model      string
)

var rootCmd = &cobra.Command{
	Use:   "reporter",
	Short: "Capture daily summaries and get AI-powered weekly recaps",
	Long:  "Reporter stores short daily work summaries in SQLite and asks a local Ollama LLM to generate a weekly recap.",
}

// Execute - executes cobra-cli
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&ollamaHost, "ollama", envOr("OLLAMA_HOST", "http://127.0.0.1:11434"), "Ollama host, e.g. http://127.0.0.1:11434")
	rootCmd.PersistentFlags().StringVar(&model, "model", envOr("OLLAMA_MODEL", "gpt-oss:20b"), "Ollama model to use for summarization")
}

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
