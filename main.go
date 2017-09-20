package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"os"
)

func main() {
	const (
		// These paths will be different on your system.
		seleniumPath     = "selenium-server-standalone-3.5.3.jar"
		chromeDriverPath = "chromedriver.exe"
		port             = 8090
	)

	opts := []selenium.ServiceOption{
		selenium.Output(os.Stderr),
	}
	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err)
	}
	defer service.Stop()

	svc, err := selenium.NewChromeDriverService(chromeDriverPath, 9515)
	if err != nil {
		panic(err)
	}
	defer svc.Stop()

	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Navigate to the simple playground interface.
	if err := wd.Get("http://play.golang.org/?simple=1"); err != nil {
		panic(err)
	}

	var buff []byte
	if buff, err = wd.Screenshot(); err != nil {
		panic(err)
	}

	f, _ := os.Create("screen.png")
	defer f.Close()

	f.Write(buff)
}
