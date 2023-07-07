
import grpc
from concurrent import futures
from proto_grpc import paddle_ocr_pb2, paddle_ocr_pb2_grpc
import PIL,io
import numpy as np
import algorithm.paddle_ocr as paddle_ocr
from grpc_reflection.v1alpha import reflection

class PaddleOcrService(paddle_ocr_pb2_grpc.PaddleOcrServiceServicer):
    def PaddleOcr(self, request, context):
        
        # 将二进制image_text.image转化PIL.Image
        pil_image = PIL.Image.open(io.BytesIO(request.image)).convert("RGB")

        # 将PIL.Image转化为ndarray，因为ocr需要ndarray格式作为输入
        array_image = np.array(pil_image)
        
        result_text, result_image = paddle_ocr.RunOcr(array_image)
        
        # 将result_image转化二进制
        binary_image = io.BytesIO()
        result_image.save(binary_image, format="JPEG")
        
        result = paddle_ocr_pb2.ImageText()
        result.image = binary_image.getvalue()
        result.text = result_text
    
        return result 


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=20))
    paddle_ocr_pb2_grpc.add_PaddleOcrServiceServicer_to_server(PaddleOcrService(), server)
    
    SERVICE_NAMES = (
        paddle_ocr_pb2.DESCRIPTOR.services_by_name['PaddleOcrService'].full_name,
        reflection.SERVICE_NAME,
    )
    reflection.enable_server_reflection(SERVICE_NAMES, server)
    
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
