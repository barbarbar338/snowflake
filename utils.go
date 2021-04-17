package snowflake

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func StringPadStart(input string, padLength int, padString string) string {
	var output string;
	inputLength := len(input);
	padStringLength := len(padString);
	if inputLength >= padLength {
		return input;
	}
	repeat := math.Ceil(float64(1) + (float64(padLength-padStringLength))/float64(padStringLength));
	output = strings.Repeat(padString, int(repeat)) + input;
	output = output[len(output)-padLength:];
	return output;
}

func BinaryToID(binary string) (string, error) {
	id := "";
	for len(binary) > 50 {
		high, err := strconv.ParseInt(binary[0:len(binary) - 32], 2, 64);
		if err != nil {
			return "", errors.New("unable to parse binary");
		}
		low, err := strconv.ParseInt(
			strconv.FormatInt((high % 10), 2) + binary[len(binary)-32:],
			2,
			64,
		);
		if err != nil {
			return "", errors.New("unable to parse binary");
		}
		id = fmt.Sprintf("%v%v", (low % 10), id);
		binary = strconv.FormatInt((high / 10), 2) + StringPadStart(strconv.FormatInt((low / 10), 2), 32, "0");
	}
	parsed, err := strconv.ParseInt(
		binary,
		2,
		64,
	);
	if err != nil {
		return "", errors.New("unable to parse binary");
	}
	for parsed > 0 {
		id = fmt.Sprintf("%v%v", (parsed % 10), id);
		parsed = parsed / 10;
	}
	return id, nil;
}

func IDToBinary(id string) (string, error) {
	binary := "";
	high, err := strconv.ParseInt(id[0:len(id) - 10], 0, 64);
	if err != nil {
		high = 0;
	}
	low, err := strconv.ParseInt(id[len(id)-10:], 0, 64);
	if err != nil {
		return "", errors.New("unable to parse binary");
	}
	for low > 0 || high > 0 {
		binary = fmt.Sprintf("%v%v", low & 1, binary);
		low = low / 2;
		if (high > 0) {
			low += 5000000000 * (high % 2);
			high = high / 2;
		}
	}
	return binary, nil;
}
