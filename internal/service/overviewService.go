package service

import (
	"fmt"
	"github.com/Aserose/CaduceusTour/internal/repository/models"
	"github.com/LeKovr/num2word"
	"github.com/araddon/dateparse"
	"log"
	"strings"
	"time"
)

var (
	overviewCounter int
	requestCounter  int
)

const (
	maxPrice     = "максимальная цена"
	minPrice     = "минимальная цена"
	offerSum     = "количество заявок"
	monthAvr     = "среднемесячная"
	avr          = "средняя"
	compResponse = "сравнение показателей организаций за %d мес. :\n1.%s\n2.%s\n\n %s"
	compDiffForm = "*список показателей, по значениям которых первая организация преобладает:\n" +
		"%s\n*список показателей, по значениям которых первая организация уступает:\n%s"
	sortDiffer = "%s (разница: %d)"
	timeLayout = "02.01.2006"
)

func (s *service) ToView() {
	s.GetDataFromDataSource(map[string]string{
		"name": s.contracts[len(s.contracts)-1].Name,
		"time": time.Now().Format("2006-01-02T15:04:05.000Z")})
}

func (s *service) Compare(id int64) {
	orgOne := s.organization
	s.General(id)

	months := reviewPeriod(len(s.organization.MaxMonth), len(orgOne.MaxMonth))

	differ := make(map[string]map[int]int)

	differ[maxPrice] = diff(maxMinMonth(orgOne.Payment.MaxMonth[months-1]), maxMinMonth(s.organization.Payment.MaxMonth[months-1]))
	differ[minPrice] = diff(maxMinMonth(orgOne.Payment.MinMonth[months-1]), maxMinMonth(s.organization.Payment.MinMonth[months-1]))
	differ[offerSum] = diff(orgOne.Payment.AverageMonths[1][months-1], s.organization.Payment.AverageMonths[1][months-1])
	differ[monthAvr] = diff(orgOne.AverageMonths[0][months-1], s.organization.AverageMonths[0][months-1])
	differ[avr] = diff(orgOne.AveragePeriod[months-1], s.organization.AveragePeriod[months-1])

	s.tg.SendMessage(id, s.CompareResponse(months, orgOne.Name, sorting(differ)))
}

func (s *service) CompareResponse(months int, nameOrgOne, differ string) string {

	return fmt.Sprintf(compResponse,
		months, nameOrgOne, s.organization.Name, differ)
}

func reviewPeriod(a, b int) int {
	if a < b {
		return b - (b - a)
	}
	return a - (a - b)
}

func diff(a, b int) map[int]int {
	if a < b {
		return map[int]int{1: b - a}
	}
	return map[int]int{0: a - b}
}

func maxMinMonth(month map[int]string) int {
	var result int
	for k := range month {
		result = k
	}
	return result
}

func sorting(rec map[string]map[int]int) string {
	var result [2][]string

	for name, maps := range rec {
		for organizationNumber, sum := range maps {
			if organizationNumber != 1 {
				result[0] = append(result[0], fmt.Sprintf(sortDiffer, name, sum))
			} else {
				result[1] = append(result[1], fmt.Sprintf(sortDiffer, name, sum))
			}
		}
	}

	return formattingCompare(result)
}

func formattingCompare(result [2][]string) string {
	return fmt.Sprintf(compDiffForm, strings.Join(result[0], "\n"), strings.Join(result[1], "\n"))
}

func (s *service) General(chatId int64) {
	s.organization = models.Organization{}
	requestCounter = 0
	overviewCounter = 0

	for {
		if len(s.contracts) < 19 {
			for _, part := range s.contracts {
				s.organization = s.iteration(part, s.organization)

				if s.organization.Payment.Timeline.StartPeriod.IsZero() {
					s.organization.Payment.Timeline.StartPeriod = s.organization.Payment.Timeline.StartDate
				}
			}
			break
		} else {
			s.organization = s.iteration(s.contracts[overviewCounter], s.organization)
		}

		if requestCounter > 5 {
			break
		}

		if overviewCounter >= len(s.contracts) {
			overviewCounter = 0
			s.Update()
			requestCounter++
		}

		if s.organization.Payment.Timeline.StartPeriod.IsZero() {
			s.organization.Payment.Timeline.StartPeriod = s.organization.Payment.Timeline.StartDate
		}
	}

	if len(s.organization.MaxMonth) < 1 {
		s.organization = addMonth(s.organization)
	}

	s.organization.Name = strings.Title(strings.ToLower(s.contracts[0].Name))
	s.organization.Region = s.contracts[len(s.contracts)-1].Region

	log.Print(s.organization)

	s.tg.SendMessage(chatId, s.formatOverview(s.organization))

}

