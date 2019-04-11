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
	Info   ccxt.ExchangeInfo
	Config ccxt.ExchangeConfig
}

// Init Exchange
func Init(conf ccxt.ExchangeConfig) (*Exchange, error) {
	var info ccxt.ExchangeInfo
	configFile := "bitmex.json"
	f, err := os.Open(configFile)
	defer f.Close()
	if err != nil {
		return nil, fmt.Errorf("Config %s missing or errored: %v", configFile, err)
	}
	json.NewDecoder(f).Decode(&info)
	timeout := 10 * time.Second
	if conf.Timeout > 0 {
		timeout = conf.Timeout
	}
	client := &http.Client{Timeout: timeout}
	exchange := Exchange{
		Config: conf,
		Client: client,
		Info:   info,
	}
	return &exchange, nil
}
