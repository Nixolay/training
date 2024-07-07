#  adb connect 127.0.0.1:62001
import subprocess, re, ipaddress
from random import randint
from time import sleep

class Resolution:
   def __init__(self, w: int, h: int) -> None:
      self.w, self.h = int(w), int(h)
   def __str__(self) -> str:
      return f"w: {self.w} h:{self.h}"

# Извлекаем разрешение экрана
def getScreenSize(addr):
   output = subprocess.check_output(['adb', "-s", addr, 'shell', 'wm', 'size'], text=True)
   print(f"Screen: {output}")
   match = re.search(r"(\d+)\s*x\s*(\d+)", output)
   if match:
      return Resolution(int(match.group(1))/2, int(match.group(2))/2)
   else:
      return Resolution(1080/2, 1920/2)

# Извлечь IP-адреса из строк вывода и получить точку середины экрана
ip_addresses = {}
for line in subprocess.check_output(["adb", "devices"]).decode().split('\n'):
   if 'device' in line and not 'List of devices attached' in line:
      addr = line.split()[0]
      ip_addresses[addr] = getScreenSize(addr)

print(f"addrs: {ip_addresses}")
if len(ip_addresses) == 0:
   addr = '127.0.0.1:5555'
   subprocess.call(["adb", "connect", addr])
   ip_addresses[addr] = getScreenSize(addr)

def isIP(addr):
   try:
      ipaddress.ip_address(addr)
      return True
   except:
      return False

def send():
   for addr, sc in ip_addresses.items():
      # print(f"{addr} {sc}")
      if isIP(addr):
         subprocess.call(["adb", "connect", addr], stdout=subprocess.DEVNULL)
      x = randint((sc.w-200), (sc.w+200))
      y = randint(sc.h, (sc.h+300))
      subprocess.call(["adb", "-s", addr, "shell", "input" , "tap", str(x), str(y)], stdout=subprocess.DEVNULL)

while True:
    for _ in range(randint(200, 320)):
        send()
        sleep(randint(2500, 3000)/1000)
    sleep(randint(20, 40))