package utils

import (
	"log"
	"sync"
	"time"
)

const (
	epoch          = int64(1622304000000)              // 设置起始时间(时间戳/毫秒)：2020-01-01 00:00:00，有效期69年
	timestampBits  = uint(51)                          // 时间戳占用位数
	sequenceBits   = uint(12)                          // 序列所占的位数
	timestampMax   = int64(-1 ^ (-1 << timestampBits)) // 时间戳最大值
	sequenceMask   = int64(-1 ^ (-1 << sequenceBits))  // 支持的最大序列id数量
	workeridShift  = int64(sequenceBits)               // 机器id左移位数
	timestampShift = int64(sequenceBits)               // 时间戳左移位数
)

type SnowFlake struct {
	sync.Mutex
	timestamp int64
	sequence  int64
}

func (s *SnowFlake) GetVal() int64 {
	s.Lock()
	defer s.Unlock()
	now := time.Now().UnixNano() / 1000000
	if now == s.timestamp {
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			// 如果当前序列超出12bit长度，则需要等待下一毫秒
			// 下一毫秒将使用sequence:0
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		// 不同时间戳（精度：毫秒）下直接使用序列号：0
		s.sequence = 0
	}
	t := now - epoch
	if t > timestampMax {
		log.Printf("epoch must be between 0 and %d", timestampMax-1)
		return 0
	}
	s.timestamp = now
	return int64((t)<<timestampShift | s.sequence)
}
