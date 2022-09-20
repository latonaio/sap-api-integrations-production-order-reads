package responses

type ToComponent struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Reservation                    string `json:"Reservation"`
			ReservationItem                string `json:"ReservationItem"`
			MaterialGroup                  string `json:"MaterialGroup"`
			Material                       string `json:"Material"`
			Plant                          string `json:"Plant"`
			ManufacturingOrderCategory     string `json:"ManufacturingOrderCategory"`
			ManufacturingOrderType         string `json:"ManufacturingOrderType"`
			ManufacturingOrder             string `json:"ManufacturingOrder"`
			ManufacturingOrderSequence     string `json:"ManufacturingOrderSequence"`
			ManufacturingOrderOperation    string `json:"ManufacturingOrderOperation"`
			ProductionPlant                string `json:"ProductionPlant"`
			OrderInternalBillOfOperations  string `json:"OrderInternalBillOfOperations"`
			MatlCompRequirementDate        string `json:"MatlCompRequirementDate"`
			MatlCompRequirementTime        string `json:"MatlCompRequirementTime"`
			ReservationIsFinallyIssued     bool   `json:"ReservationIsFinallyIssued"`
			IsBulkMaterialComponent        bool   `json:"IsBulkMaterialComponent"`
			MatlCompIsMarkedForBackflush   bool   `json:"MatlCompIsMarkedForBackflush"`
			MaterialCompIsCostRelevant     string `json:"MaterialCompIsCostRelevant"`
			SalesOrder                     string `json:"SalesOrder"`
			SalesOrderItem                 string `json:"SalesOrderItem"`
			SortField                      string `json:"SortField"`
			BillOfMaterialCategory         string `json:"BillOfMaterialCategory"`
			BOMItem                        string `json:"BOMItem"`
			BOMItemCategory                string `json:"BOMItemCategory"`
			BillOfMaterialItemNumber       string `json:"BillOfMaterialItemNumber"`
			BOMItemDescription             string `json:"BOMItemDescription"`
			StorageLocation                string `json:"StorageLocation"`
			Batch                          string `json:"Batch"`
			BatchSplitType                 string `json:"BatchSplitType"`
			GoodsMovementType              string `json:"GoodsMovementType"`
			SupplyArea                     string `json:"SupplyArea"`
			GoodsRecipientName             string `json:"GoodsRecipientName"`
			UnloadingPointName             string `json:"UnloadingPointName"`
			MaterialCompIsAlternativeItem  bool   `json:"MaterialCompIsAlternativeItem"`
			AlternativeItemGroup           string `json:"AlternativeItemGroup"`
			AlternativeItemStrategy        string `json:"AlternativeItemStrategy"`
			AlternativeItemPriority        string `json:"AlternativeItemPriority"`
			UsageProbabilityPercent        string `json:"UsageProbabilityPercent"`
			MaterialComponentIsPhantomItem bool   `json:"MaterialComponentIsPhantomItem"`
			LeadTimeOffset                 string `json:"LeadTimeOffset"`
			QuantityIsFixed                bool   `json:"QuantityIsFixed"`
			IsNetScrap                     bool   `json:"IsNetScrap"`
			ComponentScrapInPercent        string `json:"ComponentScrapInPercent"`
			OperationScrapInPercent        string `json:"OperationScrapInPercent"`
			BaseUnit                       string `json:"BaseUnit"`
			RequiredQuantity               string `json:"RequiredQuantity"`
			WithdrawnQuantity              string `json:"WithdrawnQuantity"`
			ConfirmedAvailableQuantity     string `json:"ConfirmedAvailableQuantity"`
			MaterialCompOriginalQuantity   string `json:"MaterialCompOriginalQuantity"`
			Currency                       string `json:"Currency"`
			WithdrawnQuantityAmount        string `json:"WithdrawnQuantityAmount"`
		} `json:"results"`
	} `json:"d"`
}
