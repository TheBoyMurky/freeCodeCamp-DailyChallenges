package main

import "fmt"

func complementaryDNA(dna string) string {
	var comp_dna []rune
	for i := 0; i < len(dna); i++ {
		switch dna[i] {
		case 'T':
			comp_dna = append(comp_dna, 'A')
		case 'A':
			comp_dna = append(comp_dna, 'T')
		case 'C':
			comp_dna = append(comp_dna, 'G')
		case 'G':
			comp_dna = append(comp_dna, 'C')
		}
	}
	return string(comp_dna)
}

func main() {
	fmt.Println(complementaryDNA("ATGC"))
	fmt.Println(complementaryDNA("GCAT"))
	fmt.Println(complementaryDNA("TACG"))
}
