import qrcode
import time
import random
import os,sys

def LocalTime(timestamp):

    local_time = time.localtime(timestamp)
    print(time.strftime('%Y-%m-%d %H:%M:%S', local_time))

def ToTime(InputTime):
    # time.strftime('%Y-%m-%d %H:%M:%S', local_time)
    # a = "Sat Mar 28 22:24:24 2016"
    TimeTuple = time.strptime(InputTime, "%Y-%m-%d-%H-%M-%S")
    # print(int(time.mktime(TimeTuple)))
    return int(time.mktime(TimeTuple))

def random_num():
    # 生成一个10-99的随机数
    return random.randint(10, 99)



def getnowtime():
    ToTime(int(time.time()))
    return str(int(time.time()))

def getAnothertime():
    num=random_num()
    ToTime(int(time.time())+num)
    return str(int(time.time())+num)

def img(data):
    # 创建一个QRCode对象
    qr = qrcode.QRCode(
        version=1,  # 控制二维码的大小，值越大二维码越大
        error_correction=qrcode.constants.ERROR_CORRECT_L,  # 纠错水平，L/M/Q/H四个等级
        box_size=10,  # 控制二维码中每个小格子包含的像素数
        border=4,  # 边框的格子数，默认是4
    )

    # 填充数据
    qr.add_data(data)
    qr.make(fit=True)

    # 创建一个图片对象
    img = qr.make_image(fill_color="black", back_color="white")

    # 保存图片
    name = data.split('|')[-1]
    id = data.split("|")[0]
    img.save(f"./qrcode/{id}/{name}.png")

def GetTime():
    time = input("%Y-%m-%d-%H-%M-%S\neg:2024-6-1-10-23-12=>2024年6月1日10点23分12秒\n")
    return time

if __name__ == "__main__":
    id = input("id")
    try:
        os.mkdir(f"./qrcode/{id}")
    except:
        pass
    choose = int(input("mode-1:LocalTime,Please input 1\nmode-2:InputTime,Please input 2\n"))
    match choose:
        case 1:
            print("LocalTime\n")
            QD = id + "|" + getnowtime() + "|" + "QD"
            QT = id + "|" + getAnothertime() + "|" + "QT"
            print("QD:"+QD+"\nQT:"+QT)
            img(QD)
            img(QT)
        case 2:
            print("InputTime")
            InputTime = GetTime()
            QdTime = ToTime(InputTime)
            QtTime = QdTime+random_num()
            QD = id + "|" + str(QdTime) + "|" + "QD"
            QT = id + "|" + str(QtTime) + "|" + "QT"
            print("QD:"+QD+"\nQT:"+QT)
            img(QD)
            img(QT)
        case _:
            print("error,da-ge,bie-gao")


# sys.argv[1]