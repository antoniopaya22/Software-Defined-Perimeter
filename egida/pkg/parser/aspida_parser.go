// Code generated from Aspida.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Aspida
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 52, 304,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 63, 10, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 7, 7, 90, 10,
	7, 12, 7, 14, 7, 93, 11, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 5, 8, 104, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 112,
	10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 120, 10, 10, 3,
	11, 3, 11, 5, 11, 124, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	5, 12, 132, 10, 12, 3, 13, 3, 13, 7, 13, 136, 10, 13, 12, 13, 14, 13, 139,
	11, 13, 3, 13, 3, 13, 7, 13, 143, 10, 13, 12, 13, 14, 13, 146, 11, 13,
	3, 13, 3, 13, 5, 13, 150, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5,
	14, 157, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 5, 15, 169, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 181, 10, 16, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 193, 10, 17, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18,
	205, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 5, 19, 217, 10, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 23, 3, 23, 3, 23, 7, 23, 239, 10, 23, 12, 23, 14, 23, 242, 11, 23,
	3, 24, 3, 24, 7, 24, 246, 10, 24, 12, 24, 14, 24, 249, 11, 24, 3, 25, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	5, 25, 263, 10, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3,
	27, 5, 27, 273, 10, 27, 3, 28, 3, 28, 3, 28, 3, 28, 7, 28, 279, 10, 28,
	12, 28, 14, 28, 282, 11, 28, 3, 28, 3, 28, 3, 28, 5, 28, 287, 10, 28, 3,
	29, 3, 29, 3, 29, 3, 29, 7, 29, 293, 10, 29, 12, 29, 14, 29, 296, 11, 29,
	3, 29, 3, 29, 3, 29, 3, 29, 5, 29, 302, 10, 29, 3, 29, 2, 2, 30, 2, 4,
	6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	44, 46, 48, 50, 52, 54, 56, 2, 3, 3, 2, 22, 27, 2, 307, 2, 58, 3, 2, 2,
	2, 4, 64, 3, 2, 2, 2, 6, 70, 3, 2, 2, 2, 8, 75, 3, 2, 2, 2, 10, 81, 3,
	2, 2, 2, 12, 87, 3, 2, 2, 2, 14, 103, 3, 2, 2, 2, 16, 111, 3, 2, 2, 2,
	18, 119, 3, 2, 2, 2, 20, 123, 3, 2, 2, 2, 22, 131, 3, 2, 2, 2, 24, 149,
	3, 2, 2, 2, 26, 156, 3, 2, 2, 2, 28, 168, 3, 2, 2, 2, 30, 180, 3, 2, 2,
	2, 32, 192, 3, 2, 2, 2, 34, 204, 3, 2, 2, 2, 36, 216, 3, 2, 2, 2, 38, 218,
	3, 2, 2, 2, 40, 224, 3, 2, 2, 2, 42, 230, 3, 2, 2, 2, 44, 235, 3, 2, 2,
	2, 46, 243, 3, 2, 2, 2, 48, 262, 3, 2, 2, 2, 50, 264, 3, 2, 2, 2, 52, 272,
	3, 2, 2, 2, 54, 286, 3, 2, 2, 2, 56, 301, 3, 2, 2, 2, 58, 59, 5, 4, 3,
	2, 59, 60, 5, 6, 4, 2, 60, 62, 5, 8, 5, 2, 61, 63, 5, 10, 6, 2, 62, 61,
	3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 3, 3, 2, 2, 2, 64, 65, 7, 40, 2, 2,
	65, 66, 7, 3, 2, 2, 66, 67, 7, 4, 2, 2, 67, 68, 5, 12, 7, 2, 68, 69, 7,
	5, 2, 2, 69, 5, 3, 2, 2, 2, 70, 71, 7, 41, 2, 2, 71, 72, 7, 3, 2, 2, 72,
	73, 7, 37, 2, 2, 73, 74, 7, 39, 2, 2, 74, 7, 3, 2, 2, 2, 75, 76, 7, 42,
	2, 2, 76, 77, 7, 3, 2, 2, 77, 78, 7, 4, 2, 2, 78, 79, 5, 24, 13, 2, 79,
	80, 7, 5, 2, 2, 80, 9, 3, 2, 2, 2, 81, 82, 7, 43, 2, 2, 82, 83, 7, 3, 2,
	2, 83, 84, 7, 4, 2, 2, 84, 85, 5, 46, 24, 2, 85, 86, 7, 5, 2, 2, 86, 11,
	3, 2, 2, 2, 87, 91, 5, 14, 8, 2, 88, 90, 5, 14, 8, 2, 89, 88, 3, 2, 2,
	2, 90, 93, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 13,
	3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 94, 95, 5, 16, 9, 2, 95, 96, 7, 39, 2,
	2, 96, 104, 3, 2, 2, 2, 97, 98, 5, 18, 10, 2, 98, 99, 7, 39, 2, 2, 99,
	104, 3, 2, 2, 2, 100, 101, 5, 22, 12, 2, 101, 102, 7, 39, 2, 2, 102, 104,
	3, 2, 2, 2, 103, 94, 3, 2, 2, 2, 103, 97, 3, 2, 2, 2, 103, 100, 3, 2, 2,
	2, 104, 15, 3, 2, 2, 2, 105, 106, 7, 6, 2, 2, 106, 107, 7, 3, 2, 2, 107,
	112, 5, 52, 27, 2, 108, 109, 7, 7, 2, 2, 109, 110, 7, 3, 2, 2, 110, 112,
	5, 52, 27, 2, 111, 105, 3, 2, 2, 2, 111, 108, 3, 2, 2, 2, 112, 17, 3, 2,
	2, 2, 113, 114, 7, 8, 2, 2, 114, 115, 7, 3, 2, 2, 115, 120, 5, 20, 11,
	2, 116, 117, 7, 9, 2, 2, 117, 118, 7, 3, 2, 2, 118, 120, 5, 20, 11, 2,
	119, 113, 3, 2, 2, 2, 119, 116, 3, 2, 2, 2, 120, 19, 3, 2, 2, 2, 121, 124,
	7, 44, 2, 2, 122, 124, 7, 45, 2, 2, 123, 121, 3, 2, 2, 2, 123, 122, 3,
	2, 2, 2, 124, 21, 3, 2, 2, 2, 125, 126, 7, 10, 2, 2, 126, 127, 7, 3, 2,
	2, 127, 132, 5, 52, 27, 2, 128, 129, 7, 11, 2, 2, 129, 130, 7, 3, 2, 2,
	130, 132, 5, 52, 27, 2, 131, 125, 3, 2, 2, 2, 131, 128, 3, 2, 2, 2, 132,
	23, 3, 2, 2, 2, 133, 137, 5, 26, 14, 2, 134, 136, 5, 26, 14, 2, 135, 134,
	3, 2, 2, 2, 136, 139, 3, 2, 2, 2, 137, 135, 3, 2, 2, 2, 137, 138, 3, 2,
	2, 2, 138, 150, 3, 2, 2, 2, 139, 137, 3, 2, 2, 2, 140, 144, 5, 38, 20,
	2, 141, 143, 5, 40, 21, 2, 142, 141, 3, 2, 2, 2, 143, 146, 3, 2, 2, 2,
	144, 142, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145, 147, 3, 2, 2, 2, 146,
	144, 3, 2, 2, 2, 147, 148, 5, 42, 22, 2, 148, 150, 3, 2, 2, 2, 149, 133,
	3, 2, 2, 2, 149, 140, 3, 2, 2, 2, 150, 25, 3, 2, 2, 2, 151, 157, 5, 28,
	15, 2, 152, 157, 5, 30, 16, 2, 153, 157, 5, 32, 17, 2, 154, 157, 5, 34,
	18, 2, 155, 157, 5, 36, 19, 2, 156, 151, 3, 2, 2, 2, 156, 152, 3, 2, 2,
	2, 156, 153, 3, 2, 2, 2, 156, 154, 3, 2, 2, 2, 156, 155, 3, 2, 2, 2, 157,
	27, 3, 2, 2, 2, 158, 159, 7, 12, 2, 2, 159, 160, 7, 3, 2, 2, 160, 161,
	5, 54, 28, 2, 161, 162, 7, 39, 2, 2, 162, 169, 3, 2, 2, 2, 163, 164, 7,
	13, 2, 2, 164, 165, 7, 3, 2, 2, 165, 166, 5, 54, 28, 2, 166, 167, 7, 39,
	2, 2, 167, 169, 3, 2, 2, 2, 168, 158, 3, 2, 2, 2, 168, 163, 3, 2, 2, 2,
	169, 29, 3, 2, 2, 2, 170, 171, 7, 14, 2, 2, 171, 172, 7, 3, 2, 2, 172,
	173, 5, 54, 28, 2, 173, 174, 7, 39, 2, 2, 174, 181, 3, 2, 2, 2, 175, 176,
	7, 15, 2, 2, 176, 177, 7, 3, 2, 2, 177, 178, 5, 54, 28, 2, 178, 179, 7,
	39, 2, 2, 179, 181, 3, 2, 2, 2, 180, 170, 3, 2, 2, 2, 180, 175, 3, 2, 2,
	2, 181, 31, 3, 2, 2, 2, 182, 183, 7, 16, 2, 2, 183, 184, 7, 3, 2, 2, 184,
	185, 5, 54, 28, 2, 185, 186, 7, 39, 2, 2, 186, 193, 3, 2, 2, 2, 187, 188,
	7, 17, 2, 2, 188, 189, 7, 3, 2, 2, 189, 190, 5, 54, 28, 2, 190, 191, 7,
	39, 2, 2, 191, 193, 3, 2, 2, 2, 192, 182, 3, 2, 2, 2, 192, 187, 3, 2, 2,
	2, 193, 33, 3, 2, 2, 2, 194, 195, 7, 18, 2, 2, 195, 196, 7, 3, 2, 2, 196,
	197, 5, 54, 28, 2, 197, 198, 7, 39, 2, 2, 198, 205, 3, 2, 2, 2, 199, 200,
	7, 19, 2, 2, 200, 201, 7, 3, 2, 2, 201, 202, 5, 54, 28, 2, 202, 203, 7,
	39, 2, 2, 203, 205, 3, 2, 2, 2, 204, 194, 3, 2, 2, 2, 204, 199, 3, 2, 2,
	2, 205, 35, 3, 2, 2, 2, 206, 207, 7, 20, 2, 2, 207, 208, 7, 3, 2, 2, 208,
	209, 5, 54, 28, 2, 209, 210, 7, 39, 2, 2, 210, 217, 3, 2, 2, 2, 211, 212,
	7, 21, 2, 2, 212, 213, 7, 3, 2, 2, 213, 214, 5, 54, 28, 2, 214, 215, 7,
	39, 2, 2, 215, 217, 3, 2, 2, 2, 216, 206, 3, 2, 2, 2, 216, 211, 3, 2, 2,
	2, 217, 37, 3, 2, 2, 2, 218, 219, 7, 46, 2, 2, 219, 220, 5, 44, 23, 2,
	220, 221, 7, 4, 2, 2, 221, 222, 5, 24, 13, 2, 222, 223, 7, 5, 2, 2, 223,
	39, 3, 2, 2, 2, 224, 225, 7, 47, 2, 2, 225, 226, 5, 44, 23, 2, 226, 227,
	7, 4, 2, 2, 227, 228, 5, 24, 13, 2, 228, 229, 7, 5, 2, 2, 229, 41, 3, 2,
	2, 2, 230, 231, 7, 48, 2, 2, 231, 232, 7, 4, 2, 2, 232, 233, 5, 24, 13,
	2, 233, 234, 7, 5, 2, 2, 234, 43, 3, 2, 2, 2, 235, 236, 5, 52, 27, 2, 236,
	240, 5, 50, 26, 2, 237, 239, 5, 52, 27, 2, 238, 237, 3, 2, 2, 2, 239, 242,
	3, 2, 2, 2, 240, 238, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241, 45, 3, 2,
	2, 2, 242, 240, 3, 2, 2, 2, 243, 247, 5, 48, 25, 2, 244, 246, 5, 48, 25,
	2, 245, 244, 3, 2, 2, 2, 246, 249, 3, 2, 2, 2, 247, 245, 3, 2, 2, 2, 247,
	248, 3, 2, 2, 2, 248, 47, 3, 2, 2, 2, 249, 247, 3, 2, 2, 2, 250, 251, 7,
	37, 2, 2, 251, 252, 7, 3, 2, 2, 252, 253, 5, 52, 27, 2, 253, 254, 7, 39,
	2, 2, 254, 263, 3, 2, 2, 2, 255, 256, 7, 37, 2, 2, 256, 257, 7, 3, 2, 2,
	257, 258, 7, 4, 2, 2, 258, 259, 5, 46, 24, 2, 259, 260, 7, 5, 2, 2, 260,
	261, 7, 39, 2, 2, 261, 263, 3, 2, 2, 2, 262, 250, 3, 2, 2, 2, 262, 255,
	3, 2, 2, 2, 263, 49, 3, 2, 2, 2, 264, 265, 9, 2, 2, 2, 265, 51, 3, 2, 2,
	2, 266, 273, 7, 37, 2, 2, 267, 273, 7, 38, 2, 2, 268, 273, 7, 28, 2, 2,
	269, 273, 7, 29, 2, 2, 270, 273, 7, 30, 2, 2, 271, 273, 5, 56, 29, 2, 272,
	266, 3, 2, 2, 2, 272, 267, 3, 2, 2, 2, 272, 268, 3, 2, 2, 2, 272, 269,
	3, 2, 2, 2, 272, 270, 3, 2, 2, 2, 272, 271, 3, 2, 2, 2, 273, 53, 3, 2,
	2, 2, 274, 275, 7, 31, 2, 2, 275, 280, 7, 37, 2, 2, 276, 277, 7, 32, 2,
	2, 277, 279, 7, 37, 2, 2, 278, 276, 3, 2, 2, 2, 279, 282, 3, 2, 2, 2, 280,
	278, 3, 2, 2, 2, 280, 281, 3, 2, 2, 2, 281, 283, 3, 2, 2, 2, 282, 280,
	3, 2, 2, 2, 283, 287, 7, 33, 2, 2, 284, 285, 7, 31, 2, 2, 285, 287, 7,
	33, 2, 2, 286, 274, 3, 2, 2, 2, 286, 284, 3, 2, 2, 2, 287, 55, 3, 2, 2,
	2, 288, 289, 7, 31, 2, 2, 289, 294, 5, 52, 27, 2, 290, 291, 7, 32, 2, 2,
	291, 293, 5, 52, 27, 2, 292, 290, 3, 2, 2, 2, 293, 296, 3, 2, 2, 2, 294,
	292, 3, 2, 2, 2, 294, 295, 3, 2, 2, 2, 295, 297, 3, 2, 2, 2, 296, 294,
	3, 2, 2, 2, 297, 298, 7, 33, 2, 2, 298, 302, 3, 2, 2, 2, 299, 300, 7, 31,
	2, 2, 300, 302, 7, 33, 2, 2, 301, 288, 3, 2, 2, 2, 301, 299, 3, 2, 2, 2,
	302, 57, 3, 2, 2, 2, 26, 62, 91, 103, 111, 119, 123, 131, 137, 144, 149,
	156, 168, 180, 192, 204, 216, 240, 247, 262, 272, 280, 286, 294, 301,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "':'", "'{'", "'}'", "'name'", "'NAME'", "'connection'", "'CONNECTION'",
	"'description'", "'DESCRIPTION'", "'sections'", "'SECTIONS'", "'points'",
	"'POINTS'", "'controls'", "'CONTROLS'", "'exclusions'", "'EXCLUSIONS'",
	"'tags'", "'TAGS'", "'<'", "'>'", "'=='", "'>='", "'<='", "'!='", "'true'",
	"'false'", "'null'", "'['", "','", "']'", "", "", "", "", "", "';'", "'MAIN'",
	"'HOST'", "'TASKS'", "'VARS'", "", "", "'IF'", "'ELIF'", "'ELSE'", "'OR'",
	"'AND'", "'NOT'", "'IS'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "COMMENT", "WHITESPACE",
	"NEWLINE", "STRING", "NUMBER", "NS", "MAIN_KW", "HOSTS_KW", "TASKS_KW",
	"VARS_KW", "LOCAL_KW", "SSH_KW", "IF", "ELIF", "ELSE", "OR", "AND", "NOT",
	"IS",
}

