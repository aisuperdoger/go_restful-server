import time

from LoadData import loadData
from torch.utils.data import DataLoader
import numpy as np
#from lenet import LeNet
from torch import optim
import torch
import os
import cv2
from alexnet import AlexNet,AlexNet1,AlexNet2
# from mobile_net import MobileNetV3_Small,MobileNetV3_Large
# from torch.autograd import Variable

class Test():
    def __init__(self,file_path,json_path):
        self.file_path = file_path
        self.json_path = json_path
    def test(self,save_path):
        batch_size = 16
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
        pho_pathList = os.listdir(file_path)
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
                        #save_img = cv2.imread(self.file_path +'/'+ photopath[i])
                        #cv2.imwrite('/home/xd1/zwl/yolov5-master/test_all/datasets/wrong/rf' + '/' + photopath[i],save_img)
                else:
                    neg_sample += 1
                    if predict[i]==1:
                        false_right += 1
                        #save_img = cv2.imread(self.file_path + '/'+photopath[i])
                        #cv2.imwrite('/home/xd1/zwl/yolov5-master/test_all/datasets/wrong/fr/' + photopath[i],save_img)
                    else:
                        false_false += 1

            same = (predict == label) + 0
            correct_num = correct_num + np.sum(same)
        end = time.time()
        Accuracy = (right_right + false_false)/(data_num)
        Precision = right_right/(right_right + right_false)
        Recall = right_right/(right_right + false_right)
        F1 = 2*Precision*Recall/(Precision+Recall)
        print("运行时间{0}".format(end-start))
        print("数据总个数{0}".format(data_num))
        print("正确数据个数{0}".format(correct_num))
        print("正确率{0}".format(correct_num/data_num))
        print("***************************")
        print("正样本数据个数{0}".format(pos_sample))
        print("正样本数据预测正确个数{0}".format(right_right))
        print("正样本数据预测错误个数{0}".format(right_false))
        print("正样本正确率{0}".format(right_right/pos_sample))
        print("***************************")
        print("负样本数据个数{0}".format(neg_sample))
        print("负样本数据预测正确个数{0}".format(false_false))
        print("负样本数据预测错误个数{0}".format(false_right))
        print("负样本正确率{0}".format(false_false/neg_sample))
        print("***************************")
        print("Accuracy: {0}".format(Accuracy))
        print("Precision: {0}".format(Precision))
        print("Recall: {0}".format(Recall))
        print("F1-score: {0}".format(F1))

if __name__ == "__main__":
    # file_path = 'D:/FJF/9_22_night_test/cut1_person'
    # json_path = 'D:/FJF/9_22_night_test/cut1_person_json'
    file_path = '/home/xd1/zwl/yolov5-master/test_all/data/day_image'
    json_path = '/home/xd1/zwl/yolov5-master/test_all/data/day_json'
    testdemo = Test(file_path,json_path)
    testdemo.test('/home/xd1/zwl/yolov5-master/weights/alex_11_18_mor2.pth')
