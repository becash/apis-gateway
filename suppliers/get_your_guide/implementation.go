package get_your_guide

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ ServerInterface = (*Server)(nil)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (Server) PostBook(c *gin.Context) {
	print()
}

// Booking Cancellation
// (POST /1/cancel-booking/)
func (Server) PostCancelBooking(c *gin.Context) {
	print()
}

// Reservation Cancellation
// (POST /1/cancel-reservation/)
func (Server) PostCancelReservation(c *gin.Context) {
	print()
}

// Availability Query
// (GET /1/get-availabilities/)
func (Server) GetGetAvailabilities(c *gin.Context, params GetGetAvailabilitiesParams) {
	print()
}

// Notification
// (POST /1/notify/)
func (Server) PostNotify(c *gin.Context) {
	resp := GetGetAvailabilitiesParams{
		ProductId: "piiiiong",
	}

	c.Writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(c.Writer).Encode(resp)
}

// Product Details
// (GET /1/products/{productId})
func (Server) GetProductDetails(c *gin.Context, productId string) {
	print()
}

// Fetch Addons
// (GET /1/products/{productId}/addons/)
func (Server) FetchAddons(c *gin.Context, productId string) {
	print()
}

// Pricing Categories
// (GET /1/products/{productId}/pricing-categories/)
func (Server) GetPricingCategories(c *gin.Context, productId string) {
	print()
}

// Reservation
// (POST /1/reserve/)
func (Server) PostReserve(c *gin.Context) {
	print()
}

// Products List
// (GET /1/suppliers/{supplierId}/products/)
func (Server) GetSupplierProducts(c *gin.Context, supplierId string) {
	print()
}
