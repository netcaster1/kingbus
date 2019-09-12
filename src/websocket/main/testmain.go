package main

import (
	"encoding/json"
	"fmt"
	"websocket"
)

func main() {
	client := websocket.NewWebsocketClient("192.168.150.14:6008", "")

	request := map[string]interface{}{
		"command": "r_insert",
		"secret":  "xx5F15BLC8L2YZp7hRZompuij23Mu",
		"tx_json": map[string]interface{}{
			"TransactionType": "SQLStatement",
			"Account":         "zMP1doRujrx24PcuFwsJjUQczX3hZb5sXC",
			"Owner":           "zMP1doRujrx24PcuFwsJjUQczX3hZb5sXC",
			"Tables": []interface{}{
				map[string]interface{}{
					"Table": map[string]interface{}{
						"TableName": "dc_universe1",
						"NameInDB":  "F3660EB864F5E39D91AB58CC1492C74E8ED4A2F4",
					},
				},
			},
			"Raw": []interface{}{
				map[string]interface{}{
					"id":   4,
					"name": "park",
					"age":  88,
				},
			},
			"OpType": 6,
		},
	}
	jsonstr, err := json.Marshal(request)
	fmt.Println(err)
	resp, err := client.SendMessage(jsonstr)
	var res map[string]interface{}
	json.Unmarshal(resp, &res)
	fmt.Println(res)
	fmt.Println(string(resp))
}
