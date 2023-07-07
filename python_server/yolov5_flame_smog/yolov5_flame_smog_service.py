
import sys

sys.path.append('./algorithm/yolov5_fs')
import grpc
from concurrent import futures
from proto_grpc import yolov5_flame_smog_pb2_grpc, yolov5_flame_smog_pb2
from grpc_reflection.v1alpha import reflection
import PIL,io
import numpy as np
# import algorithm.test_all.code.Test as Test
import algorithm.yolov5_fs.detect_fire_smoke_imag1_video as detect_fire_smoke_imag1_video


class Yolov5FlameSmogService(yolov5_flame_smog_pb2_grpc.Yolov5FlameSmogServiceServicer):
    def yolov5FlameSmogImage(self, request, context):
        
        # 将二进制image_text.image转化PIL.Image
        pil_image = PIL.Image.open(io.BytesIO(request.image)).convert("RGB")

        # 将PIL.Image转化为ndarray，因为ocr需要ndarray格式作为输入
        array_image = np.array(pil_image)
        
        result_image = detect_fire_smoke_imag1_video.detect_image(array_image)
        result_image = PIL.Image.fromarray(result_image)
     
        # 将result_image转化二进制
        binary_image = io.BytesIO()
        result_image.save(binary_image, format="JPEG")
        
        result = yolov5_flame_smog_pb2.Image()
        result.image = binary_image.getvalue()
    
        return result 

    def yolov5FlameSmogVideo(self, request, context):
        result = yolov5_flame_smog_pb2.VideoPath()
        # result.path = request.path[:-4] + '_result' + request.path[-4:]
        result.path = request.path[:-4] + '_result' + ".webm"
        detect_fire_smoke_imag1_video.detect_video(request.path, result.path)
        # detect_fire_smoke_imag1_video.detect_video2(request.path,detect_fire_smoke_imag1_video.detect_image, result.path)

        return result 


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=20))
    yolov5_flame_smog_pb2_grpc.add_Yolov5FlameSmogServiceServicer_to_server(Yolov5FlameSmogService(), server)
    # 使用grpcui测试时，需要的代码
    SERVICE_NAMES = (
        yolov5_flame_smog_pb2.DESCRIPTOR.services_by_name['Yolov5FlameSmogService'].full_name,
        reflection.SERVICE_NAME,
    )
    reflection.enable_server_reflection(SERVICE_NAMES, server)
    
    server.add_insecure_port('[::]:50053')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
