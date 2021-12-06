/*
* Written by Talia Bjelland, 2021
* Purpose: read madlib text out loud
 */

package main

// Copyright 2017 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"io"
	"os"

	"github.com/hajimehoshi/oto"

	"github.com/hajimehoshi/go-mp3"

	"io/ioutil"
	"log"
	"strconv"
	"strings"

	htgotts "github.com/hegedustibor/htgo-tts"
)

//run function written by Hajime Hoshe, 2017
//purpose: plays mp3 files
func run(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}
	defer c.Close()

	p := c.NewPlayer()
	defer p.Close()

	if _, err := io.Copy(p, d); err != nil {
		return err
	}
	return nil
}

//splits text into 200-character chunks to keep mp3 file size small
func textSplitter(text string) []string {

	var splitText []string
	i := 0
	j := 100
	for i < len(text) {

		if j > len(text)-1 {
			splitText = append(splitText, text[i:])
		} else {
			j = j + (strings.Index(text[j:], " "))
			splitText = append(splitText, text[i:j])
		}
		i = j
		j = j + 100
	}

	return splitText
}

//plays series of small mp3 files to provide text-to-speech functionality
func playSound(text string) {
	var splitText []string = textSplitter(text)

	speech := htgotts.Speech{Folder: "audio", Language: "en"}

	for i, element := range splitText {
		speech.Speak(element)
		files, err := ioutil.ReadDir("audio")
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			var name string = file.Name()
			if string(name[0]) == "e" {
				var num string = strconv.Itoa(i)
				if len(num) == 1 {
					num = "0" + num
				}
				var dst string = "a" + num + ".mp3"
				os.Rename("audio/"+file.Name(), "audio/"+dst)
			}
		}

	}
	files, err := ioutil.ReadDir("audio")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if err := run("audio/" + file.Name()); err != nil {
			log.Fatal(err)
		}
	}

	//remove all frm audio
	err2 := os.RemoveAll("audio")
	if err2 != nil {
		log.Fatal(err2)
	}
}
