package main

import (
	"errors"
	"fmt"
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

func (snowflake *Snowflake) Generate() (string, error) {
	now := time.Now().UnixNano() / 1000000;
	if snowflake.INCREMENT >= 4095 {
		snowflake.INCREMENT = 0;
	}
	difference_bin := StringPadStart(strconv.FormatInt(
			(now - snowflake.EPOCH),
			2,
		),
		42,
		"0",
	) 
	increment_bin := StringPadStart(strconv.FormatInt(
		snowflake.INCREMENT,
		2,
	), 12, "0");
	snowflake.INCREMENT++;
	bin := difference_bin + "0000100000" + increment_bin;
	return BinaryToID(bin);
}

func (snowflake *Snowflake) Deconstruct(id string) (DeconstructedSnowflake, error) {
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

func (snowflake *Snowflake) IsSnowflake(id string) bool {
	deconstructed, err := snowflake.Deconstruct(id);
	if err != nil {
		return false;
	}
	return deconstructed.timestamp > snowflake.EPOCH && deconstructed.timestamp <= 3619093655551;
}

func main() {
	snowflake := Snowflake {
		EPOCH: 1618606800,
	}
	fmt.Println(snowflake.Deconstruct("6782465263234318336"));
}
/*
	ID: 6782465263234318336

	Binary: 101111000100000001010000011100011001100100000100000000000000000
	Timestamp: 1617065730866 
	WorkerID: 1 
	ProcessID: 0 
	Increment: 0

	●----------------------------------------------------------------------●

	  ╔ 								Binary								   ╗
	 ║10111100010000000101000001110001100110010║ ║00001║ ║00000║ ║000000000000║
	╚				Timestamp				  ╝	╚ WID ╝	╚ PID ╝	╚  Increment ╝
*/
