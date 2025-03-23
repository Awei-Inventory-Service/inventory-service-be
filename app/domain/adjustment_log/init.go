package adjustmentlog

import adjustmentlog "github.com/inventory-service/app/resource/adjustment_log"

func NewAdjustmentLogDomain(adjustmentLogResource adjustmentlog.AdjustmentLogResource) AdjusmentLogDomain {
	return &adjustmentLogDomain{adjusmentLogResource: adjustmentLogResource}
}
