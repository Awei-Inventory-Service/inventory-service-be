package transferlog

import transferlog "github.com/inventory-service/app/resource/transfer_log"

func NewTransferLogDomain(transferLogResource transferlog.TransferLogResource) TransferLogDomain {
	return &transferLogDomain{transferLogResource: transferLogResource}
}
