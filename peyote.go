package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	promptTextColor                = "254"
	promptUserBackgroundColor      = "23"
	promptRootBackgroundColor      = "88"
	promptDirectoryBackgroundColor = "241"
	promptGitTextColor             = "2"
	promptGitBackgroundColor       = "16"
	promptErrorTextColor           = "203"
	promptErrorBackgroundColor     = "236"

	showUserAndHost    = true
	showDirectoryName  = true
	shortDirectoryName = true
	showGitStatus      = true
	showReturnCode     = true
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
		usercolor = promptUserBackgroundColor
	} else {
		usercolor = promptRootBackgroundColor
	}

	hostname, err := os.Hostname()
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	homedir := os.Getenv("HOME")
	dirname := strings.Replace(dir, homedir, "~", 1)
	if shortDirectoryName && strings.Count(dirname, "/") > 1 {
		dirname = dirname[strings.LastIndex(dirname, "/")+1:]
	}

	returncodeint, err := strconv.Atoi(returncode)
	if err != nil {
		log.Fatal(err)
	}

	prompt := NewPrompt()
	if showUserAndHost {
		prompt.AddSegment(NewSegment(promptTextColor, usercolor, fmt.Sprintf(" %s@%s ", username, hostname)))
	}
	if showDirectoryName {
		prompt.AddSegment(NewSegment(promptTextColor, promptDirectoryBackgroundColor, fmt.Sprintf(" %s ", dirname)))
	}
	if showGitStatus {
		directoryIsGitRepo := true
		gitStatus := exec.Command("git", "status")
		output, err := gitStatus.CombinedOutput()
		if err != nil {
			directoryIsGitRepo = false
		}
		if directoryIsGitRepo {
			gitHEAD := ""
			workingTreeClean := "!"
			scanner := bufio.NewScanner(bytes.NewReader(output))
			for scanner.Scan() {
				text := scanner.Text()
				if strings.HasPrefix(text, "On branch ") {
					gitHEAD = strings.Split(text, "On branch ")[1]
				}
				if strings.HasPrefix(text, "HEAD detached at ") {
					gitHEAD = strings.Split(text, "HEAD detached at ")[1]
				}
				if strings.HasPrefix(text, "nothing to commit, working tree clean") {
					workingTreeClean = ""
				}
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
			prompt.AddSegment(NewSegment(promptGitTextColor, promptGitBackgroundColor, fmt.Sprintf(" %[1]s%[2]s ", gitHEAD, workingTreeClean)))
		}
	}
	if showReturnCode && returncodeint > 0 {
		prompt.AddSegment(NewSegment(promptErrorTextColor, promptErrorBackgroundColor, fmt.Sprintf(" %s ", returncode)))
	}

	fmt.Println(prompt.ToString())
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

func (p *Prompt) AddSegment(segment Segment) *Prompt {
	p.Segments = append(p.Segments, segment)
	return p
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
