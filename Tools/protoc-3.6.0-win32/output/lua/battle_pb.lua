-- Generated By protoc-gen-lua Do not Edit
local protobuf = require "protobuf/protobuf"
local err_pb = require("PixelFarm/PBLua/err_pb")
local map_pb = require("PixelFarm/PBLua/map_pb")
local hero_pb = require("PixelFarm/PBLua/hero_pb")
local skill_pb = require("PixelFarm/PBLua/skill_pb")
local item_pb = require("PixelFarm/PBLua/item_pb")
module('battle_pb')


local BATTLECREATEREQUEST = protobuf.Descriptor();
local BATTLECREATEREQUEST_TYPE_FIELD = protobuf.FieldDescriptor();
local BATTLECREATEREQUEST_ARGS_FIELD = protobuf.FieldDescriptor();
local BATTLECREATERESPONSE = protobuf.Descriptor();
local BATTLECREATERESPONSE_CODE_FIELD = protobuf.FieldDescriptor();
local BATTLECREATERESPONSE_ERR_FIELD = protobuf.FieldDescriptor();
local BATTLECREATERESPONSE_BATTLEID_FIELD = protobuf.FieldDescriptor();
local BATTLESTARTREQUEST = protobuf.Descriptor();
local BATTLESTARTREQUEST_BATTLEID_FIELD = protobuf.FieldDescriptor();
local BATTLESTARTRESPONSE = protobuf.Descriptor();
local BATTLESTARTRESPONSE_CODE_FIELD = protobuf.FieldDescriptor();
local BATTLESTARTRESPONSE_ERR_FIELD = protobuf.FieldDescriptor();
local BATTLESTARTRESPONSE_HEROS_FIELD = protobuf.FieldDescriptor();
local BATTLESTARTRESPONSE_SKILLS_FIELD = protobuf.FieldDescriptor();
local BATTLESTARTRESPONSE_EQUIPS_FIELD = protobuf.FieldDescriptor();
local BATTLEGUANKAREQUEST = protobuf.Descriptor();
local BATTLEGUANKAREQUEST_GUANKAID_FIELD = protobuf.FieldDescriptor();
local BATTLEGUANKARESPONSE = protobuf.Descriptor();
local BATTLEGUANKARESPONSE_CODE_FIELD = protobuf.FieldDescriptor();
local BATTLEGUANKARESPONSE_ERR_FIELD = protobuf.FieldDescriptor();
local BATTLEGUANKARESPONSE_RESULT_FIELD = protobuf.FieldDescriptor();
local BATTLEGUANKARESPONSE_GUANKA_FIELD = protobuf.FieldDescriptor();
local BATTLERESULTREQUEST = protobuf.Descriptor();
local BATTLERESULTREQUEST_BATTLEID_FIELD = protobuf.FieldDescriptor();
local BATTLERESULTREQUEST_RESULT_FIELD = protobuf.FieldDescriptor();
local BATTLERESULTRESPONSE = protobuf.Descriptor();
local BATTLERESULTRESPONSE_CODE_FIELD = protobuf.FieldDescriptor();
local BATTLERESULTRESPONSE_ERR_FIELD = protobuf.FieldDescriptor();
local BATTLERESULTRESPONSE_EARN_FIELD = protobuf.FieldDescriptor();
local BATTLERESULTRESPONSE_EXP_FIELD = protobuf.FieldDescriptor();
local BATTLERESULTRESPONSE_LEVEL_FIELD = protobuf.FieldDescriptor();
local BATTLERESULTRESPONSE_LEVELUPEXP_FIELD = protobuf.FieldDescriptor();

BATTLECREATEREQUEST_TYPE_FIELD.name = "type"
BATTLECREATEREQUEST_TYPE_FIELD.full_name = ".msg.BattleCreateRequest.type"
BATTLECREATEREQUEST_TYPE_FIELD.number = 1
BATTLECREATEREQUEST_TYPE_FIELD.index = 0
BATTLECREATEREQUEST_TYPE_FIELD.label = 1
BATTLECREATEREQUEST_TYPE_FIELD.has_default_value = false
BATTLECREATEREQUEST_TYPE_FIELD.default_value = 0
BATTLECREATEREQUEST_TYPE_FIELD.type = 5
BATTLECREATEREQUEST_TYPE_FIELD.cpp_type = 1

