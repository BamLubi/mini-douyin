package configs

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

var Nacos naming_client.INamingClient

func initNacos() {
	clientConfig := constant.ClientConfig{
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "192.168.45.129",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
	}
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		Logger.Error("Nacos Connect Fail" + err.Error())
		panic(err)
	}

	Nacos = namingClient
	Logger.Info("Nacos Connect Success.")
}

func NacosRegister(ip string, port uint64, serviceName string) error {
	_, err := Nacos.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          ip,
		Port:        port,
		ServiceName: serviceName,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{},
		ClusterName: "cluster-mini-douyin", // 默认值DEFAULT
		GroupName:   "group-mini-douyin",   // 默认值DEFAULT_GROUP
	})
	return err
}

func NacosDeregister(ip string, port uint64, serviceName string) error {
	_, err := Nacos.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          ip,
		Port:        port,
		ServiceName: serviceName,
		Ephemeral:   true,
		Cluster:     "cluster-mini-douyin", // 默认值DEFAULT
		GroupName:   "group-mini-douyin",   // 默认值DEFAULT_GROUP
	})
	return err
}

func NacosSelectOneHealthyInstance(serviceName string) (string, uint64, error) {
	instance, err := Nacos.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: serviceName,
		Clusters:    []string{"cluster-mini-douyin"}, // 默认值DEFAULT
		GroupName:   "group-mini-douyin",
	})
	if err != nil {
		return "", 0, err
	}
	return instance.Ip, instance.Port, err
}
