package service

import (
	"fmt"
	"github.com/Aserose/CaduceusTour/internal/repository/models"
	"github.com/LeKovr/num2word"
	"os"
	"strconv"
	"strings"
	"time"
)

var counter int

func (s *service) ToOverview() {
	s.GetDataFromDataSource(map[string]string{
		"name": s.contracts[counter-1].Name,
		"time": time.Now().Format(timeLayoutToRequest)})
}

func (s *service) CreateDocument(chatId int64) {
	s.log.Info("service: create a document")
	file, _ := os.Create("result.txt")
	s.WriteDocument(file)
	s.tg.SendFile(chatId)
	err := os.Remove("result.txt")
	if err != nil {
		s.log.Errorf("service: document deletion error, %v", err.Error())
	}
}

func (s *service) WriteDocument(file *os.File) {

	for _, part := range s.contracts {
		_, err := file.Write([]byte(fmt.Sprintf("%s\n\n%s\n\n", s.formatView(part),
			func() string {
				var a string
				for i := 0; i <= 50; i++ {
					a += "-"
				}
				return a
			}())))
		if err != nil {
			s.log.Error(err.Error())
		}
	}
}

func (s *service) NextRow(chatId int64) {
	counter = 19
	s.Next(chatId)
}

func (s *service) Next(chatId int64) {

	if counter >= len(s.contracts) {
		s.log.Info("service: reset counter")
		counter = 0
		s.Update()
	}

	s.tg.SendMessage(chatId, s.formatView(s.contracts[counter]))
	counter++
}

func (s *service) formatView(contract models.ContractInfo) string {
	return fmt.Sprintf(s.strCfg.Response.View.Format,
		counter+1, len(s.contracts),
		strings.Title(strings.ToLower(contract.Name)),
		contract.Region,
		contract.PurchaseObjectInfo,
		s.parseTimeToString(contract.PurchaseCreatedAt),
		contract.PurchaseType,
		contract.ProcedureInfo.CollectingInfo.Order,
		contract.ProcedureInfo.CollectingInfo.Place,
		contract.MaxPrice+" руб.",
		num2word.RuMoney(func() float64 { back, _ := strconv.ParseFloat(contract.MaxPrice, 64); return back }(), false),
		contract.State,
		s.parseTimeToString(contract.PurchaseUpdatedAt))
}
