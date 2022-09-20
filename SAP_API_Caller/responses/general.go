package responses

type General struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ManufacturingOrder            string `json:"ManufacturingOrder"`
			ManufacturingOrderCategory    string `json:"ManufacturingOrderCategory"`
			ManufacturingOrderType        string `json:"ManufacturingOrderType"`
			OrderIsCreated                string `json:"OrderIsCreated"`
			OrderIsReleased               string `json:"OrderIsReleased"`
			OrderIsPrinted                string `json:"OrderIsPrinted"`
			OrderIsConfirmed              string `json:"OrderIsConfirmed"`
			OrderIsPartiallyConfirmed     string `json:"OrderIsPartiallyConfirmed"`
			OrderIsDelivered              string `json:"OrderIsDelivered"`
			OrderIsDeleted                string `json:"OrderIsDeleted"`
			OrderIsPreCosted              string `json:"OrderIsPreCosted"`
			SettlementRuleIsCreated       string `json:"SettlementRuleIsCreated"`
			OrderIsPartiallyReleased      string `json:"OrderIsPartiallyReleased"`
			OrderIsLocked                 string `json:"OrderIsLocked"`
			OrderIsTechnicallyCompleted   string `json:"OrderIsTechnicallyCompleted"`
			OrderIsClosed                 string `json:"OrderIsClosed"`
			OrderIsPartiallyDelivered     string `json:"OrderIsPartiallyDelivered"`
			OrderIsMarkedForDeletion      string `json:"OrderIsMarkedForDeletion"`
			OrderIsScheduled              string `json:"OrderIsScheduled"`
			OrderHasGeneratedOperations   string `json:"OrderHasGeneratedOperations"`
			MaterialAvailyIsNotChecked    string `json:"MaterialAvailyIsNotChecked"`
			MfgOrderCreationDate          string `json:"MfgOrderCreationDate"`
			MfgOrderCreationTime          string `json:"MfgOrderCreationTime"`
			LastChangeDateTime            string `json:"LastChangeDateTime"`
			Material                      string `json:"Material"`
			StorageLocation               string `json:"StorageLocation"`
			GoodsRecipientName            string `json:"GoodsRecipientName"`
			UnloadingPointName            string `json:"UnloadingPointName"`
			MaterialGoodsReceiptDuration  string `json:"MaterialGoodsReceiptDuration"`
			OrderInternalBillOfOperations string `json:"OrderInternalBillOfOperations"`
			ProductionPlant               string `json:"ProductionPlant"`
			Plant                         string `json:"Plant"`
			MRPArea                       string `json:"MRPArea"`
			MRPController                 string `json:"MRPController"`
			ProductionSupervisor          string `json:"ProductionSupervisor"`
			ProductionVersion             string `json:"ProductionVersion"`
			PlannedOrder                  string `json:"PlannedOrder"`
			SalesOrder                    string `json:"SalesOrder"`
			SalesOrderItem                string `json:"SalesOrderItem"`
			BasicSchedulingType           string `json:"BasicSchedulingType"`
			BusinessArea                  string `json:"BusinessArea"`
			CompanyCode                   string `json:"CompanyCode"`
			ProfitCenter                  string `json:"ProfitCenter"`
			FunctionalArea                string `json:"FunctionalArea"`
			MfgOrderPlannedStartDate      string `json:"MfgOrderPlannedStartDate"`
			MfgOrderPlannedStartTime      string `json:"MfgOrderPlannedStartTime"`
			MfgOrderPlannedEndDate        string `json:"MfgOrderPlannedEndDate"`
			MfgOrderPlannedEndTime        string `json:"MfgOrderPlannedEndTime"`
			MfgOrderScheduledStartDate    string `json:"MfgOrderScheduledStartDate"`
			MfgOrderScheduledStartTime    string `json:"MfgOrderScheduledStartTime"`
			MfgOrderScheduledEndDate      string `json:"MfgOrderScheduledEndDate"`
			MfgOrderScheduledEndTime      string `json:"MfgOrderScheduledEndTime"`
			MfgOrderActualReleaseDate     string `json:"MfgOrderActualReleaseDate"`
			ProductionUnit                string `json:"ProductionUnit"`
			TotalQuantity                 string `json:"TotalQuantity"`
			MfgOrderPlannedScrapQty       string `json:"MfgOrderPlannedScrapQty"`
			MfgOrderConfirmedYieldQty     string `json:"MfgOrderConfirmedYieldQty"`
			WBSElementExternalID          string `json:"WBSElementExternalID"`
			OrderLongText                 string `json:"OrderLongText"`
			ToComponent                   struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_ProductionOrderComponent"`
			ToItem struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_ProductionOrderItem"`
			ToOperation struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_ProductionOrderOperation"`
			ToStatus struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_ProductionOrderStatus"`
		} `json:"results"`
	} `json:"d"`
}
