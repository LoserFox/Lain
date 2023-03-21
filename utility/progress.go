package utility

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
)

type Progress struct {
	Name    string
	Hits    int
	Failed  int
	Total   int
	exit    bool
	hasExit bool
}

func (p *Progress) Println(s string) {
	fmt.Printf("\r%s\n\n", s)
}

func (p *Progress) Update(timeStart time.Time) {
	fmt.Printf("\r%s | %d Hits | %d Failed | %d Last | %s", p.Name, p.Hits, p.Failed, p.Total-p.Hits-p.Failed, humanize.Time(timeStart))
}

func (p *Progress) Run() {
	go func() {
		p.exit = false
		p.hasExit = false
		fmt.Printf("\n")
		timeStart := time.Now()
		for !p.exit {
			p.Update(timeStart)
			time.Sleep(time.Second * 1)
		}
		p.Update(timeStart)
		fmt.Printf("\r\n")
		p.hasExit = true
	}()
}

func (p *Progress) Exit() {
	p.exit = true
	for !p.hasExit {
		time.Sleep(time.Millisecond * 10)
	}
}
