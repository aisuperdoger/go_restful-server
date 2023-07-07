package paddle_ocr

import (
	"context"
	proto_grpc "example.com/go_client/yolov5_classify_life_jacket/proto_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"github.com/gin-gonic/gin"
	"github.com/mholt/archiver/v3"
	utils "example.com/go_client/utils"
)

func TestYolov5ClassifyLifeJacketClient(file_path string, json_path string) (*proto_grpc.ResultJson, error) {

	// 创建客户端连接
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("failed to connect: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := proto_grpc.NewTestYolov5ClassifyLifeJacketServiceClient(conn)
	resp, err := client.TestYolov5ClassifyLifeJacket(context.Background(), &proto_grpc.DatasetPath{FilePath: file_path, JsonPath: json_path}) // 调用服务并获取返回值

	if err != nil {
		log.Printf("failed to detect image: %v", err) // 记录错误日志但不终止程序
		// log.Fatalf("failed to detect image: %v", err)  // 记录错误日志并终止程序
		// fmt.Println("Recovered from panic:", err)
	}

	return resp, err
}



type testYolov5ClassifyLifeJacketResultJson struct {
	Data string `json:"data"`
}



func TestYolov5ClassifyLifeJacket(c *gin.Context) { // gin.Context是Gin最重要的部分。它携带请求详细信息、验证和序列化 JSON 等。
	fold_result := filepath.Join("../results/yolov5_classify_life_jacket/", utils.GetUniqueName())
	zip_file_path := filepath.Join(fold_result, "life_jacket.zip")
	unzip_file_path := filepath.Join(fold_result, "life_jacket/")

	err := os.MkdirAll(fold_result, os.ModePerm) // 创建多级目录
	if err != nil {
		c.String(http.StatusInternalServerError, "Could not create file")
		return
	}
	out, err := os.Create(zip_file_path)
	if err != nil {
		c.String(http.StatusInternalServerError, "Could not create file")
		return
	}
	defer out.Close()

	/************************************************** 数据集读取 **************************************************/
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
	log.Printf("get zip file success")
	//解压文件
	err = archiver.Unarchive(zip_file_path, unzip_file_path)
	if err != nil {
		log.Printf("failed to unzip : %v", err)
	}
	log.Printf("unzip file success")

	/************************************************** 进行yolov5_classify_life_jacket检测 **************************************************/
	log.Printf("start to test yolov5_classify_life_jacket")
	resp, err := TestYolov5ClassifyLifeJacketClient(filepath.Join(unzip_file_path, "images"), filepath.Join(unzip_file_path, "jsons"))

	// 删除life_jacket.zip文件和解压文件夹  可以优化成定时删除，或者网页刷新时删除
	// 下面代码解压速度慢
	// err = os.Remove(zipFilePath)
	// if err != nil {
	// 	log.Printf("failed to remove file : %v", err)
	// }
	// err = os.RemoveAll(unzipFilePath)
	// if err != nil {
	// 	log.Printf("failed to remove file : %v", err)
	// }

	// cmd := exec.Command("rm", zipFilePath)
	// err = cmd.Run()
	// if err != nil {
	// 	log.Printf("failed to remove file : %v", err)
	// }

	// cmd = exec.Command("rm", "-r", unzipFilePath)
	// err = cmd.Run()
	// if err != nil {
	// 	log.Printf("failed to remove file : %v", err)
	// }
	var test_yolov5_classify_life_jacket_result_json testYolov5ClassifyLifeJacketResultJson
	test_yolov5_classify_life_jacket_result_json.Data = resp.Data

	c.IndentedJSON(http.StatusOK, test_yolov5_classify_life_jacket_result_json)
	// c.String(http.StatusOK, "File uploaded successfully")
}
