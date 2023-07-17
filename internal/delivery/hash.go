package delivery

import (
	"crypto/sha256"
	"encoding/hex"
	"hash/crc64"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	hashResults = make(map[string]int)
	mutex       sync.Mutex
)

func CalcHash(c *gin.Context) {
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading request body"})
		return
	}
	defer c.Request.Body.Close()

	requestID := sha256.Sum256(data)

	go calculation(string(data), hex.EncodeToString(requestID[:]))

	c.JSON(http.StatusAccepted, gin.H{"request_id": hex.EncodeToString(requestID[:])})
}

func calculation(data string, requestID string) {

	timeout := time.After(1 * time.Minute)

	var tempHash int
	ticker := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-ticker.C:
			crcTable := crc64.MakeTable(crc64.ECMA)
			crc64Hash := crc64.Checksum([]byte(data), crcTable)

			timestamp := time.Now().UnixNano()
			hashResult := int(timestamp & int64(crc64Hash))
			numOnes := strings.Count(strconv.FormatInt(int64(hashResult), 2), "1")

			tempHash = numOnes

		case <-timeout:
			ticker.Stop()

			mutex.Lock()
			hashResults[requestID] = tempHash
			mutex.Unlock()

			return
		}
	}
}

func GetResult(c *gin.Context) {
	requestID := c.Param("request_id")

	mutex.Lock()
	result, ok := hashResults[requestID]
	mutex.Unlock()

	if !ok {

		c.JSON(http.StatusOK, gin.H{"status": "PENDING"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}