var ruleNames = []string{
	"program", "main", "hosts", "tasks", "variables", "main_content", "main_prop",
	"name", "connection", "connection_type", "description", "tasks_content",
	"tasks_prop", "sections", "points", "controls", "exclusions", "tags", "ifStat",
	"elifStat", "elseStat", "comparison", "vars_content", "vars_prop", "comp_op",
	"value", "str_array", "array",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type AspidaParser struct {
	*antlr.BaseParser
}

func NewAspidaParser(input antlr.TokenStream) *AspidaParser {
	this := new(AspidaParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Aspida.g4"

	return this
}

// AspidaParser tokens.
const (
	AspidaParserEOF        = antlr.TokenEOF
	AspidaParserT__0       = 1
	AspidaParserT__1       = 2
	AspidaParserT__2       = 3
	AspidaParserT__3       = 4
	AspidaParserT__4       = 5
	AspidaParserT__5       = 6
	AspidaParserT__6       = 7
	AspidaParserT__7       = 8
	AspidaParserT__8       = 9
	AspidaParserT__9       = 10
	AspidaParserT__10      = 11
	AspidaParserT__11      = 12
	AspidaParserT__12      = 13
	AspidaParserT__13      = 14
	AspidaParserT__14      = 15
	AspidaParserT__15      = 16
	AspidaParserT__16      = 17
	AspidaParserT__17      = 18
	AspidaParserT__18      = 19
	AspidaParserT__19      = 20
	AspidaParserT__20      = 21
	AspidaParserT__21      = 22
	AspidaParserT__22      = 23
	AspidaParserT__23      = 24
	AspidaParserT__24      = 25
	AspidaParserT__25      = 26
	AspidaParserT__26      = 27
	AspidaParserT__27      = 28
	AspidaParserT__28      = 29
	AspidaParserT__29      = 30
	AspidaParserT__30      = 31
	AspidaParserCOMMENT    = 32
	AspidaParserWHITESPACE = 33
	AspidaParserNEWLINE    = 34
	AspidaParserSTRING     = 35
	AspidaParserNUMBER     = 36
	AspidaParserNS         = 37
	AspidaParserMAIN_KW    = 38
	AspidaParserHOSTS_KW   = 39
	AspidaParserTASKS_KW   = 40
	AspidaParserVARS_KW    = 41
	AspidaParserLOCAL_KW   = 42
	AspidaParserSSH_KW     = 43
	AspidaParserIF         = 44
	AspidaParserELIF       = 45
	AspidaParserELSE       = 46
	AspidaParserOR         = 47
	AspidaParserAND        = 48
	AspidaParserNOT        = 49
	AspidaParserIS         = 50
)

// AspidaParser rules.
const (
	AspidaParserRULE_program         = 0
	AspidaParserRULE_main            = 1
	AspidaParserRULE_hosts           = 2
	AspidaParserRULE_tasks           = 3
	AspidaParserRULE_variables       = 4
	AspidaParserRULE_main_content    = 5
	AspidaParserRULE_main_prop       = 6
	AspidaParserRULE_name            = 7
	AspidaParserRULE_connection      = 8
	AspidaParserRULE_connection_type = 9
	AspidaParserRULE_description     = 10
	AspidaParserRULE_tasks_content   = 11
	AspidaParserRULE_tasks_prop      = 12
	AspidaParserRULE_sections        = 13
	AspidaParserRULE_points          = 14
	AspidaParserRULE_controls        = 15
	AspidaParserRULE_exclusions      = 16
	AspidaParserRULE_tags            = 17
	AspidaParserRULE_ifStat          = 18
	AspidaParserRULE_elifStat        = 19
	AspidaParserRULE_elseStat        = 20
	AspidaParserRULE_comparison      = 21
	AspidaParserRULE_vars_content    = 22
	AspidaParserRULE_vars_prop       = 23
	AspidaParserRULE_comp_op         = 24
	AspidaParserRULE_value           = 25
	AspidaParserRULE_str_array       = 26
	AspidaParserRULE_array           = 27
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Main() IMainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMainContext)
}

func (s *ProgramContext) Hosts() IHostsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHostsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHostsContext)
}

