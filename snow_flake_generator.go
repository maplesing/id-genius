package id_genius

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	// Sunday, January 1, 2023 12:00:00 AM
	startEpoch int64 = 1672531200000

	datacenterIdBits int64 = 5
	workerIdBits     int64 = 5
	sequenceBits     int64 = 12

	maxDatacenterId int64 = -1 ^ (-1 << datacenterIdBits)
	maxWorkerId     int64 = -1 ^ (-1 << workerIdBits)

	workerIdShift     int64 = sequenceBits
	datacenterIdShift int64 = sequenceBits + workerIdBits
	timestampShift    int64 = sequenceBits + workerIdBits + datacenterIdBits

	sequenceMask int64 = -1 ^ (-1 << sequenceBits)

	lastTimestamp int64 = -1

	sequence int64 = 0

	mutex sync.Mutex
)

type snowFlakeGenerator struct {
	datacenterId int64
	workerId     int64
}

var generator *snowFlakeGenerator
var once sync.Once

func getSnowFlakeGenerator(datacenterId int64, workerId int64) (*snowFlakeGenerator, error) {
	if datacenterId < 0 || datacenterId > maxDatacenterId {
		return nil, fmt.Errorf("datacenterId should be great than or equal to 0 and less than %v", maxDatacenterId)
	}

	if workerId < 0 || workerId > maxWorkerId {
		return nil, fmt.Errorf("workerId should be great than or equal to 0 and less than %v", maxWorkerId)
	}

	once.Do(func() {
		generator = &snowFlakeGenerator{
			datacenterId: datacenterId,
			workerId:     workerId,
		}
	})

	return generator, nil
}

func (s *snowFlakeGenerator) GenerateNewId() (string, error) {
	newId, err := s.generateNewIntId()
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(newId, 10), nil
}

func (s *snowFlakeGenerator) generateNewIntId() (int64, error) {
	mutex.Lock()
	defer mutex.Unlock()

	curTimestamp := time.Now().UnixMilli()

	if curTimestamp < lastTimestamp {
		return 0, fmt.Errorf("current time %v should be >= last time %v", curTimestamp, lastTimestamp)
	}

	if curTimestamp == lastTimestamp {
		sequence = (sequence + 1) & sequenceMask
		if sequence == 0 {
			for curTimestamp <= lastTimestamp {
				curTimestamp = time.Now().UnixMilli()
			}
		}
	} else {
		sequence = 0
	}

	lastTimestamp = curTimestamp

	newId := ((curTimestamp - startEpoch) << timestampShift) |
		(s.datacenterId << datacenterIdShift) |
		(s.workerId << workerIdShift) |
		sequence

	return newId, nil
}
