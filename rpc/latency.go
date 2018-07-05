package rpc

// Taken and modified from neo-utils
import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

type customTransport struct {
	rtp       http.RoundTripper
	dialer    *net.Dialer
	connStart time.Time
	connEnd   time.Time
	reqStart  time.Time
	reqEnd    time.Time
}

func newTransport() *customTransport {

	tr := &customTransport{
		dialer: &net.Dialer{
			Timeout:   1 * time.Second, //keep timeout low
			KeepAlive: 1 * time.Second,
		},
	}
	tr.rtp = &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		Dial:                  tr.dial,
		TLSHandshakeTimeout:   1 * time.Second,
		ResponseHeaderTimeout: 1 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	return tr
}

func (tr *customTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	tr.reqStart = time.Now()
	resp, err := tr.rtp.RoundTrip(r)
	tr.reqEnd = time.Now()
	return resp, err
}

func (tr *customTransport) dial(network, addr string) (net.Conn, error) {
	tr.connStart = time.Now()
	cn, err := tr.dialer.Dial(network, addr)
	tr.connEnd = time.Now()
	return cn, err
}

func (tr *customTransport) ReqDuration() time.Duration {
	return tr.Duration() - tr.ConnDuration()
}

func (tr *customTransport) ConnDuration() time.Duration {
	return tr.connEnd.Sub(tr.connStart)
}

func (tr *customTransport) Duration() time.Duration {
	return tr.reqEnd.Sub(tr.reqStart)
}

type BlockCountResponse struct {
	Jsonrpc      string `json:"jsonrpc"`
	ID           int    `json:"id"`
	Result       int    `json:"result"`
	ResponseTime int64  `json:"-"`
}

func fetchSeedNode(url string) *BlockCountResponse {
	transport := newTransport()
	client := http.Client{Transport: transport}
	payload := strings.NewReader(" {\"jsonrpc\": \"2.0\", \"method\": \"getblockcount\", \"params\": [], \"id\": 3}")
	res, err := client.Post(url, "application/json", payload)
	if err != nil || res == nil {
		return nil
	}
	defer res.Body.Close()
	blockResponse := BlockCountResponse{}
	err = json.NewDecoder(res.Body).Decode(&blockResponse)
	if err != nil {
		return nil
	}
	blockResponse.ResponseTime = transport.ReqDuration().Nanoseconds()
	return &blockResponse
}

func GetLatency(url string) (int64, error) {
	response := fetchSeedNode(url)
	if response == nil {
		return 0, fmt.Errorf("Response is nil")
	}

	latency := response.ResponseTime / int64(time.Millisecond)
	return latency, nil
}
