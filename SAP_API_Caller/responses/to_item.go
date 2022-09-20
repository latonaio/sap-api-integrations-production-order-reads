package responses

type ToItem struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ManufacturingOrder             string `json:"ManufacturingOrder"`
			ManufacturingOrderItem         string `json:"ManufacturingOrderItem"`
			ManufacturingOrderCategory     string `json:"ManufacturingOrderCategory"`
			ManufacturingOrderType         string `json:"ManufacturingOrderType"`
			IsCompletelyDelivered          bool   `json:"IsCompletelyDelivered"`
			Material                       string `json:"Material"`
			ProductionPlant                string `json:"ProductionPlant"`
			Plant                          string `json:"Plant"`
			MRPArea                        string `json:"MRPArea"`
			QuantityDistributionKey        string `json:"QuantityDistributionKey"`
			MaterialGoodsReceiptDuration   string `json:"MaterialGoodsReceiptDuration"`
			StorageLocation                string `json:"StorageLocation"`
			Batch                          string `json:"Batch"`
			InventoryUsabilityCode         string `json:"InventoryUsabilityCode"`
			GoodsRecipientName             string `json:"GoodsRecipientName"`
			UnloadingPointName             string `json:"UnloadingPointName"`
			MfgOrderItemPlndDeliveryDate   string `json:"MfgOrderItemPlndDeliveryDate"`
			MfgOrderItemActualDeliveryDate string `json:"MfgOrderItemActualDeliveryDate"`
			ProductionUnit                 string `json:"ProductionUnit"`
			MfgOrderItemPlannedTotalQty    string `json:"MfgOrderItemPlannedTotalQty"`
			MfgOrderItemPlannedScrapQty    string `json:"MfgOrderItemPlannedScrapQty"`
			MfgOrderItemGoodsReceiptQty    string `json:"MfgOrderItemGoodsReceiptQty"`
			MfgOrderItemActualDeviationQty string `json:"MfgOrderItemActualDeviationQty"`
		} `json:"results"`
	} `json:"d"`
}
