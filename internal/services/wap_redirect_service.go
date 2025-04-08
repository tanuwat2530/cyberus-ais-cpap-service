package services

import (
	"CyberusGolangShareLibDB/postgresql_db"
	"CyberusGolangShareLibDB/redis_db"
	"log"
	"time"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Struct to map the expected JSON fields
type WapRedirectRequest struct {
	IDPartner    string `json:"id_partner"`
	RefIDPartner string `json:"refid_partner"`
	MediaPartner string `json:"media_partner"`
	NamePartner  string `json:"name_partner"`
}

func WapRedirectProcess(r *http.Request) map[string]string {

	ip := r.RemoteAddr
	fmt.Println("ClientIP : " + ip)

	postgresql_db.ConnectDB()
	redis_db.ConnectRedis()

	key := "mykey"
	value := "This is a value with TTL"
	ttl := 1 * time.Hour // expires in 10 seconds

	// Set key with TTL
	if err := redis_db.SetWithTTL(key, value, ttl); err != nil {
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
		"code":    "0",
		"message": "success",
	}
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//http.Error(w, "Failed to read body", http.StatusBadRequest)
		return res
	}
	defer r.Body.Close()

	// Unmarshal JSON into struct
	var requestData WapRedirectRequest
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		//http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return res
	}

	// Print the data to the console
	fmt.Println("##### Received WAP Redirect Data #####")
	fmt.Println("IDPartner : " + requestData.IDPartner)
	fmt.Println("RefIDPartner : " + requestData.RefIDPartner)
	fmt.Println("MediaPartner  : " + requestData.MediaPartner)
	fmt.Println("NamePartner  : " + requestData.NamePartner)

	// Respond to client
	//w.WriteHeader(http.StatusOK)
	//w.Write([]byte("WAP Redirect received successfully"))

	return res
}
