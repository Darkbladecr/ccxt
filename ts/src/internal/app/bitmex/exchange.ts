package bitmex

// AUTOMATICALLY GENERATED, BUT NEEDS TO BE MODIFIED:
import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/ccxt/ccxt/go/pkg/ccxt"
)

// Exchange struct
type Exchange struct {
	Client         *http.Client
	Info           ccxt.ExchangeInfo
	Config         ccxt.ExchangeConfig
	Markets        map[string]ccxt.Market
	MarketsByID    map[string]ccxt.Market
	IDs            []string
	Symbols        []string
	Currencies     map[string]ccxt.Currency
	CurrenciesByID map[string]ccxt.Currency
}

// Init Exchange
func Init(conf ccxt.ExchangeConfig) (*Exchange, error) {
    var info ccxt.ExchangeInfo
	configFile := "/Users/stefan/Github/ccxt/templates/internal/app/bitmex/bitmex.json"
	f, err := os.Open(configFile)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	json.NewDecoder(f).Decode(&info)
	timeout := 10 * time.Second
	if conf.Timeout > 0 {
		timeout = conf.Timeout * time.Millisecond
	}
	client := &http.Client{Timeout: timeout}
	exchange := Exchange{
		Config: conf,
		Client: client,
		Info:   info,
	}
	return &exchange, nil
}
// FetchCurrencies returns ccxt.Currency
func (x *Exchange) FetchCurrencies() (map[string]ccxt.Currency, error) {
	currencies := x.Currencies
	return currencies, nil
}

// GetInfo returns ccxt.ExchangeInfo
func (x *Exchange) GetInfo() ccxt.ExchangeInfo {
	return x.Info
}

// GetMarkets returns []ccxt.Market
func (x *Exchange) GetMarkets() map[string]ccxt.Market {
	return x.Markets
}

// SetMarkets sets to Markets key
func (x *Exchange) SetMarkets(markets map[string]ccxt.Market) {
	x.Markets = markets
}

// GetCurrencies returns []ccxt.Currency
func (x *Exchange) GetCurrencies() map[string]ccxt.Currency {
	return x.Currencies
}

// SetCurrencies sets to Currencies key
func (x *Exchange) SetCurrencies(currencies map[string]ccxt.Currency) {
	x.Currencies = currencies
}

// SetSymbols sets to Symbols key
func (x *Exchange) SetSymbols(slice []string) {
	x.Symbols = slice
}

// SetIDs sets to Symbols key
func (x *Exchange) SetIDs(slice []string) {
	x.IDs = slice
}

// GetMarketsByID returns map[string]ccxt.Market
func (x *Exchange) GetMarketsByID() map[string]ccxt.Market {
	return x.MarketsByID
}

// SetMarketsByID sets to MarketsByID key
func (x *Exchange) SetMarketsByID(markets map[string]ccxt.Market) {
	x.MarketsByID = markets
}

// GetCurrenciesByID returns map[string]ccxt.Currency
func (x *Exchange) GetCurrenciesByID() map[string]ccxt.Currency {
	return x.CurrenciesByID
}

// SetCurrenciesByID sets to CurrenciesByID key
func (x *Exchange) SetCurrenciesByID(currencies map[string]ccxt.Currency) {
	x.CurrenciesByID = currencies
}

// FetchMarkets and insert into the Exchange
func (x *Exchange) FetchMarkets(params *url.Values) ([]ccxt.Market, error) {
	// TODO: Complete FetchMarkets
	return nil, nil
}

// FetchBalance from exchange
func (x *Exchange) FetchBalance(params *url.Values) (ccxt.Account, error) {
	// TODO: Complete FetchBalance
	return ccxt.Account{}, nil
}