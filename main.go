package main

import (
    "context"
    "log"
    "os/exec"
    "strings"
    "time"
    "github.com/redis/go-redis/v9"
    "os"
)

var rdb *redis.Client
var redisAddr string

func initRedis() {
    rdb = redis.NewClient(&redis.Options{
        Addr: redisAddr,
        Password: "",
        DB: 0,
    })
}

func readClipboard() (string, error) {
    out, err := exec.Command("wl-paste").Output()
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(string(out)), nil
}

func writeClipboard(content string) error {
    cmd := exec.Command("wl-copy")
    cmd.Stdin = strings.NewReader(content)
    return cmd.Run()
}

func publishClipboard(content string) error {
    ctx := context.Background()
    return rdb.Publish(ctx, "clipboard:updates", content).Err()
}

func subscribeClipboard(updateChan chan string) {
    ctx := context.Background()
    pubsub := rdb.Subscribe(ctx, "clipboard:updates")
    defer pubsub.Close()
    for {
        msg, err := pubsub.ReceiveMessage(ctx)
        if err != nil {
            log.Println("Error receiving message:", err)
            continue
        }
        updateChan <- msg.Payload
    }
}

func main() {
    redisAddr = "192.168.2.3:6379" // valor por defecto
    if len(os.Args) > 1 {
        arg := os.Args[1]
        if arg == "--help" || arg == "-h" {
            log.Printf("Uso: %s [redis_host:port]", os.Args[0])
            log.Printf("Por defecto: %s", redisAddr)
            return
        }
        redisAddr = arg
    }
    initRedis()
    updateChan := make(chan string)
    lastContent := ""

    go subscribeClipboard(updateChan)

    go func() {
        for {
            content, err := readClipboard()
            if err != nil {
                log.Println("Error reading clipboard:", err)
                time.Sleep(1 * time.Second)
                continue
            }
            if content != lastContent && content != "" {
                lastContent = content
                if err := publishClipboard(content); err != nil {
                    log.Println("Error publishing clipboard:", err)
                }
            }
            time.Sleep(1 * time.Second)
        }
    }()

    for content := range updateChan {
        if content != lastContent {
            lastContent = content
            if err := writeClipboard(content); err != nil {
                log.Println("Error writing to clipboard:", err)
            }
        }
    }
}