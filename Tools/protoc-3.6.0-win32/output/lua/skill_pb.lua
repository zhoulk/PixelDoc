-- Generated By protoc-gen-lua Do not Edit
local protobuf = require "protobuf/protobuf"
local err_pb = require("PixelFarm/PBLua/err_pb")
module('skill_pb')


local SKILLREQUEST = protobuf.Descriptor();
local SKILLRESPONSE = protobuf.Descriptor();
local SKILLRESPONSE_CODE_FIELD = protobuf.FieldDescriptor();
local SKILLRESPONSE_ERR_FIELD = protobuf.FieldDescriptor();
local SKILLRESPONSE_SKILLS_FIELD = protobuf.FieldDescriptor();
local SKILL = protobuf.Descriptor();
local SKILL_ID_FIELD = protobuf.FieldDescriptor();
local SKILL_NAME_FIELD = protobuf.FieldDescriptor();
local SKILL_LEVEL_FIELD = protobuf.FieldDescriptor();
local SKILL_TYPE_FIELD = protobuf.FieldDescriptor();
local SKILL_DESC_FIELD = protobuf.FieldDescriptor();

SKILLREQUEST.name = "SkillRequest"
SKILLREQUEST.full_name = ".msg.SkillRequest"
SKILLREQUEST.nested_types = {}
SKILLREQUEST.enum_types = {}
SKILLREQUEST.fields = {}
SKILLREQUEST.is_extendable = false
SKILLREQUEST.extensions = {}
SKILLRESPONSE_CODE_FIELD.name = "code"
SKILLRESPONSE_CODE_FIELD.full_name = ".msg.SkillResponse.code"
SKILLRESPONSE_CODE_FIELD.number = 1
SKILLRESPONSE_CODE_FIELD.index = 0
SKILLRESPONSE_CODE_FIELD.label = 1
SKILLRESPONSE_CODE_FIELD.has_default_value = false
SKILLRESPONSE_CODE_FIELD.default_value = nil
SKILLRESPONSE_CODE_FIELD.enum_type = ERR_PB_RESPONSECODE
SKILLRESPONSE_CODE_FIELD.type = 14
SKILLRESPONSE_CODE_FIELD.cpp_type = 8

SKILLRESPONSE_ERR_FIELD.name = "err"
SKILLRESPONSE_ERR_FIELD.full_name = ".msg.SkillResponse.err"
SKILLRESPONSE_ERR_FIELD.number = 2
SKILLRESPONSE_ERR_FIELD.index = 1
SKILLRESPONSE_ERR_FIELD.label = 1
SKILLRESPONSE_ERR_FIELD.has_default_value = false
SKILLRESPONSE_ERR_FIELD.default_value = nil
SKILLRESPONSE_ERR_FIELD.message_type = ERR_PB_ERROR
SKILLRESPONSE_ERR_FIELD.type = 11
SKILLRESPONSE_ERR_FIELD.cpp_type = 10

SKILLRESPONSE_SKILLS_FIELD.name = "skills"
SKILLRESPONSE_SKILLS_FIELD.full_name = ".msg.SkillResponse.skills"
SKILLRESPONSE_SKILLS_FIELD.number = 3
SKILLRESPONSE_SKILLS_FIELD.index = 2
SKILLRESPONSE_SKILLS_FIELD.label = 3
SKILLRESPONSE_SKILLS_FIELD.has_default_value = false
SKILLRESPONSE_SKILLS_FIELD.default_value = {}
SKILLRESPONSE_SKILLS_FIELD.message_type = _SKILL
SKILLRESPONSE_SKILLS_FIELD.type = 11
SKILLRESPONSE_SKILLS_FIELD.cpp_type = 10

SKILLRESPONSE.name = "SkillResponse"
SKILLRESPONSE.full_name = ".msg.SkillResponse"
SKILLRESPONSE.nested_types = {}
SKILLRESPONSE.enum_types = {}
SKILLRESPONSE.fields = {SKILLRESPONSE_CODE_FIELD, SKILLRESPONSE_ERR_FIELD, SKILLRESPONSE_SKILLS_FIELD}
SKILLRESPONSE.is_extendable = false
SKILLRESPONSE.extensions = {}
SKILL_ID_FIELD.name = "Id"
SKILL_ID_FIELD.full_name = ".msg.Skill.Id"
SKILL_ID_FIELD.number = 1
SKILL_ID_FIELD.index = 0
SKILL_ID_FIELD.label = 1
SKILL_ID_FIELD.has_default_value = false
SKILL_ID_FIELD.default_value = 0
SKILL_ID_FIELD.type = 5
SKILL_ID_FIELD.cpp_type = 1

SKILL_NAME_FIELD.name = "Name"
SKILL_NAME_FIELD.full_name = ".msg.Skill.Name"
SKILL_NAME_FIELD.number = 2
SKILL_NAME_FIELD.index = 1
SKILL_NAME_FIELD.label = 1
SKILL_NAME_FIELD.has_default_value = false
SKILL_NAME_FIELD.default_value = ""
SKILL_NAME_FIELD.type = 9
SKILL_NAME_FIELD.cpp_type = 9

SKILL_LEVEL_FIELD.name = "Level"
SKILL_LEVEL_FIELD.full_name = ".msg.Skill.Level"
SKILL_LEVEL_FIELD.number = 3
SKILL_LEVEL_FIELD.index = 2
SKILL_LEVEL_FIELD.label = 1
SKILL_LEVEL_FIELD.has_default_value = false
SKILL_LEVEL_FIELD.default_value = 0
SKILL_LEVEL_FIELD.type = 5
SKILL_LEVEL_FIELD.cpp_type = 1

SKILL_TYPE_FIELD.name = "Type"
SKILL_TYPE_FIELD.full_name = ".msg.Skill.Type"
SKILL_TYPE_FIELD.number = 4
SKILL_TYPE_FIELD.index = 3
SKILL_TYPE_FIELD.label = 1
SKILL_TYPE_FIELD.has_default_value = false
SKILL_TYPE_FIELD.default_value = 0
SKILL_TYPE_FIELD.type = 5
SKILL_TYPE_FIELD.cpp_type = 1

SKILL_DESC_FIELD.name = "Desc"
SKILL_DESC_FIELD.full_name = ".msg.Skill.Desc"
SKILL_DESC_FIELD.number = 5
SKILL_DESC_FIELD.index = 4
SKILL_DESC_FIELD.label = 1
SKILL_DESC_FIELD.has_default_value = false
SKILL_DESC_FIELD.default_value = ""
SKILL_DESC_FIELD.type = 9
SKILL_DESC_FIELD.cpp_type = 9

SKILL.name = "Skill"
SKILL.full_name = ".msg.Skill"
SKILL.nested_types = {}
SKILL.enum_types = {}
SKILL.fields = {SKILL_ID_FIELD, SKILL_NAME_FIELD, SKILL_LEVEL_FIELD, SKILL_TYPE_FIELD, SKILL_DESC_FIELD}
SKILL.is_extendable = false
SKILL.extensions = {}

Skill = protobuf.Message(SKILL)
SkillRequest = protobuf.Message(SKILLREQUEST)
SkillResponse = protobuf.Message(SKILLRESPONSE)

