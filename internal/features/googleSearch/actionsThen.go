package googleSearch

import (
	"fmt"
	"time"
)

func iWaitFor() error {
	u := time.Second
	fmt.Printf("Waiting for %d %s", 2, "seconds")
	time.Sleep(u * time.Duration(2))
	return nil
}

func compareTitle(expectedTitle string) error {
	var title, _ = driver.Title()
	if expectedTitle != title {
		panic("test failed")
		return nil
	}

	return driver.Quit()
}
