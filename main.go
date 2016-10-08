package main

import (
	"time"
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"github.com/yosida95/golang-jenkins"
)


func setLavaLampPower(power bool){
	fmt.Println("Modifying lava lamp power: ", power)
	err := rpio.Open()
	if err != nil {
    	fmt.Println("Failed to initialize gpio library")
    	return
	}
	pin := rpio.Pin(24)
	pin.Output()       // Output mode

	if power {
		pin.High()        // Set pin High
	} else {
		pin.Low()         // Set pin Low
	}

	rpio.Close()
	return
}


func main() {

	currentTime := time.Now()
	fmt.Println("Its ", currentTime.Weekday(), "(", int(currentTime.Weekday()), ") Current hour is: ", currentTime.Hour())

	// If it is the weekend, don't do anything. just peace
	if int(currentTime.Weekday()) > 5 {
		fmt.Println("Its the weekend, Not doing anything")
		return
	}

	// if its before 7am or after 7pm. We should make sure the lamp is turned off
	if (currentTime.Hour() < 7) || (currentTime.Hour() > 19) {
		fmt.Println("Hour is unnacceptble, turning lava lamp off")
		setLavaLampPower(false)
		return
	}

	auth := &gojenkins.Auth{
	   Username: "bclouser",
	   ApiToken: "167bff1f78bf1338f0bb21f3157744e7",
	}
	jenkins := gojenkins.NewJenkins(auth, "https://jen01.corp.tsafe.systems/")

	job, err := jenkins.GetJob("kernel-rootfs")

	fmt.Println(job)

	if err != nil {
		fmt.Println("Something bad happened")
		fmt.Errorf("error %v\n", err)
		return
	}

	// Job is buildable so we should make sure the lamp is off
	if job.Buildable {
		fmt.Println("Job is currently building")
		setLavaLampPower(false)
	} else { 
		// Uh Oh, Job is not buildable... make sure lamp is on
		fmt.Println("Job is not currently building")
		setLavaLampPower(true)
	}	
}
