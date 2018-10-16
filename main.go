package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"github.com/urfave/cli"
)

var (
	procPath = "/proc"
)

type MemInfo struct {
	MemTotal          int64
	MemFree           int64
	MemAvailable      int64
	Buffers           int64
	Cached            int64
	SwapCached        int64
	Active            int64
	Inactive          int64
	ActiveAnon        int64
	InactiveAnon      int64
	ActiveFile        int64
	InactiveFile      int64
	Unevictable       int64
	Mlocked           int64
	SwapTotal         int64
	SwapFree          int64
	Dirty             int64
	Writeback         int64
	AnonPages         int64
	Mapped            int64
	Shmem             int64
	Slab              int64
	SReclaimable      int64
	SUnreclaim        int64
	KernelStack       int64
	PageTables        int64
	NFSUnstable       int64
	Bounce            int64
	WritebackTmp      int64
	CommitLimit       int64
	CommittedAS       int64
	VmallocTotal      int64
	VmallocUsed       int64
	VmallocChunk      int64
	HardwareCorrupted int64
	AnonHugePages     int64
	ShmemHugePages    int64
	ShmemPmdMapped    int64
	CmaTotal          int64
	CmaFree           int64
	HugePages_Total   int64
	HugePages_Free    int64
	HugePages_Rsvd    int64
	HugePages_Surp    int64
	Hugepagesize      int64
	Hugetlb           int64
	DirectMap4k       int64
	DirectMap2M       int64
	DirectMap1G       int64
}

func parseMemInfo(r io.Reader) *MemInfo {
	m := &MemInfo{}
	scanner := bufio.NewScanner(r)
	// put meminfo file parsing here

	return m
}

// procFilePath joins custom paths in /proc psuedofs
func procFilePath(name string) string {
	return path.Join(procPath, name)
}

// printFile is an helper function to print an *os.File object
func printFile(file *os.File) error {
	buf := make([]byte, 1024)
	_, err := file.Read(buf)
	if err != nil && err != io.EOF {
		return err
	}
	fmt.Println(string(buf))
	return nil
}

func main() {

	app := cli.NewApp()
	app.Name = "procreader"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Giovan Battista Salinetti",
			Email: "gbsalinetti@gmail.com",
		},
	}
	app.Usage = "Proc info tool"
	app.Description = appDescription
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			// Code for meminfo here
		},
	}

	// move this in the app command section
	memInfo, err := os.Open(procFilePath("meminfo"))
	if err != nil {
		log.Fatal("Cannot open file")
	}
	defer memInfo.Close()
}
