package schedules

import "time"

type Schedule struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Cron       Cron      `json:"cron"`
	Active     bool      `json:"is_active"`
	Processing bool      `json:"is_processing"`
	LastRunAt  time.Time `json:"last_run_at"`
	NextRunAt  time.Time `json:"next_run_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Tasks      []Task    `json:"tasks"`
} 

type Task struct {
	Id         int       `json:"id"`
	Sequence   int       `json:"sequence_id"`
	Action     string    `json:"action"`
	Payload    string    `json:"payload"`
	TimeOffset int       `json:"time_offset"`
	Queued     bool      `json:"is_queued"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Cron struct {
	DayOfWeek  string `json:"day_of_week"`
	DayOfMonth string `json:"day_of_month"`
	Hour       string `json:"hour"`
	Minute     string `json:"minute"`
}
