package middleware

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func (m *middleware) getDataAboutOrganization(msg string, id int64) string {
	m.log.Info("middleware: sending request to dataSource")
	m.service.SendMessage(id, m.strCfg.Response.MsgToUser.Processing)

	if msg == "" {
		m.service.SendMessage(id, m.strCfg.Response.MsgToUser.None)
		return m.strCfg.Status.StatusSearch.None
	}

	status := m.service.GetDataFromDataSource(map[string]string{"name": msg, "time": time.Now().Format(time.RFC3339)})
	if status == m.strCfg.Status.StatusSearch.None {
		m.service.SendMessage(id, m.strCfg.Response.MsgToUser.None)
		return status
	}
	return ""
}

func (m *middleware) getOrganizationNameByAddress(address string) string {
	m.log.Info("middleware: search for an organization by address")
	resp, err := http.Get(fmt.Sprintf("%s%s&work=on", m.cfg.SearchUrl, url.QueryEscape(address)))
	if err != nil {
		m.log.Error(err.Error())
	}

	return strings.Split(m.readNameFromUrl(resp.Body), "инн")[0]
}

func (m *middleware) readNameFromUrl(r io.Reader) string {
	m.log.Info("middleware: getting name")
	var result string

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		m.log.Error(err.Error())
	}

	doc.Find("label").First().Each(func(i int, s *goquery.Selection) {
		result = fmt.Sprintf("%s", s.Find("span").Text())
	})

	return result
}

func (m middleware) getDataAboutRegion(msg string, id int64) string {
	m.log.Info("middleware: sending request to dataSource")
	m.service.SendMessage(id, m.strCfg.Response.MsgToUser.Processing)

	if msg == "" {
		m.service.GetDataFromDataSource(map[string]string{})
		return ""
	}

	if _, err := strconv.Atoi(msg); err != nil {
		m.service.SendMessage(id, m.strCfg.Response.MsgToUser.Incorrect)
		return m.strCfg.Status.StatusSearch.None
	} else {
		if m.service.GetDataFromDataSource(map[string]string{"region": msg}) == m.strCfg.Status.StatusSearch.None {
			m.service.SendMessage(id, m.strCfg.Response.MsgToUser.None)
			return m.strCfg.Status.StatusSearch.None
		}
	}
	return ""
}
