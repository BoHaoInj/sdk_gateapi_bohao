/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Batch order details
type BatchOrder struct {
	// User defined information. If not empty, must follow the rules below:  1. prefixed with `t-` 2. no longer than 28 bytes without `t-` prefix 3. can only include 0-9, A-Z, a-z, underscore(_), hyphen(-) or dot(.)
	Text string `json:"text,omitempty"`
	// Whether order succeeds
	Succeeded bool `json:"succeeded,omitempty"`
	// Error label, empty string if order succeeds
	Label string `json:"label,omitempty"`
	// Detailed error message, empty string if order succeeds
	Message string `json:"message,omitempty"`
	// Order ID
	Id string `json:"id,omitempty"`
	// Order creation time
	CreateTime string `json:"create_time,omitempty"`
	// Order last modification time
	UpdateTime string `json:"update_time,omitempty"`
	// Order status  - `open`: to be filled - `closed`: filled - `cancelled`: cancelled
	Status string `json:"status,omitempty"`
	// Currency pair
	CurrencyPair string `json:"currency_pair,omitempty"`
	// Order type. limit - limit order
	Type string `json:"type,omitempty"`
	// Account type. spot - use spot account; margin - use margin account
	Account string `json:"account,omitempty"`
	// Order side
	Side string `json:"side,omitempty"`
	// Trade amount
	Amount string `json:"amount,omitempty"`
	// Order price
	Price string `json:"price,omitempty"`
	// Time in force  - gtc: GoodTillCancelled - ioc: ImmediateOrCancelled, taker only - poc: PendingOrCancelled, makes a post-only order that always enjoys a maker fee
	TimeInForce string `json:"time_in_force,omitempty"`
	// Amount to display for the iceberg order. Null or 0 for normal orders. Set to -1 to hide the amount totally
	Iceberg string `json:"iceberg,omitempty"`
	// Used in margin trading(i.e. `account` is `margin`) to allow automatic loan of insufficient part if balance is not enough.
	AutoBorrow bool `json:"auto_borrow,omitempty"`
	// Amount left to fill
	Left string `json:"left,omitempty"`
	// Total filled in quote currency. Deprecated in favor of `filled_total`
	FillPrice string `json:"fill_price,omitempty"`
	// Total filled in quote currency
	FilledTotal string `json:"filled_total,omitempty"`
	// Fee deducted
	Fee string `json:"fee,omitempty"`
	// Fee currency unit
	FeeCurrency string `json:"fee_currency,omitempty"`
	// Point used to deduct fee
	PointFee string `json:"point_fee,omitempty"`
	// GT used to deduct fee
	GtFee string `json:"gt_fee,omitempty"`
	// Whether GT fee discount is used
	GtDiscount bool `json:"gt_discount,omitempty"`
	// Rebated fee
	RebatedFee string `json:"rebated_fee,omitempty"`
	// Rebated fee currency unit
	RebatedFeeCurrency string `json:"rebated_fee_currency,omitempty"`
}
