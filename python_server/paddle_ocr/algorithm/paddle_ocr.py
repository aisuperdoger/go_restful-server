import os
import PIL
from paddleocr import PaddleOCR, draw_ocr



def RunOcr(input_image):
    # 进行ocr识别
    ocr = PaddleOCR(use_angle_cls=True, lang="ch")
    result = ocr.ocr(input_image, cls=True)
    
    boxes = [line[0] for line in result]
    txts = [line[1][0] for line in result]
    scores = [line[1][1] for line in result]
    
    # 框出图上文字     
    result_image = draw_ocr(input_image, boxes, None, None, font_path= os.path.join(os.path.dirname(__file__) , "simfang.ttf"))
    result_image = PIL.Image.fromarray(result_image)
    
    # 保存图片
    # im_show.save('ocrResult.jpg')

    # 获取文本结果
    result_text = ""
    for line in txts:   
        result_text = result_text+ line + "\n"

    return result_text,result_image


# if __name__ == "__main__":
#      RunOcr(image)
