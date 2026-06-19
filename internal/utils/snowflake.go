package utils

import (
	"fmt"
	"sync"
	"time"
	"video_feed/config"
)

const (
	epoch          = 1609459200000 // 2021-01-01 00:00:00 UTC
	machineBits    = 10
	sequenceBits   = 12
	maxMachineID   = -1 ^ (-1 << machineBits) // 1023
	maxSequence    = -1 ^ (-1 << sequenceBits) // 4095
	machineShift   = sequenceBits
	timestampShift = sequenceBits + machineBits
)

type Snowflake struct {
	mu        sync.Mutex
	machineID int64
	sequence  int64
	lastTime  int64
}

var (
	defaultSnowflake *Snowflake
	once             sync.Once
)

// InitSnowflake 初始化全局Snowflake实例
func InitSnowflake() error {
	var initErr error
	once.Do(func() {
		cfg := config.GetConfig()
		if cfg == nil {
			initErr = fmt.Errorf("config not loaded")
			return
		}
		defaultSnowflake, initErr = NewSnowflake(cfg.Snowflake.MachineID)
	})
	return initErr
}

// GetSnowflake 获取全局Snowflake实例
func GetSnowflake() *Snowflake {
	return defaultSnowflake
}

func NewSnowflake(machineID int64) (*Snowflake, error) {
	if machineID < 0 || machineID > maxMachineID {
		return nil, fmt.Errorf("machineID must be between 0 and %d", maxMachineID)
	}
	return &Snowflake{machineID: machineID}, nil
}

func (s *Snowflake) NextID() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().UnixMilli() - epoch

	if now < s.lastTime {
		panic("clock moved backwards")
	}

	if now == s.lastTime {
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			now = s.waitNextMilli()
		}
	} else {
		s.sequence = 0
	}

	s.lastTime = now

	return (now << timestampShift) | (s.machineID << machineShift) | s.sequence
}

func (s *Snowflake) waitNextMilli() int64 {
	now := time.Now().UnixMilli() - epoch
	for now <= s.lastTime {
		now = time.Now().UnixMilli() - epoch
	}
	return now
}

// GenerateID 生成分布式ID (便捷方法)
func GenerateID() int64 {
	return defaultSnowflake.NextID()
}

func ParseID(id int64) (timestamp int64, machineID int64, sequence int64) {
	sequence = id & maxSequence
	machineID = (id >> machineShift) & maxMachineID
	timestamp = (id >> timestampShift) + epoch
	return
}
