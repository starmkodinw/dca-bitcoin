package services

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"dca-bitcoin/models"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func BuyBitcoin() {
	path := "/api/v3/market/place-bid"

	ts := strconv.FormatInt(time.Now().UnixMilli(), 10)
	reqBody := map[string]interface{}{
		"sym": SYM,
		"amt": AMOUNT,
		"rat": 0.0,
		"typ": MARKET,
	}

	payload := []string{
		ts,
		"POST",
		path,
		string(mustMarshal(reqBody)),
	}
	payloadString := strings.Join(payload, "")
	sig := genSign(API_SECRET, payloadString)

	client := &http.Client{}
	req, err := http.NewRequest("POST", URL+path, bytes.NewReader(mustMarshal(reqBody)))
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-BTK-APIKEY", API_KEY)
	req.Header.Set("X-BTK-TIMESTAMP", fmt.Sprintf("%v", ts))
	req.Header.Set("X-BTK-SIGN", sig)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var result models.Response
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v \n%+v\n", time.Now().UTC(), result.Result)
}

func genSign(apiSecret, payloadString string) string {
	mac := hmac.New(sha256.New, []byte(apiSecret))
	mac.Write([]byte(payloadString))
	return hex.EncodeToString(mac.Sum(nil))
}

func mustMarshal(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}
