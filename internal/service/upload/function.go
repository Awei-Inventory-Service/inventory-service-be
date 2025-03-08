package upload

import (
	"context"
	"fmt"
	"strconv"

	"github.com/inventory-service/internal/dto"
	"github.com/inventory-service/internal/model"
	constant "github.com/inventory-service/lib/constants"
	"github.com/inventory-service/lib/error_wrapper"
	excel "github.com/inventory-service/lib/utils"
)

func (u *uploadService) ParseTransactionExcel(ctx context.Context, fileName string) *error_wrapper.ErrorWrapper {
	requiredHeaders := []string{
		constant.Date,
		constant.ProductCode,
		constant.Quantity,
		constant.Type,
	}
	_, content, err := excel.ReadExcel(fileName, "transaction")

	if err != nil {
		fmt.Println("Error parsing excel data", err.Error())
		return error_wrapper.New(model.SErrFailParseExcel, err.Error())
	}
	fmt.Println("REQUIRED", requiredHeaders)
	fmt.Println("INI CONTENT", content)

	missingHeaders := []string{}

	for _, line := range content {
		createSales := dto.CreateSalesRequest{
			BranchID: fmt.Sprint(ctx.Value("branch_id")),
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
		fmt.Println("iNi product code", productCode)
		product, errW := u.productRespository.Find(ctx, model.GetProduct{
			Code: productCode,
		})
		fmt.Println("iNI PRODUCT", product)
		if errW != nil {
			return errW
		}

		createSales.ProductID = product[0].ID
		quantity := line[constant.Quantity]
		number, err := strconv.Atoi(quantity)

		if err != nil {
			return error_wrapper.New(model.SErrParsingExcelQuantity, fmt.Sprintf("Error parsing %s to int", quantity))
		}

		productType := line[constant.Type]

		// TO DO Validasi tipe produk yg di input

		createSales.Quantity = number
		createSales.Type = productType

		errW = u.salesService.Create(ctx, createSales)

		if errW != nil {
			return errW
		}
	}

	return nil
}
