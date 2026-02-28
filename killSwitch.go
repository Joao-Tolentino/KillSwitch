package main

// Kill Switch - Shutdown the system when used
// Used for intentional forced shutdown

//Imports, uncomment the os/exec to use commands
import (
		"log"
		//"os/exec"
)

// Kill Switch logic
func ks(){
	log.Println("Test Successful! o7")
	
	//Shutdown the machine, uncomment for actual use
	/* Windows
	cmd := exec.Command("shutdown", "/s", "/t", "0")

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	*/
}