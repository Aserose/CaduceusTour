package gosplan

import (
	"bytes"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"io"
	"log"
	"net/http"
	"time"
)

func (g *gosplan) init() (string, time.Time) {
	value, err := sjson.SetBytes([]byte(
		fmt.Sprintf(`{"grant_type":"%s","client_id":"%s","client_secret":"%s"}`,
			g.cfg.GrantType, g.cfg.ClientId, g.cfg.ClientSecret)), "", nil)
	if err != nil {
		g.log.Error(err.Error())
	}

	log.Print(string(value))

	req, err := http.NewRequest("POST", g.cfg.UrlAuth, bytes.NewBuffer(value))
	if err != nil {
		g.log.Error(err.Error())
	}

	req.Header.Add("Content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		g.log.Error(err.Error())
	}

	read, err := io.ReadAll(resp.Body)
	if err != nil {
		g.log.Error(err.Error())
	}

	g.log.Info("initGosplan ok")

	return fmt.Sprint(gjson.ParseBytes(read).Get("access_token")), time.Now()
}

func (g *gosplan) getAccess() {
	switch g.gpData.AccessData.GetToken(time.Now()) {
	case "overdue":
		g.gpData.AccessData.UpdateToken(g.init())
	case "empty":
		g.gpData.AccessData.PutToken(g.init())
	}
}
