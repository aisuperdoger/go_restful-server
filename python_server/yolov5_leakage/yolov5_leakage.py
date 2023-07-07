
import sys

sys.path.append('./algorithm/szx_system')
sys.path.append('../utils')
import grpc
from concurrent import futures
from proto_grpc import yolov5_leakage_pb2_grpc, yolov5_leakage_pb2
# from grpc_reflection.v1alpha import reflection
import PIL,io
import numpy as np
import algorithm.szx_system.heat_area_detection as heat_area_detection
import algorithm.szx_system.oil_detection as oil_detection
import utils as uitls 

class Yolov5LeakageService(yolov5_leakage_pb2_grpc.Yolov5LeakageServiceServicer):
    def heatAreaDetection(self, request, context):
        result = yolov5_leakage_pb2.VideoPath()
        # result.path = request.path[:-4] + '_result' + request.path[-4:]
        result.path = request.path[:-4] + '_result' + ".webm"
        uitls.detect_video(request.path,heat_area_detection.heat_area_detect, result.path)
        
        return result 
    
    def oilDetection(self, request, context):
        result = yolov5_leakage_pb2.VideoPath()
        # result.path = request.path[:-4] + '_result' + request.path[-4:]
        result.path = request.path[:-4] + '_result' + ".webm"
        uitls.detect_video(request.path,oil_detection.oil_detect, result.path)
        
        return result 



def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=20))
    yolov5_leakage_pb2_grpc.add_Yolov5LeakageServiceServicer_to_server(Yolov5LeakageService(), server)
    # 使用grpcui测试时，需要的代码
    # SERVICE_NAMES = (
    #     yolov5_parking_violation_pb2.DESCRIPTOR.services_by_name['Yolov5ParkingViolationService'].full_name,
    #     reflection.SERVICE_NAME,
    # )
    # reflection.enable_server_reflection(SERVICE_NAMES, server)
    print("等待请求...")
    server.add_insecure_port('[::]:50055')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
