package bot

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	URL = "https://open.larksuite.com/open-apis/message/v4/send/"
)

type LarkChatRespMsgData struct {
	MessageId  string `json:"message_id,omitempty"`
}

type LarkChatRespMsg struct {
	Code  float64 `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
	Data  *LarkChatRespMsgData `json:"data,omitempty"`
}

type Content struct {
	Text string `json:"text,omitempty"`
}

type LarkChatMsg struct {
	OpenId  string   `json:"open_id,omitempty"`
	RootId  string   `json:"root_id,omitempty"`
	ChatId  string   `json:"chat_id,omitempty"`
	UserId  string   `json:"user_id,omitempty"`
	Email   string   `json:"email,omitempty"`
	MsgType string   `json:"msg_type,omitempty"`
	Content *Content `json:"content,omitempty"`
}

func Chat(msg LarkChatMsg) (*LarkChatRespMsg, error){
	mb, err := json.Marshal(msg)
	if err != nil{
		return nil, err
	}
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(mb))
	req.Header.Set("Authorization", "Bearer t-fee42159a366c575f2cd2b2acde2ed1e94c89d5f")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	respMsg := &LarkChatRespMsg{}
	err = json.Unmarshal(body,respMsg)
	if err !=nil {
		return nil, err
	}
	return  respMsg, nil
}
