# 项目运行

需要修改的点（这可以写一个配置文件来存储）：
- 修改go_client/go_client.go中的全局变量RESULTS_PATH，RESULTS_PATH用于存储用户的请求和返回的结果。
- 修改run.sh文件中的conda_install_path换成自己anaconda3的安装路径。
运行所有算法的grpc服务：
```
bash run.sh
```


# 添加新算法的流程

添加新算法的流程（以paddle_ocr为例）：

1.在protos文件夹中创建paddle_ocr.proto文件
- 设置go的包名为`option go_package = "./paddle_ocr";`
- 使用mes方法的形参和返回值的sage定义类型
- service的名字定义为PaddleOcrService，其中的方法定义为PaddleOcr

2.在python_server文件夹中创建paddle_ocr文件夹，然后在paddle_ocr文件夹中创建algorithm和proto_grpc文件夹，以及创建paddle_ocr_service.py：
- algorithm中存放算法的代码。
- proto_grpc中存放着使用protoc生成的python代码，这些代码用于创建rpc server端。生成代码的命令如下：
```
python -m grpc_tools.protoc -I protos/ \
        --python_out=python_server/paddle_ocr/proto_grpc \
        --pyi_out=python_server/paddle_ocr/proto_grpc \
        --grpc_python_out=python_server/paddle_ocr/proto_grpc \
        paddle_ocr.proto
```
python -m grpc_tools.protoc -I protos/ \
        --python_out=python_server/yolov5_leakage/proto_grpc \
        --pyi_out=python_server/yolov5_leakage/proto_grpc \
        --grpc_python_out=python_server/yolov5_leakage/proto_grpc \
        yolov5_leakage.proto

【注】paddle_ocr_pb2_grpc.py中的代码`import paddle_ocr_pb2 as paddle__ocr__pb2`需要修改为`import proto_grpc.paddle_ocr_pb2 as paddle__ocr__pb2`，否则会发生包不存在的错误。
- paddle_ocr_service.py中使用proto_grpc中的代码来构建远程方法。paddle_ocr_service.py中使用serve()函数用于开启请求监听，定义类PaddleOcrService以及类中的成员函数PaddleOcr去调用algorithm中的接口。

3.在go_client文件夹中创建paddle_ocr文件夹，然后在paddle_ocr文件夹中创建proto_grpc文件夹paddle_ocr_client.go文件。
- proto_grpc中存放着使用protoc生成的go代码，这些代码用于创建rpc client端。生成代码的命令如下：
```
cd protos
protoc --go_out=../go_client/paddle_ocr/proto_grpc --go_opt=paths=source_relative \
   --go-grpc_out=../go_client/paddle_ocr/proto_grpc --go-grpc_opt=paths=source_relative \
   paddle_ocr.proto 
```


- paddle_ocr_client.go：创建名为PaddleOcrClient的函数，PaddleOcrClient的函数中用于连接grpc server，并向其发送远程方法调用请求。
- go_client.go：在main()中添加url对应的函数paddleOcr。
  添加json格式的数据用于作为返回值的类型。
  添加函数paddleOcr用于获取用户请求，调用PaddleOcrClient的函数，从而调用远程函数处理请求，返回处理结果。最后开启服务器监听端口。【这里可以用某些设计模式进行改进吧？go中没有类，那么有相应的设计模式？】

4.在RESULTS_PATH中添加paddle_ocr文件夹用于存储用户的请求和返回的结果。




# 改进

视频流检测，而不是上传一个视频检测一个视频。



# 杂记

如果你的视频无法在vscode上播放，只能在本地播放，那么可能视频格式本身是有问题的。


python_server中的同一个算法，如果在很多个终端中都开启了，那么就会导致rpc请求会随机向这些开启的终端中的一个发送请求。
如果A开启了算法1，B修改了算法1并开启算法1，此时B对新算法1的测试请求，可能会发送到A的旧算法1上，这就会导致测试结果不一致。













