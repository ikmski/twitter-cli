package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ChimeraCoder/anaconda"
)

type screen struct {
}

func newScreen() *screen {

	s := new(screen)

	return s
}

func (s *screen) render(tweets []interface{}) {

	s.clear()

	for i := len(tweets) - 1; i >= 0; i-- {

		var item = tweets[i]
		if item == nil {
			continue
		}

		v := item.(anaconda.Tweet)

		fmt.Printf("%s %s (%s)\n", v.User.Name, v.User.ScreenName, v.CreatedAt)
		fmt.Printf("%s\n", v.FullText)
		fmt.Printf("%s\n\n", "------------------------------------------------")
	}

}

func (s *screen) clear() {

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
