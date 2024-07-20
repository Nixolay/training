package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"os/exec"
	"os/signal"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func main() {
	// получаем сигнал системы о закрытии программы ctrl + c
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer func() { cancel() }()
	adb := ADB{cancel: cancel}
	adb.Run(ctx)
	fmt.Println("EXIT:", time.Now().String())
}

var (
	cmdStart = exec.Command("adb", "start-server")
	cmdKill  = exec.Command("adb", "kill-server")
	rgxRes   = regexp.MustCompile(`(\d+)\s*x\s*(\d+)`)
)

type Resolution struct {
	w, h int
}

func (r *Resolution) Set(h, w int) {
	if h > 0 {
		r.h = h / 2
	}

	if w > 0 {
		r.w = w / 2
	}
}

func (r Resolution) String() string {
	return fmt.Sprintf("(w=%d, h=%d)", r.w, r.h)
}

// ADB - адаптер для работы с устройствами
type ADB struct {
	data   map[string]Resolution
	cancel context.CancelFunc
}

// Add добавляет устройство в список устройств
func (adb *ADB) Add(addr string) {
	if isIP(addr) {
		exec.Command("adb", "connect", addr).Run()
	}

	adb.data[addr] = NewScreenSize(addr)
}

// Restart перезапускает сервер adb
func (ADB) Restart() error {
	for _, cmd := range []*exec.Cmd{cmdKill, cmdStart} {
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}

// Send кликает по экрану, для каждого устройства
func (adb ADB) Send() {
	for addr, sc := range adb.data {
		if !adb.TelegramOnScreen(addr) {
			fmt.Printf("Telegram на экране устройства %s, не запущен\n", addr)
			adb.cancel()

			return
		}
		x := rand.Intn(401) + (sc.w - 200)
		y := rand.Intn(301) + sc.h

		cmd := exec.Command("adb", "-s", addr, "shell", "input", "tap", strconv.Itoa(x), strconv.Itoa(y))
		cmd.Stdout = nil // Отключение вывода на консоль

		if err := cmd.Run(); err != nil {
			fmt.Printf("Ошибка выполнения команды adb для адреса %s: %v\n", addr, err)
			adb.cancel()
			return
		}
	}
}

func (adb *ADB) InitDevices() {
	cmd := exec.Command("adb", "devices")
	cmd.Stdout = nil // Отключение вывода на консоль

	adb.data = map[string]Resolution{}

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Ошибка выполнения команды adb: %v\n", err)
	}

	fmt.Printf("Devices: %s\n", output)

	for _, line := range strings.Split(string(output), "\n") {
		if strings.Contains(line, "\tdevice") {
			adb.Add(strings.Split(line, "\t")[0])
		}
	}
}

// Run запускает кликер в бесконечном цикле
func (adb *ADB) Run(ctx context.Context) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	count := 0

	adb.InitDevices()

	if len(adb.data) == 0 {
		fmt.Println("No devices found")
		return
	}

	for {
		count++
		println("COUNT:", count)
		adb.Restart()

		// цикл с рандомным количеством повторений от 200 до 300 повторений
		for i := 0; i < 200+r.Intn(300); i++ {
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Duration(r.Intn(2*int(time.Second)) + 4*int(time.Second))):
				adb.Send()
			}
		}
		<-time.After(time.Duration(r.Intn(30*int(time.Second)) + int(time.Second)*30))
	}
}

// isTelegramOnScreen проверяет, отображается ли приложение Telegram на экране устройства
func (ADB) TelegramOnScreen(addr string) bool {
	// Определяем команду в зависимости от операционной системы
	command := []string{"sh", "-c", fmt.Sprintf("adb -s %s shell dumpsys window | grep mCurrentFocus", addr)}
	if runtime.GOOS == "windows" {
		command = []string{"cmd", "/C", fmt.Sprintf("adb -s %s shell dumpsys window | findstr mCurrentFocus", addr)}
	}

	// Выполнение команды
	output, _ := exec.Command(command[0], command[1:]...).CombinedOutput()
	// fmt.Println("OUTPUT:", string(output))

	// Проверка на то что Telegram на экране
	return strings.Contains(string(output), "org.telegram.messenger")
}

func isIP(addr string) bool {
	return net.ParseIP(addr) != nil
}

func NewScreenSize(addr string) (res Resolution) {
	res.Set(1920, 1080)
	// Выполнение команды adb
	output, err := exec.Command("adb", "-s", addr, "shell", "wm", "size").Output()
	if err != nil {
		return
	}

	// Использование регулярного выражения для извлечения размеров экрана
	matches := rgxRes.FindStringSubmatch(string(output))
	if matches == nil {
		return
	}

	width, _ := strconv.Atoi(matches[1])
	height, _ := strconv.Atoi(matches[2])

	res.Set(height, width)
	fmt.Printf("Screen: %sw:%d h:%d res:%s\n", output, width, height, res)

	return
}
