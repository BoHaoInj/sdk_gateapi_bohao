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
type FuturesPriceTriggeredOrder struct {
	Initial FuturesInitialOrder `json:"initial"`
	Trigger FuturesPriceTrigger `json:"trigger"`
	// Auto order ID
	Id int64 `json:"id,omitempty"`
	// User ID
	User int32 `json:"user,omitempty"`
	// Creation time
	CreateTime float64 `json:"create_time,omitempty"`
	// Finished time
	FinishTime float64 `json:"finish_time,omitempty"`
	// ID of the newly created order on condition triggered
	TradeId int64 `json:"trade_id,omitempty"`
	// Order status.
	Status string `json:"status,omitempty"`
	// How order is finished
	FinishAs string `json:"finish_as,omitempty"`
	// Extra messages of how order is finished
	Reason string `json:"reason,omitempty"`
}
