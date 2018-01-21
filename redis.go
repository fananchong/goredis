package goredis

type IClient interface {
	Do(commandName string, args ...interface{}) (reply interface{}, err error)
}

type Client struct {
	cli IClient
}

func NewClient(dbName string, addrs []string, option *Option) *Client {
	this := &Client{}
	this.Init(dbName, addrs, option)
	return this
}

func (this *Client) Init(dbName string, addrs []string, option *Option) {
	switch option.redisType {
	case Unknow:
		{
			// TODO: 自识别
		}
	case Standalone:
		{

		}
	case Sentinel:
		{
			this.cli = NewSentinelClient(dbName, addrs, option)
		}
	case Cluster:
		{

		}
	}
}

func (this *Client) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	return this.cli.Do(commandName, args...)
}