func (s *ProgramContext) Tasks() ITasksContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITasksContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITasksContext)
}

func (s *ProgramContext) Variables() IVariablesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariablesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariablesContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AspidaParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Main()
	}
	{
		p.SetState(57)
		p.Hosts()
	}
	{
		p.SetState(58)
		p.Tasks()
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AspidaParserVARS_KW {
		{
			p.SetState(59)
			p.Variables()
		}

	}

	return localctx
}

// IMainContext is an interface to support dynamic dispatch.
type IMainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMainContext differentiates from other interfaces.
	IsMainContext()
}

type MainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainContext() *MainContext {
	var p = new(MainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_main
	return p
}

func (*MainContext) IsMainContext() {}

func NewMainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainContext {
	var p = new(MainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_main

	return p
}

func (s *MainContext) GetParser() antlr.Parser { return s.parser }

func (s *MainContext) MAIN_KW() antlr.TerminalNode {
	return s.GetToken(AspidaParserMAIN_KW, 0)
}

func (s *MainContext) Main_content() IMain_contentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMain_contentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMain_contentContext)
}

func (s *MainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterMain(s)
	}
}

func (s *MainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitMain(s)
	}
}

func (s *MainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitMain(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) Main() (localctx IMainContext) {
	localctx = NewMainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AspidaParserRULE_main)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Match(AspidaParserMAIN_KW)
	}
	{
		p.SetState(63)
		p.Match(AspidaParserT__0)
	}
	{
		p.SetState(64)
		p.Match(AspidaParserT__1)
	}
	{
		p.SetState(65)
		p.Main_content()
	}
	{
		p.SetState(66)
		p.Match(AspidaParserT__2)
	}

	return localctx
}

// IHostsContext is an interface to support dynamic dispatch.
type IHostsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHostsContext differentiates from other interfaces.
	IsHostsContext()
}

type HostsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHostsContext() *HostsContext {
	var p = new(HostsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_hosts
	return p
}

func (*HostsContext) IsHostsContext() {}

func NewHostsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HostsContext {
	var p = new(HostsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_hosts

	return p
}

func (s *HostsContext) GetParser() antlr.Parser { return s.parser }

func (s *HostsContext) HOSTS_KW() antlr.TerminalNode {
	return s.GetToken(AspidaParserHOSTS_KW, 0)
}

func (s *HostsContext) STRING() antlr.TerminalNode {
	return s.GetToken(AspidaParserSTRING, 0)
}

func (s *HostsContext) NS() antlr.TerminalNode {
	return s.GetToken(AspidaParserNS, 0)
}

func (s *HostsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HostsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HostsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterHosts(s)
	}
}

func (s *HostsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitHosts(s)
	}
}

func (s *HostsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitHosts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) Hosts() (localctx IHostsContext) {
	localctx = NewHostsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AspidaParserRULE_hosts)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Match(AspidaParserHOSTS_KW)
	}
	{
		p.SetState(69)
		p.Match(AspidaParserT__0)
	}
	{
		p.SetState(70)
		p.Match(AspidaParserSTRING)
	}
	{
		p.SetState(71)
		p.Match(AspidaParserNS)
	}

	return localctx
}

// ITasksContext is an interface to support dynamic dispatch.
type ITasksContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTasksContext differentiates from other interfaces.
	IsTasksContext()
}

type TasksContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTasksContext() *TasksContext {
	var p = new(TasksContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_tasks
	return p
}

func (*TasksContext) IsTasksContext() {}

func NewTasksContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TasksContext {
	var p = new(TasksContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_tasks

	return p
}

func (s *TasksContext) GetParser() antlr.Parser { return s.parser }

func (s *TasksContext) TASKS_KW() antlr.TerminalNode {
	return s.GetToken(AspidaParserTASKS_KW, 0)
}

func (s *TasksContext) Tasks_content() ITasks_contentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITasks_contentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITasks_contentContext)
}

func (s *TasksContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TasksContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TasksContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterTasks(s)
	}
}

func (s *TasksContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitTasks(s)
	}
}

