/*
Author Gaurav Sablok
Universitat Potsdam Library
Date: 2024-8-7

A pacbiohifi go profiler that takes the pacbiohifi reads, makes the mers according to the length,
filters the mers according to the critera. This is supported witht a desktop application to see that
if your reads have high profiled genomic mers that could hinder the graph assembly.

*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	readfile, err := os.Open("/home/gauravcodepro/Desktop/coderep/samplefile.fastq")
	if err != nil {
		panic(err)
		log.Fatal(err)
		return
	}
	readbuffer := bufio.NewScanner(readfile)
	header := []string{}
	sequences := []string{}
	header := []string{}
	sequences := []string{}
	for readbuffer.Scan() {
		line := readbuffer.Text()
		if string(line[0]) == "A" || string(line[0]) == "T" || string(line[0]) == "G" || string(line[0]) == "C" {
			sequences = append(sequences, line)
		}
		if string(line[0]) == "@" {
			header = append(header, line)
		}
	}

	seqtok := []string{}
	for i := range sequences {
		// change the profile kmer here according to your choice or check the fyne desktop application where 
		// you can do everything visuallly. 
		for j := 0; j <= len(sequences[i])-2; j++ {
			seqtok = append(seqtok, string(sequences[i][j:j+2]))
		}
	}

	kmercomp := []int{}
	for i := range seqtok {
		hold := string(seqtok[i])
		kmercomp = append(kmercomp, strings.Count(hold, "A")+strings.Count(hold, "T")+strings.Count(hold, "G")+strings.Count(hold, "C"))
	}

	kmercompGC := []int{}
	for i := range seqtok {
		hold := string(seqtok[i])
		kmercomp = append(kmercomp, strings.Count(hold, "G")+strings.Count(hold, "C"))
	}

	kmerclassify := []string{}
	kmerfilter := []int{}
	for i := range seqtok {
		kmerclassify = append(kmerclassify, seqtok[i])
		kmerfilter = append(kmerfilter, int(kmercompGC[i])/int(kmercomp[i]))
	}

	filteredKmer := []string{}
	for i := range kmerclassify {
		// mention the threshold for the filtering of the low kmers. default value set to 10 and fyne GO application allows to select based on the plotting graph. 
		if kmerfilter[i] < 10 {
			filteredKmer = append(filteredKmer, kmerclassify[i])
		}
	}

	for i := range filteredKmer {
		fmt.Println("The filtered kmers with the elective selection are %s and their length are %T:", string(filteredKmer[i]), len(string(filteredKmer[i])))
	}

}
