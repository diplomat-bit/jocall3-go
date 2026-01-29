// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/param"
)

type Address struct {
	City    string      `json:"city,required"`
	Country string      `json:"country,required"`
	Street  string      `json:"street,required"`
	State   string      `json:"state"`
	Zip     string      `json:"zip"`
	JSON    addressJSON `json:"-"`
}

// addressJSON contains the JSON metadata for the struct [Address]
type addressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Street      apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Address) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressJSON) RawJSON() string {
	return r.raw
}

type AddressParam struct {
	City    param.Field[string] `json:"city,required"`
	Country param.Field[string] `json:"country,required"`
	Street  param.Field[string] `json:"street,required"`
	State   param.Field[string] `json:"state"`
	Zip     param.Field[string] `json:"zip"`
}

func (r AddressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Transaction struct {
	ID          string          `json:"id,required"`
	Amount      float64         `json:"amount,required"`
	Currency    string          `json:"currency,required"`
	Date        time.Time       `json:"date,required" format:"date"`
	Description string          `json:"description,required"`
	Category    string          `json:"category"`
	Notes       string          `json:"notes"`
	JSON        transactionJSON `json:"-"`
}

// transactionJSON contains the JSON metadata for the struct [Transaction]
type transactionJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	Currency    apijson.Field
	Date        apijson.Field
	Description apijson.Field
	Category    apijson.Field
	Notes       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Transaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionJSON) RawJSON() string {
	return r.raw
}
