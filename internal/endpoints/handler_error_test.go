package endpoints

import (
	internalerrors "campaing/internal/internalErrors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HandlerError_Internal_Error(t *testing.T) {
	assert := assert.New(t)

	endppoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, internalerrors.ErrInternal
	}

	handlerFunc := HandlerError(endppoint)

	req, _ := http.NewRequest("GET", "/", nil)

	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), internalerrors.ErrInternal.Error())

}
