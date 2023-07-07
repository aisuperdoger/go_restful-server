import os

def delete_files_in_folder(folder_path):
    for root, dirs, files in os.walk(folder_path):
        for file in files:
            file_path = os.path.join(root, file)
            os.remove(file_path)

# 用法示例：
folder_path = 'results'
delete_files_in_folder(folder_path)