func (s *TasksContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitTasks(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) Tasks() (localctx ITasksContext) {
	localctx = NewTasksContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AspidaParserRULE_tasks)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)
		p.Match(AspidaParserTASKS_KW)
	}
	{
		p.SetState(74)
		p.Match(AspidaParserT__0)
	}
	{
		p.SetState(75)
		p.Match(AspidaParserT__1)
	}
	{
		p.SetState(76)
		p.Tasks_content()
	}
	{
		p.SetState(77)
		p.Match(AspidaParserT__2)
	}

	return localctx
}

// IVariablesContext is an interface to support dynamic dispatch.
type IVariablesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariablesContext differentiates from other interfaces.
	IsVariablesContext()
}

type VariablesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariablesContext() *VariablesContext {
	var p = new(VariablesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_variables
	return p
}

func (*VariablesContext) IsVariablesContext() {}

func NewVariablesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariablesContext {
	var p = new(VariablesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_variables

	return p
}

func (s *VariablesContext) GetParser() antlr.Parser { return s.parser }

func (s *VariablesContext) VARS_KW() antlr.TerminalNode {
	return s.GetToken(AspidaParserVARS_KW, 0)
}

func (s *VariablesContext) Vars_content() IVars_contentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVars_contentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVars_contentContext)
}

func (s *VariablesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariablesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariablesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterVariables(s)
	}
}

func (s *VariablesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitVariables(s)
	}
}

func (s *VariablesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitVariables(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) Variables() (localctx IVariablesContext) {
	localctx = NewVariablesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AspidaParserRULE_variables)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(AspidaParserVARS_KW)
	}
	{
		p.SetState(80)
		p.Match(AspidaParserT__0)
	}
	{
		p.SetState(81)
		p.Match(AspidaParserT__1)
	}
	{
		p.SetState(82)
		p.Vars_content()
	}
	{
		p.SetState(83)
		p.Match(AspidaParserT__2)
	}

	return localctx
}

// IMain_contentContext is an interface to support dynamic dispatch.
type IMain_contentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMain_contentContext differentiates from other interfaces.
	IsMain_contentContext()
}

type Main_contentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMain_contentContext() *Main_contentContext {
	var p = new(Main_contentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_main_content
	return p
}

func (*Main_contentContext) IsMain_contentContext() {}

func NewMain_contentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Main_contentContext {
	var p = new(Main_contentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_main_content

	return p
}

func (s *Main_contentContext) GetParser() antlr.Parser { return s.parser }

func (s *Main_contentContext) AllMain_prop() []IMain_propContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMain_propContext)(nil)).Elem())
	var tst = make([]IMain_propContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMain_propContext)
		}
	}

	return tst
}

func (s *Main_contentContext) Main_prop(i int) IMain_propContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMain_propContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMain_propContext)
}

func (s *Main_contentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Main_contentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Main_contentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterMain_content(s)
	}
}

func (s *Main_contentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitMain_content(s)
	}
}

func (s *Main_contentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitMain_content(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) Main_content() (localctx IMain_contentContext) {
	localctx = NewMain_contentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AspidaParserRULE_main_content)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Main_prop()
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AspidaParserT__3)|(1<<AspidaParserT__4)|(1<<AspidaParserT__5)|(1<<AspidaParserT__6)|(1<<AspidaParserT__7)|(1<<AspidaParserT__8))) != 0 {
		{
			p.SetState(86)
			p.Main_prop()
		}

		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMain_propContext is an interface to support dynamic dispatch.
type IMain_propContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMain_propContext differentiates from other interfaces.
	IsMain_propContext()
}

type Main_propContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMain_propContext() *Main_propContext {
	var p = new(Main_propContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_main_prop
	return p
}

func (*Main_propContext) IsMain_propContext() {}

func NewMain_propContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Main_propContext {
	var p = new(Main_propContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_main_prop

	return p
}

func (s *Main_propContext) GetParser() antlr.Parser { return s.parser }

func (s *Main_propContext) CopyFrom(ctx *Main_propContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Main_propContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Main_propContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NameMainContext struct {
	*Main_propContext
}

func NewNameMainContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NameMainContext {
	var p = new(NameMainContext)

	p.Main_propContext = NewEmptyMain_propContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Main_propContext))

	return p
}

func (s *NameMainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameMainContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *NameMainContext) NS() antlr.TerminalNode {
	return s.GetToken(AspidaParserNS, 0)
}

func (s *NameMainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterNameMain(s)
	}
}

func (s *NameMainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitNameMain(s)
	}
}

func (s *NameMainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitNameMain(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConnectionMainContext struct {
	*Main_propContext
}

func NewConnectionMainContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConnectionMainContext {
	var p = new(ConnectionMainContext)

	p.Main_propContext = NewEmptyMain_propContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Main_propContext))

	return p
}

func (s *ConnectionMainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConnectionMainContext) Connection() IConnectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConnectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConnectionContext)
}

func (s *ConnectionMainContext) NS() antlr.TerminalNode {
	return s.GetToken(AspidaParserNS, 0)
}

func (s *ConnectionMainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterConnectionMain(s)
	}
}

func (s *ConnectionMainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitConnectionMain(s)
	}
}

func (s *ConnectionMainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitConnectionMain(s)

	default:
		return t.VisitChildren(s)
	}
}

type DescriptionMainContext struct {
	*Main_propContext
}

func NewDescriptionMainContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DescriptionMainContext {
	var p = new(DescriptionMainContext)

	p.Main_propContext = NewEmptyMain_propContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Main_propContext))

	return p
}

func (s *DescriptionMainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionMainContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *DescriptionMainContext) NS() antlr.TerminalNode {
	return s.GetToken(AspidaParserNS, 0)
}

func (s *DescriptionMainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterDescriptionMain(s)
	}
}

func (s *DescriptionMainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitDescriptionMain(s)
	}
}

func (s *DescriptionMainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitDescriptionMain(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) Main_prop() (localctx IMain_propContext) {
	localctx = NewMain_propContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AspidaParserRULE_main_prop)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(101)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AspidaParserT__3, AspidaParserT__4:
		localctx = NewNameMainContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.Name()
		}
		{
			p.SetState(93)
			p.Match(AspidaParserNS)
		}

	case AspidaParserT__5, AspidaParserT__6:
		localctx = NewConnectionMainContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(95)
			p.Connection()
		}
		{
			p.SetState(96)
			p.Match(AspidaParserNS)
		}

	case AspidaParserT__7, AspidaParserT__8:
		localctx = NewDescriptionMainContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(98)
			p.Description()
		}
		{
			p.SetState(99)
			p.Match(AspidaParserNS)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitName(s)
	}
}

func (s *NameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AspidaParserRULE_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(109)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AspidaParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.Match(AspidaParserT__3)
		}
		{
			p.SetState(104)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(105)
			p.Value()
		}

	case AspidaParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(106)
			p.Match(AspidaParserT__4)
		}
		{
			p.SetState(107)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(108)
			p.Value()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConnectionContext is an interface to support dynamic dispatch.
type IConnectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConnectionContext differentiates from other interfaces.
	IsConnectionContext()
}

type ConnectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConnectionContext() *ConnectionContext {
	var p = new(ConnectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_connection
	return p
}

func (*ConnectionContext) IsConnectionContext() {}

func NewConnectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConnectionContext {
	var p = new(ConnectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_connection

	return p
}

func (s *ConnectionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConnectionContext) Connection_type() IConnection_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConnection_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConnection_typeContext)
}

func (s *ConnectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConnectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConnectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterConnection(s)
	}
}

func (s *ConnectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitConnection(s)
	}
}

func (s *ConnectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitConnection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) Connection() (localctx IConnectionContext) {
	localctx = NewConnectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AspidaParserRULE_connection)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(117)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AspidaParserT__5:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)
			p.Match(AspidaParserT__5)
		}
		{
			p.SetState(112)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(113)
			p.Connection_type()
		}

	case AspidaParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(114)
			p.Match(AspidaParserT__6)
		}
		{
			p.SetState(115)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(116)
			p.Connection_type()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConnection_typeContext is an interface to support dynamic dispatch.
type IConnection_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConnection_typeContext differentiates from other interfaces.
	IsConnection_typeContext()
}

