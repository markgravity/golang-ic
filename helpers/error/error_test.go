package errorhelpers_test

import (
	"errors"
	"net/http"

	errorhelpers "github.com/markgravity/golang-ic/helpers/error"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type customTimeoutError struct{}

func (e customTimeoutError) Timeout() bool {
	return true
}

func (e customTimeoutError) Error() string {
	return "timeout"
}

var _ = Describe("Error Helpers", func() {
	Describe(".IsNotFound", func() {
		Context("given error contains Not Found text", func() {
			It("returns true", func() {
				err := errors.New("Not Found")

				result := errorhelpers.IsNotFound(err)

				Expect(result).To(BeTrue())
			})
		})

		Context("given error does NOT contain Not Found text", func() {
			It("returns false", func() {
				err := errors.New("Found")

				result := errorhelpers.IsNotFound(err)

				Expect(result).To(BeFalse())
			})
		})

		Context("given error is nil", func() {
			It("returns false", func() {
				result := errorhelpers.IsNotFound(nil)

				Expect(result).To(BeFalse())
			})
		})
	})

	Describe(".IsTimeout", func() {
		Context("given a timeout error", func() {
			It("returns true", func() {
				err := customTimeoutError{}

				result := errorhelpers.IsTimeout(err)

				Expect(result).To(BeTrue())
			})
		})

		Context("given a normal error", func() {
			It("returns false", func() {
				err := errors.New("Found")

				result := errorhelpers.IsTimeout(err)

				Expect(result).To(BeFalse())
			})
		})

		Context("given error is nil", func() {
			It("returns false", func() {
				result := errorhelpers.IsTimeout(nil)

				Expect(result).To(BeFalse())
			})
		})
	})

	Describe(".GetErrorCode", func() {
		Context("given a timeout error", func() {
			It("returns timeout error code", func() {
				err := customTimeoutError{}

				result := errorhelpers.GetErrorCode(err)

				Expect(result).To(Equal(errorhelpers.TimeoutCode))
			})
		})

		Context("given a not found error", func() {
			It("returns not found error code", func() {
				err := errors.New("Not Found")

				result := errorhelpers.GetErrorCode(err)

				Expect(result).To(Equal(errorhelpers.NotFoundCode))
			})
		})

		Context("given a normal error", func() {
			It("returns internal server error code", func() {
				err := errors.New("Found")

				result := errorhelpers.GetErrorCode(err)

				Expect(result).To(Equal(errorhelpers.InternalServerErrorCode))
			})
		})

		Context("given error is nil", func() {
			It("returns internal server error code", func() {
				result := errorhelpers.GetErrorCode(nil)

				Expect(result).To(Equal(errorhelpers.InternalServerErrorCode))
			})
		})
	})

	Describe(".GetErrorStatusCode", func() {
		Context("given a timeout error", func() {
			It("returns timeout error status code", func() {
				err := customTimeoutError{}

				result := errorhelpers.GetErrorStatusCode(err)

				Expect(result).To(Equal(http.StatusServiceUnavailable))
			})
		})

		Context("given a not found error", func() {
			It("returns not found error status code", func() {
				err := errors.New("Not Found")

				result := errorhelpers.GetErrorStatusCode(err)

				Expect(result).To(Equal(http.StatusNotFound))
			})
		})

		Context("given a normal error", func() {
			It("returns internal server error status code", func() {
				err := errors.New("Found")

				result := errorhelpers.GetErrorStatusCode(err)

				Expect(result).To(Equal(http.StatusInternalServerError))
			})
		})

		Context("given error is nil", func() {
			It("returns internal server error status code", func() {
				result := errorhelpers.GetErrorStatusCode(nil)

				Expect(result).To(Equal(http.StatusInternalServerError))
			})
		})
	})
})