package main

import (
	"fmt"
	"log"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

func main() {
	out, err := exec.Command("ps", "-e", "-opid,ppid").Output()
	if err != nil {
		log.Fatal(err)
	}
	child := make(map[int][]int)
	a := strings.Split(string(out), "\n")
	for i, s := range a {
		if i == 0 {
			continue
		}
		if len(s) == 0 {
			continue
		}
		b := strings.Fields(s)
		pp, _ := strconv.Atoi(b[1])
		p, _ := strconv.Atoi(b[0])
		child[pp] = append(child[pp], p)
	}
	schild := make([]int, len(child))
	i := 0
	for k, _ := range child {
		schild[i] = k
		i++
	}
	sort.Ints(schild)
	for _, ppid := range schild {
		fmt.Printf("Pid %d has %d child", ppid, len(child[ppid]))
		if len(child[ppid]) == 1 {
			fmt.Printf(": %v\n", child[ppid])
			continue
		}
		fmt.Printf("ren: %v\n", child[ppid])
	}
}
