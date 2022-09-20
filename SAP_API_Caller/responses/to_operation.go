package responses

type ToOperation struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			OrderInternalBillOfOperations  string `json:"OrderInternalBillOfOperations"`
			OrderIntBillOfOperationsItem   string `json:"OrderIntBillOfOperationsItem"`
			ManufacturingOrder             string `json:"ManufacturingOrder"`
			ManufacturingOrderSequence     string `json:"ManufacturingOrderSequence"`
			MfgOrderSequenceText           string `json:"MfgOrderSequenceText"`
			ManufacturingOrderOperation    string `json:"ManufacturingOrderOperation"`
			ManufacturingOrderSubOperation string `json:"ManufacturingOrderSubOperation"`
			ManufacturingOrderCategory     string `json:"ManufacturingOrderCategory"`
			ManufacturingOrderType         string `json:"ManufacturingOrderType"`
			MfgOrderOperationText          string `json:"MfgOrderOperationText"`
			OperationIsCreated             string `json:"OperationIsCreated"`
			OperationIsReleased            string `json:"OperationIsReleased"`
			OperationIsPrinted             string `json:"OperationIsPrinted"`
			OperationIsConfirmed           string `json:"OperationIsConfirmed"`
			OperationIsPartiallyConfirmed  string `json:"OperationIsPartiallyConfirmed"`
			OperationIsDeleted             string `json:"OperationIsDeleted"`
			OperationIsTechlyCompleted     string `json:"OperationIsTechlyCompleted"`
			OperationIsClosed              string `json:"OperationIsClosed"`
			OperationIsScheduled           string `json:"OperationIsScheduled"`
			OperationIsPartiallyDelivered  string `json:"OperationIsPartiallyDelivered"`
			OperationIsDelivered           string `json:"OperationIsDelivered"`
			ProductionPlant                string `json:"ProductionPlant"`
			WorkCenterInternalID           string `json:"WorkCenterInternalID"`
			WorkCenterTypeCode             string `json:"WorkCenterTypeCode"`
			WorkCenter                     string `json:"WorkCenter"`
			OperationControlProfile        string `json:"OperationControlProfile"`
			OpErlstSchedldExecStrtDte      string `json:"OpErlstSchedldExecStrtDte"`
			OpErlstSchedldExecStrtTme      string `json:"OpErlstSchedldExecStrtTme"`
			OpErlstSchedldExecEndDte       string `json:"OpErlstSchedldExecEndDte"`
			OpErlstSchedldExecEndTme       string `json:"OpErlstSchedldExecEndTme"`
			OpActualExecutionStartDate     string `json:"OpActualExecutionStartDate"`
			OpActualExecutionStartTime     string `json:"OpActualExecutionStartTime"`
			OpActualExecutionEndDate       string `json:"OpActualExecutionEndDate"`
			OpActualExecutionEndTime       string `json:"OpActualExecutionEndTime"`
			ErlstSchedldExecDurnInWorkdays int    `json:"ErlstSchedldExecDurnInWorkdays"`
			OpActualExecutionDays          int    `json:"OpActualExecutionDays"`
			OperationUnit                  string `json:"OperationUnit"`
			OpPlannedTotalQuantity         string `json:"OpPlannedTotalQuantity"`
			OpTotalConfirmedYieldQty       string `json:"OpTotalConfirmedYieldQty"`
			LastChangeDateTime             string `json:"LastChangeDateTime"`
		} `json:"results"`
	} `json:"d"`
}
