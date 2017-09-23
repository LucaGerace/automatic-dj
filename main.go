package main

import "fmt"

//TODO FINISH THE GD LOGIC

type amp struct {
	left  []float64
	right []float64
	beats []int
	avE   float64
	inE   []float64
	tempE float64
}

func NewAmp() *amp {
	return &amp{
	//TODO fill in the variables
	}
}

func (a *amp) findbeats() {
	for i := 0; i < len(amp{left}); i++ {
		amp{tempE} += ((amp{left[i]} * amp{right[i]}) + (amp{right[i]} * amp{right[i]}))
		if i%1024 == 0 {
			amp{inE[i/1024]} = (amp{left[i]} * amp{left[i]}) + (amp{right[i]} * amp{right[i]})
		}

	}
	amp{avE} = (1024 / 44032) * tempE
	for i := 0; i < len(amp{left[i]}); i += 1024 {
		if (amp{inE[i]} > (amp{avE} * 1.3)) {
			amp{beat[i]} = 1
			continue
		}
		amp{beat[i]} = 0
	}
}

func main() {
	//shift buffer to the right 1024 samples, eliminate 'left 1024 samples, repeat until end of file
	//XXX can't set entire array to a single value
	//ampL[:1024] = 0
	//ampR[:1024] = 0

	for i := 0; i < (len(ampL) - 1024); i++ {
		ampL[i] = ampL[i+1024]
		ampR[i] = ampR[i+1024]
	}

	for i := (len(ampL) - 1024); i < len(ampL); i++ {
		ampL[i] = 0 //read amplitude of song
		//ampR[i] = //read amplitude of song
	}

	//TODO iterate that function for the rest of the song

	//divide number of beats by song length to determine BPM and print to screen

	var beatnum int = 0
	//var slen = //songlength in seconds

	for i := 0; i < len(beat); i++ {
		if beat[i] == 1 {
			beatnum++
		}
	}
	var BPM = (beatnum / slen) * 60
	fmt.Println("Your song's BPM is: " + BPM + ", you little twerp!")

}
