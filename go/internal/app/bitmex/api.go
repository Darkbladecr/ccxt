package bitmex

// PLEASE DO NOT EDIT THIS FILE, IT IS GENERATED AND WILL BE OVERWRITTEN:
// https://github.com/ccxt/ccxt/blob/master/CONTRIBUTING.md#how-to-contribute-code
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/ccxt/ccxt/go/internal/app/bitmex/models"
)

func handleBodyErr(body []byte, err error) error {
	var exErr models.Error
	if err2 := json.Unmarshal(body, &exErr); err2 == nil {
		return fmt.Errorf("%s: %s", exErr.Error.Name, exErr.Error.Message)
	}
	var any interface{}
	_ = json.Unmarshal(body, &any)
	if exErr, ok := any.(models.Error); ok {
		return fmt.Errorf("%s: %s", exErr.Error.Name, exErr.Error.Message)
	}
	return fmt.Errorf("Error with decoding: %+v", any)
}

func (c *Exchange) apiRequest(method string, reqURL string, params *url.Values, body bytes.Buffer) ([]byte, error) {
	baseURL, err := url.Parse(reqURL)
	if err != nil {
		return nil, err
	}
	baseURL.RawQuery = params.Encode()
	var req *http.Request
	if len(body.Bytes()) > 0 {
		req, err = http.NewRequest(method, baseURL.String(), &body)
		if err != nil {
			return nil, err
		}
		req.Header.Add("Content-Type", http.DetectContentType(body.Bytes()))
	} else {
		req, err = http.NewRequest(method, baseURL.String(), nil)
		if err != nil {
			return nil, err
		}
	}
	req.Header.Add("Accept", "application/json")
	c.SignRequest(req, method, baseURL, body.Bytes())
	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 300 {
		err = fmt.Errorf("HTTP Response %d: %s", res.StatusCode, http.StatusText(res.StatusCode))
		return nil, handleBodyErr(resBody, err)
	}
	return resBody, nil
}

