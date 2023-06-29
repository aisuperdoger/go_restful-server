from torch.utils.data import Dataset
import os
import json
import cv2
import torch
from letterbox import letterbox
import math
import numpy as np
class loadData(Dataset):
    ##########测试用
    def __init__(self,file_path,json_path,type,pho_pathList):
        super(loadData,self).__init__()
        self.file_path = file_path
        self.json_path = json_path
        self.pho_pathList = pho_pathList
        self.type = type

    #############训练用
    # def __init__(self,file_path,pos_file_path,neg_file_path,json_path,type):
    #     super(loadData,self).__init__()
    #     self.pos_file_path = pos_file_path
    #     self.neg_file_path = neg_file_path
    #     self.file_path = file_path
    #     self.json_path = json_path

    #     self.type = type
    def __len__(self):
        return len(self.pho_pathList)
    def __getitem__(self, item):
        photopath = self.pho_pathList[item]
        photojson = self.json_path + '/' + str(photopath[:-3]) + 'json'
        json_data = json.load(open(photojson, 'r', encoding='utf-8'))
        img = cv2.imread(self.file_path + '/' + photopath)
        data = json_data.get('label')


        # save_path = '/home/my/Documents/obj_detect/fjf/bad_img_test_lenet'
        # if img.shape[0]<50 or img.shape[1]<30:
        #     save_path_file = os.path.join(save_path, photopath)
        #     # print(save_path_file)
        #     cv2.imwrite(save_path_file, img)

        # img = cv2.cvtColor(img, cv2.COLOR_BGR2RGB)

        # save_path = '/home/my/Documents/obj_detect/fjf/bad_img_test_lenet'
        # if img.shape[0]<50 or img.shape[1]<30:
        #     save_path_file = os.path.join(save_path, photopath)
        #     # print(save_path_file)
        #     cv2.imwrite(save_path_file, img)


        #####图像标准化
        # mean = np.array([0.485,0.456,0.406])
        # std = np.array([0.229,0.224,0.225])
        # img = (img/255 - mean)/std
        #img = resize_image(img,228,114,photopath)

        #img = cv2.resize(img,(114,228),interpolation = cv2.INTER_LINEAR)
        #image_size = img.shape
        #if image_size[0]/image_size[1]>2 and image_size[0]/image_size[1]<=3:
        #    img = img[:math.floor(image_size[0]*3/4),:,:]
        #elif image_size[0]/image_size[1]>3 and image_size[0]/image_size[1]<=4:
        #    img = img[:math.floor(image_size[0]*2/3),:,:]
        #elif image_size[0] / image_size[1] > 4:
        #    img = img[:math.floor(image_size[0] * 1/2), :, :]
        #else:
        #    pass
        #img_size = img.shape
        #if img_size[0]/img_size[1]>=2:
        #    img = img[math.floor(img_size[0]/6):,:,:]
        # img = img /255
        img = letterbox(img,(228,114))
        if self.type == 'test':

            if data == "life-jacket":
                return img,torch.Tensor([1]),photopath
            else:
                return img,torch.Tensor([0]),photopath
        else:
            if data == "life-jacket":
                return img,torch.Tensor([1])
            else:
                return img,torch.Tensor([0])