BATTLECREATEREQUEST_ARGS_FIELD.name = "args"
BATTLECREATEREQUEST_ARGS_FIELD.full_name = ".msg.BattleCreateRequest.args"
BATTLECREATEREQUEST_ARGS_FIELD.number = 2
BATTLECREATEREQUEST_ARGS_FIELD.index = 1
BATTLECREATEREQUEST_ARGS_FIELD.label = 3
BATTLECREATEREQUEST_ARGS_FIELD.has_default_value = false
BATTLECREATEREQUEST_ARGS_FIELD.default_value = {}
BATTLECREATEREQUEST_ARGS_FIELD.type = 9
BATTLECREATEREQUEST_ARGS_FIELD.cpp_type = 9

BATTLECREATEREQUEST.name = "BattleCreateRequest"
BATTLECREATEREQUEST.full_name = ".msg.BattleCreateRequest"
BATTLECREATEREQUEST.nested_types = {}
BATTLECREATEREQUEST.enum_types = {}
BATTLECREATEREQUEST.fields = {BATTLECREATEREQUEST_TYPE_FIELD, BATTLECREATEREQUEST_ARGS_FIELD}
BATTLECREATEREQUEST.is_extendable = false
BATTLECREATEREQUEST.extensions = {}
BATTLECREATERESPONSE_CODE_FIELD.name = "code"
BATTLECREATERESPONSE_CODE_FIELD.full_name = ".msg.BattleCreateResponse.code"
BATTLECREATERESPONSE_CODE_FIELD.number = 1
BATTLECREATERESPONSE_CODE_FIELD.index = 0
BATTLECREATERESPONSE_CODE_FIELD.label = 1
BATTLECREATERESPONSE_CODE_FIELD.has_default_value = false
BATTLECREATERESPONSE_CODE_FIELD.default_value = nil
BATTLECREATERESPONSE_CODE_FIELD.enum_type = ERR_PB_RESPONSECODE
BATTLECREATERESPONSE_CODE_FIELD.type = 14
BATTLECREATERESPONSE_CODE_FIELD.cpp_type = 8

BATTLECREATERESPONSE_ERR_FIELD.name = "err"
BATTLECREATERESPONSE_ERR_FIELD.full_name = ".msg.BattleCreateResponse.err"
BATTLECREATERESPONSE_ERR_FIELD.number = 2
BATTLECREATERESPONSE_ERR_FIELD.index = 1
BATTLECREATERESPONSE_ERR_FIELD.label = 1
BATTLECREATERESPONSE_ERR_FIELD.has_default_value = false
BATTLECREATERESPONSE_ERR_FIELD.default_value = nil
BATTLECREATERESPONSE_ERR_FIELD.message_type = ERR_PB_ERROR
BATTLECREATERESPONSE_ERR_FIELD.type = 11
BATTLECREATERESPONSE_ERR_FIELD.cpp_type = 10

BATTLECREATERESPONSE_BATTLEID_FIELD.name = "battleId"
BATTLECREATERESPONSE_BATTLEID_FIELD.full_name = ".msg.BattleCreateResponse.battleId"
BATTLECREATERESPONSE_BATTLEID_FIELD.number = 3
BATTLECREATERESPONSE_BATTLEID_FIELD.index = 2
BATTLECREATERESPONSE_BATTLEID_FIELD.label = 1
BATTLECREATERESPONSE_BATTLEID_FIELD.has_default_value = false
BATTLECREATERESPONSE_BATTLEID_FIELD.default_value = ""
BATTLECREATERESPONSE_BATTLEID_FIELD.type = 9
BATTLECREATERESPONSE_BATTLEID_FIELD.cpp_type = 9

