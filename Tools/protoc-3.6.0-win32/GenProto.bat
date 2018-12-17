##protoc --csharp_out=./output/csharp ./src/login.proto

##protoc --go_out=./output/go ./src/login.proto

##protoc --lua_out=./output/lua --plugin=protoc-gen-lua="E:\\gitHub\\PixelFarm\\PixelDoc\\Tools\\protoc-gen-lua-master\\plugin/protoc-gen-lua.bat" ./src/login.proto

set curdir=E:\\gitHub\\PixelFarm\\PixelDoc\\Tools\\protoc-3.6.0-win32
cd /d %curdir%

for /r .\src %%i in (*.proto) do (
	echo "%%~nxi"
	protoc --csharp_out=./output/csharp ./src/%%~nxi
	protoc --go_out=./output/go ./src/%%~nxi
	protoc --lua_out=./output/lua --plugin=protoc-gen-lua="E:\\gitHub\\PixelFarm\\PixelDoc\\Tools\\protoc-gen-lua-master\\plugin/protoc-gen-lua.bat" ./src/%%~nxi
)

pause