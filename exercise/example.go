package exercise

import (
	"fmt"
)

func GetDataFromCMD() error {
	var name string
	var table int
	fmt.Println("Hello! What is your name?")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return fmt.Errorf("error reading name: %v", err)
	}

	fmt.Println("What table do you want?")
	_, err = fmt.Scanln(&table)
	if err != nil {
		return fmt.Errorf("error reading table: %v", err)
	}

	fmt.Printf("Hello %s! Your table is %d.", name, table)
	return nil
}
