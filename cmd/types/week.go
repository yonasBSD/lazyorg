package types

import "time"

type Week struct {
	StartDate time.Time
	EndDate   time.Time
	Days      []Day
}

func NewWeek(startDate time.Time, endDate time.Time, days []Day) *Week {
	return &Week{StartDate: startDate, EndDate: endDate, Days: days}
}

func (w *Week) InitDays() {
    w.Days = make([]Day, 7)
    
    w.Days[0] = *NewDay(time.Time{}, nil)
    w.Days[1] = *NewDay(time.Time{}, nil)
    w.Days[2] = *NewDay(time.Time{}, nil)
    w.Days[3] = *NewDay(time.Time{}, nil)
    w.Days[4] = *NewDay(time.Time{}, nil)
    w.Days[5] = *NewDay(time.Time{}, nil)
    w.Days[6] = *NewDay(time.Time{}, nil)
}

