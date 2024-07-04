package main


// connectAndTap выполняет подключение и команду tap для указанного устройства
func connectAndTap(device string) {
    defer wg.Done()

    connectCmd := exec.Command("adb", "connect", device)
    connectOutput, err := connectCmd.CombinedOutput()
    if err != nil {
        fmt.Printf("Ошибка подключения к устройству %s: %s\n", device, err)
        fmt.Printf("Вывод команды: %s\n", connectOutput)
        return
    }
    fmt.Printf("Успешное подключение к устройству %s: %s\n", device, connectOutput)
    connectedDevices = append(connectedDevices, device)

    tapCmd := exec.Command("adb", "shell", "input", "tap", "10", "20")
    tapOutput, err := tapCmd.CombinedOutput()
    if err != nil {
        fmt.Printf("Ошибка выполнения команды tap на устройстве %s: %s\n", device, err)
        fmt.Printf("Вывод команды: %s\n", tapOutput)
        return
    }
    fmt.Printf("Результат выполнения команды tap на устройстве %s: %s\n", device, tapOutput)
}

// gracefulShutdown корректно завершает все подключения adb
func gracefulShutdown() {
    for _, device := range connectedDevices {
        disconnectCmd := exec.Command("adb", "disconnect", device)
        disconnectOutput, err := disconnectCmd.CombinedOutput()
        if err != nil {
            fmt.Printf("Ошибка отключения от устройства %s: %s\n", device, err)
            fmt.Printf("Вывод команды: %s\n", disconnectOutput)
            continue
        }
        fmt.Printf("Отключение от устройства %s: %s\n", device, disconnectOutput)
    }
}

// extractIPsAndPorts извлекает IP-адреса и порты из вывода команды adb devices
func extractIPsAndPorts(output []byte) []string {
    var devices []string
    scanner := bufio.NewScanner(bytes.NewReader(output))
    re := regexp.MustCompile(`^([\d.]+:\d+)\s+device$`)

    for scanner.Scan() {
        line := scanner.Text()
        match := re.FindStringSubmatch(line)
        if len(match) > 1 {
            devices = append(devices, match[1])
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("Ошибка чтения вывода команды adb devices: %s\n", err)
    }

    return devices
}
