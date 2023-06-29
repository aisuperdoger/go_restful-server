import sys
import os
CURRENT_DIR = os.path.split(os.path.abspath(__file__))[0]  # 当前目录
sys.path.append(CURRENT_DIR)

import time
import json
from LoadData import loadData
from torch.utils.data import DataLoader
import numpy as np
#from lenet import LeNet
from torch import optim
import torch
import cv2
from alexnet import AlexNet,AlexNet1,AlexNet2

#from mobile_net import MobileNetV3_Small,MobileNetV3_Large
# from torch.autograd import Variable

class Test():
    def __init__(self,file_path,json_path):
        self.file_path = file_path
        self.json_path = json_path
    def test(self,save_path):
        batch_size = 64
        load_path = None
        if os.path.exists(save_path):
            load_path = save_path
        model = AlexNet()
        # model = LeNet()
        # model = MobileNetV3_Small(num_classes=2)
        # model = MobileNetV3_Large(num_classes=2)
        if load_path != None:
            model.load_state_dict(torch.load(load_path))
            print("load_successful")
        if torch.cuda.is_available():
            device = torch.device('cuda')
            model = model.to(device)

        data_num = 0
        correct_num = 0

        pos_sample = 0
        neg_sample = 0

        right_right = 0
        right_false = 0
        false_false = 0
        false_right = 0
        start = time.time()
        pho_pathList = os.listdir(self.file_path)
        dataloader = DataLoader(loadData(self.file_path,self.json_path,'test',pho_pathList),batch_size=batch_size,shuffle=True)
        model.eval()
        for imagedata,label,photopath in dataloader:
            imagedata = imagedata.permute(0,3,1,2)
            imagedata = imagedata.float().cuda()
            pre_label = model(imagedata)
            label = label.reshape(label.shape[0] * label.shape[1]).long()
            label = label.numpy()
            data_num = data_num + label.shape[0]
            predict = np.array(pre_label.argmax(axis=1).cpu())

            ############保存识别的结果是否正确
            for i in range(len(label)):
                if label[i]==1:
                    pos_sample += 1
                    if predict[i]==1:
                        right_right += 1
                    else:
                        right_false += 1
                        # save_img = cv2.imread(self.file_path +'/'+ photopath[i])
                        # cv2.imwrite('/home/xd2/hhx2022/projects/trt/tensorrt-integrate/restful-server/src/python_scripts/yolov5_classify_life_jacket/test_all/data/wrong/p1/' + photopath[i],save_img)
                else:
                    neg_sample += 1
                    if predict[i]==1:
                        false_right += 1
                        # save_img = cv2.imread(self.file_path + '/'+photopath[i])
                        # cv2.imwrite('/home/xd2/hhx2022/projects/trt/tensorrt-integrate/restful-server/src/python_scripts/yolov5_classify_life_jacket/test_all/data/wrong/n1/' + photopath[i],save_img)
                    else:
                        false_false += 1

            same = (predict == label) + 0
            correct_num = correct_num + np.sum(same)
        end = time.time()
        Accuracy = (right_right + false_false)/(data_num)
        right_Recall = right_right/(right_right + right_false)
        right_Precision = right_right/(right_right + false_right)
        right_F1 = 2*right_Precision*right_Recall/(right_Precision+right_Recall)

        false_Recall = false_false/(false_false + false_right)
        false_Precision = false_false/(false_false + right_false)
        false_F1 = 2*false_Precision*false_Recall/(false_Precision+false_Recall)

        # f = open("lifejacketResult.txt", "w")

        result_data = {
            "运行时间": end-start, 
            "数据总个数": data_num,
            "正确数据个数": correct_num,
            "正确率": correct_num/data_num,
            "正样本数据个数": pos_sample,
            "正样本数据预测正确个数": right_right,
            "正样本数据预测错误个数": right_false,
            "正样本正确率": right_right/pos_sample,
            "负样本数据个数": neg_sample,
            "负样本数据预测正确个数": false_false,
            "负样本数据预测错误个数": false_right,
            "负样本正确率": false_false/neg_sample,
            "Accuracy": Accuracy,
            "穿救生衣Precision": right_Precision,
            "穿救生衣Recall": right_Recall,
            "穿救生衣F1-score": right_F1,
            "未穿救生衣Precision": false_Precision,
            "未穿救生衣Recall": false_Recall,
            "未穿救生衣F1-score": false_F1,
            }
        class JsonEncoder(json.JSONEncoder):
            """Convert numpy classes to JSON serializable objects."""
            def default(self, obj):
                if isinstance(obj, (np.integer, np.floating, np.bool_)):
                    return obj.item()
                elif isinstance(obj, np.ndarray):
                    return obj.tolist()
                else:
                    return super(JsonEncoder, self).default(obj)
        
        return json.dumps(result_data,ensure_ascii=False, cls=JsonEncoder)

        # f.write("运行时间：{0}\n".format(end-start))
        # f.write("数据总个数：{0}\n".format(data_num))
        # f.write("正确数据个数：{0}\n".format(correct_num))
        # f.write("正确率：{0}\n".format(correct_num/data_num))
        # f.write("***************************\n")
        # f.write("正样本数据个数：{0}\n".format(pos_sample))
        # f.write("正样本数据预测正确个数：{0}\n".format(right_right))
        # f.write("正样本数据预测错误个数：{0}\n".format(right_false))
        # f.write("正样本正确率：{0}\n".format(right_right/pos_sample))
        # f.write("***************************\n")
        # f.write("负样本数据个数：{0}\n".format(neg_sample))
        # f.write("负样本数据预测正确个数：{0}\n".format(false_false))
        # f.write("负样本数据预测错误个数{0}\n".format(false_right))
        # f.write("负样本正确率{0}\n".format(false_false/neg_sample))
        # f.write("***************************\n")
        # f.write("Accuracy: {0}\n".format(Accuracy))
        # f.write("穿救生衣Precision: {0}\n".format(right_Precision))
        # f.write("穿救生衣Recall: {0}\n".format(right_Recall))
        # f.write("穿救生衣F1-score: {0}\n".format(right_F1))

        # f.write("未穿救生衣Precision: {0}\n".format(false_Precision))
        # f.write("未穿救生衣Recall: {0}\n".format(false_Recall))
        # f.write("未穿救生衣F1-score: {0}\n".format(false_F1))


def RunTest(file_path, json_path):
    # file_path = 'D:/FJF/9_22_night_test/cut1_person'
    # json_path = 'D:/FJF/9_22_night_test/cut1_person_json'
    # file_path = sys.argv[1]
    # json_path = sys.argv[2]
    testdemo = Test(file_path,json_path)
    return testdemo.test('../../weights/alex_11_18_mor2.pth')
