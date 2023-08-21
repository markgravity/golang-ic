package jsonhelpers

import (
	"net/http"

	"github.com/markgravity/golang-ic/helpers"
	"github.com/markgravity/golang-ic/lib/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/jsonapi"
	"github.com/samber/lo"
)

const (
	codeBadRequestError          string = "bad_request_error"
	codeForbiddenError           string = "forbidden_error"
	codeInternalServerError      string = "internal_server_error"
	codeNotFoundError            string = "not_found_error"
	codeUnauthorizedError        string = "unauthorized_error"
	codeUnprocessableEntityError string = "unprocessable_entity_error"
	codeServiceUnavailable       string = "service_unavailable_error"
	codeValidationError          string = "validation_error"
)

func RenderError(ctx *gin.Context, statusCode int, errMessage string, errorCode string) {
	errorObjects := []*jsonapi.ErrorObject{{
		Title:  http.StatusText(statusCode),
		Detail: errMessage,
		Code:   errorCode,
	}}
	payload := buildErrorPayload(errorObjects)

	renderJSONError(ctx, statusCode, payload)
}

func RenderErrorWithDefaultCode(ctx *gin.Context, statusCode int, err error) {
	validationErrors, ok := err.(validator.ValidationErrors)
	if ok {
		statusCode = http.StatusUnprocessableEntity
		translator := validators.GetTranslator()
		var errorObjects []*jsonapi.ErrorObject

		for _, f := range validationErrors {
			errorObjects = append(errorObjects, &jsonapi.ErrorObject{
				Title:  http.StatusText(statusCode),
				Detail: f.Translate(translator),
				Code:   codeValidationError,
				Meta: &map[string]interface{}{
					"parameter": f.Namespace(),
				},
			})
		}

		payload := buildErrorPayload(errorObjects)
		renderJSONError(ctx, statusCode, payload)

		return
	}

	RenderError(ctx, statusCode, err.Error(), defaultErrorResponseCode(statusCode))
}

func RenderErrorWithMeta(ctx *gin.Context, statusCode int, errMessage string, errorCode string, meta *map[string]interface{}) {
	errorObjects := []*jsonapi.ErrorObject{{
		Title:  http.StatusText(statusCode),
		Detail: errMessage,
		Code:   errorCode,
		Meta:   snakeCaseMetaKeys(meta),
	}}
	payload := buildErrorPayload(errorObjects)

	renderJSONError(ctx, statusCode, payload)
}

func RenderJSONWithMeta(ctx *gin.Context, code int, data interface{}, meta *jsonapi.Meta) {
	if data == nil {
		renderJSON(ctx, code, make(map[string]string))
		return
	}

	payload, err := jsonapi.Marshal(data)
	if err != nil {
		RenderErrorWithDefaultCode(ctx, http.StatusInternalServerError, err)
		return
	}

	manyPayload, valid := payload.(*jsonapi.ManyPayload)
	if valid && len(*meta) > 0 {
		manyPayload.Meta = meta
		renderJSON(ctx, code, manyPayload)
	} else {
		renderJSON(ctx, code, payload)
	}
}

func RenderJSON(ctx *gin.Context, code int, data interface{}) {
	RenderJSONWithMeta(ctx, code, data, &jsonapi.Meta{})
}

func buildErrorPayload(errorObjects []*jsonapi.ErrorObject) (payload *jsonapi.ErrorsPayload) {
	payload = &jsonapi.ErrorsPayload{
		Errors: errorObjects,
	}
	return payload
}

func renderJSON(ctx *gin.Context, code int, payload interface{}) {
	ctx.JSON(code, payload)
}

func renderJSONError(ctx *gin.Context, code int, payload interface{}) {
	ctx.AbortWithStatusJSON(code, payload)
}

func defaultErrorResponseCode(status int) string {
	switch status {
	case http.StatusBadRequest:
		return codeBadRequestError
	case http.StatusForbidden:
		return codeForbiddenError
	case http.StatusNotFound:
		return codeNotFoundError
	case http.StatusUnauthorized:
		return codeUnauthorizedError
	case http.StatusUnprocessableEntity:
		return codeUnprocessableEntityError
	case http.StatusInternalServerError:
		return codeInternalServerError
	case http.StatusServiceUnavailable:
		return codeServiceUnavailable
	default:
		return http.StatusText(status)
	}
}

func snakeCaseMetaKeys(meta *map[string]interface{}) *map[string]interface{} {
	metaData := lo.MapKeys(*meta, func(_ interface{}, v string) string {
		return helpers.ToSnakeCase(v)
	})

	return &metaData
}
