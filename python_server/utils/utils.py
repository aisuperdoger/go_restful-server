import cv2 
import os

def detect_video(source, get_image_func,save_path):
    vid_path, vid_writer = None, None
    directory = os.path.dirname(save_path)
    # 检查目录是否存在，如果不存在则创建
    if not os.path.exists(directory):
        os.makedirs(directory)
    
    vid_cap = cv2.VideoCapture(source)

    # count = 0
    # print(source)
    while vid_cap.isOpened():
        ret, input_im = vid_cap.read()

        if not ret:
            break

        # 检测视频中每一帧
        # ouput_im = input_im
        ouput_im = get_image_func(input_im)
        
         # 保存图片为视频
        if vid_path != save_path:  # new video
            vid_path = save_path
            if isinstance(vid_writer, cv2.VideoWriter):
                vid_writer.release()  # release previous video writer
        
            fps = vid_cap.get(cv2.CAP_PROP_FPS)
            w = int(vid_cap.get(cv2.CAP_PROP_FRAME_WIDTH))
            h = int(vid_cap.get(cv2.CAP_PROP_FRAME_HEIGHT))
            vid_writer = cv2.VideoWriter(save_path, cv2.VideoWriter_fourcc(*'VP80'), fps, (w, h))
        vid_writer.write(ouput_im)

    vid_cap.release()
    print("results saved to %s" % save_path)
    return True
    
    