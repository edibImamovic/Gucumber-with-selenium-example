package googleSearch

import (
	"fmt"
	"github.com/gucumber/gucumber/internal/features/googleSearch/pages"
	"github.com/tebeka/selenium"
)

var driver selenium.WebDriver
var page pages.Page

func actionsGiven(website string) {

	var err error
	// set browser as chrome
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})

	// remote to selenium server
	if driver, err = selenium.NewRemote(caps, ""); err != nil {
		fmt.Printf("Failed to open session: %s\n", err)
		return
	}

	page = pages.Page{Driver: driver}
	err = driver.Get(website)
	if err != nil {
		fmt.Printf("Failed to load page: %s\n", err)
		return
	}

}
