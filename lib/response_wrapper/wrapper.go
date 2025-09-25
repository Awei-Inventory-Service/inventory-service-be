package response_wrapper

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/inventory-service/lib/error_wrapper"
)

type errorArr [3]string
type Response struct {
	Header   Header      `json:"header,omitempty"`
	Body     interface{} `json:"body,omitempty"`
	Errors   *errorArr   `json:"errors,omitempty"`
	ErrorsV2 *ErrorV2    `json:"errors_v2,omitempty"`
}

type (
	Header struct {
		ProcessTime string `json:"process_time"`
		IsSuccess   bool   `json:"is_success"`
	}

	ErrorV2 struct {
		Message        string    `json:"message"`
		Error          string    `json:"error"`
		Code           int       `json:"code"`
		Classification [3]string `json:"classification"`
	}
)

var (
	defaultCategory   = error_wrapper.NewCategory(500, "Errors at New Errors")
	defaultDefinition = error_wrapper.NewDefinition(0, "an Errors", true, defaultCategory)
	defaultErr        = error_wrapper.New(defaultDefinition, "an Errors")
)

func New(w *gin.ResponseWriter, context context.Context, isSuccess bool, body interface{}, errW *error_wrapper.ErrorWrapper) {
	now := time.Now().UnixNano() / int64(time.Millisecond)

	startTime, err := strconv.Atoi(fmt.Sprint(context.Value("start_time")))
	requestId := fmt.Sprint(context.Value("request_id"))
	if err != nil {
		New(w, context, false, nil, defaultErr)
		return
	}
	start := int64(startTime)
	start /= int64(time.Millisecond)

	header := Header{
		ProcessTime: fmt.Sprintf(`%v ms`, now-start),
		IsSuccess:   isSuccess,
	}

	code, errors := BuildErrors(requestId, errW)
	writeHeader(*w, code)
	json.NewEncoder(*w).Encode(Response{
		Header: header,
		Body:   body,
		Errors: errors,
	})

	return
}

func BuildErrors(requestId string, errW *error_wrapper.ErrorWrapper) (int, *errorArr) {
	if errW == nil {
		return 200, nil
	}

	errors := &errorArr{
		fmt.Sprintf(`%s (%d) (%s)`, errW.Error(), errW.Code(), requestId),
		fmt.Sprintf(`%s (%d)`, errW.Error(), errW.Code()),
		fmt.Sprintf(`%s`, errW.ActualError()),
	}

	return errW.StatusCode(), errors
}

func writeHeader(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
}
