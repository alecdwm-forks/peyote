package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	promptTextColor            = "254"
	promptUserColor            = "23"
	promptRootColor            = "88"
	promptDirColor             = "241"
	promptErrorTextColor       = "203"
	promptErrorBackgroundColor = "236"
	shortDirName               = true
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: peyote $? $EUID")
		os.Exit(1)
	}

	returncode := os.Args[1]
	username := os.Getenv("USER")
	euid := os.Args[2]
	euidnum, err := strconv.Atoi(euid)
	if err != nil {
		log.Fatal(err)
	}

	var usercolor string
	if euidnum > 0 {
		usercolor = promptUserColor
	} else {
		usercolor = promptRootColor
	}

	hostname, err := os.Hostname()
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	homedir := os.Getenv("HOME")
	dirname := strings.Replace(dir, homedir, "~", 1)
	if shortDirName && strings.Count(dirname, "/") > 1 {
		dirname = dirname[strings.LastIndex(dirname, "/")+1:]
	}

	returncodeint, err := strconv.Atoi(returncode)
	if err != nil {
		log.Fatal(err)
	}
	if returncodeint > 0 {
		prompt := NewPrompt(
			[]string{promptTextColor, usercolor, fmt.Sprintf(" %s@%s ", username, hostname)},
			[]string{promptTextColor, promptDirColor, fmt.Sprintf(" %s ", dirname)},
			[]string{promptErrorTextColor, promptErrorBackgroundColor, fmt.Sprintf(" %s ", returncode)},
		)
		fmt.Println(prompt.ToString())
	} else {
		prompt := NewPrompt(
			[]string{promptTextColor, usercolor, fmt.Sprintf(" %s@%s ", username, hostname)},
			[]string{promptTextColor, promptDirColor, fmt.Sprintf(" %s ", dirname)},
		)
		fmt.Println(prompt.ToString())
	}
}

type Prompt struct {
	Segments []Segment
}

func NewPrompt(segmentsParams ...[]string) Prompt {
	segments := []Segment{}
	for _, segmentParams := range segmentsParams {
		segments = append(segments, NewSegment(segmentParams[0], segmentParams[1], segmentParams[2]))
	}
	return Prompt{
		Segments: segments,
	}
}

func (p *Prompt) ToString() string {
	fmtString := ""
	for i := range p.Segments {
		nextBackgroundColor := ""
		if i < len(p.Segments)-1 {
			nextBackgroundColor = p.Segments[i+1].BackgroundColor
		}
		fmtString += p.Segments[i].ToString(nextBackgroundColor)
	}
	fmtString += "%f%k "
	return fmtString
}

type Segment struct {
	TextColor       string
	BackgroundColor string
	Text            string
}

func NewSegment(textColor, backgroundColor, text string) Segment {
	return Segment{
		TextColor:       textColor,
		BackgroundColor: backgroundColor,
		Text:            text,
	}
}

func (s *Segment) ToString(nextBackgroundColor string) string {
	var powerlineArrow string
	if nextBackgroundColor == "" {
		powerlineArrow = fmt.Sprintf("%%F{%[1]s}", s.BackgroundColor)
	} else {
		powerlineArrow = fmt.Sprintf("%%K{%[1]s}%%F{%[2]s}", nextBackgroundColor, s.BackgroundColor)
	}
	return fmt.Sprintf("%%F{%[1]s}%%K{%[2]s}%[3]s%%f%%k%[4]s", s.TextColor, s.BackgroundColor, s.Text, powerlineArrow)
}
