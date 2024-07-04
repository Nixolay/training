#  adb connect 127.0.0.1:62001
import subprocess
from random import randint
from time import sleep

output = subprocess.check_output(["adb", "devices"])

# Извлечь IP-адреса из строк вывода
ip_addresses = []
for line in output.decode().split('\n'):
  if 'device' in line and not 'List of devices attached' in line:
    ip_addresses.append(line.split()[0])

print(ip_addresses)
if len(ip_addresses) == 0:
   ip_addresses.append('127.0.0.1:62001')

def send():
    for addr in ip_addresses:
        x = randint(310, 723)
        y = randint(912, 1290)
        subprocess.call(["adb", "-s", addr, "shell", "input", "tap", str(x), str(y)], stdout=subprocess.DEVNULL)

while True:
    for _ in range(randint(200, 320)):
        send()
        sleep(randint(1100, 1200)/1000)
    sleep(randint(20, 40))