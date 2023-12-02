package regulation

import "time"

type Request struct {
	countRequest map[int64]int
}

func NewRequestRegulation() *Request {
	return &Request{countRequest: make(map[int64]int)}
}

func (reg *Request) GetCountRequestsNow(now time.Time) int {
	return reg.countRequest[now.Unix()]
}

func (reg *Request) NewRequest(now time.Time) {
	count := reg.countRequest[now.Unix()]

	if count == 0 {
		reg.countRequest[now.Unix()] = 1
	} else {
		reg.countRequest[now.Unix()] = count + 1
	}

}

func (reg *Request) ClearInfo(now time.Time) {
	delete(reg.countRequest, now.Unix())
}
