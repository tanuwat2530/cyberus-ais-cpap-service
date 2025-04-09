package services

import (
	"CyberusGolangShareLibDB/postgresql_db"
	"CyberusGolangShareLibDB/redis_db"
	"log"
	"strconv"
	"time"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
)

// Struct to map the expected JSON fields
type SubscriptionCallbackRequest struct {
	Msisdn    string `json:"msisdn"`
	Shortcode string `json:"shortcode"`
	Operator  string `json:"operator"`
	Action    string `json:"action"`
	Code      string `json:"code"`
	Desc      string `json:"desc"`
	Timestamp string `json:"timestamp"`
	TranRef   string `json:"tranref"`
	RefId     string `json:"refid"`
	Media     string `json:"media"`
	Token     string `json:"token"`
	ClientId  string `json:"clientid"`
}

func SubscriptionCallbackProcessRequest(r *http.Request) map[string]string {

	// Get current time
	now := time.Now()
	// Unix timestamp in nanoseconds
	timestamp := (now.UnixNano())
	nano_timestamp := strconv.FormatInt(timestamp, 10)

	// Generate a random UUID (UUID v4)
	transaction_id := uuid.New().String()

	// Get a Client IP address
	ip := r.RemoteAddr

	fmt.Println("ClientIP : " + ip)

	postgresql_db.ConnectDB()
	redis_db.ConnectRedis()

	var payload map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err == nil {
		// Example: print the values
		fmt.Printf("Received: %+v\n", payload)
	}

	// Add ClientIP to Payload
	payload["ClientIP"] = string(ip)
	payload["ProviderUrl"] = "https://cpdomain.com/cp/aoc/subscription/callback/url"
	payload["RedisKey"] = transaction_id

	// Convert the struct to JSON string
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Failed to convert payload to JSON:  %+v\n ", http.StatusInternalServerError)
	}

	payloadString := string(payloadBytes)
	redis_key := "subscription-callback:" + transaction_id
	ttl := 24 * time.Hour // expires in 1 Hour

	// Set key with TTL
	if err := redis_db.SetWithTTL(redis_key, payloadString, ttl); err != nil {
		//write to file if Redis problem or forward request to AIS
		log.Fatalf("SetWithTTL error: %v", err)
	}
	fmt.Println("Key set successfully with TTL")

	// Get the key
	// val, err := redis_db.GetValue(key)
	// if err != nil {
	// 	log.Printf("GetValue error: %v", err)
	// } else {
	// 	fmt.Printf("Retrieved value: %s\n", val)
	// }

	//redis_db.Set("aaa", "AAA", 300)
	res := map[string]string{
		"code":           "0",
		"message":        "retrieved",
		"timestamp":      nano_timestamp,
		"transaction_id": transaction_id,
	}

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//http.Error(w, "Failed to read body", http.StatusBadRequest)
		return res
	}
	defer r.Body.Close()

	// Unmarshal JSON into struct
	var requestData SubscriptionCallbackRequest
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		//http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return res
	}

	// Print the data to the console
	fmt.Println("##### Received Subscription Callback Data #####")
	fmt.Println("IDPartner : " + requestData.Msisdn)
	fmt.Println("Shortcode : " + requestData.Shortcode)
	fmt.Println("Operator  : " + requestData.Operator)
	fmt.Println("Action  : " + requestData.Action)
	fmt.Println("Code  : " + requestData.Code)
	fmt.Println("Desc  : " + requestData.Desc)
	fmt.Println("Timestamp  : " + requestData.Timestamp)
	fmt.Println("TranRef  : " + requestData.TranRef)
	fmt.Println("Action  : " + requestData.Action)
	fmt.Println("RefId  : " + requestData.RefId)
	fmt.Println("Media  : " + requestData.Media)
	fmt.Println("Token  : " + requestData.Token)

	// Respond to client
	//w.WriteHeader(http.StatusOK)
	//w.Write([]byte("WAP Redirect received successfully"))

	return res
}
