package paddle_ocr

import (
	"context"
	"net/http"
	configs "example.com/go_client/configs"
	proto_grpc "example.com/go_client/yolov5_parking_violation/proto_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"github.com/gin-gonic/gin"
	utils "example.com/go_client/utils"
	"path/filepath"
	"strings"
)


func yolov5ParkingViolationRPC(file_path string) (*proto_grpc.VideoPath, error) {
	// 判断输入是否为图片

	// 创建客户端连接
	conn, err := grpc.Dial("localhost:50054", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("failed to connect: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := proto_grpc.NewYolov5ParkingViolationServiceClient(conn)                                       // 创建服务客户端
	resp, err := client.Yolov5ParkingViolation(context.Background(), &proto_grpc.VideoPath{Path: file_path}) // 调用服务并获取返回值

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

func Yolov5ParkingViolation(c *gin.Context) { // gin.Context是Gin最重要的部分。它携带请求详细信息、验证和序列化 JSON 等。
	result_file := filepath.Join("yolov5_parking_violation", utils.GetUniqueName()+".mp4")
	result_path := filepath.Join(configs.RESULT_PATH,result_file)
	var result_json Yolov5FlameSmogResultJson
	result_json.URL = ""
	result_json.Status = ""

	// 接收数据
	utils.ReceiveFile(c, result_path)

	log.Printf("start to test yolov5_flame_smog")
	resp, err := yolov5ParkingViolationRPC(result_path)
	if err != nil {
		log.Printf("failed to  test yolov5_flame_smog")
		result_json.Status = "failed to  test yolov5_flame_smog"
		c.IndentedJSON(http.StatusOK, result_json)
	}

	// 返回结果
	// 取出 resp.Path中以斜杆作为分隔符的最后两部分
	split := strings.Split(resp.Path, "/") 
	result_json.URL = configs.PROXY +configs.RESULT_URL + split[len(split)-2] + "/" + split[len(split)-1]  // input和output使用同一个文件名
	result_json.Status = "success"
	log.Printf("dsd %v", result_json.URL )
	c.IndentedJSON(http.StatusOK, result_json) // 也可以使用c.JSON，只不过c.IndentedJSON有缩进，更具可读性。
	// c.String(http.StatusOK, "File uploaded successfully")
}