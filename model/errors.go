package model

import (
	"net/http"

	errW "github.com/inventory-service/lib/error_wrapper"
)

var (
	CategoryInternalServerError = errW.NewCategory(http.StatusInternalServerError, "Sorry sedang ada gangguan")
	CategoryUnAuthorized        = errW.NewCategory(http.StatusUnauthorized, "Sorry sedang ada gangguan")
	CategoryBadRequest          = errW.NewCategory(http.StatusBadRequest, "Sorry anda tidak bisa melakukan action tersebut")
	CategoryNotFound            = errW.NewCategory(http.StatusNotFound, "Sorry resource tidak ditemukan")
)

var (

	//	Default
	ErrIoUtilReadFile         = errW.NewDefinition(0, "Error Ioutil Read File", true, CategoryInternalServerError)
	ErrUnmarshalYaml          = errW.NewDefinition(1, "Error Unmarshal Yaml", true, CategoryInternalServerError)
	ErrNewMonitoring          = errW.NewDefinition(2, "Error New Monitoring", true, CategoryInternalServerError)
	ErrInitDB                 = errW.NewDefinition(3, "Error Init Db", true, CategoryInternalServerError)
	ErrInitRabbitMQ           = errW.NewDefinition(3, "Error Init Rabbit MQ", true, CategoryInternalServerError)
	ErrEncryptAES             = errW.NewDefinition(4, "Error Encrypt AES", true, CategoryInternalServerError)
	ErrEncryptToUrlEncoding   = errW.NewDefinition(5, "Error Encrypt to URL Encoding", true, CategoryInternalServerError)
	ErrQueryUnescape          = errW.NewDefinition(6, "Error Query unescape", true, CategoryInternalServerError)
	ErrNewCipher              = errW.NewDefinition(7, "Error New Cipher", true, CategoryInternalServerError)
	ErrBase64DecodeString     = errW.NewDefinition(8, "Error Base64 decode string", true, CategoryInternalServerError)
	ErrJsonUnmarshal          = errW.NewDefinition(9, "Error Json Unmarshal", true, CategoryInternalServerError)
	ErrJsonMarshal            = errW.NewDefinition(9, "Error Json Marshal", true, CategoryInternalServerError)
	ErrHeader                 = errW.NewDefinition(10, "Error Header", true, CategoryInternalServerError)
	ErrPrepareX               = errW.NewDefinition(11, "Error PrepareX %v", true, CategoryInternalServerError)
	ErrBeginTxx               = errW.NewDefinition(12, "Error Begin Transaction %v", true, CategoryInternalServerError)
	ErrExecContext            = errW.NewDefinition(13, "Error Exec Context %v", true, CategoryInternalServerError)
	ErrRowsAffected           = errW.NewDefinition(14, "Error Rows Affected %v", true, CategoryInternalServerError)
	ErrSelectContext          = errW.NewDefinition(15, "Error Select Context %v", true, CategoryInternalServerError)
	ErrStructToMap            = errW.NewDefinition(16, "Error Struct To Map not struct", true, CategoryInternalServerError)
	ErrNewRelicNewApplication = errW.NewDefinition(17, "Error New Relic New Application", true, CategoryInternalServerError)

	//	- Default -

	//	Controller
	CErrJsonDecode        = errW.NewDefinition(100000, "Error JSON Decode", true, CategoryInternalServerError)
	CErrPayloadIncomplete = errW.NewDefinition(100001, "Error Payload Incomplete. Payload %s", true, CategoryBadRequest)
	CErrHeaderIncomplete  = errW.NewDefinition(100002, "Error Header Incomplete", true, CategoryUnAuthorized)
	CErrJsonBind          = errW.NewDefinition(100003, "Error JSON Bind", true, CategoryInternalServerError)
	CErrFileUpload        = errW.NewDefinition(100004, "Error uploading file", true, CategoryInternalServerError)
	//	- Handler -

	// Usecase
	UErrInvalidItemCategory = errW.NewDefinition(300000, "Error Invalid Item Category", true, CategoryBadRequest)
	//	Service
	SErrDataExist       = errW.NewDefinition(200000, "Error Data Already Exist", false, CategoryBadRequest)
	SErrUnableToProceed = errW.NewDefinition(200001, "Error Unable To Proceed", false, CategoryBadRequest)

	SErrConfigApproverKeyNotFound = errW.NewDefinition(202000, "Error Approver Key Not Found. Approver Key: %s", false, CategoryBadRequest)

	SErrUserNotBranchManager     = errW.NewDefinition(205000, "Error User Not Branch Manager", false, CategoryBadRequest)
	SErrAuthInvalidCredentials   = errW.NewDefinition(202000, "Error Invalid Credentials %s", false, CategoryUnAuthorized)
	SErrAuthGenerateToken        = errW.NewDefinition(202001, "Error generating JWT token", true, CategoryInternalServerError)
	SErrBranchNotExist           = errW.NewDefinition(203000, "Error Branch Not Found. Branch ID: %s", true, CategoryBadRequest)
	SErrItemNotExist             = errW.NewDefinition(204000, "Error Item Not Found. Item ID: %s", true, CategoryBadRequest)
	SErrFailParseExcel           = errW.NewDefinition(206000, "Error parsing excel data", false, CategoryInternalServerError)
	SErrExcelMissingRequiredData = errW.NewDefinition(206001, "Error missing data from excel", false, CategoryBadRequest)
	SErrParsingExcelQuantity     = errW.NewDefinition(206002, "Error parsing quantity from excel to int", true, CategoryInternalServerError)

	//	Resource
	RErrMongoDBCollection      = errW.NewDefinition(400000, "Error MongoDB Collection", true, CategoryInternalServerError)
	RErrMongoDBQuery           = errW.NewDefinition(400001, "Error MongoDB Query", true, CategoryInternalServerError)
	RErrMongoDBReadDocument    = errW.NewDefinition(400002, "Error MongoDB Read Document", true, CategoryInternalServerError)
	RErrMongoDBCreateDocument  = errW.NewDefinition(400003, "Error MongoDB Create Document", true, CategoryInternalServerError)
	RErrMongoDBUpdateDocument  = errW.NewDefinition(400004, "Error MongoDB Update Document", true, CategoryInternalServerError)
	RErrMongoDBDeleteDocument  = errW.NewDefinition(400005, "Error MongoDB Delete Document", true, CategoryInternalServerError)
	RErrDecodeStringToObjectID = errW.NewDefinition(400006, "Error Unable to Decode String ID to Object ID", true, CategoryInternalServerError)
	RErrPostgresCreateDocument = errW.NewDefinition(400007, "Error PostgreSQL Create Document", true, CategoryInternalServerError)
	RErrPostgresReadDocument   = errW.NewDefinition(400008, "Error PostgreSQL Read Document", true, CategoryInternalServerError)
	RErrPostgresUpdateDocument = errW.NewDefinition(400009, "Error Postgres Update Document", true, CategoryInternalServerError)
	RErrPostgresDeleteDocument = errW.NewDefinition(400010, "Error Postgres Delete Document", true, CategoryInternalServerError)

	RErrDataNotFound = errW.NewDefinition(401000, "Error Data Not Found", true, CategoryNotFound)

	RErrJsonMarshal   = errW.NewDefinition(402000, "Error JSON Marshal", true, CategoryInternalServerError)
	RErrJsonUnmarshal = errW.NewDefinition(402001, "Error JSON Unmarshal", true, CategoryInternalServerError)
	RErrIoReadAll     = errW.NewDefinition(402002, "Error IO Read All", true, CategoryInternalServerError)

	// - Unit Test
	UErrUnitTest = errW.NewDefinition(999999, "Error Unit Test", true, CategoryInternalServerError)
)
