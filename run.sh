# conda_install_path=/home/zwl/software/anaconda3
# source ${conda_install_path}/etc/profile.d/conda.sh

# server端运行
cd go_client
go  run .



# 算法运行
cd python_server/paddle_ocr/
conda activate paddle_ocr
python3 paddle_ocr_service.py

# conda activate yolov5Env  
# python3 python_server/yolov5_classify_life_jacket/yolov5_classify_life_jacket_service.py

cd python_server/yolov5_flame_smog/
conda activate yolov5_flame_smog
python yolov5_flame_smog_service.py 


conda activate  yolov5_parking_violation 
python yolov5_parking_violation.py 