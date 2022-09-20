package responses

type ToStatus struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ManufacturingOrder string `json:"ManufacturingOrder"`
			StatusCode         string `json:"StatusCode"`
			IsUserStatus       bool   `json:"IsUserStatus"`
			StatusShortName    string `json:"StatusShortName"`
			StatusName         string `json:"StatusName"`
		} `json:"results"`
	} `json:"d"`
}
