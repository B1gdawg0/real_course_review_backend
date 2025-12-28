package background

import (
	"log"

	"github.com/B1gdawg0/real_course_review_backend/internal/dtos"
)

func StartEventWorker(queue <-chan dtos.Event) {
	log.Println("[worker] event worker started")

    go func() {
        for e := range queue {
            process(e)
        }
    }()
}


func process(e dtos.Event) {
	log.Printf(
        "[process] type=%d object=%d user=%d ts=%d\n",
        e.T, e.I, e.U, e.Ts,
    )
    // example using Redis sorted set
    // key := "hot:course"
    // redis.ZIncrBy(ctx, key, 1, strconv.Itoa(int(e.I)))
}