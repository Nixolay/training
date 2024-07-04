#  adb connect 127.0.0.1:62001
import subprocess, re
from random import randint
from time import sleep

output = subprocess.check_output(["adb", "devices"])

def getScreenSize(addr):
   output = subprocess.check_output(['adb', "-s", addr, 'shell', 'wm', 'size'], text=True)
   print(f"Screen: {output}")
   match = re.search(r"(\d+)\s*x\s*(\d+)", output)
   if match:
      return int(match.group(1)), int(match.group(2))
   else:
      return 1080, 1920

# Извлечь IP-адреса из строк вывода
ip_addresses = {}
for line in output.decode().split('\n'):
   if 'device' in line and not 'List of devices attached' in line:
      addr = line.split()[0]
      w, h = getScreenSize(addr)
      ip_addresses[addr] = [int(w/2), int(h/2)]

print(ip_addresses)
if len(ip_addresses) == 0:
   addr = '127.0.0.1:62001'
   w, h = getScreenSize(addr)
   ip_addresses[addr] = [int(w/2), int(h/2)]

def send():
    for addr, sc in ip_addresses.items():
        x = randint((sc[0]-200), (sc[0]+200))
        y = randint(sc[1], (sc[1]+300))
        subprocess.call(["adb", "-s", addr, "shell", "input", "tap", str(x), str(y)], stdout=subprocess.DEVNULL)

while True:
    for _ in range(randint(200, 320)):
        send()
        sleep(randint(2500, 3000)/1000)
    sleep(randint(20, 40))