type Connection_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConnection_typeContext() *Connection_typeContext {
	var p = new(Connection_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_connection_type
	return p
}

func (*Connection_typeContext) IsConnection_typeContext() {}

func NewConnection_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Connection_typeContext {
	var p = new(Connection_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_connection_type

	return p
}

func (s *Connection_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Connection_typeContext) CopyFrom(ctx *Connection_typeContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Connection_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Connection_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ConnectionSSHContext struct {
	*Connection_typeContext
}

func NewConnectionSSHContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConnectionSSHContext {
	var p = new(ConnectionSSHContext)

	p.Connection_typeContext = NewEmptyConnection_typeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Connection_typeContext))

	return p
}

func (s *ConnectionSSHContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConnectionSSHContext) SSH_KW() antlr.TerminalNode {
	return s.GetToken(AspidaParserSSH_KW, 0)
}

func (s *ConnectionSSHContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterConnectionSSH(s)
	}
}

func (s *ConnectionSSHContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitConnectionSSH(s)
	}
}

func (s *ConnectionSSHContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitConnectionSSH(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConnectionLocalContext struct {
	*Connection_typeContext
}

func NewConnectionLocalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConnectionLocalContext {
	var p = new(ConnectionLocalContext)

	p.Connection_typeContext = NewEmptyConnection_typeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Connection_typeContext))

	return p
}

func (s *ConnectionLocalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConnectionLocalContext) LOCAL_KW() antlr.TerminalNode {
	return s.GetToken(AspidaParserLOCAL_KW, 0)
}

func (s *ConnectionLocalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterConnectionLocal(s)
	}
}

func (s *ConnectionLocalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitConnectionLocal(s)
	}
}

func (s *ConnectionLocalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitConnectionLocal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) Connection_type() (localctx IConnection_typeContext) {
	localctx = NewConnection_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AspidaParserRULE_connection_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(121)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AspidaParserLOCAL_KW:
		localctx = NewConnectionLocalContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.Match(AspidaParserLOCAL_KW)
		}

	case AspidaParserSSH_KW:
		localctx = NewConnectionSSHContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.Match(AspidaParserSSH_KW)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDescriptionContext is an interface to support dynamic dispatch.
type IDescriptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDescriptionContext differentiates from other interfaces.
	IsDescriptionContext()
}

type DescriptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescriptionContext() *DescriptionContext {
	var p = new(DescriptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_description
	return p
}

func (*DescriptionContext) IsDescriptionContext() {}

func NewDescriptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescriptionContext {
	var p = new(DescriptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_description

	return p
}

func (s *DescriptionContext) GetParser() antlr.Parser { return s.parser }

func (s *DescriptionContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *DescriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DescriptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterDescription(s)
	}
}

func (s *DescriptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitDescription(s)
	}
}

func (s *DescriptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitDescription(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) Description() (localctx IDescriptionContext) {
	localctx = NewDescriptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AspidaParserRULE_description)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(129)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AspidaParserT__7:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.Match(AspidaParserT__7)
		}
		{
			p.SetState(124)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(125)
			p.Value()
		}

	case AspidaParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(126)
			p.Match(AspidaParserT__8)
		}
		{
			p.SetState(127)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(128)
			p.Value()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITasks_contentContext is an interface to support dynamic dispatch.
type ITasks_contentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTasks_contentContext differentiates from other interfaces.
	IsTasks_contentContext()
}

type Tasks_contentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTasks_contentContext() *Tasks_contentContext {
	var p = new(Tasks_contentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_tasks_content
	return p
}

func (*Tasks_contentContext) IsTasks_contentContext() {}

func NewTasks_contentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tasks_contentContext {
	var p = new(Tasks_contentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_tasks_content

	return p
}

func (s *Tasks_contentContext) GetParser() antlr.Parser { return s.parser }

func (s *Tasks_contentContext) CopyFrom(ctx *Tasks_contentContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Tasks_contentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tasks_contentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IfStatementContext struct {
	*Tasks_contentContext
}

func NewIfStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStatementContext {
	var p = new(IfStatementContext)

	p.Tasks_contentContext = NewEmptyTasks_contentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Tasks_contentContext))

	return p
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) IfStat() IIfStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatContext)
}

func (s *IfStatementContext) ElseStat() IElseStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElseStatContext)
}

func (s *IfStatementContext) AllElifStat() []IElifStatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElifStatContext)(nil)).Elem())
	var tst = make([]IElifStatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElifStatContext)
		}
	}

	return tst
}

func (s *IfStatementContext) ElifStat(i int) IElifStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElifStatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElifStatContext)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type TContentContext struct {
	*Tasks_contentContext
}

func NewTContentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TContentContext {
	var p = new(TContentContext)

	p.Tasks_contentContext = NewEmptyTasks_contentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Tasks_contentContext))

	return p
}

func (s *TContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TContentContext) AllTasks_prop() []ITasks_propContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITasks_propContext)(nil)).Elem())
	var tst = make([]ITasks_propContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITasks_propContext)
		}
	}

	return tst
}

func (s *TContentContext) Tasks_prop(i int) ITasks_propContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITasks_propContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITasks_propContext)
}

func (s *TContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterTContent(s)
	}
}

func (s *TContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitTContent(s)
	}
}

func (s *TContentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitTContent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) Tasks_content() (localctx ITasks_contentContext) {
	localctx = NewTasks_contentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, AspidaParserRULE_tasks_content)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(147)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AspidaParserT__9, AspidaParserT__10, AspidaParserT__11, AspidaParserT__12, AspidaParserT__13, AspidaParserT__14, AspidaParserT__15, AspidaParserT__16, AspidaParserT__17, AspidaParserT__18:
		localctx = NewTContentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)
			p.Tasks_prop()
		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AspidaParserT__9)|(1<<AspidaParserT__10)|(1<<AspidaParserT__11)|(1<<AspidaParserT__12)|(1<<AspidaParserT__13)|(1<<AspidaParserT__14)|(1<<AspidaParserT__15)|(1<<AspidaParserT__16)|(1<<AspidaParserT__17)|(1<<AspidaParserT__18))) != 0 {
			{
				p.SetState(132)
				p.Tasks_prop()
			}

			p.SetState(137)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case AspidaParserIF:
		localctx = NewIfStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(138)
			p.IfStat()
		}
		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AspidaParserELIF {
			{
				p.SetState(139)
				p.ElifStat()
			}

			p.SetState(144)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(145)
			p.ElseStat()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITasks_propContext is an interface to support dynamic dispatch.
type ITasks_propContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTasks_propContext differentiates from other interfaces.
	IsTasks_propContext()
}

type Tasks_propContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTasks_propContext() *Tasks_propContext {
	var p = new(Tasks_propContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_tasks_prop
	return p
}

func (*Tasks_propContext) IsTasks_propContext() {}

func NewTasks_propContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tasks_propContext {
	var p = new(Tasks_propContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_tasks_prop

	return p
}

func (s *Tasks_propContext) GetParser() antlr.Parser { return s.parser }

func (s *Tasks_propContext) CopyFrom(ctx *Tasks_propContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Tasks_propContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tasks_propContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TPointsContext struct {
	*Tasks_propContext
}

func NewTPointsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TPointsContext {
	var p = new(TPointsContext)

	p.Tasks_propContext = NewEmptyTasks_propContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Tasks_propContext))

	return p
}

func (s *TPointsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TPointsContext) Points() IPointsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPointsContext)
}

func (s *TPointsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterTPoints(s)
	}
}

func (s *TPointsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitTPoints(s)
	}
}

func (s *TPointsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitTPoints(s)

	default:
		return t.VisitChildren(s)
	}
}

type TTagsContext struct {
	*Tasks_propContext
}

func NewTTagsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TTagsContext {
	var p = new(TTagsContext)

	p.Tasks_propContext = NewEmptyTasks_propContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Tasks_propContext))

	return p
}

func (s *TTagsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TTagsContext) Tags() ITagsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagsContext)
}

func (s *TTagsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterTTags(s)
	}
}

func (s *TTagsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitTTags(s)
	}
}

func (s *TTagsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitTTags(s)

	default:
		return t.VisitChildren(s)
	}
}

type TControlsContext struct {
	*Tasks_propContext
}

func NewTControlsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TControlsContext {
	var p = new(TControlsContext)

	p.Tasks_propContext = NewEmptyTasks_propContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Tasks_propContext))

	return p
}

func (s *TControlsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TControlsContext) Controls() IControlsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IControlsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IControlsContext)
}

func (s *TControlsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterTControls(s)
	}
}

