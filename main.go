package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

var cmd string
var duration_tracker int
var start_time time.Time
var end_time time.Time

// #helpers
func flt_to_str(flt float64) string {
	return strconv.FormatFloat(flt, 'f', 1, 64)
}
func get_hour_min() (hour, min int) {
	//tui
	var input_hour string
	var input_min string
	fmt.Println("-------------------")
	fmt.Println("Session Started")
	fmt.Println("------------------")
	fmt.Print("Enter duration hour: ")
	fmt.Scanln(&input_hour)
	/* fmt.Println("")
	fmt.Println("") */
	fmt.Print("Enter duration min: ")
	fmt.Scanln(&input_min)

	//i will be needing them in int
	hour, _ = strconv.Atoi(input_hour)
	min, _ = strconv.Atoi(input_min)
	return hour, min

}

func fmt_pomodoro(duration int) time.Duration {
	total_duration_str, err := time.ParseDuration(strconv.Itoa(duration) + "s")
	if err != nil {
		log.Fatalln(err)
	}
	return total_duration_str
}

func into_seconds(hour, min int) (total_duration int) {
	hour = hour * 60 * 60
	min = min * 60
	total_duration = hour + min
	return
}

func pomodoro(total_duration int) error {
	// converting into seconds
	// Deducting one from duration per sec
	for total_duration != 0 {
		time.Sleep(time.Second)
		total_duration -= 1
		fmt.Println("")
		total_duration_str := fmt_pomodoro(total_duration)
		fmt.Print(total_duration_str.String())
		go fmt.Scanln(&cmd)
		if cmd == "p" {
			duration_tracker = total_duration
			fmt.Println("Paused")
			break
		}
		if total_duration == 0 {
			if runtime.GOOS == "lunux" {
				cmd := exec.Command("paplay", "/usr/share/sounds/freedesktop/stereo/alarm-clock-elapsed.oga")
				cmd.Run()
			}
			end_time = time.Now()
			fmt.Printf("\nstart time: %v \nend time: %v", start_time.Format(time.Kitchen), end_time.Format(time.Kitchen))
			return nil
		}
	}
	return nil
}

func pomodoro_start() error {
	var err error
	if duration_tracker == 0 {
		fmt.Print(duration_tracker)
		total_duration := into_seconds(get_hour_min())
		err = pomodoro(total_duration)
	} else {
		fmt.Scanln(&cmd)
		if cmd == "s" {
			pomodoro(duration_tracker)
		}
	}
	return err
}

func main() {
	start_time = time.Now()
	for {
		err := pomodoro_start()
		if err != nil {
			log.Fatal(err) // to end the program on total_duration == 0
		}
		if err == nil {
			break
		}
	}
}
