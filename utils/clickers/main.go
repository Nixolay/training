package main

import (
    "bufio"
    "bytes"
    "fmt"
    "os"
    "os/exec"
    "os/signal"
    "regexp"
    "strings"
    "sync"
    "syscall"
)



func main() {
	var connectedDevices []string
var wg sync.WaitGroup
    // Установка каналов для сигналов завершения
    shutdownChan := make(chan os.Signal, 1)
    signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)
    
    // Шаг 1: Получение списка устройств
    devicesCmd := exec.Command("adb", "devices")
    devicesOutput, err := devicesCmd.Output()
    if err != nil {
        fmt.Printf("Ошибка выполнения команды adb devices: %s\n", err)
        return
    }

    // Шаг 2: Извлечение IP-адресов и портов
    devices := extractIPsAndPorts(devicesOutput)
    if len(devices) == 0 {
        fmt.Println("Нет доступных устройств для подключения.")
        return
    }

    fmt.Printf("Найденные устройства: %v\n", devices)

    // Канал для получения статуса завершения всех команд
    doneChan := make(chan struct{})

    // Запуск горутины для выполнения команд
    go func() {
        for _, device := range devices {
            wg.Add(1)
            go connectAndTap(device)
        }
        wg.Wait()
        close(doneChan)
    }()

    // Ожидание завершения или получения сигнала
    select {
    case <-shutdownChan:
        fmt.Println("Получен сигнал завершения. Отключение...")
        gracefulShutdown()
    case <-doneChan:
        fmt.Println("Все команды выполнены.")
    }
}
