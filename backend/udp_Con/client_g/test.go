// Copyright 2017 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

const BUFFERSIZE = 1024

func main() {
	walk.Resources.SetRootDirPath("./")

	checkFile := fileExist()
	if checkFile == 1 {
		textBased()
	} else {
		imageViewer()
	}
}

func retrieveImage() {
	connection, err := net.Dial("tcp", "localhost:27001")
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	fmt.Println("Connected to server, start receiving the file name and file size")
	bufferFileName := make([]byte, 64)
	bufferFileSize := make([]byte, 10)

	connection.Read(bufferFileSize)
	fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)

	connection.Read(bufferFileName)
	fileName := strings.Trim(string(bufferFileName), ":")

	newFile, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	var receivedBytes int64

	for {
		if (fileSize - receivedBytes) < BUFFERSIZE {
			io.CopyN(newFile, connection, (fileSize - receivedBytes))
			connection.Read(make([]byte, (receivedBytes+BUFFERSIZE)-fileSize))
			break
		}
		io.CopyN(newFile, connection, BUFFERSIZE)
		receivedBytes += BUFFERSIZE
	}
	fmt.Println("Received file completely!")
}

func fileExist() int {
	if _, err := os.Stat("a.jpg"); os.IsNotExist(err) {
		// path/to/whatever does not exist
		return 1
	} else {
		return 0
	}
}

func imageViewer() {
	type Mode struct {
		Name  string
		Value ImageViewMode
	}

	modes := []Mode{
		{"ImageViewModeIdeal", ImageViewModeIdeal},
	}

	var widgets []Widget

	for _, mode := range modes {
		widgets = append(widgets,
			ImageView{
				Background: SolidColorBrush{Color: walk.RGB(255, 191, 0)},
				Image:      "a.jpg",
				Margin:     10,
				Mode:       mode.Value,
			},
		)
	}

	MainWindow{
		Title:    "Walk ImageView Example",
		Size:     Size{1000, 1000},
		Layout:   Grid{Columns: 1},
		Children: widgets,
	}.Run()
}

func textBased() {
	var inTE, outTE *walk.TextEdit

	MainWindow{
		Title:   "SCREAMO",
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "SCREAM",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
		},
	}.Run()
}
