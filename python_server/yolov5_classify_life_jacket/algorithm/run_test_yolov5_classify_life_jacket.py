import os

def run():
    # file_path = '/home/xd2/hhx2022/projects/trt/tensorrt-integrate/restful-server/src/python_scripts/yolov5_classify_life_jacket/test_all/data/all_image'
    # json_path = '/home/xd2/hhx2022/projects/trt/tensorrt-integrate/restful-server/src/python_scripts/yolov5_classify_life_jacket/test_all/data/all_json'
    file_path = "lifejacket/images"
    json_path = "lifejacket/jsons"

    os.system('/home/xd2/anaconda3/envs/yolov5_classify_life_jacket/bin/python3 ../src/python_scripts/yolov5_classify_life_jacket/test_all/code/Test.py ' +file_path +' '+json_path)

    
    