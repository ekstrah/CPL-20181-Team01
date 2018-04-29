from random import randint
import PIL.Image
import time

def genValues(min, max):
    return randint(min, max)

    



if __name__ == "__main__":
    while True:
        temperature = genValues(0, 100)
        humidity = genValues(0, 100)
        numofppl = genValues(0, 10)
        print(temperature, humidity, numofppl)
        img = PIL.Image.open("./img/one.jpg")
        img.show()
        time.sleep(4)