/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Futures order details
type FuturesOrder struct {
	// Futures order ID
	Id int64 `json:"id,omitempty"`
	// User ID
	User int32 `json:"user,omitempty"`
	// Creation time of order
	CreateTime float64 `json:"create_time,omitempty"`
	// Order finished time. Not returned if order is open
	FinishTime float64 `json:"finish_time,omitempty"`
	// How the order was finished.  - filled: all filled - cancelled: manually cancelled - liquidated: cancelled because of liquidation - ioc: time in force is `IOC`, finish immediately - auto_deleveraged: finished by ADL - reduce_only: cancelled because of increasing position while `reduce-only` set- position_closed: cancelled because of position close
	FinishAs string `json:"finish_as,omitempty"`
	// Order status  - `open`: waiting to be traded - `finished`: finished
	Status string `json:"status,omitempty"`
	// Futures contract
	Contract string `json:"contract"`
	// Order size. Specify positive number to make a bid, and negative number to ask
	Size int64 `json:"size"`
	// Display size for iceberg order. 0 for non-iceberg. Note that you will have to pay the taker fee for the hidden size
	Iceberg int64 `json:"iceberg,omitempty"`
	// Order price. 0 for market order with `tif` set as `ioc`
	Price string `json:"price,omitempty"`
	// Set as `true` to close the position, with `size` set to 0
	Close bool `json:"close,omitempty"`
	// Is the order to close position
	IsClose bool `json:"is_close,omitempty"`
	// Set as `true` to be reduce-only order
	ReduceOnly bool `json:"reduce_only,omitempty"`
	// Is the order reduce-only
	IsReduceOnly bool `json:"is_reduce_only,omitempty"`
	// Is the order for liquidation
	IsLiq bool `json:"is_liq,omitempty"`
	// Time in force  - gtc: GoodTillCancelled - ioc: ImmediateOrCancelled, taker only - poc: PendingOrCancelled, reduce-only
	Tif string `json:"tif,omitempty"`
	// Size left to be traded
	Left int64 `json:"left,omitempty"`
	// Fill price of the order
	FillPrice string `json:"fill_price,omitempty"`
	// User defined information. If not empty, must follow the rules below:  1. prefixed with `t-` 2. no longer than 28 bytes without `t-` prefix 3. can only include 0-9, A-Z, a-z, underscore(_), hyphen(-) or dot(.) Besides user defined information, reserved contents are listed below, denoting how the order is created:  - web: from web - api: from API - app: from mobile phones - auto_deleveraging: from ADL - liquidation: from liquidation - insurance: from insurance
	Text string `json:"text,omitempty"`
	// Taker fee
	Tkfr string `json:"tkfr,omitempty"`
	// Maker fee
	Mkfr string `json:"mkfr,omitempty"`
	// Reference user ID
	Refu int32 `json:"refu,omitempty"`
	// Set side to close dual-mode position. `close_long` closes the long side; while `close_short` the short one. Note `size` also needs to be set to 0
	AutoSize string `json:"auto_size,omitempty"`
}
