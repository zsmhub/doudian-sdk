package doudian_message

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sync"
)

type Response struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

var (
	SuccessResponse        = Response{Code: 0, Msg: "success"}
	FailResponse           = Response{Code: 1, Msg: "fail"}
	MessageDataParseHandle = new(MessageDataParse)
)

type Message struct {
	Tag        string      `json:"tag"`
	MsgID      string      `json:"msg_id"`
	Data       string      `json:"data"`
	DataStruct MessageData `json:"-"`
}

func (m Message) ParseMessageFromJson() (Message, error) {
	data, err := MessageDataParseHandle.ParseMessageData(m)
	if err != nil {
		return m, err
	}
	m.DataStruct = data

	return m, nil
}

type MessageData interface {
	Tag() string
	Parse(data []byte) (MessageData, error)
}

type MessageDataParse struct {
	sync.Map
}

func (m *MessageDataParse) RegisterMessage(msgData MessageData) {
	m.Store(msgData.Tag(), msgData)
}

func (m *MessageDataParse) ParseMessageData(msg Message) (MessageData, error) {
	msgData, ok := m.Load(msg.Tag)
	if !ok {
		msgByte, _ := json.Marshal(msg)
		return nil, errors.New("回调消息解析失败，请去SDK生成对应的回调消息：" + string(msgByte))
	}
	return msgData.(MessageData).Parse([]byte(msg.Data))
}

type GetMessageResp struct {
	Message Message
	Error   error
}

// 解析并获取回调消息
func GetMessage(r *http.Request) ([]GetMessageResp, error) {
	var resp []GetMessageResp
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return resp, err
	}

	// 验签
	eventSign := r.Header.Get("event-sign")
	if !VerifySign(body, eventSign) {
		return resp, errors.New("invalid signature")
	}

	var msgs []Message
	if err := json.Unmarshal(body, &msgs); err != nil {
		return resp, err
	}

	for _, m := range msgs {
		tmp, err := m.ParseMessageFromJson()
		if err != nil {
			resp = append(resp, GetMessageResp{Message: Message{}, Error: err})
			continue
		}
		resp = append(resp, GetMessageResp{Message: tmp, Error: nil})
	}

	return resp, nil
}
