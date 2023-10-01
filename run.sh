# conda_install_path=/home/zwl/software/anaconda3
# source ${conda_install_path}/etc/profile.d/conda.sh

cd go_client/
go run .


# 算法运行
cd python_server/paddle_ocr/
conda activate paddle_ocr
python3 paddle_ocr_service.py

cd python_server/yolov5_flame_smog/
conda activate yolov5_flame_smog
python yolov5_flame_smog_service.py 

cd python_server/yolov5_parking_violation
conda activate  yolov5_parking_violation 
python yolov5_parking_violation.py 

cd python_server/yolov5_leakage
conda activate yolov5_flame_smog
python yolov5_leakage.py