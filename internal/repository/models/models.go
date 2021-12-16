package models

import "time"

type GP struct {
	Token    string
	TokenTTL time.Time
}

type Organization struct {
	Name    string `json:"name"`
	Region  int    `json:"region"`
	Payment `json:"payment"`
}

type Payment struct {
	Amount        int              `json:"amount"`
	NumberOffer   int              `json:"numberOffer"`
	Average       int              `json:"average"`
	AveragePeriod []int            `json:"averagePeriod"`
	Max           int              `json:"max"`
	Min           int              `json:"min"`
	InfoMax       string           `json:"infoMax"`
	InfoMin       string           `json:"infoMin"`
	MaxMonth      []map[int]string `json:"monthMax"`
	MinMonth      []map[int]string `json:"monthMin"`
	AverageMonths [2][]int         `json:"averageMonths"`
	Timeline      `json:"timeline"`
}

type Timeline struct {
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	StartPeriod time.Time `json:"startPeriod"`
}

type ContractInfo struct {
	Name               string `json:"responsible_name"`
	Region             int    `json:"region"`
	PurchaseObjectInfo string `json:"purchase_object_info"`
	PurchaseCreatedAt  string `json:"purchase_created_at"`
	PurchaseType       string `json:"purchase_type"`
	ProcedureInfo      struct {
		CollectingInfo struct {
			Place string `json:"place"`
			Order string `json:"order"`
		} `json:"collecting"`
	} `json:"procedure_info"`
	MaxPrice          string `json:"max_price"`
	State             string `json:"state"`
	PurchaseUpdatedAt string `json:"purchase_updated_at"`
}
