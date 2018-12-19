-- Generated By protoc-gen-lua Do not Edit
local protobuf = require "protobuf/protobuf"
local err_pb = require("PixelFarm/PBLua/err_pb")
module('login_pb')


local RESPONSECODE = protobuf.EnumDescriptor();
local RESPONSECODE_FAIL_ENUM = protobuf.EnumValueDescriptor();
local RESPONSECODE_SUCCESS_ENUM = protobuf.EnumValueDescriptor();
local LOGINREQUEST = protobuf.Descriptor();
local LOGINREQUEST_ACCOUNT_FIELD = protobuf.FieldDescriptor();
local LOGINREQUEST_PASSWORD_FIELD = protobuf.FieldDescriptor();
local LOGINRESPONSE = protobuf.Descriptor();
local LOGINRESPONSE_CODE_FIELD = protobuf.FieldDescriptor();
local LOGINRESPONSE_UID_FIELD = protobuf.FieldDescriptor();
local LOGINRESPONSE_ERR_FIELD = protobuf.FieldDescriptor();
local REGISTEREQUEST = protobuf.Descriptor();
local REGISTEREQUEST_ACCOUNT_FIELD = protobuf.FieldDescriptor();
local REGISTEREQUEST_PASSWORD_FIELD = protobuf.FieldDescriptor();
local REGISTERESPONSE = protobuf.Descriptor();
local REGISTERESPONSE_CODE_FIELD = protobuf.FieldDescriptor();
local REGISTERESPONSE_UID_FIELD = protobuf.FieldDescriptor();
local REGISTERESPONSE_ERR_FIELD = protobuf.FieldDescriptor();

RESPONSECODE_FAIL_ENUM.name = "FAIL"
RESPONSECODE_FAIL_ENUM.index = 0
RESPONSECODE_FAIL_ENUM.number = 0
RESPONSECODE_SUCCESS_ENUM.name = "SUCCESS"
RESPONSECODE_SUCCESS_ENUM.index = 1
RESPONSECODE_SUCCESS_ENUM.number = 1
RESPONSECODE.name = "ResponseCode"
RESPONSECODE.full_name = ".msg.ResponseCode"
RESPONSECODE.values = {RESPONSECODE_FAIL_ENUM,RESPONSECODE_SUCCESS_ENUM}
LOGINREQUEST_ACCOUNT_FIELD.name = "account"
LOGINREQUEST_ACCOUNT_FIELD.full_name = ".msg.LoginRequest.account"
LOGINREQUEST_ACCOUNT_FIELD.number = 1
LOGINREQUEST_ACCOUNT_FIELD.index = 0
LOGINREQUEST_ACCOUNT_FIELD.label = 1
LOGINREQUEST_ACCOUNT_FIELD.has_default_value = false
LOGINREQUEST_ACCOUNT_FIELD.default_value = ""
LOGINREQUEST_ACCOUNT_FIELD.type = 9
LOGINREQUEST_ACCOUNT_FIELD.cpp_type = 9

LOGINREQUEST_PASSWORD_FIELD.name = "password"
LOGINREQUEST_PASSWORD_FIELD.full_name = ".msg.LoginRequest.password"
LOGINREQUEST_PASSWORD_FIELD.number = 2
LOGINREQUEST_PASSWORD_FIELD.index = 1
LOGINREQUEST_PASSWORD_FIELD.label = 1
LOGINREQUEST_PASSWORD_FIELD.has_default_value = false
LOGINREQUEST_PASSWORD_FIELD.default_value = ""
LOGINREQUEST_PASSWORD_FIELD.type = 9
LOGINREQUEST_PASSWORD_FIELD.cpp_type = 9

LOGINREQUEST.name = "LoginRequest"
LOGINREQUEST.full_name = ".msg.LoginRequest"
LOGINREQUEST.nested_types = {}
LOGINREQUEST.enum_types = {}
LOGINREQUEST.fields = {LOGINREQUEST_ACCOUNT_FIELD, LOGINREQUEST_PASSWORD_FIELD}
LOGINREQUEST.is_extendable = false
LOGINREQUEST.extensions = {}
LOGINRESPONSE_CODE_FIELD.name = "code"
LOGINRESPONSE_CODE_FIELD.full_name = ".msg.LoginResponse.code"
LOGINRESPONSE_CODE_FIELD.number = 1
LOGINRESPONSE_CODE_FIELD.index = 0
LOGINRESPONSE_CODE_FIELD.label = 1
LOGINRESPONSE_CODE_FIELD.has_default_value = false
LOGINRESPONSE_CODE_FIELD.default_value = nil
LOGINRESPONSE_CODE_FIELD.enum_type = RESPONSECODE
LOGINRESPONSE_CODE_FIELD.type = 14
LOGINRESPONSE_CODE_FIELD.cpp_type = 8

LOGINRESPONSE_UID_FIELD.name = "uid"
LOGINRESPONSE_UID_FIELD.full_name = ".msg.LoginResponse.uid"
LOGINRESPONSE_UID_FIELD.number = 2
LOGINRESPONSE_UID_FIELD.index = 1
LOGINRESPONSE_UID_FIELD.label = 1
LOGINRESPONSE_UID_FIELD.has_default_value = false
LOGINRESPONSE_UID_FIELD.default_value = ""
LOGINRESPONSE_UID_FIELD.type = 9
LOGINRESPONSE_UID_FIELD.cpp_type = 9

