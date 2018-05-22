package mns

import (
	"encoding/xml"
)

type XmlError struct {
	XMLName xml.Name `xml:"Error"`
	Code    string   `xml:"Code"`
	Message string   `xml:"Message"`
}

type XmlCreateQueueReq struct {
	XMLName xml.Name `xml:"Queue"`
}

type XmlSendMessageReq struct {
	XMLName      xml.Name `xml:"Message"`
	MessageBody  string   `xml:"MessageBody"`
	DelaySeconds int      `xml:"DelaySeconds,omitempty"`
	Priority     int      `xml:"Priority,omitempty"`
}

// 消费消息
type XmlReceiveMessage struct {
	XMLName          xml.Name `xml:"Message"`
	MessageId        string   `xml:"MessageId"`        // 消息编号，在一个 Queue 中唯一
	ReceiptHandle    string   `xml:"ReceiptHandle"`    // 本次获取消息产生的临时句柄，用于删除和修改处于 Inactive 消息，NextVisibleTime 之前有效。
	MessageBody      string   `xml:"MessageBody"`      // 消息正文
	MessageBodyMD5   string   `xml:"MessageBodyMD5"`   // 消息正文的 MD5 值
	EnqueueTime      string   `xml:"EnqueueTime"`      // 消息发送到队列的时间，从 1970年1月1日 00:00:00 000 开始的毫秒数
	NextVisibleTime  string   `xml:"NextVisibleTime"`  // 下次可被再次消费的时间，从1970年1月1日 00:00:00 000 开始的毫秒数
	FirstDequeueTime string   `xml:"FirstDequeueTime"` // 第一次被消费的时间，从1970年1月1日 00:00:00 000 开始的毫秒数
	DequeueCount     string   `xml:"DequeueCount"`     // 总共被消费的次数
	Priority         string   `xml:"Priority"`         // 消息的优先级权值
}

func GetXmlStr(v interface{}) (str string, err error) {
	var b []byte
	b, err = xml.Marshal(v)
	if err != nil {
		return
	}
	str = xml.Header + string(b)
	return
}
