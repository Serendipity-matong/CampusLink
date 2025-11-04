package mq

import (
	"context"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

// Config NATS 配置
type Config struct {
	URL             string        `yaml:"url"`
	MaxReconnects   int           `yaml:"max_reconnects"`
	ReconnectWait   time.Duration `yaml:"reconnect_wait"`
	Timeout         time.Duration `yaml:"timeout"`
	StreamName      string        `yaml:"stream_name"`
	StreamSubjects  []string      `yaml:"stream_subjects"`
}

// NATS NATS 客户端封装
type NATS struct {
	conn *nats.Conn
	js   nats.JetStreamContext
}

// NewNATS 创建 NATS 客户端
func NewNATS(config *Config) (*NATS, error) {
	opts := []nats.Option{
		nats.MaxReconnects(config.MaxReconnects),
		nats.ReconnectWait(config.ReconnectWait),
		nats.Timeout(config.Timeout),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			fmt.Printf("NATS disconnected: %v\n", err)
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			fmt.Printf("NATS reconnected: %s\n", nc.ConnectedUrl())
		}),
	}

	conn, err := nats.Connect(config.URL, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect NATS: %w", err)
	}

	// 创建 JetStream 上下文
	js, err := conn.JetStream()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to create JetStream context: %w", err)
	}

	// 创建或更新 Stream
	if config.StreamName != "" && len(config.StreamSubjects) > 0 {
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     config.StreamName,
			Subjects: config.StreamSubjects,
			Storage:  nats.FileStorage,
			MaxAge:   7 * 24 * time.Hour, // 消息保留7天
		})
		if err != nil {
			conn.Close()
			return nil, fmt.Errorf("failed to add stream: %w", err)
		}
	}

	return &NATS{
		conn: conn,
		js:   js,
	}, nil
}

// Publish 发布消息
func (n *NATS) Publish(ctx context.Context, subject string, data []byte) error {
	_, err := n.js.Publish(subject, data)
	return err
}

// PublishAsync 异步发布消息
func (n *NATS) PublishAsync(subject string, data []byte) (nats.PubAckFuture, error) {
	return n.js.PublishAsync(subject, data)
}

// Subscribe 订阅消息
func (n *NATS) Subscribe(subject, queue string, handler nats.MsgHandler) (*nats.Subscription, error) {
	if queue != "" {
		return n.js.QueueSubscribe(subject, queue, handler)
	}
	return n.js.Subscribe(subject, handler)
}

// SubscribeDurable 持久订阅
func (n *NATS) SubscribeDurable(subject, durableName string, handler nats.MsgHandler) (*nats.Subscription, error) {
	return n.js.Subscribe(subject, handler, nats.Durable(durableName))
}

// Close 关闭连接
func (n *NATS) Close() {
	if n.conn != nil {
		n.conn.Close()
	}
}

// Request 请求-响应模式
func (n *NATS) Request(subject string, data []byte, timeout time.Duration) (*nats.Msg, error) {
	return n.conn.Request(subject, data, timeout)
}


