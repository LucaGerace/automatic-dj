//slice song into segments, 44032/second

//load 1s of song into L/R buffer B[0,n], B[1,n]

//every 1024 segments, determine amplitude on left and right channels (a[n], b[n)

//compute instant energy, e = sum(a[n]^2 + b [n]^2) 

//compute local average energy E = sum (B[0,n]^2 + B[1,n]^2) for n=0 to 44032

//shift buffer to "right" 1024 samples, eliminate 'left' 1024 samples

//compare 'e' to 'E*C' C=1.3 (apparently a reasonable value?????)

//if e > E*C, beat detected!

//iterate over entire song, determine total number of beats

//divide number of beats by song length, BAMMO there's the BPM you little TWERP