BATTLECREATERESPONSE.name = "BattleCreateResponse"
BATTLECREATERESPONSE.full_name = ".msg.BattleCreateResponse"
BATTLECREATERESPONSE.nested_types = {}
BATTLECREATERESPONSE.enum_types = {}
BATTLECREATERESPONSE.fields = {BATTLECREATERESPONSE_CODE_FIELD, BATTLECREATERESPONSE_ERR_FIELD, BATTLECREATERESPONSE_BATTLEID_FIELD}
BATTLECREATERESPONSE.is_extendable = false
BATTLECREATERESPONSE.extensions = {}
BATTLESTARTREQUEST_BATTLEID_FIELD.name = "battleId"
BATTLESTARTREQUEST_BATTLEID_FIELD.full_name = ".msg.BattleStartRequest.battleId"
BATTLESTARTREQUEST_BATTLEID_FIELD.number = 1
BATTLESTARTREQUEST_BATTLEID_FIELD.index = 0
BATTLESTARTREQUEST_BATTLEID_FIELD.label = 1
BATTLESTARTREQUEST_BATTLEID_FIELD.has_default_value = false
BATTLESTARTREQUEST_BATTLEID_FIELD.default_value = ""
BATTLESTARTREQUEST_BATTLEID_FIELD.type = 9
BATTLESTARTREQUEST_BATTLEID_FIELD.cpp_type = 9

BATTLESTARTREQUEST.name = "BattleStartRequest"
BATTLESTARTREQUEST.full_name = ".msg.BattleStartRequest"
BATTLESTARTREQUEST.nested_types = {}
BATTLESTARTREQUEST.enum_types = {}
BATTLESTARTREQUEST.fields = {BATTLESTARTREQUEST_BATTLEID_FIELD}
BATTLESTARTREQUEST.is_extendable = false
BATTLESTARTREQUEST.extensions = {}
BATTLESTARTRESPONSE_CODE_FIELD.name = "code"
BATTLESTARTRESPONSE_CODE_FIELD.full_name = ".msg.BattleStartResponse.code"
BATTLESTARTRESPONSE_CODE_FIELD.number = 1
BATTLESTARTRESPONSE_CODE_FIELD.index = 0
BATTLESTARTRESPONSE_CODE_FIELD.label = 1
BATTLESTARTRESPONSE_CODE_FIELD.has_default_value = false
BATTLESTARTRESPONSE_CODE_FIELD.default_value = nil
BATTLESTARTRESPONSE_CODE_FIELD.enum_type = ERR_PB_RESPONSECODE
BATTLESTARTRESPONSE_CODE_FIELD.type = 14
BATTLESTARTRESPONSE_CODE_FIELD.cpp_type = 8

BATTLESTARTRESPONSE_ERR_FIELD.name = "err"
BATTLESTARTRESPONSE_ERR_FIELD.full_name = ".msg.BattleStartResponse.err"
BATTLESTARTRESPONSE_ERR_FIELD.number = 2
BATTLESTARTRESPONSE_ERR_FIELD.index = 1
BATTLESTARTRESPONSE_ERR_FIELD.label = 1
BATTLESTARTRESPONSE_ERR_FIELD.has_default_value = false
BATTLESTARTRESPONSE_ERR_FIELD.default_value = nil
BATTLESTARTRESPONSE_ERR_FIELD.message_type = ERR_PB_ERROR
BATTLESTARTRESPONSE_ERR_FIELD.type = 11
BATTLESTARTRESPONSE_ERR_FIELD.cpp_type = 10

BATTLESTARTRESPONSE_HEROS_FIELD.name = "heros"
BATTLESTARTRESPONSE_HEROS_FIELD.full_name = ".msg.BattleStartResponse.heros"
BATTLESTARTRESPONSE_HEROS_FIELD.number = 3
BATTLESTARTRESPONSE_HEROS_FIELD.index = 2
BATTLESTARTRESPONSE_HEROS_FIELD.label = 3
BATTLESTARTRESPONSE_HEROS_FIELD.has_default_value = false
BATTLESTARTRESPONSE_HEROS_FIELD.default_value = {}
BATTLESTARTRESPONSE_HEROS_FIELD.message_type = HERO_PB_HERO
BATTLESTARTRESPONSE_HEROS_FIELD.type = 11
BATTLESTARTRESPONSE_HEROS_FIELD.cpp_type = 10

