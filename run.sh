conda_install_path=/home/zwl/software/anaconda3
source ${conda_install_path}/etc/profile.d/conda.sh

conda activate paddle_ocr
python3 python_server/paddle_ocr/paddle_ocr_service.py

