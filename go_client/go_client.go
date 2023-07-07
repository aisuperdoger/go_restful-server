package main

import (
	
	paddle_ocr_client "example.com/go_client/paddle_ocr"
	// yolov5_classify_life_jacket_client "example.com/go_client/yolov5_classify_life_jacket"
	yolov5_flame_smog_client "example.com/go_client/yolov5_flame_smog"
	yolov5_parking_violation_client "example.com/go_client/yolov5_parking_violation"
	yolov5_leakage_client "example.com/go_client/yolov5_leakage"
	"github.com/gin-gonic/gin"
	configs "example.com/go_client/configs"
)



func main() {// 先不修改C++服务器中的代码  go服务器使用8080端口
	router := gin.Default()
	router.POST("/api2/yolov5FlameSmogImage", yolov5_flame_smog_client.Yolov5FlameSmogImage)         
	router.POST("/api2/yolov5FlameSmogVideo", yolov5_flame_smog_client.Yolov5FlameSmogVideo)  
	router.POST("/api2/yolov5ParkingViolation", yolov5_parking_violation_client.Yolov5ParkingViolation)
	router.POST("/api2/oilDetection", yolov5_leakage_client.RpcRequestOilDetection)
	router.POST("/api2/heatAreaDetection", yolov5_leakage_client.RpcRequestHeatAreaDetection)  
	router.POST("/api2/paddleOcr", paddle_ocr_client.PaddleOcr)                                       // multipart forms的内存限制默认是 32 MiB
	// router.POST("/api2/testYolov5ClassifyLifeJacket", yolov5_classify_life_jacket_client.TestYolov5ClassifyLifeJacket) // multipart forms的内存限制默认是 32 MiB
	router.Static("/results", configs.RESULT_PATH)
	router.Run(":8080")
}
