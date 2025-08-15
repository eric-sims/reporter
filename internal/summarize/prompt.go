// Package summarize summarizes the week's summaries
package summarize

import (
	"fmt"
	"strings"
	"time"

	"github.com/eric-sims/reporter/internal/model"
)

// WeeklyPrompt - custom prompt for aggregating the week's daily summaries
func WeeklyPrompt(entries []model.Summary, start, end time.Time) string {
	var b strings.Builder
	fmt.Fprintf(&b, "You are an assistant helping write a concise weekly work recap for a software engineer.\n")
	fmt.Fprintf(&b, "Summarize the following daily notes from %s to %s.\n", start.Format(time.DateOnly), end.Format(time.DateOnly))
	fmt.Fprintf(&b, "Produce a readable summary with:\n- Highlights\n- Completed work\n- In-progress / blockers\n- Notable metrics or PRs\n\n")
	for _, e := range entries {
		fmt.Fprintf(&b, "## %s\n", e.Date.Format("Mon 2006-01-02"))
		fmt.Fprintf(&b, "%s\n\n", strings.TrimSpace(e.Text))
	}
	fmt.Fprintf(&b, "Keep it under 250 words if possible. Use bullet points when appropriate. Try to not make things up and stick to the information given.")
	return b.String()
}