func (s *service) iteration(contract models.ContractInfo, organization models.Organization) models.Organization {

	if contract.State == "закупка завершена" {
		if organization.Payment.Timeline.StartDate.IsZero() == true {
			organization.Payment.Timeline.StartDate, _ = dateparse.ParseAny(contract.PurchaseUpdatedAt)
		}

		organization.Payment.Amount += parseFloatToInt(contract.MaxPrice)
		organization.Payment.NumberOffer++
		organization.Payment.Average = organization.Payment.Amount / organization.Payment.NumberOffer // calculate the average amount of expenses

		if organization.Payment.Max < parseFloatToInt(contract.MaxPrice) {
			organization.Payment.Max = parseFloatToInt(contract.MaxPrice)
			organization.Payment.InfoMax = contract.PurchaseObjectInfo
		}

		if organization.Payment.Min == 0 || organization.Payment.Min > parseFloatToInt(contract.MaxPrice) {
			organization.Payment.Min = parseFloatToInt(contract.MaxPrice)
			organization.Payment.InfoMin = contract.PurchaseObjectInfo
		}

		organization.Payment.Timeline.EndDate, _ = dateparse.ParseAny(contract.PurchaseUpdatedAt)

		if organization.Payment.Timeline.StartDate.Sub(organization.Payment.Timeline.EndDate) >= 720*time.Hour {
			for t := organization.Payment.Timeline.EndDate; t.After(organization.Payment.Timeline.StartDate) == false; t = t.AddDate(0, 1, 0) {
				organization = addMonth(organization)
			}
			organization.Payment.Timeline.StartDate, _ = dateparse.ParseAny(contract.PurchaseUpdatedAt)
		}
	}

	overviewCounter++

	return organization
}

func addMonth(organization models.Organization) models.Organization {
	organization.Payment.AverageMonths[0] =
		append(organization.Payment.AverageMonths[0], organization.Payment.Amount/(len(organization.Payment.AverageMonths[0])+1))
	organization.Payment.AverageMonths[1] =
		append(organization.AverageMonths[1], organization.Payment.NumberOffer)

	organization.Payment.MaxMonth =
		append(organization.Payment.MaxMonth, map[int]string{organization.Payment.Max: organization.Payment.InfoMax})
	organization.Payment.MinMonth =
		append(organization.Payment.MinMonth, map[int]string{organization.Payment.Min: organization.Payment.InfoMin})

	organization.Payment.AveragePeriod = append(organization.Payment.AveragePeriod, organization.Payment.Average)

	return organization
}

func (s *service) formatOverview(organization models.Organization) string {
	if organization.Payment.NumberOffer < 1 {
		return fmt.Sprintf("нет завершенных закупок в заданный период")
	}

	return fmt.Sprintf(s.strCfg.Response.Overview.WithMonth,
		strings.Title(strings.ToLower(s.contracts[0].Name)),
		s.contracts[0].Region,
		organization.Payment.Timeline.EndDate.Format(timeLayout), organization.Payment.Timeline.StartPeriod.Format(timeLayout),
		len(organization.Payment.AverageMonths[0]), organization.Payment.AverageMonths[0][len(organization.Payment.AverageMonths[0])-1],
		num2word.RuMoney(float64(organization.Payment.AverageMonths[0][len(organization.Payment.AverageMonths[0])-1]), false),
		organization.Payment.NumberOffer, organization.Payment.Average, num2word.RuMoney(float64(organization.Payment.Average), false),
		organization.Payment.Max, num2word.RuMoney(float64(organization.Payment.Max), false), organization.Payment.InfoMax,
		organization.Payment.Min, num2word.RuMoney(float64(organization.Payment.Min), false), organization.Payment.InfoMin)
}
