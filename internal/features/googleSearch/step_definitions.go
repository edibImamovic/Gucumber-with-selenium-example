package googleSearch

import (
	. "github.com/gucumber/gucumber"
)

func init() {
	Given(`^I navigate to "(.+?)"$`, func(website string) {
		actionsGiven(website)
	})

	When(`^I write "(.+?)" to the search bar$`, func(search string) {
		inputText(search)
	})

	And(`^I click  button search$`, func() {
		searchBtn()
	})

	Then(`^I should see "(.+?)" in link text$`, func(expectedTitle string) {
		//iWaitFor()
		compareTitle(expectedTitle)
	})

}