func (s *TControlsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitTControls(s)
	}
}

func (s *TControlsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitTControls(s)

	default:
		return t.VisitChildren(s)
	}
}

type TExclusionsContext struct {
	*Tasks_propContext
}

func NewTExclusionsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TExclusionsContext {
	var p = new(TExclusionsContext)

	p.Tasks_propContext = NewEmptyTasks_propContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Tasks_propContext))

	return p
}

func (s *TExclusionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TExclusionsContext) Exclusions() IExclusionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExclusionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExclusionsContext)
}

func (s *TExclusionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterTExclusions(s)
	}
}

func (s *TExclusionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitTExclusions(s)
	}
}

func (s *TExclusionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitTExclusions(s)

	default:
		return t.VisitChildren(s)
	}
}

type TSectionsContext struct {
	*Tasks_propContext
}

func NewTSectionsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TSectionsContext {
	var p = new(TSectionsContext)

	p.Tasks_propContext = NewEmptyTasks_propContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Tasks_propContext))

	return p
}

func (s *TSectionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TSectionsContext) Sections() ISectionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISectionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISectionsContext)
}

func (s *TSectionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterTSections(s)
	}
}

func (s *TSectionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitTSections(s)
	}
}

func (s *TSectionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitTSections(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) Tasks_prop() (localctx ITasks_propContext) {
	localctx = NewTasks_propContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, AspidaParserRULE_tasks_prop)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(154)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AspidaParserT__9, AspidaParserT__10:
		localctx = NewTSectionsContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(149)
			p.Sections()
		}

	case AspidaParserT__11, AspidaParserT__12:
		localctx = NewTPointsContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(150)
			p.Points()
		}

	case AspidaParserT__13, AspidaParserT__14:
		localctx = NewTControlsContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(151)
			p.Controls()
		}

	case AspidaParserT__15, AspidaParserT__16:
		localctx = NewTExclusionsContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(152)
			p.Exclusions()
		}

	case AspidaParserT__17, AspidaParserT__18:
		localctx = NewTTagsContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(153)
			p.Tags()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISectionsContext is an interface to support dynamic dispatch.
type ISectionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSectionsContext differentiates from other interfaces.
	IsSectionsContext()
}

type SectionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySectionsContext() *SectionsContext {
	var p = new(SectionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_sections
	return p
}

func (*SectionsContext) IsSectionsContext() {}

func NewSectionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SectionsContext {
	var p = new(SectionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_sections

	return p
}

func (s *SectionsContext) GetParser() antlr.Parser { return s.parser }

func (s *SectionsContext) Str_array() IStr_arrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStr_arrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStr_arrayContext)
}

func (s *SectionsContext) NS() antlr.TerminalNode {
	return s.GetToken(AspidaParserNS, 0)
}

func (s *SectionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SectionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterSections(s)
	}
}

func (s *SectionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitSections(s)
	}
}

func (s *SectionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitSections(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) Sections() (localctx ISectionsContext) {
	localctx = NewSectionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, AspidaParserRULE_sections)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(166)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AspidaParserT__9:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(156)
			p.Match(AspidaParserT__9)
		}
		{
			p.SetState(157)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(158)
			p.Str_array()
		}
		{
			p.SetState(159)
			p.Match(AspidaParserNS)
		}

	case AspidaParserT__10:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(161)
			p.Match(AspidaParserT__10)
		}
		{
			p.SetState(162)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(163)
			p.Str_array()
		}
		{
			p.SetState(164)
			p.Match(AspidaParserNS)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPointsContext is an interface to support dynamic dispatch.
type IPointsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPointsContext differentiates from other interfaces.
	IsPointsContext()
}

type PointsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPointsContext() *PointsContext {
	var p = new(PointsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_points
	return p
}

func (*PointsContext) IsPointsContext() {}

func NewPointsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointsContext {
	var p = new(PointsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_points

	return p
}

func (s *PointsContext) GetParser() antlr.Parser { return s.parser }

func (s *PointsContext) Str_array() IStr_arrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStr_arrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStr_arrayContext)
}

func (s *PointsContext) NS() antlr.TerminalNode {
	return s.GetToken(AspidaParserNS, 0)
}

func (s *PointsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterPoints(s)
	}
}

func (s *PointsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitPoints(s)
	}
}

func (s *PointsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitPoints(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) Points() (localctx IPointsContext) {
	localctx = NewPointsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, AspidaParserRULE_points)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(178)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AspidaParserT__11:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(168)
			p.Match(AspidaParserT__11)
		}
		{
			p.SetState(169)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(170)
			p.Str_array()
		}
		{
			p.SetState(171)
			p.Match(AspidaParserNS)
		}

	case AspidaParserT__12:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(173)
			p.Match(AspidaParserT__12)
		}
		{
			p.SetState(174)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(175)
			p.Str_array()
		}
		{
			p.SetState(176)
			p.Match(AspidaParserNS)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IControlsContext is an interface to support dynamic dispatch.
type IControlsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsControlsContext differentiates from other interfaces.
	IsControlsContext()
}

type ControlsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyControlsContext() *ControlsContext {
	var p = new(ControlsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_controls
	return p
}

func (*ControlsContext) IsControlsContext() {}

func NewControlsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ControlsContext {
	var p = new(ControlsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_controls

	return p
}

func (s *ControlsContext) GetParser() antlr.Parser { return s.parser }

func (s *ControlsContext) Str_array() IStr_arrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStr_arrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStr_arrayContext)
}

func (s *ControlsContext) NS() antlr.TerminalNode {
	return s.GetToken(AspidaParserNS, 0)
}

func (s *ControlsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ControlsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ControlsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterControls(s)
	}
}

func (s *ControlsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitControls(s)
	}
}

func (s *ControlsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitControls(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) Controls() (localctx IControlsContext) {
	localctx = NewControlsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, AspidaParserRULE_controls)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(190)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AspidaParserT__13:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.Match(AspidaParserT__13)
		}
		{
			p.SetState(181)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(182)
			p.Str_array()
		}
		{
			p.SetState(183)
			p.Match(AspidaParserNS)
		}

	case AspidaParserT__14:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(185)
			p.Match(AspidaParserT__14)
		}
		{
			p.SetState(186)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(187)
			p.Str_array()
		}
		{
			p.SetState(188)
			p.Match(AspidaParserNS)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExclusionsContext is an interface to support dynamic dispatch.
type IExclusionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExclusionsContext differentiates from other interfaces.
	IsExclusionsContext()
}

type ExclusionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExclusionsContext() *ExclusionsContext {
	var p = new(ExclusionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_exclusions
	return p
}

func (*ExclusionsContext) IsExclusionsContext() {}

func NewExclusionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExclusionsContext {
	var p = new(ExclusionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_exclusions

	return p
}

func (s *ExclusionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExclusionsContext) Str_array() IStr_arrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStr_arrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStr_arrayContext)
}

func (s *ExclusionsContext) NS() antlr.TerminalNode {
	return s.GetToken(AspidaParserNS, 0)
}

func (s *ExclusionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExclusionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExclusionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterExclusions(s)
	}
}

func (s *ExclusionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitExclusions(s)
	}
}

func (s *ExclusionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitExclusions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) Exclusions() (localctx IExclusionsContext) {
	localctx = NewExclusionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, AspidaParserRULE_exclusions)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(202)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AspidaParserT__15:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(192)
			p.Match(AspidaParserT__15)
		}
		{
			p.SetState(193)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(194)
			p.Str_array()
		}
		{
			p.SetState(195)
			p.Match(AspidaParserNS)
		}

	case AspidaParserT__16:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(197)
			p.Match(AspidaParserT__16)
		}
		{
			p.SetState(198)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(199)
			p.Str_array()
		}
		{
			p.SetState(200)
			p.Match(AspidaParserNS)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITagsContext is an interface to support dynamic dispatch.
type ITagsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTagsContext differentiates from other interfaces.
	IsTagsContext()
}

type TagsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagsContext() *TagsContext {
	var p = new(TagsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_tags
	return p
}

func (*TagsContext) IsTagsContext() {}

func NewTagsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagsContext {
	var p = new(TagsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_tags

	return p
}

func (s *TagsContext) GetParser() antlr.Parser { return s.parser }

func (s *TagsContext) Str_array() IStr_arrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStr_arrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStr_arrayContext)
}

