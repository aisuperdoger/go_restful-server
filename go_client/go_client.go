package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"

	paddle_ocr_client "example.com/go_client/paddle_ocr"
	proto_grpc "example.com/go_client/paddle_ocr/proto_grpc"
)

// 定义一个名为PROXY的常量，用于指定代理服务器的地址
const PROXY = "http://10.50.1.124:8080/"

// json定义
type PaddleOcrResultJson struct {
	URL    string `json:"url"`
	Data   string `json:"data"`
	Status string `json:"status"`
}

func paddleOcr(c *gin.Context) { // gin.Context是Gin最重要的部分。它携带请求详细信息、验证和序列化 JSON 等。

	// 构建PaddleOcrResultJson变量
	var paddle_ocr_result_json PaddleOcrResultJson
	paddle_ocr_result_json.URL = ""
	paddle_ocr_result_json.Data = ""
	paddle_ocr_result_json.Status = ""

	// 获取上传的图片
	file, _ := c.FormFile("file")
	file_content, _ := file.Open()
	input_image, err := ioutil.ReadAll(file_content) // 将file_content转换为二进制
	if err != nil {
		log.Printf("failed to read image file: %v", err)
		paddle_ocr_result_json.Status = "failed to read image file"
		c.IndentedJSON(http.StatusOK, paddle_ocr_result_json)
	}
	defer file_content.Close()

	// dst := "./" + file.Filename
	// 上传文件至指定的完整文件路径
	// c.SaveUploadedFile(file, "./image.jpg")

	// 进行ocr检测
	var resp *proto_grpc.ImageText
	resp, err = paddle_ocr_client.PaddleOcrClient(input_image)
	if err != nil {
		log.Printf("rpc failed: %v", err)
		paddle_ocr_result_json.Status = "failed, please try again."
		c.IndentedJSON(http.StatusOK, paddle_ocr_result_json)
	}

	// 保存resp.Image
	var result_image_path = "results/paddle_ocr/result_image.jpg"
	err = ioutil.WriteFile(result_image_path, resp.Image, 0644)
	if err != nil {
		log.Printf("failed to save image file: %v", err)
		paddle_ocr_result_json.Status = "failed to save image file"
		c.IndentedJSON(http.StatusOK, paddle_ocr_result_json)
	}

	// 返回结果
	paddle_ocr_result_json.URL = PROXY + result_image_path
	paddle_ocr_result_json.Data = resp.Text
	paddle_ocr_result_json.Status = "success"

	c.IndentedJSON(http.StatusOK, paddle_ocr_result_json) // 也可以使用c.JSON，只不过c.IndentedJSON有缩进，更具可读性。
}

func main() {
	router := gin.Default()
	router.POST("/api/paddleOcr", paddleOcr) // multipart forms的内存限制默认是 32 MiB
	router.Static("/results", "./results")
	router.Run("0.0.0.0:8080")
}
