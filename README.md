# 1BRC

1️⃣🐝🏎️ [The One Billion Row Challenge](https://github.com/gunnarmorling/1brc) -- A fun exploration of how quickly 1B rows from a text file can be aggregated. The challenge was primarily foces on Java but I decided to solve it in Golang, while learning the language! 

## Record of iterations
#### Attempt 1 - Naive Approach
As far as the algorithm goes for this problem, I thought it was fairly trivial.

I used a hashmap to store each station name and allocated a size 4 array to store [min, max, sum, count] for each respective station. I knew I had a long way to go to optimize performance, but I wanted to see where this would get me.

I passed some of the test cases after messing with rounding. My program was functionally correct, but I ran into to a bunch of OOM (Out of Memory) errors when running the 1 billion row file.

My suspicions are on not using a pointer as a reference for the array and making a copy each time (crazy large input)
