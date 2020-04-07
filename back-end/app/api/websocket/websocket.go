package websocket

import (
	//"encoding/json"
	//"fmt"
	//"github.com/friday182/ttm-admin/app/model"
	//"github.com/gogf/gf/container/gmap"
	//"github.com/gogf/gf/container/garray"
	//"github.com/gogf/gf/container/gset"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"

	//"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/os/glog"

	//"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/grand"
	"github.com/gogf/gf/util/gvalid"
	//"time"
)

// Msg 消息结构体
type Msg struct {
	Sq    int         `json:"sq" v:""`
	Type  string      `json:"type" v:"type@required#消息类型不能为空"`
	Data  interface{} `json:"data" v:""`
	From  string      `json:"from" v:""`
	Token string      `json:"token" v:""`
}

type Worker struct {
	name       string
	ws         *ghttp.WebSocket
	source     chan interface{}
	sleepCount int
	quit       bool
}

const (
	// SendInterval  = time.Second
	WS_TIMEOUT_COUNT = 20 // 10 minutes based on 30 s task
)

var (
	//users = gmap.New(true)
	// cache   = gcache.New()
	//atest = garray.NewIntArray(true)
	wsList  []*Worker
	squence int
)

func init() {
	// Periodic task to handle websocket clients
	// * means every second minutes hour ..., when match the number, run function
	_, err := gcron.Add("0 */12 * * * *", PeriodSend, "ws-task")
	if err != nil {
		panic("[WebSocket] Start Timer Failed ")
	}
	// gcron.Start("ws-task")
}

func Handler(r *ghttp.Request) {
	worker := &Worker{}

	// Create new WebSocket
	s, err := r.WebSocket()
	if err != nil {
		g.Log().Error(err)
		return
	} else {
		worker.ws = s
		worker.quit = false
		worker.name = grand.Digits(6)
		glog.Println("websocket name is: ", worker.name)
	}

	worker.Start()
	wsList = append(wsList, worker)
}

func (w *Worker) Start() {
	w.source = make(chan interface{})
	w.SendMsg(Msg{0, "CONNECTED", "", "S", ""})	

	go w.readMsg()
}

func (w *Worker) readMsg() {
	msg := &Msg{}
	for {
		_, msgByte, err := w.ws.ReadMessage()
		if err != nil {
			w.Close()
			break
		}

		w.sleepCount = 0

		if err := gjson.DecodeTo(msgByte, msg); err != nil {
			w.SendMsg(Msg{0, "ERROR", "消息格式不正确: " + err.Error(), "", ""})
			continue
		}

		if e := gvalid.CheckStruct(msg, nil); e != nil {
			w.SendMsg(Msg{0, "ERROR", e.String(), "", ""})
			continue
		}

		glog.Println("Add student: ", msg)

		if msg.Type == "HEARTBEAT" {
			w.SendMsg(Msg{0, "HEARTBEAT", "pong", "S", ""})
			continue
		}

		rMsg := Msg{}
		switch msg.From {
		case "MENTOR":
			//rMsg = model.ProcessMentorData(*msg)
			break
		case "STUDENT":
			//rMsg = model.ProcessStudentData(*msg)
			break
		default:
			break
		}
		if rMsg.From != "" {
			w.SendMsg(rMsg)
		}
	}
}

func (w *Worker) SendMsg(msg Msg) error {
	b, err := gjson.Encode(msg)
	if err != nil {
		return err
	}

	return w.ws.WriteMessage(ghttp.WS_MSG_TEXT, []byte(b))
}

func (w *Worker) Close() {
	// disconnected
	w.quit = true
	// remove itself from the list
	var tmp []*Worker
	for i, w := range wsList {
		if w.quit == false {
			tmp = append(tmp, wsList[i])
		}
	}
	wsList = tmp
	w.ws.Close()
}

func BroadcastMsg(msg Msg) error {
	squence++
	msg.Sq = squence

	b, err := gjson.Encode(msg)
	if err != nil {
		return err
	}

	glog.Println("WS Sending Msg Size: ", len(b))
	for _, w := range wsList {
		glog.Println("WS name: ", w.name)
		w.ws.WriteMessage(ghttp.WS_MSG_TEXT, []byte(b))
	}

	return nil
}

// Function to send message out every 10 second
func PeriodSend() {
	glog.Println("Current WS list size: ", len(wsList))
	for _, w := range wsList {
		w.sleepCount++
		if w.sleepCount > WS_TIMEOUT_COUNT {
			w.Close()
		}
		glog.Println("WS name: ", w.name)
	}
}
