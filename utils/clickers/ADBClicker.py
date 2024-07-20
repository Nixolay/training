#  adb connect 127.0.0.1:62001
import subprocess, re, ipaddress, platform, datetime
from random import randint
from time import sleep

subprocess.run("adb kill-server && adb start-server", shell=True, check=True)


class Resolution:
    def __init__(self, w=1080 / 2, h=1920 / 2) -> None:
        self.w, self.h = int(w), int(h)

    def __str__(self) -> str:
        return f"(w={self.w}, h={self.h})"

    def __format__(self, format_spec: str) -> str:
        return self.__str__()

    def __repr__(self) -> str:
        return self.__str__()


# Извлекаем разрешение экрана
def getScreenSize(addr):
    output = subprocess.check_output(
        ["adb", "-s", addr, "shell", "wm", "size"], text=True
    )
    print(f"Screen: {output}")
    match = re.search(r"(\d+)\s*x\s*(\d+)", output)
    if match:
        return Resolution(int(match.group(1)) / 2, int(match.group(2)) / 2)
    return Resolution()


def isIP(addr):
    try:
        ipaddress.ip_address(addr)
        subprocess.call(["adb", "connect", addr])
    except:
        return


ip_addresses = {}


def addAddr(addr: str):
    isIP(addr)
    try:
        ip_addresses[addr] = getScreenSize(addr)
    except:
        ip_addresses[addr] = Resolution()


# Извлечь IP-адреса из строк вывода и получить точку середины экрана
for line in subprocess.check_output(["adb", "devices"]).decode().split("\n"):
    if "device" in line.split():
        print(f"LINE: {line}")
        addAddr(line.split()[0])

if len(ip_addresses) == 0:
    print("ip_addresses is empty")
    addAddr("127.0.0.1:5555")


def send():
    for addr, sc in ip_addresses.items():
        x = randint((sc.w - 200), (sc.w + 200))
        y = randint(sc.h, (sc.h + 300))
        subprocess.call(
            ["adb", "-s", addr, "shell", "input", "tap", str(x), str(y)],
            stdout=subprocess.DEVNULL,
        )


def is_telegram_on_screen():
    try:
        # Определяем команду в зависимости от операционной системы
        if platform.system() == "Windows":
            command = [
                "adb",
                "shell",
                "dumpsys",
                "window",
                "|",
                "findstr",
                "mCurrentFocus",
            ]
        else:
            command = [
                "adb",
                "shell",
                "dumpsys",
                "window",
                "|",
                "grep",
                "mCurrentFocus",
            ]
        # Выполнение команды
        result = subprocess.check_output(
            " ".join(command), shell=True, stderr=subprocess.STDOUT
        )
        result = result.decode("utf-8")
        # Проверка на наличие Telegram
        if "org.telegram.messenger" in result:
            return True
        else:
            return False
    except subprocess.CalledProcessError as e:
        print("Ошибка выполнения команды adb:", e.output.decode("utf-8"))
        return False


print(f"ADDR: {ip_addresses}")
count = 0
while True:
    count += 1
    print(f"\nCOUNT: {count}")
    for _ in range(randint(400, 500)):
        send()
        while not is_telegram_on_screen():
            now = datetime.datetime.now()
            print(f"Telegram is not on screen: {now.strftime('%H:%M:%S')}")
            exit(1)
        sleep(randint(2500, 3000) / 1000)
    for addr in ip_addresses.keys():
        isIP(addr)
    sleep(randint(20, 40))
    try:
        subprocess.run(
            "adb kill-server && adb start-server",
            shell=True,
            check=True,
            stdout=subprocess.DEVNULL,
        )
    except:
        print("Error while restarting adb")
