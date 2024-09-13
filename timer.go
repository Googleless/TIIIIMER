package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/go-toast/toast"
	"github.com/hajimehoshi/go-mp3"
)

func main() {
	var seconds int
	fmt.Print("Введите количество секунд для таймера: ")
	_, err := fmt.Scan(&seconds)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}
	switch seconds % 10 {
	case 1:
		fmt.Printf("Таймер на %d секунду начат...\n", seconds)
		notification := toast.Notification{
			AppID:   "Таймер на Go",
			Title:   "Таймер установлен!",
			Message: fmt.Sprintf("Вы успешно установили таймер на %d секунду!", seconds),
		}
		errpush1 := notification.Push()
		if errpush1 != nil {
			log.Fatalln(errpush1)
		}
	case 2, 3, 4:
		fmt.Printf("Таймер на %d секунды начат...\n", seconds)
		notification := toast.Notification{
			AppID:   "Таймер на Go",
			Title:   "Таймер установлен!",
			Message: fmt.Sprintf("Вы успешно установили таймер на %d секунды!", seconds),
		}
		errpush1 := notification.Push()
		if errpush1 != nil {
			log.Fatalln(errpush1)
		}
	default:
		fmt.Printf("Таймер на %d секунд начат...\n", seconds)
		notification := toast.Notification{
			AppID:   "Таймер на Go",
			Title:   "Таймер установлен!",
			Message: fmt.Sprintf("Вы успешно установили таймер на %d секунд!", seconds),
		}
		errpush1 := notification.Push()
		if errpush1 != nil {
			log.Fatalln(errpush1)
		}
	}
	fileBytes, err := os.ReadFile("02.mp3")
	if err != nil {
		panic("reading my-file.mp3 failed: " + err.Error())
	}
	fileBytesReader := bytes.NewReader(fileBytes)
	decodedMp3, err := mp3.NewDecoder(fileBytesReader)
	if err != nil {
		panic("mp3.NewDecoder failed: " + err.Error())
	}
	op := &oto.NewContextOptions{}
	op.SampleRate = 44100
	op.ChannelCount = 2
	op.Format = oto.FormatSignedInt16LE
	otoCtx, readyChan, err := oto.NewContext(op)
	if err != nil {
		panic("oto.NewContext failed: " + err.Error())
	}
	<-readyChan
	player := otoCtx.NewPlayer(decodedMp3)

	for seconds > 0 {
		minutes := seconds / 60
		remainingSeconds := seconds % 60
		fmt.Printf("%02d:%02d\n", minutes, remainingSeconds)
		time.Sleep(1 * time.Second)
		seconds--
	}
	fmt.Println("Время вышло!")
	notificationend := toast.Notification{
		AppID:   "Таймер на Go",
		Title:   "Время вышло!",
		Message: "Таймер подошёл к концу!",
	}
	player.Play()
	errpush := notificationend.Push()
	if errpush != nil {
		log.Fatalln(errpush)
	}
	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}
	err = player.Close()
	if err != nil {
		panic("player.Close failed: " + err.Error())
	}
}


