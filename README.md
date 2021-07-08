# crc-32 hash collider

CRC-32 collision finder

[CRC-32 algorithm](https://en.wikipedia.org/wiki/Cyclic_redundancy_check#CRC-32_algorithm) outputs a 32-bit unsigned value and therefore can be easily bruteforced to find hash collisions.

This code uses the IEEE polynomial, however can be easily modified to other polynomials.

Written because I was looking for a tool to generate CRC-32 collisions during a CTF but couldn't find any on github.

## Usage

To use this tool, simply modify the crc-32 target value and run:

```bash
go run collide.go
```

## Example

Looking to find a collision for CRC-32 value: -432570933

_Note: old python versions generated signed integers and therefore allowed negative CRC-32 values_

```golang
// target CRC-32
// & 0xffffffff is to convert to uint
// required since old python versions allowed negative values to be produced
// hence its needed if you want to find a collision for a "negative" crc hash value
const target = -432570933 & 0xffffffff

// max string length
maxLen := 5
```

Running it produces the following output:
```
$ go run collide.go
Collision found: 4iSg@
```

Which can be verified using `ipython`:
```
In [1]: import binascii
In [2]: print(binascii.crc32("4iSg@"))
-432570933
```

