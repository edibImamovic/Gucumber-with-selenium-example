## work in progress 

## Gucumber with selenium web driver

An implementation of Gucumber with Selenium for GO. 
This repo are implementation of selenium web driver with Gucumber through simple example 

## Installing of Gucumber

```sh
$ go get github.com/gucumber/gucumber/cmd/gucumber
```
For more details on Gucumber install visit `https://github.com/gucumber/gucumber`

 Usage

1. Create cucumber features in `/features` folder eg googleSearch
1. Make sure that `step_definitions` are in same folder as your features
1. Execute Selenium Server `selenium-standalone start` (https://www.seleniumhq.org/download/)
1. Run `gucumber`.


## Acknowledgements

* [tebeka/selenium](https://github.com/tebeka/selenium) project for Selenium client for Golang.


