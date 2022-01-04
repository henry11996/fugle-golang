package fugle

import (
	"time"

	"github.com/shopspring/decimal"
)

type Response struct {
	APIVersion string
	StatusCode int
	Data       Data
}

type Data struct {
	Info   Info
	Chart  Chart
	Quote  Quote
	Meta   Meta
	Dealts []Dealt
}

type Info struct {
	Date          string
	Type          string
	Exchange      string
	Market        string
	SymbolID      string
	CountryCode   string
	Timezone      string
	LastUpdatedAt time.Time
}

type Chart struct {
	O []decimal.Decimal
	H []decimal.Decimal
	I []decimal.Decimal
	C []decimal.Decimal
	V int
	T int
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
	PriceAvg       QuotePrice
	PriceLimit     int
}

type QuoteTotal struct {
	At               time.Time
	Transaction      int
	TradeValue       decimal.Decimal
	TradeVolume      int
	TradeVolumeAtBid int
	TradeVolumeAtAsk int
	BidOrders        int
	AskOrders        int
	BidVolume        int
	AskVolume        int
	Serial           int
}

type QuoteTrial struct {
	At     time.Time
	Bid    decimal.Decimal
	Ask    decimal.Decimal
	Price  decimal.Decimal
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
	At   time.Time
	Bids []BestPrice
	Asks []BestPrice
}

type BestPrice struct {
	Price  decimal.Decimal
	Volume int
}

type QuotePrice struct {
	At    time.Time
	Price decimal.Decimal
}

type Meta struct {
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
