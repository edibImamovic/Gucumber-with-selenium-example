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
	title, err := driver.Title()
	if err != nil {
		return err
	}
	if expectedTitle !=title{
		panic(title + " and " + expectedTitle + "are not equal")
		return driver.Close()
	}

	return driver.Close()
}

