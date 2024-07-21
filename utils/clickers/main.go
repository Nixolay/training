package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	log "log/slog"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const deviceTemplate = "\tdevice"

var (
	cmdStart = exec.Command("adb", "start-server")
	cmdKill  = exec.Command("adb", "kill-server")
	rgxRes   = regexp.MustCompile(`(\d+)\s*x\s*(\d+)`)
)

func main() {
	// получаем сигнал системы о закрытии программы ctrl + c
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer func() { cancel() }()

	adb := ADB{cancel: cancel}
	if err := adb.InitLogger(); err != nil {
		panic(err)
	}

	adb.Run(ctx)
	log.Info("EXIT")
}

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
func (adb *ADB) Add(deviceID string) {
	if isIP(deviceID) {
		exec.Command("adb", "connect", deviceID).Run()
	}

	if isDeviceAvailable(deviceID) {
		adb.data[deviceID] = NewScreenSize(deviceID)
	}
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
	for deviceID, sc := range adb.data {
		if !adb.TelegramOnScreen(deviceID) {
			log.Error("Telegram на экране устройства, не запущен", "addr", deviceID)
			adb.cancel()
			return
		}

		x := rand.Intn(401) + (sc.w - 200)
		y := rand.Intn(301) + sc.h

		cmd := exec.Command("adb", "-s", deviceID, "shell", "input", "tap", strconv.Itoa(x), strconv.Itoa(y))
		cmd.Stdout = nil // Отключение вывода на консоль

		if err := cmd.Run(); err != nil {
			log.Error("Ошибка выполнения команды adb для адреса "+deviceID, "err", err)
			adb.cancel()
			return
		}
	}
}

func (adb *ADB) InitLogger() error {
	file, err := os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		return err
	}

	log.SetDefault(log.New(log.NewJSONHandler(file, nil)))

	return nil
}

func (adb *ADB) InitDevices() {
	cmd := exec.Command("adb", "devices")
	cmd.Stdout = nil // Отключение вывода на консоль

	adb.data = map[string]Resolution{}

	output, err := cmd.Output()
	if err != nil {
		log.Error("Ошибка выполнения команды adb", "err", err)
	}

	log.Info("Devices: " + string(output))

	for _, line := range append(strings.Split(string(output), "\n"), readAddresses("devices.txt")...) {
		if strings.Contains(line, deviceTemplate) {
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
		log.Error("Устройства не обнаружены")
		return
	}

	for {
		count++
		log.Info("STEP", "count", count)
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
func (ADB) TelegramOnScreen(deviceID string) bool {
	// Определяем команду в зависимости от операционной системы
	command := []string{"sh", "-c", fmt.Sprintf("adb -s %s shell dumpsys window | grep mCurrentFocus", deviceID)}
	if runtime.GOOS == "windows" {
		command = []string{"cmd", "/C", fmt.Sprintf("adb -s %s shell dumpsys window | findstr mCurrentFocus", deviceID)}
	}

	// Выполнение команды
	output, _ := exec.Command(command[0], command[1:]...).CombinedOutput()
	// log.Info("OUTPUT: " +string(output))

	// Проверка на то что Telegram на экране
	return strings.Contains(string(output), "org.telegram.messenger")
}

func isIP(addr string) bool {
	return net.ParseIP(addr) != nil
}

func NewScreenSize(deviceID string) (res Resolution) {
	res.Set(1920, 1080)
	// Выполнение команды adb
	output, err := exec.Command("adb", "-s", deviceID, "shell", "wm", "size").Output()
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
	log.Info("Screen", "output", string(output), "w", width, "h", height, "res", res)

	return
}

func readAddresses(filename string) []string {
	var addrs []string

	// Открываем файл
	file, err := os.Open(filename)
	if err != nil {
		log.Error("не удалось открыть файл", "err", err)
		return nil
	}
	defer file.Close()

	// Используем bufio.Scanner для построчного чтения файла
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		addrs = append(addrs, scanner.Text()+deviceTemplate)
	}

	// Проверяем на ошибки, которые могут произойти во время сканирования
	if err := scanner.Err(); err != nil {
		log.Error("ошибка при чтении файла", "err", err)
		return nil
	}

	return addrs
}

func isDeviceAvailable(deviceID string) bool {
	var out bytes.Buffer

	cmd := exec.Command("adb", "-s", deviceID, "shell", "echo", "hello")
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Error("ошибка выполнения команды adb", "err", err)
		return false
	}

	return strings.Contains(out.String(), "hello")
}
