package dtm

import (
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
)

// Config DTM 配置
type Config struct {
	Server string `yaml:"server"`
}

// DTM DTM 客户端封装
type DTM struct {
	config *Config
}

// NewDTM 创建 DTM 客户端
func NewDTM(config *Config) *DTM {
	return &DTM{config: config}
}

// NewSaga 创建 Saga 事务
func (d *DTM) NewSaga() *dtmgrpc.SagaGrpc {
	return dtmgrpc.NewSagaGrpc(d.config.Server, dtmcli.MustGenGid(d.config.Server))
}

// NewTcc 创建 TCC 事务
func (d *DTM) NewTcc() *dtmgrpc.TccGrpc {
	return dtmgrpc.NewTccGrpc(d.config.Server, dtmcli.MustGenGid(d.config.Server))
}

// NewMsg 创建消息事务
func (d *DTM) NewMsg() *dtmgrpc.MsgGrpc {
	return dtmgrpc.NewMsgGrpc(d.config.Server, dtmcli.MustGenGid(d.config.Server))
}


