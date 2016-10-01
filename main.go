package main

import (
	"fmt"
	"github.com/yosida95/golang-jenkins"
	"github.com/stianeikeland/go-rpio"
)



func main() {

	// Lets see if we can get our gpio space opened up
	err := rpio.Open()
	if err != nil {
		fmt.Println("Gpio failed to open");
	}

	// GPIO 24 (Header pin 18)
	pin := rpio.Pin(24)

	pin.Output()       // Output mode
	fmt.Println("setting pin high");
	pin.High()         // Set pin High
	// pin.Low()          // Set pin Low
	// pin.Toggle()       // Toggle pin (Low -> High -> Low)


	auth := & gojenkins.Auth{
		Username: SecretUsername,
		ApiToken: SecretApiToken,
	}

	jenkins := gojenkins.NewJenkins(auth, "https://jen01.corp.tsafe.systems/")


	/*jobs, err := jenkins.GetJobs()

	if err != nil {
		fmt.Errorf("error %v\n", err)
	}

	if len(jobs) == 0 {
		fmt.Errorf("return no jobs\n")
	}*/

	job, err := jenkins.GetJob("kernel-rootfs")

	if err != nil {
		fmt.Errorf("error %v\n", err)
	}


	if job.Buildable {
		fmt.Println("Job is currently buildable!!!")
	} else {
		fmt.Println("Job is not currently buildable")
	}






}