func (s *TagsContext) NS() antlr.TerminalNode {
	return s.GetToken(AspidaParserNS, 0)
}

func (s *TagsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterTags(s)
	}
}

func (s *TagsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitTags(s)
	}
}

func (s *TagsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitTags(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) Tags() (localctx ITagsContext) {
	localctx = NewTagsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, AspidaParserRULE_tags)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(214)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AspidaParserT__17:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(204)
			p.Match(AspidaParserT__17)
		}
		{
			p.SetState(205)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(206)
			p.Str_array()
		}
		{
			p.SetState(207)
			p.Match(AspidaParserNS)
		}

	case AspidaParserT__18:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(209)
			p.Match(AspidaParserT__18)
		}
		{
			p.SetState(210)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(211)
			p.Str_array()
		}
		{
			p.SetState(212)
			p.Match(AspidaParserNS)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIfStatContext is an interface to support dynamic dispatch.
type IIfStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStatContext differentiates from other interfaces.
	IsIfStatContext()
}

type IfStatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatContext() *IfStatContext {
	var p = new(IfStatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_ifStat
	return p
}

func (*IfStatContext) IsIfStatContext() {}

func NewIfStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatContext {
	var p = new(IfStatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_ifStat

	return p
}

func (s *IfStatContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatContext) IF() antlr.TerminalNode {
	return s.GetToken(AspidaParserIF, 0)
}

func (s *IfStatContext) Comparison() IComparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonContext)
}

func (s *IfStatContext) Tasks_content() ITasks_contentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITasks_contentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITasks_contentContext)
}

func (s *IfStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterIfStat(s)
	}
}

func (s *IfStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitIfStat(s)
	}
}

func (s *IfStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitIfStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) IfStat() (localctx IIfStatContext) {
	localctx = NewIfStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, AspidaParserRULE_ifStat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.Match(AspidaParserIF)
	}
	{
		p.SetState(217)
		p.Comparison()
	}
	{
		p.SetState(218)
		p.Match(AspidaParserT__1)
	}
	{
		p.SetState(219)
		p.Tasks_content()
	}
	{
		p.SetState(220)
		p.Match(AspidaParserT__2)
	}

	return localctx
}

// IElifStatContext is an interface to support dynamic dispatch.
type IElifStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElifStatContext differentiates from other interfaces.
	IsElifStatContext()
}

type ElifStatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElifStatContext() *ElifStatContext {
	var p = new(ElifStatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_elifStat
	return p
}

func (*ElifStatContext) IsElifStatContext() {}

func NewElifStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElifStatContext {
	var p = new(ElifStatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_elifStat

	return p
}

func (s *ElifStatContext) GetParser() antlr.Parser { return s.parser }

func (s *ElifStatContext) ELIF() antlr.TerminalNode {
	return s.GetToken(AspidaParserELIF, 0)
}

func (s *ElifStatContext) Comparison() IComparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonContext)
}

func (s *ElifStatContext) Tasks_content() ITasks_contentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITasks_contentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITasks_contentContext)
}

func (s *ElifStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElifStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElifStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterElifStat(s)
	}
}

func (s *ElifStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitElifStat(s)
	}
}

func (s *ElifStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitElifStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) ElifStat() (localctx IElifStatContext) {
	localctx = NewElifStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, AspidaParserRULE_elifStat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		p.Match(AspidaParserELIF)
	}
	{
		p.SetState(223)
		p.Comparison()
	}
	{
		p.SetState(224)
		p.Match(AspidaParserT__1)
	}
	{
		p.SetState(225)
		p.Tasks_content()
	}
	{
		p.SetState(226)
		p.Match(AspidaParserT__2)
	}

	return localctx
}

// IElseStatContext is an interface to support dynamic dispatch.
type IElseStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseStatContext differentiates from other interfaces.
	IsElseStatContext()
}

type ElseStatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseStatContext() *ElseStatContext {
	var p = new(ElseStatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_elseStat
	return p
}

func (*ElseStatContext) IsElseStatContext() {}

func NewElseStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseStatContext {
	var p = new(ElseStatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_elseStat

	return p
}

func (s *ElseStatContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseStatContext) ELSE() antlr.TerminalNode {
	return s.GetToken(AspidaParserELSE, 0)
}

func (s *ElseStatContext) Tasks_content() ITasks_contentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITasks_contentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITasks_contentContext)
}

func (s *ElseStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterElseStat(s)
	}
}

func (s *ElseStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitElseStat(s)
	}
}

func (s *ElseStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitElseStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) ElseStat() (localctx IElseStatContext) {
	localctx = NewElseStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, AspidaParserRULE_elseStat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.Match(AspidaParserELSE)
	}
	{
		p.SetState(229)
		p.Match(AspidaParserT__1)
	}
	{
		p.SetState(230)
		p.Tasks_content()
	}
	{
		p.SetState(231)
		p.Match(AspidaParserT__2)
	}

	return localctx
}

// IComparisonContext is an interface to support dynamic dispatch.
type IComparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonContext differentiates from other interfaces.
	IsComparisonContext()
}

type ComparisonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonContext() *ComparisonContext {
	var p = new(ComparisonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_comparison
	return p
}

func (*ComparisonContext) IsComparisonContext() {}

func NewComparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonContext {
	var p = new(ComparisonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_comparison

	return p
}

func (s *ComparisonContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *ComparisonContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ComparisonContext) Comp_op() IComp_opContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComp_opContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComp_opContext)
}

func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterComparison(s)
	}
}

func (s *ComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitComparison(s)
	}
}

func (s *ComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) Comparison() (localctx IComparisonContext) {
	localctx = NewComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, AspidaParserRULE_comparison)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(233)
		p.Value()
	}
	{
		p.SetState(234)
		p.Comp_op()
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-26)&-(0x1f+1)) == 0 && ((1<<uint((_la-26)))&((1<<(AspidaParserT__25-26))|(1<<(AspidaParserT__26-26))|(1<<(AspidaParserT__27-26))|(1<<(AspidaParserT__28-26))|(1<<(AspidaParserSTRING-26))|(1<<(AspidaParserNUMBER-26)))) != 0 {
		{
			p.SetState(235)
			p.Value()
		}

		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVars_contentContext is an interface to support dynamic dispatch.
type IVars_contentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVars_contentContext differentiates from other interfaces.
	IsVars_contentContext()
}

type Vars_contentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVars_contentContext() *Vars_contentContext {
	var p = new(Vars_contentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_vars_content
	return p
}

func (*Vars_contentContext) IsVars_contentContext() {}

func NewVars_contentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vars_contentContext {
	var p = new(Vars_contentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_vars_content

	return p
}

func (s *Vars_contentContext) GetParser() antlr.Parser { return s.parser }

func (s *Vars_contentContext) AllVars_prop() []IVars_propContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVars_propContext)(nil)).Elem())
	var tst = make([]IVars_propContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVars_propContext)
		}
	}

	return tst
}

func (s *Vars_contentContext) Vars_prop(i int) IVars_propContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVars_propContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVars_propContext)
}

func (s *Vars_contentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vars_contentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Vars_contentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterVars_content(s)
	}
}

func (s *Vars_contentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitVars_content(s)
	}
}

func (s *Vars_contentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitVars_content(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) Vars_content() (localctx IVars_contentContext) {
	localctx = NewVars_contentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, AspidaParserRULE_vars_content)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.Vars_prop()
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AspidaParserSTRING {
		{
			p.SetState(242)
			p.Vars_prop()
		}

		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVars_propContext is an interface to support dynamic dispatch.
type IVars_propContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVars_propContext differentiates from other interfaces.
	IsVars_propContext()
}

type Vars_propContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVars_propContext() *Vars_propContext {
	var p = new(Vars_propContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_vars_prop
	return p
}

func (*Vars_propContext) IsVars_propContext() {}

func NewVars_propContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vars_propContext {
	var p = new(Vars_propContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_vars_prop

	return p
}

func (s *Vars_propContext) GetParser() antlr.Parser { return s.parser }

func (s *Vars_propContext) CopyFrom(ctx *Vars_propContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Vars_propContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vars_propContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VarDefContext struct {
	*Vars_propContext
}

func NewVarDefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarDefContext {
	var p = new(VarDefContext)

	p.Vars_propContext = NewEmptyVars_propContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Vars_propContext))

	return p
}

func (s *VarDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDefContext) STRING() antlr.TerminalNode {
	return s.GetToken(AspidaParserSTRING, 0)
}

func (s *VarDefContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *VarDefContext) NS() antlr.TerminalNode {
	return s.GetToken(AspidaParserNS, 0)
}

func (s *VarDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterVarDef(s)
	}
}

func (s *VarDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitVarDef(s)
	}
}

