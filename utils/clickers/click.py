import pyautogui
from random import randint
from time import sleep

sleep(1)

while True:
    for _ in range(randint(1000, 1200)):
        pyautogui.click()
        sleep(randint(10, 30)/100)
    sleep(randint(10, 20))