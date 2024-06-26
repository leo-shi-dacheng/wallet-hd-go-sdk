package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Server struct {
	Post string `yaml:"post"`
}

type RPC struct {
	RPCURL  string `yaml:"rpc_url"`
	RPCUser string `yaml:"rpc_user"`
	RPCPass string `yaml:"rpc_pass"`
}
type Node struct {
	//*号什么意思  加了的区别
	//这是一个指向RPC结构体的指针切片。每个元素都是一个指向RPC的指针
	RPCs          []*RPC `yaml:"rpcs"`
	TpApiUrl      string `yaml:"tp_api_url"`
	TpApiKey      string `yaml:"tp_api_key"`
	Confirmations uint64 `yaml:"confirmations"`
	ApiToken      string `yaml:"apiToken"`
}

type Fullnode struct {
	Eth     Node `yaml:"eth"`
	Op      Node `yaml:"op"`
	Polygon Node `yaml:"polygon"`
}

type Config struct {
	Server   Server   `yaml:"server"`
	Fullnode Fullnode `yaml:"fullnode"`
	NetWork  string   `yaml:"network"`
	Chains   []string `yaml:"chains"`
}

type NetWorkType int

const (
	MainNet NetWorkType = iota
	TestNet
	RegTest
)

// setup init config
func New(path string) (*Config, error) {
	// config global config instance
	var config = new(Config)
	//h := log.NewTerminalHandler(os.Stdout, true)
	//log.Root()

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	println(data, "config data")

	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
