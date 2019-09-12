package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://192.168.150.14:5008"
	request := map[string]interface{}{
		"method": "r_get",
		"params": []interface{}{
			map[string]interface{}{
				"tx_json": map[string]interface{}{
					"Account": "zMP1doRujrx24PcuFwsJjUQczX3hZb5sXC",
					"Owner":   "zMP1doRujrx24PcuFwsJjUQczX3hZb5sXC",
					"Tables": []interface{}{
						map[string]interface{}{
							"Table": map[string]interface{}{
								"TableName": "dc_universe1",
							},
						},
					},
					"Raw": []interface{}{
						[]interface{}{},
						map[string]interface{}{
							"id": 1,
						},
					},
				},
			},
		},
	}

	jsonstr, err := json.Marshal(request)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonstr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}
