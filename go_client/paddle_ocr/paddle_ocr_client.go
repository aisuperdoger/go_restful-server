package paddle_ocr

import (
	"context"
	proto_grpc "example.com/go_client/paddle_ocr/proto_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func PaddleOcrClient(input_image []byte) (*proto_grpc.ImageText, error) {
	// 判断输入是否为图片

	// 创建客户端连接
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("failed to connect: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := proto_grpc.NewPaddleOcrServiceClient(conn)                                       // 创建服务客户端
	resp, err := client.PaddleOcr(context.Background(), &proto_grpc.Image{Image: input_image}) // 调用服务并获取返回值

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
