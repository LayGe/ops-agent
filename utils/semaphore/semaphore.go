package semaphore

/*
	Semaphore 限制协程并发数量
*/
type Semaphore struct {
	bufSize int
	channel chan int8
}

func NewSemaphore(concurrencyNumber int) *Semaphore {
	return &Semaphore{bufSize: concurrencyNumber, channel: make(chan int8, concurrencyNumber)}
}

func (s *Semaphore) Acquire() {
	s.channel <- int8(0)
}

func (s *Semaphore) Release() {
	<-s.channel
}

func (s *Semaphore) AvailablePermits() int {
	return s.bufSize - len(s.channel)
}
