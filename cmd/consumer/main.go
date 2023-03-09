package main

import (
	"context"
	"fmt"
	config "mini-douyin/pkg/configs"
	"mini-douyin/pkg/consts"
	"mini-douyin/pkg/entity"
	"strconv"
	"sync"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

var MQC rocketmq.PushConsumer

func Init() {
	config.InitLoggerOnly("/root/mini-douyin/logs/consumer.log", "debug")
	config.InitDB()
	config.InitRedis()
}

func initMQConsumer() {
	// 订阅主题、消费
	endPoint := []string{consts.MQEndPoint}
	// 创建一个consumer实例
	c, err := rocketmq.NewPushConsumer(consumer.WithNameServer(endPoint),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName("mini-douyin"),
	)
	if err != nil {
		panic(err)
	}
	MQC = c
}

func main() {
	// 初始化
	Init()
	// 初始化消费者
	initMQConsumer()
	// 订阅消息
	SubcribeMessage()
	// 启动消费
	// 启动consumer
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("consumer start error: %s\n", err.(error))
		}
		MQC.Shutdown()
	}()
	MQC.Start()
	MQC.Resume()

	for {
		time.Sleep(time.Second)
	}
}

func convertISlice2StringSlice(ifaceSlice []interface{}) []string {
	stringSlice := make([]string, len(ifaceSlice))
	for i, v := range ifaceSlice {
		str := ""
		for _, n := range v.([]uint8) {
			str += string(n)
		}
		stringSlice[i] = str
	}
	return stringSlice
}

func SubcribeMessage() {
	// 订阅用户点赞视频信息，读取redis，然后写入数据库
	err := MQC.Subscribe("mini-douyin-user-favorite", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		var wg sync.WaitGroup
		for i := range msgs {
			wg.Add(1)
			go func(userId string) {
				defer wg.Done()
				UserFavorite(userId)
			}(string(msgs[i].Body))
		}
		wg.Wait()
		return consumer.ConsumeSuccess, nil
	})

	if err != nil {
		panic(err)
	}
}

func UserFavorite(userId string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("UserFavorite Error", err)
		}
	}()
	var err error
	// 1. 读取user_favorite_userid的所有信息
	config.RD.Send("SELECT", 10)
	config.RD.Send("HGETALL", "user_favorite_"+userId)
	config.RD.Flush()
	if _, err = config.RD.Receive(); err != nil {
		panic(err)
	}
	v, err := config.RD.Receive()
	if err != nil {
		panic(err)
	}
	stringSlice := convertISlice2StringSlice(v.([]interface{}))

	// 2. 构造video map和插入mysql的结构体
	var favoriteList []entity.Favorite
	var favoriteIdList []int64
	videoActionMap := make(map[int64]int32)
	userIdInt64, _ := strconv.ParseInt(userId, 10, 64)
	for i := 0; i < len(stringSlice); i++ {
		videoId, _ := strconv.ParseInt(stringSlice[i], 10, 64)
		action, _ := strconv.ParseInt(stringSlice[i+1], 10, 32)
		videoActionMap[videoId] = int32(action)
		id, _ := strconv.ParseInt(userId[:9]+stringSlice[i][:9], 10, 64)
		favoriteList = append(favoriteList, entity.Favorite{Id: id, UserId: userIdInt64, VideoId: videoId, ActionType: int32(action)})
		favoriteIdList = append(favoriteIdList, id)
		i += 1
	}
	config.Logger.Info("UserId : " + userId + " favorite length : " + strconv.Itoa(len(favoriteIdList)))

	// 3. 查询哪些已经在数据库中了要执行update，那些执行create
	var remoteFavoriteList []entity.Favorite
	existMap := make(map[int64]int32) // 存储了已经在数据库中的id
	err = config.DB.Model(&entity.Favorite{}).Where("user_id = ? AND id in ?", userIdInt64, favoriteIdList).Find(&remoteFavoriteList).Error
	if err != nil {
		panic(err)
	}
	for _, favorite := range remoteFavoriteList {
		existMap[favorite.Id] = favorite.ActionType
	}

	// 4. 启动事务执行数据库插入和修改, 并
	tx := config.DB.Begin()
	for _, favorite := range favoriteList {
		if value, exist := existMap[favorite.Id]; !exist {
			// 执行增加
			config.Logger.Info("ADD UserId : " + strconv.Itoa(int(favorite.UserId)) + " VideoId : " + strconv.Itoa(int(favorite.VideoId)) + " Action : " + strconv.Itoa(int(favorite.ActionType)))
			err = tx.Create(&favorite).Error
		} else if value != favorite.ActionType {
			// 执行修改
			config.Logger.Info("UPDATE UserId : " + strconv.Itoa(int(favorite.UserId)) + " VideoId : " + strconv.Itoa(int(favorite.VideoId)) + " Action : " + strconv.Itoa(int(favorite.ActionType)))
			err = config.DB.Exec("UPDATE favorites SET action_type = ? where user_id = ? AND id = ?", favorite.ActionType, userIdInt64, favorite.Id).Error
		} else {
			config.Logger.Info("DONOTHING UserId : " + strconv.Itoa(int(favorite.UserId)) + " VideoId : " + strconv.Itoa(int(favorite.VideoId)) + " Action : " + strconv.Itoa(int(favorite.ActionType)))
		}
		if err != nil {
			tx.Rollback()
		}
	}
	tx.Commit()
}
