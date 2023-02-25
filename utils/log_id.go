package utils

import (
	"encoding/hex"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	// 目前版本为 02
	version    = "02"
	length     = 53
	maxRandNum = 1<<24 - 1<<20
)

// LogID represents a logID generator
type LogID struct{}

// NewLogID create a new LogID instance
func NewLogID() LogID {
	return LogID{}
}

// GenLogID return a new logID string
func (l LogID) GenLogID() string {
	rand.Uint32()
	r := randUint32n(maxRandNum) + 1<<20
	sb := strings.Builder{}
	sb.Grow(length)
	sb.WriteString(version)
	sb.WriteString(strconv.FormatUint(uint64(getMSTimestamp()), 10))
	sb.Write(localIP)
	sb.WriteString(strconv.FormatUint(uint64(r), 16))
	return sb.String()
}

var defaultLogID LogID

func init() {
	defaultLogID = NewLogID()
}

// GenLogID return a new logID
func GenLogID() string {
	return defaultLogID.GenLogID()
}

func randUint32n(n uint32) uint32 {
	return rand.Uint32() % n
}

const (
	// IPUnknown represents unknown ip
	// 32 * 0
	IPUnknown = "00000000000000000000000000000000"
)

var localIP []byte

func init() {
	localIP = formatIP(getLocalIp())
}

// getMSTimestamp return the millisecond timestamp
func getMSTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func getLocalIp() net.IP {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP
			}
		}
	}
	return nil
}

func formatIP(ip net.IP) []byte {
	if ip == nil {
		return []byte(IPUnknown)
	}
	dst := make([]byte, 32)
	i := ip.To16()
	hex.Encode(dst, i)
	return dst
}
