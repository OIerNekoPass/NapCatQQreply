package qq_reply

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

var url = "http://127.0.0.0:11451/send_group_msg"
var bot_token = "bot_token"

func Set_url(ip string, port string) {
	url = "http://" + ip + ":" + port + "/send_group_msg"
}

fun Set_token(new_token string) {
	bot_token = new_token
}

func send_payload(payload *strings.Reader) {
	client := &http.Client {
	}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bot_token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
   fmt.Println(string(body))
}

func Group_text(group_id string, text string) {
	text = strings.ReplaceAll(text, "\n", `\n`)
	payload := strings.NewReader(`{
	"group_id": "` + group_id + `",
	"message": [
		{
			"type": "text",
			"data": {
				"text": "` + text + `"
			}
		}
	]
}`)

	send_payload(payload)
}

func Luck_reply(group_id string, user_id string, pic_route string) {
	payload := strings.NewReader(`{
	"group_id": "` + group_id + `",
	"message": [
		{
			"type": "at",
			"data": {
				"qq": "` + user_id + `"
			}
		},
		{
			"type": "image",
			"data": {
				"file": "` + pic_route + `"
			}
		}
	]
}`)
	
	send_payload(payload)
}

func Group_record(group_id string, record_path string) {
	payload := strings.NewReader(`{
	"group_id": "` + group_id + `",
	"message": [
		{
			"type": "record",
			"data": {
				"file": "` + record_path + `"
			}
		}
	]
}`)

	send_payload(payload)
}
