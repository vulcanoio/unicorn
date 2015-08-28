package unicorn
// For Unicorn Engine. AUTO-GENERATED FILE, DO NOT EDIT [arm64_const.go]
const (

// ARM64 registers

	UC_ARM64_REG_INVALID = 0
	UC_ARM64_REG_X29 = 1
	UC_ARM64_REG_X30 = 2
	UC_ARM64_REG_NZCV = 3
	UC_ARM64_REG_SP = 4
	UC_ARM64_REG_WSP = 5
	UC_ARM64_REG_WZR = 6
	UC_ARM64_REG_XZR = 7
	UC_ARM64_REG_B0 = 8
	UC_ARM64_REG_B1 = 9
	UC_ARM64_REG_B2 = 10
	UC_ARM64_REG_B3 = 11
	UC_ARM64_REG_B4 = 12
	UC_ARM64_REG_B5 = 13
	UC_ARM64_REG_B6 = 14
	UC_ARM64_REG_B7 = 15
	UC_ARM64_REG_B8 = 16
	UC_ARM64_REG_B9 = 17
	UC_ARM64_REG_B10 = 18
	UC_ARM64_REG_B11 = 19
	UC_ARM64_REG_B12 = 20
	UC_ARM64_REG_B13 = 21
	UC_ARM64_REG_B14 = 22
	UC_ARM64_REG_B15 = 23
	UC_ARM64_REG_B16 = 24
	UC_ARM64_REG_B17 = 25
	UC_ARM64_REG_B18 = 26
	UC_ARM64_REG_B19 = 27
	UC_ARM64_REG_B20 = 28
	UC_ARM64_REG_B21 = 29
	UC_ARM64_REG_B22 = 30
	UC_ARM64_REG_B23 = 31
	UC_ARM64_REG_B24 = 32
	UC_ARM64_REG_B25 = 33
	UC_ARM64_REG_B26 = 34
	UC_ARM64_REG_B27 = 35
	UC_ARM64_REG_B28 = 36
	UC_ARM64_REG_B29 = 37
	UC_ARM64_REG_B30 = 38
	UC_ARM64_REG_B31 = 39
	UC_ARM64_REG_D0 = 40
	UC_ARM64_REG_D1 = 41
	UC_ARM64_REG_D2 = 42
	UC_ARM64_REG_D3 = 43
	UC_ARM64_REG_D4 = 44
	UC_ARM64_REG_D5 = 45
	UC_ARM64_REG_D6 = 46
	UC_ARM64_REG_D7 = 47
	UC_ARM64_REG_D8 = 48
	UC_ARM64_REG_D9 = 49
	UC_ARM64_REG_D10 = 50
	UC_ARM64_REG_D11 = 51
	UC_ARM64_REG_D12 = 52
	UC_ARM64_REG_D13 = 53
	UC_ARM64_REG_D14 = 54
	UC_ARM64_REG_D15 = 55
	UC_ARM64_REG_D16 = 56
	UC_ARM64_REG_D17 = 57
	UC_ARM64_REG_D18 = 58
	UC_ARM64_REG_D19 = 59
	UC_ARM64_REG_D20 = 60
	UC_ARM64_REG_D21 = 61
	UC_ARM64_REG_D22 = 62
	UC_ARM64_REG_D23 = 63
	UC_ARM64_REG_D24 = 64
	UC_ARM64_REG_D25 = 65
	UC_ARM64_REG_D26 = 66
	UC_ARM64_REG_D27 = 67
	UC_ARM64_REG_D28 = 68
	UC_ARM64_REG_D29 = 69
	UC_ARM64_REG_D30 = 70
	UC_ARM64_REG_D31 = 71
	UC_ARM64_REG_H0 = 72
	UC_ARM64_REG_H1 = 73
	UC_ARM64_REG_H2 = 74
	UC_ARM64_REG_H3 = 75
	UC_ARM64_REG_H4 = 76
	UC_ARM64_REG_H5 = 77
	UC_ARM64_REG_H6 = 78
	UC_ARM64_REG_H7 = 79
	UC_ARM64_REG_H8 = 80
	UC_ARM64_REG_H9 = 81
	UC_ARM64_REG_H10 = 82
	UC_ARM64_REG_H11 = 83
	UC_ARM64_REG_H12 = 84
	UC_ARM64_REG_H13 = 85
	UC_ARM64_REG_H14 = 86
	UC_ARM64_REG_H15 = 87
	UC_ARM64_REG_H16 = 88
	UC_ARM64_REG_H17 = 89
	UC_ARM64_REG_H18 = 90
	UC_ARM64_REG_H19 = 91
	UC_ARM64_REG_H20 = 92
	UC_ARM64_REG_H21 = 93
	UC_ARM64_REG_H22 = 94
	UC_ARM64_REG_H23 = 95
	UC_ARM64_REG_H24 = 96
	UC_ARM64_REG_H25 = 97
	UC_ARM64_REG_H26 = 98
	UC_ARM64_REG_H27 = 99
	UC_ARM64_REG_H28 = 100
	UC_ARM64_REG_H29 = 101
	UC_ARM64_REG_H30 = 102
	UC_ARM64_REG_H31 = 103
	UC_ARM64_REG_Q0 = 104
	UC_ARM64_REG_Q1 = 105
	UC_ARM64_REG_Q2 = 106
	UC_ARM64_REG_Q3 = 107
	UC_ARM64_REG_Q4 = 108
	UC_ARM64_REG_Q5 = 109
	UC_ARM64_REG_Q6 = 110
	UC_ARM64_REG_Q7 = 111
	UC_ARM64_REG_Q8 = 112
	UC_ARM64_REG_Q9 = 113
	UC_ARM64_REG_Q10 = 114
	UC_ARM64_REG_Q11 = 115
	UC_ARM64_REG_Q12 = 116
	UC_ARM64_REG_Q13 = 117
	UC_ARM64_REG_Q14 = 118
	UC_ARM64_REG_Q15 = 119
	UC_ARM64_REG_Q16 = 120
	UC_ARM64_REG_Q17 = 121
	UC_ARM64_REG_Q18 = 122
	UC_ARM64_REG_Q19 = 123
	UC_ARM64_REG_Q20 = 124
	UC_ARM64_REG_Q21 = 125
	UC_ARM64_REG_Q22 = 126
	UC_ARM64_REG_Q23 = 127
	UC_ARM64_REG_Q24 = 128
	UC_ARM64_REG_Q25 = 129
	UC_ARM64_REG_Q26 = 130
	UC_ARM64_REG_Q27 = 131
	UC_ARM64_REG_Q28 = 132
	UC_ARM64_REG_Q29 = 133
	UC_ARM64_REG_Q30 = 134
	UC_ARM64_REG_Q31 = 135
	UC_ARM64_REG_S0 = 136
	UC_ARM64_REG_S1 = 137
	UC_ARM64_REG_S2 = 138
	UC_ARM64_REG_S3 = 139
	UC_ARM64_REG_S4 = 140
	UC_ARM64_REG_S5 = 141
	UC_ARM64_REG_S6 = 142
	UC_ARM64_REG_S7 = 143
	UC_ARM64_REG_S8 = 144
	UC_ARM64_REG_S9 = 145
	UC_ARM64_REG_S10 = 146
	UC_ARM64_REG_S11 = 147
	UC_ARM64_REG_S12 = 148
	UC_ARM64_REG_S13 = 149
	UC_ARM64_REG_S14 = 150
	UC_ARM64_REG_S15 = 151
	UC_ARM64_REG_S16 = 152
	UC_ARM64_REG_S17 = 153
	UC_ARM64_REG_S18 = 154
	UC_ARM64_REG_S19 = 155
	UC_ARM64_REG_S20 = 156
	UC_ARM64_REG_S21 = 157
	UC_ARM64_REG_S22 = 158
	UC_ARM64_REG_S23 = 159
	UC_ARM64_REG_S24 = 160
	UC_ARM64_REG_S25 = 161
	UC_ARM64_REG_S26 = 162
	UC_ARM64_REG_S27 = 163
	UC_ARM64_REG_S28 = 164
	UC_ARM64_REG_S29 = 165
	UC_ARM64_REG_S30 = 166
	UC_ARM64_REG_S31 = 167
	UC_ARM64_REG_W0 = 168
	UC_ARM64_REG_W1 = 169
	UC_ARM64_REG_W2 = 170
	UC_ARM64_REG_W3 = 171
	UC_ARM64_REG_W4 = 172
	UC_ARM64_REG_W5 = 173
	UC_ARM64_REG_W6 = 174
	UC_ARM64_REG_W7 = 175
	UC_ARM64_REG_W8 = 176
	UC_ARM64_REG_W9 = 177
	UC_ARM64_REG_W10 = 178
	UC_ARM64_REG_W11 = 179
	UC_ARM64_REG_W12 = 180
	UC_ARM64_REG_W13 = 181
	UC_ARM64_REG_W14 = 182
	UC_ARM64_REG_W15 = 183
	UC_ARM64_REG_W16 = 184
	UC_ARM64_REG_W17 = 185
	UC_ARM64_REG_W18 = 186
	UC_ARM64_REG_W19 = 187
	UC_ARM64_REG_W20 = 188
	UC_ARM64_REG_W21 = 189
	UC_ARM64_REG_W22 = 190
	UC_ARM64_REG_W23 = 191
	UC_ARM64_REG_W24 = 192
	UC_ARM64_REG_W25 = 193
	UC_ARM64_REG_W26 = 194
	UC_ARM64_REG_W27 = 195
	UC_ARM64_REG_W28 = 196
	UC_ARM64_REG_W29 = 197
	UC_ARM64_REG_W30 = 198
	UC_ARM64_REG_X0 = 199
	UC_ARM64_REG_X1 = 200
	UC_ARM64_REG_X2 = 201
	UC_ARM64_REG_X3 = 202
	UC_ARM64_REG_X4 = 203
	UC_ARM64_REG_X5 = 204
	UC_ARM64_REG_X6 = 205
	UC_ARM64_REG_X7 = 206
	UC_ARM64_REG_X8 = 207
	UC_ARM64_REG_X9 = 208
	UC_ARM64_REG_X10 = 209
	UC_ARM64_REG_X11 = 210
	UC_ARM64_REG_X12 = 211
	UC_ARM64_REG_X13 = 212
	UC_ARM64_REG_X14 = 213
	UC_ARM64_REG_X15 = 214
	UC_ARM64_REG_X16 = 215
	UC_ARM64_REG_X17 = 216
	UC_ARM64_REG_X18 = 217
	UC_ARM64_REG_X19 = 218
	UC_ARM64_REG_X20 = 219
	UC_ARM64_REG_X21 = 220
	UC_ARM64_REG_X22 = 221
	UC_ARM64_REG_X23 = 222
	UC_ARM64_REG_X24 = 223
	UC_ARM64_REG_X25 = 224
	UC_ARM64_REG_X26 = 225
	UC_ARM64_REG_X27 = 226
	UC_ARM64_REG_X28 = 227
	UC_ARM64_REG_V0 = 228
	UC_ARM64_REG_V1 = 229
	UC_ARM64_REG_V2 = 230
	UC_ARM64_REG_V3 = 231
	UC_ARM64_REG_V4 = 232
	UC_ARM64_REG_V5 = 233
	UC_ARM64_REG_V6 = 234
	UC_ARM64_REG_V7 = 235
	UC_ARM64_REG_V8 = 236
	UC_ARM64_REG_V9 = 237
	UC_ARM64_REG_V10 = 238
	UC_ARM64_REG_V11 = 239
	UC_ARM64_REG_V12 = 240
	UC_ARM64_REG_V13 = 241
	UC_ARM64_REG_V14 = 242
	UC_ARM64_REG_V15 = 243
	UC_ARM64_REG_V16 = 244
	UC_ARM64_REG_V17 = 245
	UC_ARM64_REG_V18 = 246
	UC_ARM64_REG_V19 = 247
	UC_ARM64_REG_V20 = 248
	UC_ARM64_REG_V21 = 249
	UC_ARM64_REG_V22 = 250
	UC_ARM64_REG_V23 = 251
	UC_ARM64_REG_V24 = 252
	UC_ARM64_REG_V25 = 253
	UC_ARM64_REG_V26 = 254
	UC_ARM64_REG_V27 = 255
	UC_ARM64_REG_V28 = 256
	UC_ARM64_REG_V29 = 257
	UC_ARM64_REG_V30 = 258
	UC_ARM64_REG_V31 = 259

// pseudo registers
	UC_ARM64_REG_PC = 260
	UC_ARM64_REG_ENDING = 261

// alias registers
	UC_ARM64_REG_IP1 = 215
	UC_ARM64_REG_IP0 = 216
	UC_ARM64_REG_FP = 1
	UC_ARM64_REG_LR = 2
)