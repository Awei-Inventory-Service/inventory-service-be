package upload

import (
	"context"
	"fmt"
	"strconv"

	"github.com/inventory-service/dto"
	constant "github.com/inventory-service/lib/constants"
	"github.com/inventory-service/lib/error_wrapper"
	excel "github.com/inventory-service/lib/utils"
	"github.com/inventory-service/model"
)

func (u *uploadService) ParseTransactionExcel(ctx context.Context, fileName string, branchId string) *error_wrapper.ErrorWrapper {
	requiredHeaders := []string{
		constant.Date,
		constant.ProductCode,
		constant.Quantity,
		constant.Type,
	}
	_, content, err := excel.ReadExcel(fileName, "transaction")

	if err != nil {
		return error_wrapper.New(model.SErrFailParseExcel, err.Error())
	}

	missingHeaders := []string{}

	for _, line := range content {
		createSales := dto.CreateSalesRequest{
			BranchID: branchId,
		}

		for _, requiredHeader := range requiredHeaders {
			if line[requiredHeader] == "" {
				missingHeaders = append(missingHeaders, requiredHeader)
			}
		}

		if len(missingHeaders) > 0 {
			return error_wrapper.New(model.SErrExcelMissingRequiredData, fmt.Sprintf("%s is required", missingHeaders[0]))
		}

		productCode := line[constant.ProductCode]
		product, errW := u.productRespository.FindByID(ctx, productCode)

		if errW != nil {
			return errW
		}

		createSales.ProductID = product.UUID
		quantity := line[constant.Quantity]
		number, err := strconv.Atoi(quantity)

		if err != nil {
			return error_wrapper.New(model.SErrParsingExcelQuantity, fmt.Sprintf("Error parsing %s to int", quantity))
		}

		productType := line[constant.Type]

		// TO DO Validasi tipe produk yg di input

		createSales.Quantity = float64(number)
		createSales.Type = productType

		errW = u.salesService.Create(ctx, createSales, "")

		if errW != nil {
			return errW
		}
	}

	return nil
}
