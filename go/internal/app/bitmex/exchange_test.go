package bitmex

// AUTOMATICALLY GENERATED, BUT NEEDS TO BE MODIFIED:
import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/ccxt/ccxt/go/pkg/ccxt"
)

func TestExchangeInit(t *testing.T) {
	x, err := Init(ccxt.ExchangeConfig{})
	if err != nil {
		t.Fatal(err)
	}
	if x.Info.ID != "bitmex" {
		t.Fatal("Did not load ID bitmex correctly")
	}
}

func TestExchangeInitPrivate(t *testing.T) {
	keys, err := ioutil.ReadFile("../../../keys.json")
	if err != nil {
		t.Fatal(err)
	}
	var configKeys map[string]ccxt.ExchangeConfig
	err = json.Unmarshal(keys, &configKeys)
	if err != nil {
		t.Fatal(err)
	}
	config := configKeys["bitmex"]
	x, err := Init(config)
	if err != nil {
		t.Fatal(err)
	}
	if x.Config.APIKey != config.APIKey {
		t.Fatal("Did not load bitmex APIKey correctly")
	}
}
