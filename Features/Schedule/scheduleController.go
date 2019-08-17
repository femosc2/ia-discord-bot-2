package schedule

import (
	scraperutils "github.com/femosc2/ia-discord-bot-2/Utilities/ScheduleScraper"
)

// GetExams Get upcoming exams
func GetExams(url string) []string {
	return scraperutils.GetExams(url)
}
