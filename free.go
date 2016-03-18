package main

import (
	"fmt"
	"flag"
	"os/exec"
	"strings"
	"strconv"
)

func main() {
	var human_readable bool
	flag.BoolVar(&human_readable, "h", false, "display in human readable format")
	flag.Parse()

	installed_memory, err := get_installed_bytes()
	if err != nil {
		panic(err)
	}

	total_consumed := float64(0)
	options := []string{"wired down", "active", "inactive",}

	for i := range options {
		total_consumed = total_consumed + get_used_page_bytes(options[i], DefaultPageScalar)
	}

	free_bytes := installed_memory - total_consumed

	if human_readable {
		fmt.Printf("%.2f gb free\n", free_bytes / 1024 / 1024 / 1000)
	} else {
		fmt.Printf("%.0f bytes free\n", free_bytes)
	}
}

const DefaultPageScalar = 4096

func get_installed_bytes() (float64, error) {
	raw, err := exec.Command("sysctl", "-n", "hw.memsize").Output()
	if err != nil {
		return 0, err
	}

	hwmem, err := strconv.ParseFloat(strings.TrimSpace(string(raw)), 64)
	if err != nil {
		return 0, err
	}

	max_mem := hwmem
	return max_mem, nil
}

func get_used_page_bytes(match_string string, paging int) float64 {
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

	gigabyte_val := (paged_val_float64 * float64(paging))
	return gigabyte_val
}
