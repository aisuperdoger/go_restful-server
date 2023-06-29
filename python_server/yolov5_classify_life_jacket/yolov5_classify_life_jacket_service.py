import grpc
from concurrent import futures
from proto_grpc import yolov5_classify_life_jacket_pb2_grpc, yolov5_classify_life_jacket_pb2
from grpc_reflection.v1alpha import reflection
import PIL,io
import numpy as np
import algorithm.test_all.code.Test as Test


class TestYolov5ClassifyLifeJacketService(yolov5_classify_life_jacket_pb2_grpc.TestYolov5ClassifyLifeJacketServiceServicer):
    def testYolov5ClassifyLifeJacket(self, request, context):
        
        result = yolov5_classify_life_jacket_pb2.ResultJson()
        result.data = Test.RunTest(request.file_path, request.json_path)
    
        return result


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=20))
    yolov5_classify_life_jacket_pb2_grpc.add_TestYolov5ClassifyLifeJacketServiceServicer_to_server(TestYolov5ClassifyLifeJacketService(), server)
    
    # 使用grpcui测试时，需要的代码
    SERVICE_NAMES = (
        yolov5_classify_life_jacket_pb2.DESCRIPTOR.services_by_name['TestYolov5ClassifyLifeJacketService'].full_name,
        reflection.SERVICE_NAME,
    )
    reflection.enable_server_reflection(SERVICE_NAMES, server)
    
    server.add_insecure_port('[::]:50052')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
