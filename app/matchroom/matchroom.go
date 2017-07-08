package matchroom

import (
	"container/list"
	"fmt"
	"time"
)

type Event struct {
	UserID    string `json:"user"`
	VP        int    `json:"vp"`
	Timestamp int    // Unix timestamp (secs)
}

type Matching struct {
	Host  Event `json:"host"`
	Guest Event `json:"guest"`
}

type Subscription struct {
	New <-chan Matching
}

func newEvent(user string, vp int) Event {
	return Event{user, vp, int(time.Now().Unix())}
}

func Subscribe() Subscription {
	resp := make(chan Subscription)
	//subscribeの引数がチャネルなので、待たずに次に進む
	subscribe <- resp
	//subscribeを処理した結果をrespへ受信 受信したSubscriptionをReturn
	return <-resp
}

func Join(user string, vp int) {
	publish <- newEvent(user, vp)
}

const archiveSize = 10

var (
	// Send a channel here to get room events back.  It will send the entire
	// archive initially, and then new messages as they come in.
	subscribe = make(chan (chan<- Subscription), 10)
	// Send a channel here to unsubscribe.
	unsubscribe = make(chan (<-chan Event), 10)
	// Send events here to publish them.
	publish = make(chan Event, 10)
)

// This function loops forever, handling the chat room pubsub
func matchroom() {

	subscribers := list.New()
	waitings := list.New()

	for {
		select {
		case ch := <-subscribe:
			fmt.Println("subscribe")

			//イベントの受信チャンネル、Opponentに代入
			subscriber := make(chan Matching, 10)
			subscribers.PushBack(subscriber)
			ch <- Subscription{subscriber}

		case event := <-publish:
			if waitings.Len() > 0 {
				fmt.Println("publish", event.UserID)
				for ch := subscribers.Front(); ch != nil; ch = ch.Next() {
					host := waitings.Front().Value.(Event)
					matching := Matching{host, event}
					//subscriber = Newチャンネルに送信
					ch.Value.(chan Matching) <- matching
				}
				waitings.Remove(waitings.Front())
			} else {
				fmt.Println("waiting", event.UserID)
				waitings.PushBack(event)
			}
		}
	}
}

func init() {
	go matchroom()
}

// Helpers

// Drains a given channel of any messages.
func drain(ch <-chan Event) {
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				return
			}
		default:
			return
		}
	}
}
