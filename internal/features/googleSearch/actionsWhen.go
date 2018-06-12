package googleSearch

var (
	searchFieldID    = "lst-ib"
	searchButtonName = "btnK"
)

func inputText(searchText string) {
	page.FindElementById(searchFieldID).SendKeys(searchText)
}

func searchBtn() {
	page.FindElementByName(searchButtonName).Click()
}