BATTLESTARTRESPONSE_SKILLS_FIELD.name = "skills"
BATTLESTARTRESPONSE_SKILLS_FIELD.full_name = ".msg.BattleStartResponse.skills"
BATTLESTARTRESPONSE_SKILLS_FIELD.number = 4
BATTLESTARTRESPONSE_SKILLS_FIELD.index = 3
BATTLESTARTRESPONSE_SKILLS_FIELD.label = 3
BATTLESTARTRESPONSE_SKILLS_FIELD.has_default_value = false
BATTLESTARTRESPONSE_SKILLS_FIELD.default_value = {}
BATTLESTARTRESPONSE_SKILLS_FIELD.message_type = SKILL_PB_SKILL
BATTLESTARTRESPONSE_SKILLS_FIELD.type = 11
BATTLESTARTRESPONSE_SKILLS_FIELD.cpp_type = 10

BATTLESTARTRESPONSE_EQUIPS_FIELD.name = "equips"
BATTLESTARTRESPONSE_EQUIPS_FIELD.full_name = ".msg.BattleStartResponse.equips"
BATTLESTARTRESPONSE_EQUIPS_FIELD.number = 5
BATTLESTARTRESPONSE_EQUIPS_FIELD.index = 4
BATTLESTARTRESPONSE_EQUIPS_FIELD.label = 3
BATTLESTARTRESPONSE_EQUIPS_FIELD.has_default_value = false
BATTLESTARTRESPONSE_EQUIPS_FIELD.default_value = {}
BATTLESTARTRESPONSE_EQUIPS_FIELD.message_type = ITEM_PB_EQUIP
BATTLESTARTRESPONSE_EQUIPS_FIELD.type = 11
BATTLESTARTRESPONSE_EQUIPS_FIELD.cpp_type = 10

BATTLESTARTRESPONSE.name = "BattleStartResponse"
BATTLESTARTRESPONSE.full_name = ".msg.BattleStartResponse"
BATTLESTARTRESPONSE.nested_types = {}
BATTLESTARTRESPONSE.enum_types = {}
BATTLESTARTRESPONSE.fields = {BATTLESTARTRESPONSE_CODE_FIELD, BATTLESTARTRESPONSE_ERR_FIELD, BATTLESTARTRESPONSE_HEROS_FIELD, BATTLESTARTRESPONSE_SKILLS_FIELD, BATTLESTARTRESPONSE_EQUIPS_FIELD}
BATTLESTARTRESPONSE.is_extendable = false
BATTLESTARTRESPONSE.extensions = {}
BATTLEGUANKAREQUEST_GUANKAID_FIELD.name = "guanKaId"
BATTLEGUANKAREQUEST_GUANKAID_FIELD.full_name = ".msg.BattleGuanKaRequest.guanKaId"
BATTLEGUANKAREQUEST_GUANKAID_FIELD.number = 1
BATTLEGUANKAREQUEST_GUANKAID_FIELD.index = 0
BATTLEGUANKAREQUEST_GUANKAID_FIELD.label = 1
BATTLEGUANKAREQUEST_GUANKAID_FIELD.has_default_value = false
BATTLEGUANKAREQUEST_GUANKAID_FIELD.default_value = 0
BATTLEGUANKAREQUEST_GUANKAID_FIELD.type = 5
BATTLEGUANKAREQUEST_GUANKAID_FIELD.cpp_type = 1

BATTLEGUANKAREQUEST.name = "BattleGuanKaRequest"
BATTLEGUANKAREQUEST.full_name = ".msg.BattleGuanKaRequest"
BATTLEGUANKAREQUEST.nested_types = {}
BATTLEGUANKAREQUEST.enum_types = {}
BATTLEGUANKAREQUEST.fields = {BATTLEGUANKAREQUEST_GUANKAID_FIELD}
BATTLEGUANKAREQUEST.is_extendable = false
BATTLEGUANKAREQUEST.extensions = {}
BATTLEGUANKARESPONSE_CODE_FIELD.name = "code"
BATTLEGUANKARESPONSE_CODE_FIELD.full_name = ".msg.BattleGuanKaResponse.code"
BATTLEGUANKARESPONSE_CODE_FIELD.number = 1
BATTLEGUANKARESPONSE_CODE_FIELD.index = 0
BATTLEGUANKARESPONSE_CODE_FIELD.label = 1
BATTLEGUANKARESPONSE_CODE_FIELD.has_default_value = false
BATTLEGUANKARESPONSE_CODE_FIELD.default_value = nil
BATTLEGUANKARESPONSE_CODE_FIELD.enum_type = ERR_PB_RESPONSECODE
BATTLEGUANKARESPONSE_CODE_FIELD.type = 14
BATTLEGUANKARESPONSE_CODE_FIELD.cpp_type = 8

