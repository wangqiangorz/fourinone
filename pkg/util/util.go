package util

import (
	"fmt"
	"hash/crc32"
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetTimeStamp() int64 {
	return time.Now().Unix()
}

func GetHostFromString(server string) (string, int) {
	host := strings.Split(server, ":")
	port, _ := strconv.Atoi(host[1])
	return host[0], port
}

func GetStringFromHost(host string, port int) string {
	return host + fmt.Sprintf(":%d", port)
}

func GetDomainNodeKey(domain, node string) string {
	if node == "" {
		return domain
	}
	return domain + "." + node
}

func GetDomainNodeFromKey(domainNodeKey string) (string, string) {
	domainNode := strings.Split(domainNodeKey, ".")
	return domainNode[0], domainNode[1]
}

func GetUUid() string {
	return uuid.NewV4().String()
}

func HashCode(key string) int {
	v := int(crc32.ChecksumIEEE([]byte(key)))
	if v > 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	return 0
}
