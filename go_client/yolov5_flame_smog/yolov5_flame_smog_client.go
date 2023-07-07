package paddle_ocr

import (
	"context"
	"io/ioutil"
	"net/http"
	configs "example.com/go_client/configs"
	proto_grpc "example.com/go_client/yolov5_flame_smog/proto_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"github.com/gin-gonic/gin"
	utils "example.com/go_client/utils"
	"path/filepath"
	"strings"
)

func yolov5FlameSmogImageRPC(input_image []byte) (*proto_grpc.Image, error) {
	// 判断输入是否为图片

	// 创建客户端连接
	conn, err := grpc.Dial("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("failed to connect: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := proto_grpc.NewYolov5FlameSmogServiceClient(conn)                                       // 创建服务客户端
	resp, err := client.Yolov5FlameSmogImage(context.Background(), &proto_grpc.Image{Image: input_image}) // 调用服务并获取返回值

	if err != nil {
		log.Printf("failed to detect image: %v", err) // 记录错误日志但不终止程序
		// log.Fatalf("failed to detect image: %v", err)  // 记录错误日志并终止程序
		// fmt.Println("Recovered from panic:", err)
	}

	// 在这里处理返回的图像数据
	// fmt.Printf("Received image data with size: %d\n", len(resp.Text))

	// 保存resp.Image
	// err = ioutil.WriteFile("./result_image.jpg", resp.Image, 0644)
	// if err != nil {
	// 	log.Fatalf("failed to save image file: %v", err)
	// }

	return resp, err
}


func yolov5FlameSmogVideoRPC(file_path string) (*proto_grpc.VideoPath, error) {
	// 判断输入是否为图片

	// 创建客户端连接
	conn, err := grpc.Dial("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("failed to connect: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := proto_grpc.NewYolov5FlameSmogServiceClient(conn)                                       // 创建服务客户端
	resp, err := client.Yolov5FlameSmogVideo(context.Background(), &proto_grpc.VideoPath{Path: file_path}) // 调用服务并获取返回值

	if err != nil {
		log.Printf("failed to detect image: %v", err) // 记录错误日志但不终止程序
		// log.Fatalf("failed to detect image: %v", err)  // 记录错误日志并终止程序
		// fmt.Println("Recovered from panic:", err)
	}

	// 在这里处理返回的图像数据
	// fmt.Printf("Received image data with size: %d\n", len(resp.Text))

	// 保存resp.Image
	// err = ioutil.WriteFile("./result_image.jpg", resp.Image, 0644)
	// if err != nil {
	// 	log.Fatalf("failed to save image file: %v", err)
	// }

	return resp, err
}


// json定义
type Yolov5FlameSmogResultJson struct {
	URL    string `json:"url"`
	// Data   string `json:"data"`
	Status string `json:"status"`
}

func Yolov5FlameSmogImage(c *gin.Context) { // gin.Context是Gin最重要的部分。它携带请求详细信息、验证和序列化 JSON 等。

	// 构建PaddleOcrResultJson变量
	var result_json Yolov5FlameSmogResultJson
	result_json.URL = ""
	result_json.Status = ""

	// 获取上传的图片
	file, _ := c.FormFile("file")
	file_content, _ := file.Open()
	input_image, err := ioutil.ReadAll(file_content) // 将file_content转换为二进制
	if err != nil {
		log.Printf("failed to read image file: %v", err)
		result_json.Status = "failed to read image file"
		c.IndentedJSON(http.StatusOK, result_json)
	}
	defer file_content.Close()

	// dst := "./" + file.Filename
	// 上传文件至指定的完整文件路径
	// c.SaveUploadedFile(file, "./image.jpg")

	// 进行ocr检测
	var resp *proto_grpc.Image
	resp, err = yolov5FlameSmogImageRPC(input_image)
	if err != nil {
		log.Printf("rpc failed: %v", err)
		result_json.Status = "failed, please try again."
		c.IndentedJSON(http.StatusOK, result_json)
	}

	// 保存resp.Image
	result_image_file := filepath.Join("yolov5_flame_smog_images", utils.GetUniqueName()+".jpg")
	result_image_path :=filepath.Join(configs.RESULT_PATH,result_image_file)
	// log.Printf(configs.RESULT_PATH)
	// log.Printf(result_image_path)
	// result_image_path := "results/paddle_ocr/result_image.jpg"
	// err := os.MkdirAll(fold_result, os.ModePerm) // 创建多级目录
	// if err != nil {
	// 	c.String(http.StatusInternalServerError, "Could not create file")
	// 	return
	// }
	err = ioutil.WriteFile(result_image_path, resp.Image, 0644)
	if err != nil {
		log.Printf("failed to save image file: %v", err)
		result_json.Status = "failed to save image file"
		c.IndentedJSON(http.StatusOK, result_json)
	}

	// 返回结果
	result_json.URL = configs.PROXY +configs.RESULT_URL + result_image_file
	result_json.Status = "success"
	// log.Printf("dsd %v", result_json.URL )
	c.IndentedJSON(http.StatusOK, result_json) // 也可以使用c.JSON，只不过c.IndentedJSON有缩进，更具可读性。
}

// func Yolov5FlameSmogVideo2(c *gin.Context) { // gin.Context是Gin最重要的部分。它携带请求详细信息、验证和序列化 JSON 等。

// 	// 构建PaddleOcrResultJson变量
// 	var result_json Yolov5FlameSmogResultJson
// 	result_json.URL = ""
// 	result_json.Status = ""

// 	// 获取上传的图片
// 	file, _ := c.FormFile("file")
// 	file_content, _ := file.Open()
// 	input_image, err := ioutil.ReadAll(file_content) // 将file_content转换为二进制
// 	if err != nil {
// 		log.Printf("failed to read image file: %v", err)
// 		result_json.Status = "failed to read image file"
// 		c.IndentedJSON(http.StatusOK, result_json)
// 	}
// 	defer file_content.Close()

// 	// dst := "./" + file.Filename
// 	// 上传文件至指定的完整文件路径
// 	// c.SaveUploadedFile(file, "./image.jpg")

// 	// 进行ocr检测
// 	var resp *proto_grpc.Image
// 	resp, err = yolov5FlameSmogImageRPC(input_image)
// 	if err != nil {
// 		log.Printf("rpc failed: %v", err)
// 		result_json.Status = "failed, please try again."
// 		c.IndentedJSON(http.StatusOK, result_json)
// 	}

// 	// 保存resp.Image
// 	result_image_file := filepath.Join("yolov5_flame_smog_images", utils.GetUniqueName()+".jpg")
// 	result_image_path :=filepath.Join(configs.RESULT_PATH,result_image_file)
// 	// log.Printf(configs.RESULT_PATH)
// 	// log.Printf(result_image_path)
// 	// result_image_path := "results/paddle_ocr/result_image.jpg"
// 	// err := os.MkdirAll(fold_result, os.ModePerm) // 创建多级目录
// 	// if err != nil {
// 	// 	c.String(http.StatusInternalServerError, "Could not create file")
// 	// 	return
// 	// }
// 	err = ioutil.WriteFile(result_image_path, resp.Image, 0644)
// 	if err != nil {
// 		log.Printf("failed to save image file: %v", err)
// 		result_json.Status = "failed to save image file"
// 		c.IndentedJSON(http.StatusOK, result_json)
// 	}

// 	// 返回结果
// 	result_json.URL = configs.PROXY +configs.RESULT_URL + result_image_file
// 	result_json.Status = "success"
// 	// log.Printf("dsd %v", result_json.URL )
// 	c.IndentedJSON(http.StatusOK, result_json) // 也可以使用c.JSON，只不过c.IndentedJSON有缩进，更具可读性。
// }


func Yolov5FlameSmogVideo(c *gin.Context) { // gin.Context是Gin最重要的部分。它携带请求详细信息、验证和序列化 JSON 等。
	result_file := filepath.Join("yolov5_flame_smog_videos", utils.GetUniqueName()+".mp4")
	result_path := filepath.Join(configs.RESULT_PATH,result_file)
	var result_json Yolov5FlameSmogResultJson
	result_json.URL = ""
	result_json.Status = ""

	// 接收数据
	utils.ReceiveFile(c, result_path)

	log.Printf("start to test yolov5_flame_smog")
	resp, err := yolov5FlameSmogVideoRPC(result_path)
	if err != nil {
		log.Printf("failed to  test yolov5_flame_smog")
		result_json.Status = "failed to  test yolov5_flame_smog"
		c.IndentedJSON(http.StatusOK, result_json)
	}

	// 返回结果
	// 取出 resp.Path中以斜杆作为分隔符的最后两部分
	split := strings.Split(resp.Path, "/") 
	result_json.URL = configs.PROXY +configs.RESULT_URL + split[len(split)-2] + "/" + split[len(split)-1]  
	result_json.Status = "success"
	log.Printf("url: %v", result_json.URL )
	c.IndentedJSON(http.StatusOK, result_json) // 也可以使用c.JSON，只不过c.IndentedJSON有缩进，更具可读性。
	// c.String(http.StatusOK, "File uploaded successfully")
}