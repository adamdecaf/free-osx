package main

import (
	"fmt"
	"os/exec"
	"strings"
	"strconv"
)

func main() {
	installed_memory, err := exec.Command("sysctl", "-n", "hw.memsize").Output()
	if err != nil {
		panic(err)
	}

	installed_memory_float64, err := strconv.ParseFloat(strings.TrimSpace(string(installed_memory)), 64)
	if err != nil {
		panic(err)
	}

	max_mem := installed_memory_float64 / float64(1024) / float64(1024) / float64(1000)
	total_consumed := float64(0)
	options := []string{"wired down", "active", "inactive",}

	for i := range options {
		total_consumed = total_consumed + get_paged_memory_usage(options[i], DefaultPageScalar)
	}

	fmt.Printf("%.2fGB\n", max_mem - total_consumed)
}

const DefaultPageScalar = 4096

func get_paged_memory_usage(match_string string, paging int) float64 {
	mvar := "3"
	if strings.Contains(match_string, " ") {
		mvar = "4"
	}

	b := "vm_stat | grep \"Pages " + match_string + ":\" | awk '{ print $" + mvar + " }'"
	paged_val, err := exec.Command("bash", "-c", b).Output()
	if err != nil {
		panic(err)
	}

	trimmed := strings.TrimSpace(string(paged_val))
	clean := strings.Replace(trimmed, ".", "", -1)
	paged_val_float64, err := strconv.ParseFloat(clean, 64)
	if err != nil {
		panic(err)
	}

	gigabyte_val := (paged_val_float64 * float64(paging)) / 1024 / 1024 / 1000
	return gigabyte_val
}
