# go-ex
Go experience

# Initialize module.

go mod main

# %d: Decimal integer
# %s: String
# %f: Floating-point number
# %t: Boolean
# %v: Value in a default format
# %g: read in the internet

Basic types in GO:

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

# T(v) - converts v to T type

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)