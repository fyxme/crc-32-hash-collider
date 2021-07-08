# crc-32 hash collider

CRC-32 collision finder

[CRC-32 algorithm](https://en.wikipedia.org/wiki/Cyclic_redundancy_check#CRC-32_algorithm) outputs a 32-bit unsigned value and therefore can be easily bruteforce to find hash collisions.


To use this tool, simply modify the crc-32 target value and run:

```bash
go run collide.go
```

This code uses the IEEE polynomial, however can be easily modified to other polynomials.

Written because I was looking for a tool to generate CRC-32 collisions during a CTF but couldn't find any on github.
