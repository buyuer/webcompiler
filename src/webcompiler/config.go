package webcompiler

const (
	gccPath       = `F:\Development\SDK\LLVM\clang-9.0.0\bin\clang.exe`
	gppPath       = `F:\\Development\\SDK\\GCC\\mingw-w64\\x86_64-8.1.0-posix-seh-rt_v6-rev0\\mingw64\\bin\\g++.exe`
	javaPath      = ""
	goPath        = ""
	pythonPath    = ""
	cacheFilePath = "D:\\temp.c"
)

var Compiler = map[string]string{
	"gcc":    gccPath,
	"g++":    gppPath,
	"java":   javaPath,
	"go":     goPath,
	"python": pythonPath,
}
