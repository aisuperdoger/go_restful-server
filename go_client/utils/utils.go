package utils

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/gin-gonic/gin"
	"net/http"
	"io"
	"os"
	"log"
)

func GetUniqueName() string {
	timestamp := time.Now().Unix()

	// 生成4位随机数
	rand.Seed(time.Now().UnixNano()) // 可以再做复杂一点
	// random := rand.Intn(10000)

	// 组合时间戳和随机数作为文件夹名称
	// folderName := fmt.Sprintf("%d-%04d", timestamp, random)
	folderName := fmt.Sprintf("%d", timestamp) // 为了更好调试，这里暂时只使用时间戳

	return folderName
}


func ReceiveFile(c *gin.Context, result_path string){
	out, err := os.Create(result_path)
	if err != nil {
		c.String(http.StatusInternalServerError, "Could not create file")
		return
	}
	defer out.Close()

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "Could not get file")
		return
	}
	defer file.Close()

	fileSize := header.Size
	chunkSize := 64 * 1024 // 64 KB

	// Read file in chunks
	var offset int64 = 0
	for offset < fileSize {
		endOffset := offset + int64(chunkSize)
		if endOffset > fileSize {
			endOffset = fileSize
		}
		chunk := make([]byte, endOffset-offset)
		_, err := file.Read(chunk)
		if err != nil && err != io.EOF {
			c.String(http.StatusInternalServerError, "Error reading file")
			return
		}
		_, err = out.Write(chunk)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error writing file")
			return
		}
		offset = endOffset
	}
	log.Printf("get file success")


}