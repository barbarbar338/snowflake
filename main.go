package snowflake

import (
	"errors"
	"strconv"
	"time"
)

type Snowflake struct {
	EPOCH int64
	INCREMENT int64
}

type DeconstructedSnowflake struct {
	timestamp int64
	workerID int64
	processID int64
	increment int64
	binary string
}

func (s *Snowflake) Generate() (string, error) {
	now := time.Now().UnixNano() / 1000000;
	if s.INCREMENT >= 4095 {
		s.INCREMENT = 0;
	}
	difference_bin := StringPadStart(strconv.FormatInt(
			(now - s.EPOCH),
			2,
		),
		42,
		"0",
	) 
	increment_bin := StringPadStart(strconv.FormatInt(
		s.INCREMENT,
		2,
	), 12, "0");
	s.INCREMENT++;
	bin := difference_bin + "0000100000" + increment_bin;
	return BinaryToID(bin);
}

func (s *Snowflake) Deconstruct(id string) (DeconstructedSnowflake, error) {
	binary, err := IDToBinary(id);
	if err != nil {
		return DeconstructedSnowflake {}, errors.New("unable to parse binary from ID");
	}
	timestamp, err := strconv.ParseInt(binary[0:41], 2, 64);
	if err != nil {
		return DeconstructedSnowflake {}, errors.New("unable to parse timestamp");
	}
	workerID, err := strconv.ParseInt(binary[41:46], 2, 64);
	if err != nil {
		return DeconstructedSnowflake {}, errors.New("unable to parse workerID");
	}
	processID, err := strconv.ParseInt(binary[46:51], 2, 64);
	if err != nil {
		return DeconstructedSnowflake {}, errors.New("unable to parse processID");
	}
	increment, err := strconv.ParseInt(binary[51:63], 2, 64);
	if err != nil {
		return DeconstructedSnowflake {}, errors.New("unable to parse increment");
	}
	deconstructed := DeconstructedSnowflake {
		timestamp: timestamp,
		workerID: workerID,
		processID: processID,
		increment: increment,
		binary: binary,
	}
	return deconstructed, nil;
}

func (s *Snowflake) IsSnowflake(id string) bool {
	deconstructed, err := s.Deconstruct(id);
	if err != nil {
		return false;
	}
	return deconstructed.timestamp > s.EPOCH && deconstructed.timestamp <= 3619093655551;
}