BATTLEGUANKARESPONSE_ERR_FIELD.name = "err"
BATTLEGUANKARESPONSE_ERR_FIELD.full_name = ".msg.BattleGuanKaResponse.err"
BATTLEGUANKARESPONSE_ERR_FIELD.number = 2
BATTLEGUANKARESPONSE_ERR_FIELD.index = 1
BATTLEGUANKARESPONSE_ERR_FIELD.label = 1
BATTLEGUANKARESPONSE_ERR_FIELD.has_default_value = false
BATTLEGUANKARESPONSE_ERR_FIELD.default_value = nil
BATTLEGUANKARESPONSE_ERR_FIELD.message_type = ERR_PB_ERROR
BATTLEGUANKARESPONSE_ERR_FIELD.type = 11
BATTLEGUANKARESPONSE_ERR_FIELD.cpp_type = 10

BATTLEGUANKARESPONSE_RESULT_FIELD.name = "result"
BATTLEGUANKARESPONSE_RESULT_FIELD.full_name = ".msg.BattleGuanKaResponse.result"
BATTLEGUANKARESPONSE_RESULT_FIELD.number = 3
BATTLEGUANKARESPONSE_RESULT_FIELD.index = 2
BATTLEGUANKARESPONSE_RESULT_FIELD.label = 1
BATTLEGUANKARESPONSE_RESULT_FIELD.has_default_value = false
BATTLEGUANKARESPONSE_RESULT_FIELD.default_value = 0
BATTLEGUANKARESPONSE_RESULT_FIELD.type = 5
BATTLEGUANKARESPONSE_RESULT_FIELD.cpp_type = 1

BATTLEGUANKARESPONSE_GUANKA_FIELD.name = "guanka"
BATTLEGUANKARESPONSE_GUANKA_FIELD.full_name = ".msg.BattleGuanKaResponse.guanka"
BATTLEGUANKARESPONSE_GUANKA_FIELD.number = 4
BATTLEGUANKARESPONSE_GUANKA_FIELD.index = 3
BATTLEGUANKARESPONSE_GUANKA_FIELD.label = 1
BATTLEGUANKARESPONSE_GUANKA_FIELD.has_default_value = false
BATTLEGUANKARESPONSE_GUANKA_FIELD.default_value = nil
BATTLEGUANKARESPONSE_GUANKA_FIELD.message_type = MAP_PB_GUANKA
BATTLEGUANKARESPONSE_GUANKA_FIELD.type = 11
BATTLEGUANKARESPONSE_GUANKA_FIELD.cpp_type = 10

BATTLEGUANKARESPONSE.name = "BattleGuanKaResponse"
BATTLEGUANKARESPONSE.full_name = ".msg.BattleGuanKaResponse"
BATTLEGUANKARESPONSE.nested_types = {}
BATTLEGUANKARESPONSE.enum_types = {}
BATTLEGUANKARESPONSE.fields = {BATTLEGUANKARESPONSE_CODE_FIELD, BATTLEGUANKARESPONSE_ERR_FIELD, BATTLEGUANKARESPONSE_RESULT_FIELD, BATTLEGUANKARESPONSE_GUANKA_FIELD}
BATTLEGUANKARESPONSE.is_extendable = false
BATTLEGUANKARESPONSE.extensions = {}
BATTLERESULTREQUEST_BATTLEID_FIELD.name = "battleId"
BATTLERESULTREQUEST_BATTLEID_FIELD.full_name = ".msg.BattleResultRequest.battleId"
BATTLERESULTREQUEST_BATTLEID_FIELD.number = 1
BATTLERESULTREQUEST_BATTLEID_FIELD.index = 0
BATTLERESULTREQUEST_BATTLEID_FIELD.label = 1
BATTLERESULTREQUEST_BATTLEID_FIELD.has_default_value = false
BATTLERESULTREQUEST_BATTLEID_FIELD.default_value = ""
BATTLERESULTREQUEST_BATTLEID_FIELD.type = 9
BATTLERESULTREQUEST_BATTLEID_FIELD.cpp_type = 9

