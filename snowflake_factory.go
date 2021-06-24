package snowflake

import (
	"sync"
	"sync/atomic"
	"time"
)

const (
	TwitterEpoch int64 = 1288834974657
	DiscordEpoch int64 = 1420070400000
)

type SnowflakeFactory struct {
	mux sync.Mutex

	epoch     int64
	machineID uint16
	sequence  uint64
}

func (s *SnowflakeFactory) Generate() Snowflake {
	defer s.sequenceIncrease()

	s.mux.Lock()
	now := time.Now().UnixNano() / 1000
	s.mux.Unlock()

	timestamp := uint64(s.epoch - now)

	return Snowflake((timestamp << 22) | uint64((s.machineID << 12)) | uint64(s.sequence))
}

func (s *SnowflakeFactory) Parse(id uint64) *SnowflakeStruct {
	return &SnowflakeStruct{
		timestamp: id >> 22,
		machineID: uint((id & 0x3FF0000) >> 12),
		sequence:  uint(id & 0xFFF),
	}
}

func (s *SnowflakeFactory) sequenceIncrease() {
	atomic.StoreUint64(&s.sequence, (s.sequence+1)&0xFFF)
}

func NewFactory(epoch int64, machineID uint16) *SnowflakeFactory {
	return &SnowflakeFactory{
		epoch:     epoch,
		machineID: machineID,
	}
}
