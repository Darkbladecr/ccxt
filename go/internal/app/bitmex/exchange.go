package bitmex

// AUTOMATICALLY GENERATED, BUT NEEDS TO BE MODIFIED:
import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ccxt/ccxt/go/pkg/ccxt"
)

// Exchange struct
type Exchange struct {
	Client *http.Client
    Info ccxt.ExchangeInfo
}

// Init Exchange
func Init() (*Exchange, error) {
    var info ccxt.ExchangeInfo
	configFile := "bitmex.json"
	f, err := os.Open(configFile)
	defer f.Close()
	if err != nil {
		return nil, fmt.Errorf("Config %s missing or errored: %v", configFile, err)
	}
	json.NewDecoder(f).Decode(&info)
	client := &http.Client{Timeout: time.Second * 10}	
	exchange := Exchange{
		Client: client,
		Info: info,
	}
	return &exchange, nil
}