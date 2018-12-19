##protoc --csharp_out=./output/csharp ./src/login.proto

##protoc --go_out=./output/go ./src/login.proto

##protoc --lua_out=./output/lua --plugin=protoc-gen-lua="E:\\gitHub\\PixelFarm\\PixelDoc\\Tools\\protoc-gen-lua-master\\plugin/protoc-gen-lua.bat" ./src/login.proto

set curdir=E:\\gitHub\\PixelFarm\\PixelDoc\\Tools\\protoc-3.6.0-win32
cd /d %curdir%

for /r .\src %%i in (*.proto) do (
	echo "%%~nxi"
	protoc --proto_path=./src/ --csharp_out=./output/csharp %%~nxi
	protoc --proto_path=./src/ --go_out=./output/go %%~nxi
	protoc --proto_path=./src/ --lua_out=./output/lua --plugin=protoc-gen-lua="E:\\gitHub\\PixelFarm\\PixelDoc\\Tools\\protoc-gen-lua-master\\plugin/protoc-gen-lua.bat" %%~nxi
	protoc --proto_path=./src/ --descriptor_set_out=./output/pb/%%~ni.pb %%~nxi
)

pause