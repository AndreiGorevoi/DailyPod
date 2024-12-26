package formatter

import (
	"DailyPod/dto"
	"fmt"
	"strings"
	"time"
)

const outputFormat = "Mon,Jan 2 - 3:04PM(MST)"

func FormatGamesToString(games []dto.Game) string {
	b := strings.Builder{}

	// Estimate total size
	estimatedSize := len(games) * 70
	b.Grow(estimatedSize)

	for i, g := range games {
		l, _ := time.LoadLocation("Europe/Warsaw")
		date := g.Date.Start.In(l).Format(outputFormat)
		b.WriteString(fmt.Sprintf("%d: %s - %s. | %s\n", i+1, g.Teams.Home.Name, g.Teams.Visitors.Name, date))
	}

	return b.String()
}
