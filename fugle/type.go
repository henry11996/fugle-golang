package fugle

import (
	"time"

	"github.com/shopspring/decimal"
)

type Response struct {
	APIVersion string
	StatusCode int
	Data       struct {
		Info   Info
		Chart  map[time.Time]Deal
		Quote  Quote
		Meta   Meta
		Dealts []Dealt
	}
}

type Info struct {
	Date          string
	Mode          string
	SymbolID      string
	CountryCode   string
	Timezone      string
	LastUpdatedAt time.Time
}

type Deal struct {
	Open   decimal.Decimal
	High   decimal.Decimal
	Low    decimal.Decimal
	Close  decimal.Decimal
	Volume decimal.Decimal
	Unit   decimal.Decimal
}

type Quote struct {
	IsCurbing      bool
	IsCurbingRise  bool
	IsCurbingFall  bool
	IsTrial        bool
	IsOpenDelayed  bool
	IsCloseDelayed bool
	IsHalting      bool
	IsClosed       bool
	Change         decimal.Decimal
	ChangePercent  float64
	Amplitude      float64
	Total          QuoteTotal
	Trial          QuoteTrial
	Trade          QuoteTrade
	Order          QuoteOrder
	PriceHigh      QuotePrice
	PriceLow       QuotePrice
	PriceOpen      QuotePrice
}

type QuoteTotal struct {
	At     time.Time
	Order  float32
	Price  float32
	Unit   int
	Volume int
}

type QuoteTrial struct {
	At     time.Time
	Price  decimal.Decimal
	Unit   int
	Volume int
}

type QuoteTrade struct {
	At     time.Time
	Price  decimal.Decimal
	Unit   int
	Volume int
	Serial int
}

type QuoteOrder struct {
	At       time.Time
	Bestbids []BestPrice
	Bestasks []BestPrice
}

type BestPrice struct {
	Price  decimal.Decimal
	Unit   int
	Volume int
}

type QuotePrice struct {
	At    time.Time
	Price decimal.Decimal
}

type Meta struct {
	IsIndex        bool
	NameZhTw       string
	IndustryZhTw   string
	PriceReference decimal.Decimal
	PriceHighLimit decimal.Decimal
	PriceLowLimit  decimal.Decimal
	CanDayBuySell  bool
	CanDaySellBuy  bool
	CanShortMargin bool
	CanShortLend   bool
	VolumePerUnit  int
	Currency       string
	IsTerminated   bool
	IsSuspended    bool
	IsWarrant      bool
	TypeZhTw       string
	Abnormal       string
}

type Dealt struct {
	At     time.Time
	Price  decimal.Decimal
	Unit   int
	Volume int
	Serial int
}
