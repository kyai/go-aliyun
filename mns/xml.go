package mns

import (
	"encoding/xml"
)

type XmlError struct {
	XMLName xml.Name `xml:"Error"`
	Code    string   `xml:"Code"`
	Message string   `xml:"Message"`
}

// 创建队列
type XmlCreateQueueReq struct {
	XMLName                xml.Name `xml:"Queue"`
	DelaySeconds           int      `xml:"DelaySeconds,omitempty"`           // 发送到该 Queue 的所有消息默认将以DelaySeconds参数指定的秒数延后可被消费，单位为秒。	0-604800秒（7天）范围内某个整数值，默认值为0
	MaximumMessageSize     int      `xml:"MaximumMessageSize,omitempty"`     // 发送到该Queue的消息体的最大长度，单位为byte。	1024(1KB)-65536（64KB）范围内的某个整数值，默认值为65536（64KB）。
	MessageRetentionPeriod int      `xml:"MessageRetentionPeriod,omitempty"` // 消息在该 Queue 中最长的存活时间，从发送到该队列开始经过此参数指定的时间后，不论消息是否被取出过都将被删除，单位为秒。	60 (1分钟)-1296000 (15 天)范围内某个整数值，默认值345600 (4 天)
	VisibilityTimeout      int      `xml:"VisibilityTimeout,omitempty"`      // 消息从该 Queue 中取出后从Active状态变成Inactive状态后的持续时间，单位为秒。	1-43200(12小时)范围内的某个值整数值，默认为30（秒）
	PollingWaitSeconds     int      `xml:"PollingWaitSeconds,omitempty"`     // 当 Queue 中没有消息时，针对该 Queue 的 ReceiveMessage 请求最长的等待时间，单位为秒。	0-30秒范围内的某个整数值，默认为0（秒）
	LoggingEnabled         bool     `xml:"LoggingEnabled,omitempty"`         // 是否开启日志管理功能，True表示启用，False表示停用	True/False，默认为False
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
