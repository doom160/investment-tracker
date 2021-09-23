package equity

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "strings"
    "github.com/fxtlabs/date"
)

func GetEquityInfo(ticker string) (equityInfo EquityInfo, err error) {
    var equity Equity
    ticker = strings.ToUpper(ticker)
    epoch := date.Today().UTC().AddDate(0,-6,0).Unix()
    
    var resp *http.Response
    resp, err = http.Get(fmt.Sprintf("https://query2.finance.yahoo.com/v8/finance/chart/%s?period1=%d&period2=99999999999999&interval=1d",ticker, epoch))

    if err != nil {
        return equity, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return equity, err
    }

    json.Unmarshal([]byte(string(body)), &equity)
    //fmt.Printf("%f", myStoredVariable.Chart.Result[0].Meta.RegularMarketPrice)
    equityInfo = EquityInfo{ Symbol: equity.Chart.Result[0].Meta.Symbol , 
                             Currency: equity.Chart.Result[0].Meta.Currency,
                             RegularMarketPrice: equity.Chart.Result[0].Meta.RegularMarketPrice,
                             ChartPreviousClose: equity.Chart.Result[0].Meta.ChartPreviousClose,
                             Timestamp: equity.Chart.Result[0].Timestamp,
                             Open: equity.Chart.Result[0].Indicators.[0].Open,
                             High: equity.Chart.Result[0].Indicators.[0].High,
                             Low: equity.Chart.Result[0].Indicators.[0].Low ,
                             Close: equity.Chart.Result[0].Indicators.[0].Close,
                             Volume: equity.Chart.Result[0].Indicators.[0].Volume
                             }
    return equity, nil
}

type Equity struct {
	Chart Chart `json:"chart"`
}
 
type Chart struct {
	Result []Result `json: "result"`
	Error  string   `json: "error"`
}

type Result struct {
	Meta Meta `json: "meta"`
	Timestamp  []int32    `json: "timestamp"`
    Indicators Indicators `json: "indicators"`
}

type Meta struct {
	Currency string `json: "currency"`
	Symbol  string `json: "symbol"`
    RegularMarketPrice float32 `json: "regularMarketPrice"`
    ChartPreviousClose float32 `json: "chartPreviousClose"`
}

type Indicators struct {
    Quote []Quote `json: "quote"`
}

type Quote struct {
    Open []float32 `json: "open"`
    High []float32 `json: "high"`
    Low []float32 `json: "low"`
    Close []float32 `json: "close"`
    Volume []int32 `json: "volume"`
}

type EquityInfo struct {
    Symbol  string `json: "symbol"`
    Currency string `json: "currency"`
    RegularMarketPrice float32 `json: "regularMarketPrice"`
    ChartPreviousClose float32 `json: "chartPreviousClose"`
    Timestamp  []int32    `json: "timestamp"`
    Open []float32 `json: "open"`
    High []float32 `json: "high"`
    Low []float32 `json: "low"`
    Close []float32 `json: "close"`
    Volume []int32 `json: "volume"`
}

