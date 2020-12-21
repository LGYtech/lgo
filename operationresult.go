package lgo

import (
	"encoding/json"
	"net/http"
)

const (
	// lgoOrSuccess OperationResult has 'Success' result
	lgoOrSuccess = iota
	// lgoOrLogicError OperationResult has 'Logic Error' result
	lgoOrLogicError = iota
	// lgoOrFailure OperationResult has 'Failure' result
	lgoOrFailure = iota
	// lgoOrAuthError OperationResult has 'Authentication Failure' result
	lgoOrAuthError = iota
	// lgoOrAutoError OperationResult has 'Authorization Failure' result
	lgoOrAutoError = iota
)

// OperationResult Defines result of the operation
type OperationResult struct {
	Result       uint8       `json:"r" mapstructure:"r"`
	ReturnObject interface{} `json:"ro,omitempty" mapstructure:"ro"`
	ErrorMessage string      `json:"em,omitempty" mapstructure:"em"`
	ErrorCode    uint8       `json:"ec" mapstructure:"ec"`
}

// NewSuccess Returns a new lgoOrSuccess result with given return object
func NewSuccess(returnObject interface{}) *OperationResult {
	return &OperationResult{Result: lgoOrSuccess, ReturnObject: returnObject, ErrorMessage: "", ErrorCode: 0}
}

// NewFailure Returns a new lgoOrFailure result
func NewFailure() *OperationResult {
	return &OperationResult{Result: lgoOrFailure, ReturnObject: nil, ErrorMessage: "", ErrorCode: 0}
}

// NewFailureWithReturnObject Returns a new lgoOrFailure result with given return object
func NewFailureWithReturnObject(returnObject interface{}) *OperationResult {
	return &OperationResult{Result: lgoOrFailure, ReturnObject: returnObject, ErrorMessage: "", ErrorCode: 0}
}

// NewLogicError Returns a new lgoOrLogicError result with given error message and return object
func NewLogicError(errorMessage string, returnObject interface{}) *OperationResult {
	return &OperationResult{Result: lgoOrLogicError, ReturnObject: returnObject, ErrorMessage: errorMessage, ErrorCode: 0}
}

// NewAuthError Returns a new lgoOrAuthError result
func NewAuthError() *OperationResult {
	return &OperationResult{Result: lgoOrAuthError, ReturnObject: nil, ErrorMessage: "Authentication error", ErrorCode: 0}
}

// NewAutoError Returns a new lgoOrAutoError result
func NewAutoError() *OperationResult {
	return &OperationResult{Result: lgoOrAutoError, ReturnObject: nil, ErrorMessage: "Authorization error", ErrorCode: 0}
}

// WriteResult Writes operation result values to http.ResponseWriter
func (operationResult OperationResult) WriteResult(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(operationResult)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Write(js)
	}
}

// IsSuccess Returns if operation has lgoOrSuccess result
func (operationResult OperationResult) IsSuccess() bool {
	return operationResult.Result == lgoOrSuccess
}
