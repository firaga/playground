// fibonacci_amd64.s
#include "textflag.h"

TEXT ·fibonacciAwesome(SB), NOSPLIT, $0
MOVQ n+0(FP), AX
CMPQ AX, $2
JLE SMALLN   //n <= 2, jump to the end to return 1
MOVQ $1, CX
MOVQ $1, DX
JMP COND
LOOP:
DECQ AX
XADDQ CX, DX // DX, CX <= DX+CX, DX
COND:
CMPQ AX, $2
JGT LOOP
MOVQ DX, r+8(FP)
RET
SMALLN:
MOVQ $1, r+8(FP)
RET