LOGINRESPONSE_ERR_FIELD.name = "err"
LOGINRESPONSE_ERR_FIELD.full_name = ".msg.LoginResponse.err"
LOGINRESPONSE_ERR_FIELD.number = 3
LOGINRESPONSE_ERR_FIELD.index = 2
LOGINRESPONSE_ERR_FIELD.label = 1
LOGINRESPONSE_ERR_FIELD.has_default_value = false
LOGINRESPONSE_ERR_FIELD.default_value = nil
LOGINRESPONSE_ERR_FIELD.message_type = ERR_PB_ERROR
LOGINRESPONSE_ERR_FIELD.type = 11
LOGINRESPONSE_ERR_FIELD.cpp_type = 10

LOGINRESPONSE.name = "LoginResponse"
LOGINRESPONSE.full_name = ".msg.LoginResponse"
LOGINRESPONSE.nested_types = {}
LOGINRESPONSE.enum_types = {}
LOGINRESPONSE.fields = {LOGINRESPONSE_CODE_FIELD, LOGINRESPONSE_UID_FIELD, LOGINRESPONSE_ERR_FIELD}
LOGINRESPONSE.is_extendable = false
LOGINRESPONSE.extensions = {}
REGISTEREQUEST_ACCOUNT_FIELD.name = "account"
REGISTEREQUEST_ACCOUNT_FIELD.full_name = ".msg.RegisteRequest.account"
REGISTEREQUEST_ACCOUNT_FIELD.number = 1
REGISTEREQUEST_ACCOUNT_FIELD.index = 0
REGISTEREQUEST_ACCOUNT_FIELD.label = 1
REGISTEREQUEST_ACCOUNT_FIELD.has_default_value = false
REGISTEREQUEST_ACCOUNT_FIELD.default_value = ""
REGISTEREQUEST_ACCOUNT_FIELD.type = 9
REGISTEREQUEST_ACCOUNT_FIELD.cpp_type = 9

REGISTEREQUEST_PASSWORD_FIELD.name = "password"
REGISTEREQUEST_PASSWORD_FIELD.full_name = ".msg.RegisteRequest.password"
REGISTEREQUEST_PASSWORD_FIELD.number = 2
REGISTEREQUEST_PASSWORD_FIELD.index = 1
REGISTEREQUEST_PASSWORD_FIELD.label = 1
REGISTEREQUEST_PASSWORD_FIELD.has_default_value = false
REGISTEREQUEST_PASSWORD_FIELD.default_value = ""
REGISTEREQUEST_PASSWORD_FIELD.type = 9
REGISTEREQUEST_PASSWORD_FIELD.cpp_type = 9

REGISTEREQUEST.name = "RegisteRequest"
REGISTEREQUEST.full_name = ".msg.RegisteRequest"
REGISTEREQUEST.nested_types = {}
REGISTEREQUEST.enum_types = {}
REGISTEREQUEST.fields = {REGISTEREQUEST_ACCOUNT_FIELD, REGISTEREQUEST_PASSWORD_FIELD}
REGISTEREQUEST.is_extendable = false
REGISTEREQUEST.extensions = {}
REGISTERESPONSE_CODE_FIELD.name = "code"
REGISTERESPONSE_CODE_FIELD.full_name = ".msg.RegisteResponse.code"
REGISTERESPONSE_CODE_FIELD.number = 1
REGISTERESPONSE_CODE_FIELD.index = 0
REGISTERESPONSE_CODE_FIELD.label = 1
REGISTERESPONSE_CODE_FIELD.has_default_value = false
REGISTERESPONSE_CODE_FIELD.default_value = nil
REGISTERESPONSE_CODE_FIELD.enum_type = RESPONSECODE
REGISTERESPONSE_CODE_FIELD.type = 14
REGISTERESPONSE_CODE_FIELD.cpp_type = 8

REGISTERESPONSE_UID_FIELD.name = "uid"
REGISTERESPONSE_UID_FIELD.full_name = ".msg.RegisteResponse.uid"
REGISTERESPONSE_UID_FIELD.number = 2
REGISTERESPONSE_UID_FIELD.index = 1
REGISTERESPONSE_UID_FIELD.label = 1
REGISTERESPONSE_UID_FIELD.has_default_value = false
REGISTERESPONSE_UID_FIELD.default_value = ""
REGISTERESPONSE_UID_FIELD.type = 9
REGISTERESPONSE_UID_FIELD.cpp_type = 9

REGISTERESPONSE_ERR_FIELD.name = "err"
REGISTERESPONSE_ERR_FIELD.full_name = ".msg.RegisteResponse.err"
REGISTERESPONSE_ERR_FIELD.number = 3
REGISTERESPONSE_ERR_FIELD.index = 2
REGISTERESPONSE_ERR_FIELD.label = 1
REGISTERESPONSE_ERR_FIELD.has_default_value = false
REGISTERESPONSE_ERR_FIELD.default_value = nil
REGISTERESPONSE_ERR_FIELD.message_type = ERR_PB_ERROR
REGISTERESPONSE_ERR_FIELD.type = 11
REGISTERESPONSE_ERR_FIELD.cpp_type = 10

REGISTERESPONSE.name = "RegisteResponse"
REGISTERESPONSE.full_name = ".msg.RegisteResponse"
REGISTERESPONSE.nested_types = {}
REGISTERESPONSE.enum_types = {}
REGISTERESPONSE.fields = {REGISTERESPONSE_CODE_FIELD, REGISTERESPONSE_UID_FIELD, REGISTERESPONSE_ERR_FIELD}
REGISTERESPONSE.is_extendable = false
REGISTERESPONSE.extensions = {}

FAIL = 0
LoginRequest = protobuf.Message(LOGINREQUEST)
LoginResponse = protobuf.Message(LOGINRESPONSE)
RegisteRequest = protobuf.Message(REGISTEREQUEST)
RegisteResponse = protobuf.Message(REGISTERESPONSE)
SUCCESS = 1

