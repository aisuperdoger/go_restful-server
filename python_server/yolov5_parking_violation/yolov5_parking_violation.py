
import sys

sys.path.append('./algorithm/yolov7_illegal_parking')
import grpc
from concurrent import futures
from proto_grpc import yolov5_parking_violation_pb2_grpc, yolov5_parking_violation_pb2
# from grpc_reflection.v1alpha import reflection
import PIL,io
import numpy as np
# import algorithm.test_all.code.Test as Test
import algorithm.yolov7_illegal_parking.illegal_parking_interface as detect


class Yolov5ParkingViolationService(yolov5_parking_violation_pb2_grpc.Yolov5ParkingViolationServiceServicer):
    def yolov5ParkingViolation(self, request, context):
        result = yolov5_parking_violation_pb2.VideoPath()
        # result.path = request.path[:-4] + '_result' + request.path[-4:]
        result.path = request.path[:-4] + '_result' + ".webm"
        detect.detect(request.path, result.path)
        
        return result 


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=20))
    yolov5_parking_violation_pb2_grpc.add_Yolov5ParkingViolationServiceServicer_to_server(Yolov5ParkingViolationService(), server)
    # 使用grpcui测试时，需要的代码
    # SERVICE_NAMES = (
    #     yolov5_parking_violation_pb2.DESCRIPTOR.services_by_name['Yolov5ParkingViolationService'].full_name,
    #     reflection.SERVICE_NAME,
    # )
    # reflection.enable_server_reflection(SERVICE_NAMES, server)
    print("等待请求...")
    server.add_insecure_port('[::]:50054')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
