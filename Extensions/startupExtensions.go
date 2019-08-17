package extensions

import (
	"strconv"

	schedule "github.com/femosc2/ia-discord-bot-2/Features/Schedule"
	viewmodels "github.com/femosc2/ia-discord-bot-2/viewmodels"
)

// StartupScrape Creates the exams on startup
func StartupScrape() viewmodels.ScheduleViewModel {
	url := "http://schema.mah.se/setup/jsp/Schema.jsp?startDatum=idag&intervallTyp=m&intervallAntal=6&sprak=SV&sokMedAND=true&forklaringar=true&resurser=p.TGIAA"
	Schedules := viewmodels.ScheduleViewModel{
		IA17: schedule.GetExams(url + strconv.Itoa(17) + "h"),
		IA18: schedule.GetExams(url + strconv.Itoa(18) + "h"),
		IA19: schedule.GetExams(url + strconv.Itoa(19) + "h"),
	}
	return Schedules
}
