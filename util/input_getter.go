package util

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func GetPuzzleInput(day string, sessioncookie string) []string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://adventofcode.com/2021/day/"+day+"/input", nil)
	req.Header.Add("Cookie", "session="+sessioncookie)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Failed to get puzzle input")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to parse puzzle input")
	}
	text := string(body)
	return strings.Split(text, "\n")
}