BATTLERESULTREQUEST_RESULT_FIELD.name = "result"
BATTLERESULTREQUEST_RESULT_FIELD.full_name = ".msg.BattleResultRequest.result"
BATTLERESULTREQUEST_RESULT_FIELD.number = 2
BATTLERESULTREQUEST_RESULT_FIELD.index = 1
BATTLERESULTREQUEST_RESULT_FIELD.label = 1
BATTLERESULTREQUEST_RESULT_FIELD.has_default_value = false
BATTLERESULTREQUEST_RESULT_FIELD.default_value = 0
BATTLERESULTREQUEST_RESULT_FIELD.type = 5
BATTLERESULTREQUEST_RESULT_FIELD.cpp_type = 1

BATTLERESULTREQUEST.name = "BattleResultRequest"
BATTLERESULTREQUEST.full_name = ".msg.BattleResultRequest"
BATTLERESULTREQUEST.nested_types = {}
BATTLERESULTREQUEST.enum_types = {}
BATTLERESULTREQUEST.fields = {BATTLERESULTREQUEST_BATTLEID_FIELD, BATTLERESULTREQUEST_RESULT_FIELD}
BATTLERESULTREQUEST.is_extendable = false
BATTLERESULTREQUEST.extensions = {}
BATTLERESULTRESPONSE_CODE_FIELD.name = "code"
BATTLERESULTRESPONSE_CODE_FIELD.full_name = ".msg.BattleResultResponse.code"
BATTLERESULTRESPONSE_CODE_FIELD.number = 1
BATTLERESULTRESPONSE_CODE_FIELD.index = 0
BATTLERESULTRESPONSE_CODE_FIELD.label = 1
BATTLERESULTRESPONSE_CODE_FIELD.has_default_value = false
BATTLERESULTRESPONSE_CODE_FIELD.default_value = nil
BATTLERESULTRESPONSE_CODE_FIELD.enum_type = ERR_PB_RESPONSECODE
BATTLERESULTRESPONSE_CODE_FIELD.type = 14
BATTLERESULTRESPONSE_CODE_FIELD.cpp_type = 8

BATTLERESULTRESPONSE_ERR_FIELD.name = "err"
BATTLERESULTRESPONSE_ERR_FIELD.full_name = ".msg.BattleResultResponse.err"
BATTLERESULTRESPONSE_ERR_FIELD.number = 2
BATTLERESULTRESPONSE_ERR_FIELD.index = 1
BATTLERESULTRESPONSE_ERR_FIELD.label = 1
BATTLERESULTRESPONSE_ERR_FIELD.has_default_value = false
BATTLERESULTRESPONSE_ERR_FIELD.default_value = nil
BATTLERESULTRESPONSE_ERR_FIELD.message_type = ERR_PB_ERROR
BATTLERESULTRESPONSE_ERR_FIELD.type = 11
BATTLERESULTRESPONSE_ERR_FIELD.cpp_type = 10

BATTLERESULTRESPONSE_EARN_FIELD.name = "earn"
BATTLERESULTRESPONSE_EARN_FIELD.full_name = ".msg.BattleResultResponse.earn"
BATTLERESULTRESPONSE_EARN_FIELD.number = 3
BATTLERESULTRESPONSE_EARN_FIELD.index = 2
BATTLERESULTRESPONSE_EARN_FIELD.label = 1
BATTLERESULTRESPONSE_EARN_FIELD.has_default_value = false
BATTLERESULTRESPONSE_EARN_FIELD.default_value = nil
BATTLERESULTRESPONSE_EARN_FIELD.message_type = MAP_PB_EARN
BATTLERESULTRESPONSE_EARN_FIELD.type = 11
BATTLERESULTRESPONSE_EARN_FIELD.cpp_type = 10

