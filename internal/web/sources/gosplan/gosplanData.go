package gosplan

import (
	"encoding/json"
	"fmt"
	"github.com/Aserose/CaduceusTour/internal/repository/models"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func (g *gosplan) RequestToDataSource(params map[string]string) ([]models.ContractInfo, string) {
	g.getAccess()
	g.log.Info("source: requisitioning")

	resp := g.sendRequest(g.urlRequest(params))

	return g.readResponse(resp)
}

func (g *gosplan) urlRequest(params map[string]string) string {
	g.log.Info("source: getting request url")
	if len(params["name"]) >= 1 && len(params["time"]) >= 1 {
		return fmt.Sprintf("%spurchases?responsible_name=%s&purchase_created_at_before=%s&sorted_by=purchase_created_at_desc",
			g.cfg.UrlRequest, url.QueryEscape(params["name"]), params["time"])
	}
	if len(params["region"]) >= 1 {
		return fmt.Sprintf("%spurchases?region=%s&sorted_by=max_price_desc",
			g.cfg.UrlRequest, params["region"])
	}
	return fmt.Sprintf("%spurchases?sorted_by=max_price_desc",
		g.cfg.UrlRequest)
}

func (g *gosplan) sendRequest(urlReq string) *http.Response {
	g.log.Info("source: sending a request")
	req, err := http.NewRequest("GET", urlReq, nil); if err != nil {
		g.log.Errorf("source: %v", err.Error())
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GP_TOKEN")))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		g.log.Errorf("source: error on response, %v", err.Error())
	}

	return resp
}

func (g *gosplan) readResponse(resp *http.Response) ([]models.ContractInfo, string) {
	g.log.Info("source: response decoding")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		g.log.Errorf("Error while reading the response bytes: %v", err.Error())
	}

	if string(body) == "[]" {
		return nil, "not found"
	}

	defer resp.Body.Close()

	return g.respCompletion(body), " "
}

func (g *gosplan) respCompletion(body []byte) []models.ContractInfo {
	g.log.Info("source: end of response decoding")

	var contract []models.ContractInfo

	if err := json.Unmarshal(body, &contract); err != nil {
		g.log.Errorf("source: error while unmarshalling: %v", err.Error())
	}

	for i, part := range gjson.ParseBytes(body).Array() {
		if contract[i].ProcedureInfo.CollectingInfo.Place == "" {
			contract[i].ProcedureInfo.CollectingInfo.Place = part.Get("procedure_info.collecting_info.place").Str
		}
		if contract[i].ProcedureInfo.CollectingInfo.Order == "" {
			contract[i].ProcedureInfo.CollectingInfo.Order = part.Get("procedure_info.collecting_info.order").Str
		}
	}

	for i, part := range contract {
		contract[i].PurchaseType = g.transcription.Type[part.PurchaseType]
		contract[i].State = g.transcription.State[part.State]
	}

	return contract
}