func (s *VarDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitVarDef(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarObjDefContext struct {
	*Vars_propContext
}

func NewVarObjDefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarObjDefContext {
	var p = new(VarObjDefContext)

	p.Vars_propContext = NewEmptyVars_propContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Vars_propContext))

	return p
}

func (s *VarObjDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarObjDefContext) STRING() antlr.TerminalNode {
	return s.GetToken(AspidaParserSTRING, 0)
}

func (s *VarObjDefContext) Vars_content() IVars_contentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVars_contentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVars_contentContext)
}

func (s *VarObjDefContext) NS() antlr.TerminalNode {
	return s.GetToken(AspidaParserNS, 0)
}

func (s *VarObjDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterVarObjDef(s)
	}
}

func (s *VarObjDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitVarObjDef(s)
	}
}

func (s *VarObjDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitVarObjDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) Vars_prop() (localctx IVars_propContext) {
	localctx = NewVars_propContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, AspidaParserRULE_vars_prop)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVarDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(248)
			p.Match(AspidaParserSTRING)
		}
		{
			p.SetState(249)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(250)
			p.Value()
		}
		{
			p.SetState(251)
			p.Match(AspidaParserNS)
		}

	case 2:
		localctx = NewVarObjDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(253)
			p.Match(AspidaParserSTRING)
		}
		{
			p.SetState(254)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(255)
			p.Match(AspidaParserT__1)
		}
		{
			p.SetState(256)
			p.Vars_content()
		}
		{
			p.SetState(257)
			p.Match(AspidaParserT__2)
		}
		{
			p.SetState(258)
			p.Match(AspidaParserNS)
		}

	}

	return localctx
}

// IComp_opContext is an interface to support dynamic dispatch.
type IComp_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComp_opContext differentiates from other interfaces.
	IsComp_opContext()
}

type Comp_opContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComp_opContext() *Comp_opContext {
	var p = new(Comp_opContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_comp_op
	return p
}

func (*Comp_opContext) IsComp_opContext() {}

func NewComp_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Comp_opContext {
	var p = new(Comp_opContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_comp_op

	return p
}

func (s *Comp_opContext) GetParser() antlr.Parser { return s.parser }
func (s *Comp_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comp_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Comp_opContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterComp_op(s)
	}
}

func (s *Comp_opContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitComp_op(s)
	}
}

func (s *Comp_opContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitComp_op(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) Comp_op() (localctx IComp_opContext) {
	localctx = NewComp_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, AspidaParserRULE_comp_op)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(262)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AspidaParserT__19)|(1<<AspidaParserT__20)|(1<<AspidaParserT__21)|(1<<AspidaParserT__22)|(1<<AspidaParserT__23)|(1<<AspidaParserT__24))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) CopyFrom(ctx *ValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TrueValContext struct {
	*ValueContext
}

func NewTrueValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueValContext {
	var p = new(TrueValContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *TrueValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterTrueVal(s)
	}
}

func (s *TrueValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitTrueVal(s)
	}
}

func (s *TrueValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitTrueVal(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberValContext struct {
	*ValueContext
}

func NewNumberValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberValContext {
	var p = new(NumberValContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *NumberValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberValContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(AspidaParserNUMBER, 0)
}

func (s *NumberValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterNumberVal(s)
	}
}

func (s *NumberValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitNumberVal(s)
	}
}

func (s *NumberValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitNumberVal(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayValContext struct {
	*ValueContext
}

func NewArrayValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayValContext {
	var p = new(ArrayValContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *ArrayValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayValContext) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *ArrayValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterArrayVal(s)
	}
}

func (s *ArrayValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitArrayVal(s)
	}
}

func (s *ArrayValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitArrayVal(s)

	default:
		return t.VisitChildren(s)
	}
}

type FalseValContext struct {
	*ValueContext
}

func NewFalseValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FalseValContext {
	var p = new(FalseValContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *FalseValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FalseValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterFalseVal(s)
	}
}

func (s *FalseValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitFalseVal(s)
	}
}

func (s *FalseValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitFalseVal(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringValContext struct {
	*ValueContext
}

func NewStringValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringValContext {
	var p = new(StringValContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *StringValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValContext) STRING() antlr.TerminalNode {
	return s.GetToken(AspidaParserSTRING, 0)
}

func (s *StringValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterStringVal(s)
	}
}

func (s *StringValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitStringVal(s)
	}
}

func (s *StringValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitStringVal(s)

	default:
		return t.VisitChildren(s)
	}
}

type NullValContext struct {
	*ValueContext
}

func NewNullValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullValContext {
	var p = new(NullValContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *NullValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterNullVal(s)
	}
}

func (s *NullValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitNullVal(s)
	}
}

func (s *NullValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitNullVal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, AspidaParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(270)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AspidaParserSTRING:
		localctx = NewStringValContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(264)
			p.Match(AspidaParserSTRING)
		}

	case AspidaParserNUMBER:
		localctx = NewNumberValContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(265)
			p.Match(AspidaParserNUMBER)
		}

	case AspidaParserT__25:
		localctx = NewTrueValContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(266)
			p.Match(AspidaParserT__25)
		}

	case AspidaParserT__26:
		localctx = NewFalseValContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(267)
			p.Match(AspidaParserT__26)
		}

	case AspidaParserT__27:
		localctx = NewNullValContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(268)
			p.Match(AspidaParserT__27)
		}

	case AspidaParserT__28:
		localctx = NewArrayValContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(269)
			p.Array()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStr_arrayContext is an interface to support dynamic dispatch.
type IStr_arrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStr_arrayContext differentiates from other interfaces.
	IsStr_arrayContext()
}

type Str_arrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStr_arrayContext() *Str_arrayContext {
	var p = new(Str_arrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_str_array
	return p
}

func (*Str_arrayContext) IsStr_arrayContext() {}

func NewStr_arrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Str_arrayContext {
	var p = new(Str_arrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_str_array

	return p
}

func (s *Str_arrayContext) GetParser() antlr.Parser { return s.parser }

func (s *Str_arrayContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(AspidaParserSTRING)
}

func (s *Str_arrayContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(AspidaParserSTRING, i)
}

func (s *Str_arrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Str_arrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Str_arrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterStr_array(s)
	}
}

func (s *Str_arrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitStr_array(s)
	}
}

func (s *Str_arrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitStr_array(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) Str_array() (localctx IStr_arrayContext) {
	localctx = NewStr_arrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, AspidaParserRULE_str_array)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(272)
			p.Match(AspidaParserT__28)
		}
		{
			p.SetState(273)
			p.Match(AspidaParserSTRING)
		}
		p.SetState(278)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AspidaParserT__29 {
			{
				p.SetState(274)
				p.Match(AspidaParserT__29)
			}
			{
				p.SetState(275)
				p.Match(AspidaParserSTRING)
			}

			p.SetState(280)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(281)
			p.Match(AspidaParserT__30)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(282)
			p.Match(AspidaParserT__28)
		}
		{
			p.SetState(283)
			p.Match(AspidaParserT__30)
		}

	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *ArrayContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AspidaListener); ok {
		listenerT.ExitArray(s)
	}
}

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AspidaParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, AspidaParserRULE_array)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(286)
			p.Match(AspidaParserT__28)
		}
		{
			p.SetState(287)
			p.Value()
		}
		p.SetState(292)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AspidaParserT__29 {
			{
				p.SetState(288)
				p.Match(AspidaParserT__29)
			}
			{
				p.SetState(289)
				p.Value()
			}

			p.SetState(294)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(295)
			p.Match(AspidaParserT__30)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(297)
			p.Match(AspidaParserT__28)
		}
		{
			p.SetState(298)
			p.Match(AspidaParserT__30)
		}

	}

	return localctx
}
