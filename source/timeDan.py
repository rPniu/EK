import time

def LocalTime(timestamp):

    local_time = time.localtime(timestamp)
    print(time.strftime('%Y-%m-%d %H:%M:%S', local_time))

def ToTime(InputTime):
    # time.strftime('%Y-%m-%d %H:%M:%S', local_time)
    # a = "Sat Mar 28 22:24:24 2016"
    TimeTuple = time.strptime(InputTime, "%Y-%m-%d-%H-%M-%S")
    # print(int(time.mktime(TimeTuple)))
    return int(time.mktime(TimeTuple))


if __name__ == '__main__':
    # while(True):
    #     timestamp = int(input("time:"))
    #     if(timestamp==0):
    #         break
    #     Totime(timestamp)
    # print("End_Up")

    InputTime = input("time:\n")
    ToTime(InputTime)

