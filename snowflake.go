package snowflake

import (
	"strconv"
)

type Snowflake uint64

type SnowflakeStruct struct {
	timestamp uint64
	machineID uint
	sequence  uint
}

func (s Snowflake) String() string {
	return strconv.FormatUint(s.Number(), 10)
}

func (s Snowflake) Number() uint64 {
	return uint64(s)
}

func (s Snowflake) Bytes() []byte {
	return []byte(s.String())
}

func (s Snowflake) Timestamp() uint {
	return uint(s.Number() >> 22)
}

func (s Snowflake) MachineID() uint {
	return uint((s.Number() & 0x3FF0000) >> 12)
}

func (s Snowflake) Sequence() uint {
	return uint(s.Number() & 0xFFF)
}

func (s *Snowflake) Parse() *SnowflakeStruct {
	sNum := s.Number()

	return &SnowflakeStruct{
		timestamp: sNum >> 22,
		machineID: uint((s.Number() & 0x3FF0000) >> 12),
		sequence:  uint(sNum & 0xFFF),
	}
}
