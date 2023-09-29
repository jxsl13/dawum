package dto

import (
	"encoding/json"
	"time"
)

type Survey struct {
	Date            string             `json:"Date"`
	SurveyPeriod    SurveyPeriod       `json:"Survey_Period"`
	SurveyedPersons string             `json:"Surveyed_Persons"`
	ParliamentID    string             `json:"Parliament_ID"`
	InstituteID     string             `json:"Institute_ID"`
	TaskerID        string             `json:"Tasker_ID"`
	MethodID        string             `json:"Method_ID"`
	Results         map[string]float64 `json:"Results"`
}

type SurveyPeriod struct {
	DateStart time.Time `json:"Date_Start"`
	DateEnd   time.Time `json:"Date_End"`
}

func (sp *SurveyPeriod) UnmarshalJSON(b []byte) error {
	var primary struct {
		DateStart date `json:"Date_Start"`
		DateEnd   date `json:"Date_End"`
	}
	err := json.Unmarshal(b, &primary)
	if err != nil {
		return err
	}

	*sp = SurveyPeriod{
		DateStart: time.Time(primary.DateStart),
		DateEnd:   time.Time(primary.DateEnd),
	}

	return nil
}

func (sp SurveyPeriod) MarshalJSON() ([]byte, error) {
	primary := struct {
		DateStart date `json:"Date_Start"`
		DateEnd   date `json:"Date_End"`
	}{
		DateStart: date(sp.DateStart),
		DateEnd:   date(sp.DateEnd),
	}
	return json.Marshal(primary)
}
