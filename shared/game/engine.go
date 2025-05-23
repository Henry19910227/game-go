package game

import (
	"errors"
	"github.com/gorilla/websocket"
	"github.com/robfig/cron/v3"
	"net/url"
	"sync"
	"time"
)

type Engine struct {
	conn      *websocket.Conn
	countdown int
	stageChin *StageLinkedList
	timer     *cron.Cron
	mu        sync.Mutex
}

func New() *Engine {
	return &Engine{
		stageChin: &StageLinkedList{},
		timer:     cron.New(cron.WithSeconds()),
	}
}

func (e *Engine) AddCronFunc(spec string, f func(ctx *Context) func()) {
	_, _ = e.timer.AddFunc(spec, f(&Context{engine: e}))
}

func (e *Engine) AddStage(stage *Stage) {
	if stage == nil {
		return
	}
	if e.stageChin.head == nil {
		e.stageChin.head = stage
		e.stageChin.current = stage
		e.stageChin.tail = stage
		return
	}
	stage.next = e.stageChin.head
	e.stageChin.tail.next = stage
	e.stageChin.tail = stage
}

func (e *Engine) Run(scheme string, host string, path string) error {
	if e.stageChin.head == nil {
		return errors.New("尚未設置 stage")
	}
	// WebSocket 服務器 URL
	serverURL := url.URL{Scheme: scheme, Host: host, Path: path}

	// 建立連接
	conn, _, err := websocket.DefaultDialer.Dial(serverURL.String(), nil)
	if err != nil {
		return err
	}
	e.conn = conn
	defer conn.Close()

	// 初始化
	e.countdown = e.stageChin.Current().Countdown
	e.stageChin.Current().Handler(&Context{engine: e})

	// 啟動 cron 排程
	e.timer.Start()

	// 主循環
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case _ = <-ticker.C:
			e.countdown--
			if e.countdown >= 0 {
				continue
			}
			e.stageChin.Next()
			e.countdown = e.stageChin.Current().Countdown
			e.stageChin.Current().Handler(&Context{engine: e})
		}
	}
}
