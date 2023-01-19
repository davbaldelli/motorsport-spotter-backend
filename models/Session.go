package models

type Session struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	EventId      int    `json:"eventId"`
	Date         string `json:"date"`
	Time         string `json:"time"`
	Timezone     string `json:"timezone"`
	DurationMin  int    `json:"durationMin"`
	DurationLaps int    `json:"durationLaps"`
}