BATTLERESULTRESPONSE_EXP_FIELD.name = "Exp"
BATTLERESULTRESPONSE_EXP_FIELD.full_name = ".msg.BattleResultResponse.Exp"
BATTLERESULTRESPONSE_EXP_FIELD.number = 4
BATTLERESULTRESPONSE_EXP_FIELD.index = 3
BATTLERESULTRESPONSE_EXP_FIELD.label = 1
BATTLERESULTRESPONSE_EXP_FIELD.has_default_value = false
BATTLERESULTRESPONSE_EXP_FIELD.default_value = 0
BATTLERESULTRESPONSE_EXP_FIELD.type = 5
BATTLERESULTRESPONSE_EXP_FIELD.cpp_type = 1

BATTLERESULTRESPONSE_LEVEL_FIELD.name = "Level"
BATTLERESULTRESPONSE_LEVEL_FIELD.full_name = ".msg.BattleResultResponse.Level"
BATTLERESULTRESPONSE_LEVEL_FIELD.number = 5
BATTLERESULTRESPONSE_LEVEL_FIELD.index = 4
BATTLERESULTRESPONSE_LEVEL_FIELD.label = 1
BATTLERESULTRESPONSE_LEVEL_FIELD.has_default_value = false
BATTLERESULTRESPONSE_LEVEL_FIELD.default_value = 0
BATTLERESULTRESPONSE_LEVEL_FIELD.type = 5
BATTLERESULTRESPONSE_LEVEL_FIELD.cpp_type = 1

BATTLERESULTRESPONSE_LEVELUPEXP_FIELD.name = "LevelUpExp"
BATTLERESULTRESPONSE_LEVELUPEXP_FIELD.full_name = ".msg.BattleResultResponse.LevelUpExp"
BATTLERESULTRESPONSE_LEVELUPEXP_FIELD.number = 6
BATTLERESULTRESPONSE_LEVELUPEXP_FIELD.index = 5
BATTLERESULTRESPONSE_LEVELUPEXP_FIELD.label = 1
BATTLERESULTRESPONSE_LEVELUPEXP_FIELD.has_default_value = false
BATTLERESULTRESPONSE_LEVELUPEXP_FIELD.default_value = 0
BATTLERESULTRESPONSE_LEVELUPEXP_FIELD.type = 5
BATTLERESULTRESPONSE_LEVELUPEXP_FIELD.cpp_type = 1

BATTLERESULTRESPONSE.name = "BattleResultResponse"
BATTLERESULTRESPONSE.full_name = ".msg.BattleResultResponse"
BATTLERESULTRESPONSE.nested_types = {}
BATTLERESULTRESPONSE.enum_types = {}
BATTLERESULTRESPONSE.fields = {BATTLERESULTRESPONSE_CODE_FIELD, BATTLERESULTRESPONSE_ERR_FIELD, BATTLERESULTRESPONSE_EARN_FIELD, BATTLERESULTRESPONSE_EXP_FIELD, BATTLERESULTRESPONSE_LEVEL_FIELD, BATTLERESULTRESPONSE_LEVELUPEXP_FIELD}
BATTLERESULTRESPONSE.is_extendable = false
BATTLERESULTRESPONSE.extensions = {}

BattleCreateRequest = protobuf.Message(BATTLECREATEREQUEST)
BattleCreateResponse = protobuf.Message(BATTLECREATERESPONSE)
BattleGuanKaRequest = protobuf.Message(BATTLEGUANKAREQUEST)
BattleGuanKaResponse = protobuf.Message(BATTLEGUANKARESPONSE)
BattleResultRequest = protobuf.Message(BATTLERESULTREQUEST)
BattleResultResponse = protobuf.Message(BATTLERESULTRESPONSE)
BattleStartRequest = protobuf.Message(BATTLESTARTREQUEST)
BattleStartResponse = protobuf.Message(BATTLESTARTRESPONSE)

