package main

import "fmt"

//TODO FINISH THE GD LOGIC

type amp struct {
	left  []float64
	right []float64
}
var inE[] float64
var tempE float64
var avE float64
var beatnum int = 0
var beat[] int


//func NewAmp(x int) *amp {
	//TODO reads wav/mp3 and fills in amp left and right

	
	//shift buffer to the right 1024 samples, eliminate 'left 1024 samples, repeat until end of file
	//XXX can't set entire array to a single value
	//ampL[:1024] = 0
	//ampR[:1024] = 0

//	for i := 0; i < (len(amp{left}) - 1024); i++ {
//		amp{left[i]} = amp{left[i+1024]}
//		amp{right[i]} = amp{right[i+1024]}
//	}

//	for i := (len(ampL) - 1024); i < len(ampL); i++ {
		//ampL[i] = //read amplitude of song
		//ampR[i] = //read amplitude of song
//	}
//}


func findbeats() {
	for i := 0; i < len(amp{left}); i++ {
		tempE += ((amp{left[i]} * amp{right[i]}) + (amp{right[i]} * amp{right[i]}))
		if i%1024 == 0 {
			inE[i/1024] = (amp{left[i]} * amp{left[i]}) + (amp{right[i]} * amp{right[i]})
		}

	}
	avE = (1024 / 44032) * tempE
	for i := 0; i < len(amp{left[i]}); i += 1024 {
		if (amp{inE[i]} > (avE * 1.3)) {
			beat[i] = 1
			continue
		}
		beat[i] = 0
	}
}

func main() {


	//TODO iterate that function for the rest of the song

	//divide number of beats by song length to determine BPM and print to screen

	//var slen = //songlength in seconds

	for i := 0; i < len(beat); i++ {
		if beat[i] == 1 {
			beatnum++
		}
	}
	var BPM = (beatnum / slen) * 60
	fmt.Println("Your song's BPM is: " + BPM + ", you little twerp!")

}
