package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Masukkan waktu hitungan mundur (detik): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		countdown, err := strconv.Atoi(input)
		if err != nil || countdown <= 0 {
			fmt.Println("Waktu tidak valid. Gunakan angka positif.")
			return
		}

		fmt.Printf("Memulai hitungan mundur dari %d detik...\n", countdown)
		startCountdown(countdown)
	} else if len(os.Args) == 2 {
		countdownStr := os.Args[1]
		countdown, err := strconv.Atoi(countdownStr)
		if err != nil || countdown <= 0 {
			fmt.Println("Waktu tidak valid. Gunakan angka positif.")
			return
		}

		fmt.Printf("Memulai hitungan mundur dari %d detik...\n", countdown)
		startCountdown(countdown)
	} else {
		fmt.Println("Gunakan: countdown [waktu]")
	}
}

func startCountdown(countdown int) {
	go func() {
		for countdown > 0 {
			fmt.Printf("\rSisa waktu: %d detik...", countdown)
			time.Sleep(time.Second)
			countdown--
		}
		fmt.Println("\nHitungan mundur selesai.")

		cmd := exec.Command("shutdown", "/s", "/t", "0")
		err := cmd.Run()
		if err != nil {
			fmt.Println("Gagal mematikan komputer:", err)
		}
	}()

	select {}
}
