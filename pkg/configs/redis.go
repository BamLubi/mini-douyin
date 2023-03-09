package configs

import (
	"math/rand"
	"strconv"
	"strings"
	"unsafe"

	"github.com/gomodule/redigo/redis"
)

var RD redis.Conn
var PS redis.PubSubConn

func InitRedis() {
	var err error
	RD, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		Logger.Error("Redis Connect Fail" + err.Error())
		panic(err)
	}
	Logger.Info("Redis Connect Success.")
}

// 监听订阅事件
func InitPubSub() {
	PS = redis.PubSubConn{Conn: RD}

	channel := "__keyevent@10__:expired"
	PS.PSubscribe(channel)

	go expiredKeyListen()
}

func PubSubCallbackController(pattern string, channel string, message string) error {
	Logger.Debug("PubSubCallbackController, message is " + message)
	msgs := strings.Split(message, "_")
	if msgs[0] == "user" && msgs[1] == "favorite" && msgs[2] == "ex" {
		userId := msgs[3]
		// 将用户id等信息发送给消息队列
		err := SendSyncMessage("mini-douyin-user-favorite", userId)
		if err != nil {
			Logger.Error(err.Error())
		}
	}
	return nil
}

func expiredKeyListen() {
	Logger.Info("Start Listen Redis")
	for {
		switch res := PS.Receive().(type) {
		case redis.Message:
			pattern := (*string)(unsafe.Pointer(&res.Pattern))
			channel := (*string)(unsafe.Pointer(&res.Channel))
			message := (*string)(unsafe.Pointer(&res.Data))
			if PubSubCallbackController(*pattern, *channel, *message) != nil {
				// 重新设置计时器，并且为了防止缓存雪崩，使用基础时间+随机事件
				// 这些事件默认存在10号数据库
				RD.Send("SELECT", 10)
				RD.Send("SET", message, 1, "EX", 10+rand.Intn(20))
				RD.Flush()
				if _, err := RD.Receive(); err != nil {
					Logger.Error("callback error and reset error, message is " + *message + "; channel is " + *channel)
				}
				if _, err := RD.Receive(); err != nil {
					Logger.Error("callback error and reset error, message is " + *message + "; channel is " + *channel)
				}
				Logger.Info("callback error and reset success, message is  " + *message + "; channel is " + *channel)
			}
		case redis.Subscription:
			Logger.Info("Subscription: " + res.Channel + ": " + res.Kind + strconv.Itoa(res.Count))
		case error:
			Logger.Error("error handle..." + res.Error())
			continue
		}
	}
}

// CONFIG SET notify-keyspace-events "KEx"