// PublicGetAnnouncement method for https://www.bitmex.com/api/v1/announcement
func (c *Exchange) PublicGetAnnouncement(params *url.Values) (data []models.Announcement, err error) {
	reqURL := "https://www.bitmex.com/api/v1/announcement"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PublicGetAnnouncementUrgent method for https://www.bitmex.com/api/v1/announcement/urgent
func (c *Exchange) PublicGetAnnouncementUrgent(params *url.Values) (data []models.Announcement, err error) {
	reqURL := "https://www.bitmex.com/api/v1/announcement/urgent"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PublicGetFunding method for https://www.bitmex.com/api/v1/funding
func (c *Exchange) PublicGetFunding(params *url.Values) (data []models.Funding, err error) {
	reqURL := "https://www.bitmex.com/api/v1/funding"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PublicGetInstrument method for https://www.bitmex.com/api/v1/instrument
func (c *Exchange) PublicGetInstrument(params *url.Values) (data []models.Instrument, err error) {
	reqURL := "https://www.bitmex.com/api/v1/instrument"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PublicGetInstrumentActive method for https://www.bitmex.com/api/v1/instrument/active
func (c *Exchange) PublicGetInstrumentActive(params *url.Values) (data []models.Instrument, err error) {
	reqURL := "https://www.bitmex.com/api/v1/instrument/active"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PublicGetInstrumentActiveAndIndices method for https://www.bitmex.com/api/v1/instrument/activeAndIndices
func (c *Exchange) PublicGetInstrumentActiveAndIndices(params *url.Values) (data []models.Instrument, err error) {
	reqURL := "https://www.bitmex.com/api/v1/instrument/activeAndIndices"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PublicGetInstrumentActiveIntervals method for https://www.bitmex.com/api/v1/instrument/activeIntervals
func (c *Exchange) PublicGetInstrumentActiveIntervals(params *url.Values) (data models.InstrumentInterval, err error) {
	reqURL := "https://www.bitmex.com/api/v1/instrument/activeIntervals"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PublicGetInstrumentCompositeIndex method for https://www.bitmex.com/api/v1/instrument/compositeIndex
func (c *Exchange) PublicGetInstrumentCompositeIndex(params *url.Values) (data []models.Instrument, err error) {
	reqURL := "https://www.bitmex.com/api/v1/instrument/compositeIndex"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PublicGetInstrumentIndices method for https://www.bitmex.com/api/v1/instrument/indices
func (c *Exchange) PublicGetInstrumentIndices(params *url.Values) (data []models.Instrument, err error) {
	reqURL := "https://www.bitmex.com/api/v1/instrument/indices"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PublicGetInsurance method for https://www.bitmex.com/api/v1/insurance
func (c *Exchange) PublicGetInsurance(params *url.Values) (data []models.Insurance, err error) {
	reqURL := "https://www.bitmex.com/api/v1/insurance"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PublicGetLeaderboard method for https://www.bitmex.com/api/v1/leaderboard
func (c *Exchange) PublicGetLeaderboard(params *url.Values) (data []models.Leaderboard, err error) {
	reqURL := "https://www.bitmex.com/api/v1/leaderboard"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PublicGetLeaderboardName method for https://www.bitmex.com/api/v1/leaderboard/name
func (c *Exchange) PublicGetLeaderboardName(params *url.Values) (data models.Leaderboard, err error) {
	reqURL := "https://www.bitmex.com/api/v1/leaderboard/name"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PublicGetLiquidation method for https://www.bitmex.com/api/v1/liquidation
func (c *Exchange) PublicGetLiquidation(params *url.Values) (data []models.Liquidation, err error) {
	reqURL := "https://www.bitmex.com/api/v1/liquidation"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PublicGetOrderBookL2 method for https://www.bitmex.com/api/v1/orderBook/L2
func (c *Exchange) PublicGetOrderBookL2(params *url.Values) (data []models.OrderBookL2, err error) {
	reqURL := "https://www.bitmex.com/api/v1/orderBook/L2"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PublicGetQuote method for https://www.bitmex.com/api/v1/quote
func (c *Exchange) PublicGetQuote(params *url.Values) (data []models.Quote, err error) {
	reqURL := "https://www.bitmex.com/api/v1/quote"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PublicGetQuoteBucketed method for https://www.bitmex.com/api/v1/quote/bucketed
func (c *Exchange) PublicGetQuoteBucketed(params *url.Values) (data []models.Quote, err error) {
	reqURL := "https://www.bitmex.com/api/v1/quote/bucketed"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PublicGetSchema method for https://www.bitmex.com/api/v1/schema
func (c *Exchange) PublicGetSchema(params *url.Values) (data interface{}, err error) {
	reqURL := "https://www.bitmex.com/api/v1/schema"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PublicGetSchemaWebsocketHelp method for https://www.bitmex.com/api/v1/schema/websocketHelp
func (c *Exchange) PublicGetSchemaWebsocketHelp(params *url.Values) (data interface{}, err error) {
	reqURL := "https://www.bitmex.com/api/v1/schema/websocketHelp"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PublicGetSettlement method for https://www.bitmex.com/api/v1/settlement
func (c *Exchange) PublicGetSettlement(params *url.Values) (data []models.Settlement, err error) {
	reqURL := "https://www.bitmex.com/api/v1/settlement"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PublicGetStats method for https://www.bitmex.com/api/v1/stats
func (c *Exchange) PublicGetStats(params *url.Values) (data []models.Stats, err error) {
	reqURL := "https://www.bitmex.com/api/v1/stats"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PublicGetStatsHistory method for https://www.bitmex.com/api/v1/stats/history
func (c *Exchange) PublicGetStatsHistory(params *url.Values) (data []models.StatsHistory, err error) {
	reqURL := "https://www.bitmex.com/api/v1/stats/history"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PublicGetTrade method for https://www.bitmex.com/api/v1/trade
func (c *Exchange) PublicGetTrade(params *url.Values) (data []models.Trade, err error) {
	reqURL := "https://www.bitmex.com/api/v1/trade"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PublicGetTradeBucketed method for https://www.bitmex.com/api/v1/trade/bucketed
func (c *Exchange) PublicGetTradeBucketed(params *url.Values) (data []models.TradeBin, err error) {
	reqURL := "https://www.bitmex.com/api/v1/trade/bucketed"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivateGetApiKey method for https://www.bitmex.com/api/v1/apiKey
func (c *Exchange) PrivateGetApiKey(params *url.Values) (data []models.APIKey, err error) {
	reqURL := "https://www.bitmex.com/api/v1/apiKey"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivateGetChat method for https://www.bitmex.com/api/v1/chat
func (c *Exchange) PrivateGetChat(params *url.Values) (data []models.Chat, err error) {
	reqURL := "https://www.bitmex.com/api/v1/chat"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivateGetChatChannels method for https://www.bitmex.com/api/v1/chat/channels
func (c *Exchange) PrivateGetChatChannels(params *url.Values) (data []models.ChatChannel, err error) {
	reqURL := "https://www.bitmex.com/api/v1/chat/channels"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivateGetChatConnected method for https://www.bitmex.com/api/v1/chat/connected
func (c *Exchange) PrivateGetChatConnected(params *url.Values) (data models.ConnectedUsers, err error) {
	reqURL := "https://www.bitmex.com/api/v1/chat/connected"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivateGetExecution method for https://www.bitmex.com/api/v1/execution
func (c *Exchange) PrivateGetExecution(params *url.Values) (data []models.Execution, err error) {
	reqURL := "https://www.bitmex.com/api/v1/execution"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivateGetExecutionTradeHistory method for https://www.bitmex.com/api/v1/execution/tradeHistory
func (c *Exchange) PrivateGetExecutionTradeHistory(params *url.Values) (data []models.Execution, err error) {
	reqURL := "https://www.bitmex.com/api/v1/execution/tradeHistory"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivateGetGlobalNotification method for https://www.bitmex.com/api/v1/globalNotification
func (c *Exchange) PrivateGetGlobalNotification(params *url.Values) (data []models.GlobalNotification, err error) {
	reqURL := "https://www.bitmex.com/api/v1/globalNotification"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivateGetOrder method for https://www.bitmex.com/api/v1/order
func (c *Exchange) PrivateGetOrder(params *url.Values) (data []models.Order, err error) {
	reqURL := "https://www.bitmex.com/api/v1/order"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivateGetPosition method for https://www.bitmex.com/api/v1/position
func (c *Exchange) PrivateGetPosition(params *url.Values) (data []models.Position, err error) {
	reqURL := "https://www.bitmex.com/api/v1/position"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivateGetUser method for https://www.bitmex.com/api/v1/user
func (c *Exchange) PrivateGetUser(params *url.Values) (data models.User, err error) {
	reqURL := "https://www.bitmex.com/api/v1/user"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivateGetUserAffiliateStatus method for https://www.bitmex.com/api/v1/user/affiliateStatus
func (c *Exchange) PrivateGetUserAffiliateStatus(params *url.Values) (data models.Affiliate, err error) {
	reqURL := "https://www.bitmex.com/api/v1/user/affiliateStatus"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivateGetUserCheckReferralCode method for https://www.bitmex.com/api/v1/user/checkReferralCode
func (c *Exchange) PrivateGetUserCheckReferralCode(params *url.Values) (data float64, err error) {
	reqURL := "https://www.bitmex.com/api/v1/user/checkReferralCode"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivateGetUserCommission method for https://www.bitmex.com/api/v1/user/commission
func (c *Exchange) PrivateGetUserCommission(params *url.Values) (data models.UserCommissionsBySymbol, err error) {
	reqURL := "https://www.bitmex.com/api/v1/user/commission"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivateGetUserDepositAddress method for https://www.bitmex.com/api/v1/user/depositAddress
func (c *Exchange) PrivateGetUserDepositAddress(params *url.Values) (data string, err error) {
	reqURL := "https://www.bitmex.com/api/v1/user/depositAddress"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivateGetUserExecutionHistory method for https://www.bitmex.com/api/v1/user/executionHistory
func (c *Exchange) PrivateGetUserExecutionHistory(params *url.Values) (data []models.Execution, err error) {
	reqURL := "https://www.bitmex.com/api/v1/user/executionHistory"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivateGetUserMargin method for https://www.bitmex.com/api/v1/user/margin
func (c *Exchange) PrivateGetUserMargin(params *url.Values) (data []models.Margin, err error) {
	reqURL := "https://www.bitmex.com/api/v1/user/margin"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivateGetUserMinWithdrawalFee method for https://www.bitmex.com/api/v1/user/minWithdrawalFee
func (c *Exchange) PrivateGetUserMinWithdrawalFee(params *url.Values) (data interface{}, err error) {
	reqURL := "https://www.bitmex.com/api/v1/user/minWithdrawalFee"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivateGetUserWallet method for https://www.bitmex.com/api/v1/user/wallet
func (c *Exchange) PrivateGetUserWallet(params *url.Values) (data models.Wallet, err error) {
	reqURL := "https://www.bitmex.com/api/v1/user/wallet"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivateGetUserWalletHistory method for https://www.bitmex.com/api/v1/user/walletHistory
func (c *Exchange) PrivateGetUserWalletHistory(params *url.Values) (data []models.Transaction, err error) {
	reqURL := "https://www.bitmex.com/api/v1/user/walletHistory"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivateGetUserWalletSummary method for https://www.bitmex.com/api/v1/user/walletSummary
func (c *Exchange) PrivateGetUserWalletSummary(params *url.Values) (data []models.Transaction, err error) {
	reqURL := "https://www.bitmex.com/api/v1/user/walletSummary"
	res, err := c.apiRequest("GET", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivatePostApiKey method for https://www.bitmex.com/api/v1/apiKey
func (c *Exchange) PrivatePostApiKey(params *url.Values, body bytes.Buffer) (data models.APIKey, err error) {
	reqURL := "https://www.bitmex.com/api/v1/apiKey"
	res, err := c.apiRequest("POST", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivatePostApiKeyDisable method for https://www.bitmex.com/api/v1/apiKey/disable
func (c *Exchange) PrivatePostApiKeyDisable(params *url.Values, body bytes.Buffer) (data models.APIKey, err error) {
	reqURL := "https://www.bitmex.com/api/v1/apiKey/disable"
	res, err := c.apiRequest("POST", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivatePostApiKeyEnable method for https://www.bitmex.com/api/v1/apiKey/enable
func (c *Exchange) PrivatePostApiKeyEnable(params *url.Values, body bytes.Buffer) (data models.APIKey, err error) {
	reqURL := "https://www.bitmex.com/api/v1/apiKey/enable"
	res, err := c.apiRequest("POST", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivatePostChat method for https://www.bitmex.com/api/v1/chat
func (c *Exchange) PrivatePostChat(params *url.Values, body bytes.Buffer) (data models.Chat, err error) {
	reqURL := "https://www.bitmex.com/api/v1/chat"
	res, err := c.apiRequest("POST", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivatePostOrder method for https://www.bitmex.com/api/v1/order
func (c *Exchange) PrivatePostOrder(params *url.Values, body bytes.Buffer) (data models.Order, err error) {
	reqURL := "https://www.bitmex.com/api/v1/order"
	res, err := c.apiRequest("POST", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivatePostOrderBulk method for https://www.bitmex.com/api/v1/order/bulk
func (c *Exchange) PrivatePostOrderBulk(params *url.Values, body bytes.Buffer) (data []models.Order, err error) {
	reqURL := "https://www.bitmex.com/api/v1/order/bulk"
	res, err := c.apiRequest("POST", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivatePostOrderCancelAllAfter method for https://www.bitmex.com/api/v1/order/cancelAllAfter
func (c *Exchange) PrivatePostOrderCancelAllAfter(params *url.Values, body bytes.Buffer) (data interface{}, err error) {
	reqURL := "https://www.bitmex.com/api/v1/order/cancelAllAfter"
	res, err := c.apiRequest("POST", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivatePostPositionIsolate method for https://www.bitmex.com/api/v1/position/isolate
func (c *Exchange) PrivatePostPositionIsolate(params *url.Values, body bytes.Buffer) (data models.Position, err error) {
	reqURL := "https://www.bitmex.com/api/v1/position/isolate"
	res, err := c.apiRequest("POST", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivatePostPositionLeverage method for https://www.bitmex.com/api/v1/position/leverage
func (c *Exchange) PrivatePostPositionLeverage(params *url.Values, body bytes.Buffer) (data models.Position, err error) {
	reqURL := "https://www.bitmex.com/api/v1/position/leverage"
	res, err := c.apiRequest("POST", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivatePostPositionRiskLimit method for https://www.bitmex.com/api/v1/position/riskLimit
func (c *Exchange) PrivatePostPositionRiskLimit(params *url.Values, body bytes.Buffer) (data models.Position, err error) {
	reqURL := "https://www.bitmex.com/api/v1/position/riskLimit"
	res, err := c.apiRequest("POST", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivatePostPositionTransferMargin method for https://www.bitmex.com/api/v1/position/transferMargin
func (c *Exchange) PrivatePostPositionTransferMargin(params *url.Values, body bytes.Buffer) (data models.Position, err error) {
	reqURL := "https://www.bitmex.com/api/v1/position/transferMargin"
	res, err := c.apiRequest("POST", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivatePostUserCancelWithdrawal method for https://www.bitmex.com/api/v1/user/cancelWithdrawal
func (c *Exchange) PrivatePostUserCancelWithdrawal(params *url.Values, body bytes.Buffer) (data models.Transaction, err error) {
	reqURL := "https://www.bitmex.com/api/v1/user/cancelWithdrawal"
	res, err := c.apiRequest("POST", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivatePostUserCommunicationToken method for https://www.bitmex.com/api/v1/user/communicationToken
func (c *Exchange) PrivatePostUserCommunicationToken(params *url.Values, body bytes.Buffer) (data []models.CommunicationToken, err error) {
	reqURL := "https://www.bitmex.com/api/v1/user/communicationToken"
	res, err := c.apiRequest("POST", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivatePostUserConfirmEmail method for https://www.bitmex.com/api/v1/user/confirmEmail
func (c *Exchange) PrivatePostUserConfirmEmail(params *url.Values, body bytes.Buffer) (data models.AccessToken, err error) {
	reqURL := "https://www.bitmex.com/api/v1/user/confirmEmail"
	res, err := c.apiRequest("POST", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivatePostUserConfirmWithdrawal method for https://www.bitmex.com/api/v1/user/confirmWithdrawal
func (c *Exchange) PrivatePostUserConfirmWithdrawal(params *url.Values, body bytes.Buffer) (data models.Transaction, err error) {
	reqURL := "https://www.bitmex.com/api/v1/user/confirmWithdrawal"
	res, err := c.apiRequest("POST", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivatePostUserLogout method for https://www.bitmex.com/api/v1/user/logout
func (c *Exchange) PrivatePostUserLogout(params *url.Values, body bytes.Buffer) (data interface{}, err error) {
	reqURL := "https://www.bitmex.com/api/v1/user/logout"
	res, err := c.apiRequest("POST", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivatePostUserPreferences method for https://www.bitmex.com/api/v1/user/preferences
func (c *Exchange) PrivatePostUserPreferences(params *url.Values, body bytes.Buffer) (data models.UserPreferences, err error) {
	reqURL := "https://www.bitmex.com/api/v1/user/preferences"
	res, err := c.apiRequest("POST", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivatePostUserRequestWithdrawal method for https://www.bitmex.com/api/v1/user/requestWithdrawal
func (c *Exchange) PrivatePostUserRequestWithdrawal(params *url.Values, body bytes.Buffer) (data models.Transaction, err error) {
	reqURL := "https://www.bitmex.com/api/v1/user/requestWithdrawal"
	res, err := c.apiRequest("POST", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivatePutOrder method for https://www.bitmex.com/api/v1/order
func (c *Exchange) PrivatePutOrder(params *url.Values, body bytes.Buffer) (data models.Order, err error) {
	reqURL := "https://www.bitmex.com/api/v1/order"
	res, err := c.apiRequest("PUT", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivatePutOrderBulk method for https://www.bitmex.com/api/v1/order/bulk
func (c *Exchange) PrivatePutOrderBulk(params *url.Values, body bytes.Buffer) (data []models.Order, err error) {
	reqURL := "https://www.bitmex.com/api/v1/order/bulk"
	res, err := c.apiRequest("PUT", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivateDeleteApiKey method for https://www.bitmex.com/api/v1/apiKey
func (c *Exchange) PrivateDeleteApiKey(params *url.Values) (data interface{}, err error) {
	reqURL := "https://www.bitmex.com/api/v1/apiKey"
	res, err := c.apiRequest("DELETE", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivateDeleteOrder method for https://www.bitmex.com/api/v1/order
func (c *Exchange) PrivateDeleteOrder(params *url.Values) (data []models.Order, err error) {
	reqURL := "https://www.bitmex.com/api/v1/order"
	res, err := c.apiRequest("DELETE", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}

// PrivateDeleteOrderAll method for https://www.bitmex.com/api/v1/order/all
func (c *Exchange) PrivateDeleteOrderAll(params *url.Values) (data []models.Order, err error) {
	reqURL := "https://www.bitmex.com/api/v1/order/all"
	res, err := c.apiRequest("DELETE", reqURL, params, bytes.Buffer{})
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return data, handleBodyErr(res, err)
	}
	return data, nil
}
