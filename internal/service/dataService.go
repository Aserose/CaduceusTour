package service

import (
	"github.com/araddon/dateparse"
	"strconv"
	"time"
)

const (
	timeLayoutToRequest = "2006-01-02T15:04:05.000Z"
)

func (s *service) GetDataFromDataSource(params map[string]string) string {
	counter = 0
	overviewCounter = 0
	var status string
	s.contracts, status = s.sources.RequestToDataSource(params)
	return status
}

func (s *service) Update() {
	s.log.Info("service: updating data")
	s.contracts, _ = s.sources.RequestToDataSource(map[string]string{
		"name": s.contracts[len(s.contracts)-1].Name,
		"time": func() string {
			t1, err := time.Parse(
				time.RFC3339,
				s.contracts[len(s.contracts)-1].PurchaseCreatedAt)
			if err != nil {
				s.log.Error(err.Error())
			}
			return t1.Format(timeLayoutToRequest)
		}()})
}

func (s service) parseTimeToString(time string) string {
	a, err := dateparse.ParseAny(time)
	if err != nil {
		s.log.Error(err.Error())
	}

	return a.Format(timeLayout)
}

func (s service) parseStringToTime(time string) time.Time {
	a, err := dateparse.ParseAny(time)
	if err != nil {
		s.log.Error(err.Error())
	}

	return a
}

func parseFloatToInt(str string) int {
	a, _:= strconv.ParseFloat(str, 64)

	return int(a)
}
