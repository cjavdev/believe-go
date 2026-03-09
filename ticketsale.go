// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package believe

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cjavdev/believe-go/internal/apijson"
	"github.com/cjavdev/believe-go/internal/apiquery"
	"github.com/cjavdev/believe-go/internal/requestconfig"
	"github.com/cjavdev/believe-go/option"
	"github.com/cjavdev/believe-go/packages/pagination"
	"github.com/cjavdev/believe-go/packages/param"
	"github.com/cjavdev/believe-go/packages/respjson"
)

// Ticket sales with 300 records for practicing pagination, filtering, and
// financial data
//
// TicketSaleService contains methods and other services that help with interacting
// with the believe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTicketSaleService] method instead.
type TicketSaleService struct {
	Options []option.RequestOption
}

// NewTicketSaleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTicketSaleService(opts ...option.RequestOption) (r TicketSaleService) {
	r = TicketSaleService{}
	r.Options = opts
	return
}

// Record a new ticket sale.
func (r *TicketSaleService) New(ctx context.Context, body TicketSaleNewParams, opts ...option.RequestOption) (res *TicketSale, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ticket-sales"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve detailed information about a specific ticket sale.
func (r *TicketSaleService) Get(ctx context.Context, ticketSaleID string, opts ...option.RequestOption) (res *TicketSale, err error) {
	opts = slices.Concat(r.Options, opts)
	if ticketSaleID == "" {
		err = errors.New("missing required ticket_sale_id parameter")
		return
	}
	path := fmt.Sprintf("ticket-sales/%s", url.PathEscape(ticketSaleID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update specific fields of an existing ticket sale.
func (r *TicketSaleService) Update(ctx context.Context, ticketSaleID string, body TicketSaleUpdateParams, opts ...option.RequestOption) (res *TicketSale, err error) {
	opts = slices.Concat(r.Options, opts)
	if ticketSaleID == "" {
		err = errors.New("missing required ticket_sale_id parameter")
		return
	}
	path := fmt.Sprintf("ticket-sales/%s", url.PathEscape(ticketSaleID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a paginated list of all ticket sales with optional filtering. With 300
// records, this endpoint is ideal for practicing pagination.
func (r *TicketSaleService) List(ctx context.Context, query TicketSaleListParams, opts ...option.RequestOption) (res *pagination.SkipLimitPage[TicketSale], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ticket-sales"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get a paginated list of all ticket sales with optional filtering. With 300
// records, this endpoint is ideal for practicing pagination.
func (r *TicketSaleService) ListAutoPaging(ctx context.Context, query TicketSaleListParams, opts ...option.RequestOption) *pagination.SkipLimitPageAutoPager[TicketSale] {
	return pagination.NewSkipLimitPageAutoPager(r.List(ctx, query, opts...))
}

// Remove a ticket sale from the database.
func (r *TicketSaleService) Delete(ctx context.Context, ticketSaleID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if ticketSaleID == "" {
		err = errors.New("missing required ticket_sale_id parameter")
		return
	}
	path := fmt.Sprintf("ticket-sales/%s", url.PathEscape(ticketSaleID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// How the ticket was purchased.
type PurchaseMethod string

const (
	PurchaseMethodOnline    PurchaseMethod = "online"
	PurchaseMethodBoxOffice PurchaseMethod = "box_office"
	PurchaseMethodWillCall  PurchaseMethod = "will_call"
	PurchaseMethodPhone     PurchaseMethod = "phone"
)

// Full ticket sale model with ID.
type TicketSale struct {
	// Unique identifier
	ID string `json:"id" api:"required"`
	// Name of the ticket buyer
	BuyerName string `json:"buyer_name" api:"required"`
	// Currency code (GBP, USD, or EUR)
	Currency string `json:"currency" api:"required"`
	// Discount amount applied from coupon
	Discount string `json:"discount" api:"required"`
	// ID of the match
	MatchID string `json:"match_id" api:"required"`
	// How the ticket was purchased
	//
	// Any of "online", "box_office", "will_call", "phone".
	PurchaseMethod PurchaseMethod `json:"purchase_method" api:"required"`
	// Number of tickets purchased
	Quantity int64 `json:"quantity" api:"required"`
	// Subtotal before discount and tax (unit_price \* quantity)
	Subtotal string `json:"subtotal" api:"required"`
	// Tax amount (20% UK VAT on discounted subtotal)
	Tax string `json:"tax" api:"required"`
	// Final total (subtotal - discount + tax)
	Total string `json:"total" api:"required"`
	// Price per ticket (decimal string)
	UnitPrice string `json:"unit_price" api:"required"`
	// Email of the ticket buyer
	BuyerEmail string `json:"buyer_email" api:"nullable" format:"email"`
	// Coupon code applied, if any
	CouponCode string `json:"coupon_code" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		BuyerName      respjson.Field
		Currency       respjson.Field
		Discount       respjson.Field
		MatchID        respjson.Field
		PurchaseMethod respjson.Field
		Quantity       respjson.Field
		Subtotal       respjson.Field
		Tax            respjson.Field
		Total          respjson.Field
		UnitPrice      respjson.Field
		BuyerEmail     respjson.Field
		CouponCode     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TicketSale) RawJSON() string { return r.JSON.raw }
func (r *TicketSale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TicketSaleNewParams struct {
	// Name of the ticket buyer
	BuyerName string `json:"buyer_name" api:"required"`
	// Currency code (GBP, USD, or EUR)
	Currency string `json:"currency" api:"required"`
	// Discount amount applied from coupon
	Discount string `json:"discount" api:"required"`
	// ID of the match
	MatchID string `json:"match_id" api:"required"`
	// How the ticket was purchased
	//
	// Any of "online", "box_office", "will_call", "phone".
	PurchaseMethod PurchaseMethod `json:"purchase_method,omitzero" api:"required"`
	// Number of tickets purchased
	Quantity int64 `json:"quantity" api:"required"`
	// Subtotal before discount and tax (unit_price \* quantity)
	Subtotal string `json:"subtotal" api:"required"`
	// Tax amount (20% UK VAT on discounted subtotal)
	Tax string `json:"tax" api:"required"`
	// Final total (subtotal - discount + tax)
	Total string `json:"total" api:"required"`
	// Price per ticket (decimal string)
	UnitPrice string `json:"unit_price" api:"required"`
	// Email of the ticket buyer
	BuyerEmail param.Opt[string] `json:"buyer_email,omitzero" format:"email"`
	// Coupon code applied, if any
	CouponCode param.Opt[string] `json:"coupon_code,omitzero"`
	paramObj
}

func (r TicketSaleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TicketSaleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TicketSaleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TicketSaleUpdateParams struct {
	BuyerEmail param.Opt[string] `json:"buyer_email,omitzero" format:"email"`
	BuyerName  param.Opt[string] `json:"buyer_name,omitzero"`
	CouponCode param.Opt[string] `json:"coupon_code,omitzero"`
	Currency   param.Opt[string] `json:"currency,omitzero"`
	Discount   param.Opt[string] `json:"discount,omitzero"`
	MatchID    param.Opt[string] `json:"match_id,omitzero"`
	Quantity   param.Opt[int64]  `json:"quantity,omitzero"`
	Subtotal   param.Opt[string] `json:"subtotal,omitzero"`
	Tax        param.Opt[string] `json:"tax,omitzero"`
	Total      param.Opt[string] `json:"total,omitzero"`
	UnitPrice  param.Opt[string] `json:"unit_price,omitzero"`
	// How the ticket was purchased.
	//
	// Any of "online", "box_office", "will_call", "phone".
	PurchaseMethod PurchaseMethod `json:"purchase_method,omitzero"`
	paramObj
}

func (r TicketSaleUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TicketSaleUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TicketSaleUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TicketSaleListParams struct {
	// Filter by coupon code (use 'none' for sales without coupons)
	CouponCode param.Opt[string] `query:"coupon_code,omitzero" json:"-"`
	// Filter by currency (GBP, USD, EUR)
	Currency param.Opt[string] `query:"currency,omitzero" json:"-"`
	// Filter by match ID
	MatchID param.Opt[string] `query:"match_id,omitzero" json:"-"`
	// Maximum number of items to return (max: 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip (offset)
	Skip param.Opt[int64] `query:"skip,omitzero" json:"-"`
	// Filter by purchase method
	//
	// Any of "online", "box_office", "will_call", "phone".
	PurchaseMethod PurchaseMethod `query:"purchase_method,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TicketSaleListParams]'s query parameters as `url.Values`.
func (r TicketSaleListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
