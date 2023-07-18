package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfItGetsAnErroIfIDIsBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "id is required")
}

func TestIfItGetsAnErroIfPriceIsLessThanZero(t *testing.T) {
	order := Order{ID: "1", Price: -10}
	assert.Error(t, order.Validate(), "price must be greater than zero")
}

func TestIfItGetsAnErroIfTaxIsLessThanZero(t *testing.T) {
	order := Order{ID: "1", Price: 10, Tax: -1}
	assert.Error(t, order.Validate(), "invalid tax value")
}

func TestFinalPrice(t *testing.T) {
	order := Order{ID: "1", Price: 10, Tax: 1}
	assert.NoError(t, order.Validate())
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 11.0, order.FinalPrice)
}
