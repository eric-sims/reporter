// Package summarize summarizes the week's summaries
package summarize

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/eric-sims/reporter/internal/model"
)

// WeeklyPrompt - custom prompt for aggregating the week's daily summaries
func WeeklyPrompt(entries []model.Summary, start, end time.Time) string {
	var b strings.Builder
	fmt.Fprintf(&b, "You are an assistant helping write a concise weekly work recap for a software engineer.\n")
	fmt.Fprintf(&b, "Summarize the following daily notes from %s to %s.\n", start.Format("2006-01-02"), end.Format("2006-01-02"))
	fmt.Fprintf(&b, "Produce a readable summary with:\n- Highlights\n- Completed work\n- In-progress / blockers\n- Notable metrics or PRs\n- Next week focus\n\n")
	for _, e := range entries {
		fmt.Fprintf(&b, "## %s\n", e.Date.Format("Mon 2006-01-02"))
		if len(e.Projects) > 0 {
			var ps []string
			for _, p := range e.Projects {
				ps = append(ps, p.Name)
			}
			fmt.Fprintf(&b, "Projects: %s\n", strings.Join(ps, ", "))
		}
		if len(e.Tags) > 0 {
			var ts []string
			for _, t := range e.Tags {
				ts = append(ts, t.Name)
			}
			fmt.Fprintf(&b, "Tags: %s\n", strings.Join(ts, ", "))
		}
		fmt.Fprintf(&b, "%s\n\n", strings.TrimSpace(e.Text))
	}
	fmt.Fprintf(&b, "Keep it under 250 words if possible. Use bullet points when appropriate.")
	return b.String()
}

// WeeklyPromptByProject - filters by project
func WeeklyPromptByProject(entries []model.Summary, start, end time.Time) string {
	groups := map[string][]model.Summary{}
	for _, e := range entries {
		if len(e.Projects) == 0 {
			groups["(No Project)"] = append(groups["(No Project)"], e)
			continue
		}
		for _, p := range e.Projects {
			groups[p.Name] = append(groups[p.Name], e)
		}
	}
	keys := make([]string, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var b strings.Builder
	fmt.Fprintf(&b, "You are an assistant creating a weekly report grouped by project for a software engineer.\n")
	fmt.Fprintf(&b, "Consider entries from %s to %s. For each project, write a short section with:\n- Highlights\n- Completed\n- In-progress / blockers\n- Next week focus\nKeep each project section under 120 words.\n\n", start.Format("2006-01-02"), end.Format("2006-01-02"))
	for _, k := range keys {
		fmt.Fprintf(&b, "### Project: %s\n", k)
		for _, e := range groups[k] {
			fmt.Fprintf(&b, "- %s: %s\n", e.Date.Format("Mon 2006-01-02"), strings.TrimSpace(e.Text))
		}
		b.WriteString("\n")
	}
	return b.String()
}
