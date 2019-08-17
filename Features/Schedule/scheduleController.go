package schedule

import (
	"fmt"

	scraperutils "github.com/femosc2/ia-discord-bot-2/Utilities/ScheduleScraper"
)

func GetExams(url string) []string {
	fmt.Println("controller runs")
	return scraperutils.GetExams(url)
}
