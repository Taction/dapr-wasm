/**
 * @Author: zhangchao
 * @Description:
 * @Date: 2023/3/2 8:21 PM
 */
package runtime

type Config struct {
	Port     int
	WASMPath string
}

type subscription struct {
	PubsubName      string            `json:"pubsubname"`
	Topic           string            `json:"topic"`
	Route           string            `json:"route"`
	DeadLetterTopic string            `json:"deadLetterTopic"`
	Metadata        map[string]string `json:"metadata"`
}
