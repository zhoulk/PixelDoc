// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: skill.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Msg {

  /// <summary>Holder for reflection information generated from skill.proto</summary>
  public static partial class SkillReflection {

    #region Descriptor
    /// <summary>File descriptor for skill.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static SkillReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "Cgtza2lsbC5wcm90bxIDbXNnGgllcnIucHJvdG8iDgoMU2tpbGxSZXF1ZXN0",
            "ImUKDVNraWxsUmVzcG9uc2USHwoEY29kZRgBIAEoDjIRLm1zZy5SZXNwb25z",
            "ZUNvZGUSFwoDZXJyGAIgASgLMgoubXNnLkVycm9yEhoKBnNraWxscxgDIAMo",
            "CzIKLm1zZy5Ta2lsbCImChNTa2lsbFVwZ3JhZGVSZXF1ZXN0Eg8KB3NraWxs",
            "SWQYASABKAkiawoUU2tpbGxVcGdyYWRlUmVzcG9uc2USHwoEY29kZRgBIAEo",
            "DjIRLm1zZy5SZXNwb25zZUNvZGUSFwoDZXJyGAIgASgLMgoubXNnLkVycm9y",
            "EhkKBXNraWxsGAMgASgLMgoubXNnLlNraWxsIpABCgVTa2lsbBIKCgJJZBgB",
            "IAEoBRIMCgROYW1lGAIgASgJEg0KBUxldmVsGAMgASgFEgwKBFR5cGUYBCAB",
            "KAUSDAoERGVzYxgFIAEoCRIOCgZJc09wZW4YBiABKAgSDwoHU2tpbGxJZBgH",
            "IAEoCRIOCgZIZXJvSWQYCCABKAkSEQoJTGV2ZWxEZXNjGAkgAygJYgZwcm90",
            "bzM="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Msg.ErrReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Msg.SkillRequest), global::Msg.SkillRequest.Parser, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Msg.SkillResponse), global::Msg.SkillResponse.Parser, new[]{ "Code", "Err", "Skills" }, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Msg.SkillUpgradeRequest), global::Msg.SkillUpgradeRequest.Parser, new[]{ "SkillId" }, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Msg.SkillUpgradeResponse), global::Msg.SkillUpgradeResponse.Parser, new[]{ "Code", "Err", "Skill" }, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Msg.Skill), global::Msg.Skill.Parser, new[]{ "Id", "Name", "Level", "Type", "Desc", "IsOpen", "SkillId", "HeroId", "LevelDesc" }, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class SkillRequest : pb::IMessage<SkillRequest> {
    private static readonly pb::MessageParser<SkillRequest> _parser = new pb::MessageParser<SkillRequest>(() => new SkillRequest());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<SkillRequest> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Msg.SkillReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public SkillRequest() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public SkillRequest(SkillRequest other) : this() {
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public SkillRequest Clone() {
      return new SkillRequest(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as SkillRequest);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(SkillRequest other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(SkillRequest other) {
      if (other == null) {
        return;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
        }
      }
    }

  }

  public sealed partial class SkillResponse : pb::IMessage<SkillResponse> {
    private static readonly pb::MessageParser<SkillResponse> _parser = new pb::MessageParser<SkillResponse>(() => new SkillResponse());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<SkillResponse> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Msg.SkillReflection.Descriptor.MessageTypes[1]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public SkillResponse() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public SkillResponse(SkillResponse other) : this() {
      code_ = other.code_;
      err_ = other.err_ != null ? other.err_.Clone() : null;
      skills_ = other.skills_.Clone();
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public SkillResponse Clone() {
      return new SkillResponse(this);
    }

    /// <summary>Field number for the "code" field.</summary>
    public const int CodeFieldNumber = 1;
    private global::Msg.ResponseCode code_ = 0;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Msg.ResponseCode Code {
      get { return code_; }
      set {
        code_ = value;
      }
    }

    /// <summary>Field number for the "err" field.</summary>
    public const int ErrFieldNumber = 2;
    private global::Msg.Error err_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Msg.Error Err {
      get { return err_; }
      set {
        err_ = value;
      }
    }

    /// <summary>Field number for the "skills" field.</summary>
    public const int SkillsFieldNumber = 3;
    private static readonly pb::FieldCodec<global::Msg.Skill> _repeated_skills_codec
        = pb::FieldCodec.ForMessage(26, global::Msg.Skill.Parser);
    private readonly pbc::RepeatedField<global::Msg.Skill> skills_ = new pbc::RepeatedField<global::Msg.Skill>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<global::Msg.Skill> Skills {
      get { return skills_; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as SkillResponse);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(SkillResponse other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Code != other.Code) return false;
      if (!object.Equals(Err, other.Err)) return false;
      if(!skills_.Equals(other.skills_)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Code != 0) hash ^= Code.GetHashCode();
      if (err_ != null) hash ^= Err.GetHashCode();
      hash ^= skills_.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (Code != 0) {
        output.WriteRawTag(8);
        output.WriteEnum((int) Code);
      }
      if (err_ != null) {
        output.WriteRawTag(18);
        output.WriteMessage(Err);
      }
      skills_.WriteTo(output, _repeated_skills_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Code != 0) {
        size += 1 + pb::CodedOutputStream.ComputeEnumSize((int) Code);
      }
      if (err_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Err);
      }
      size += skills_.CalculateSize(_repeated_skills_codec);
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(SkillResponse other) {
      if (other == null) {
        return;
      }
      if (other.Code != 0) {
        Code = other.Code;
      }
      if (other.err_ != null) {
        if (err_ == null) {
          err_ = new global::Msg.Error();
        }
        Err.MergeFrom(other.Err);
      }
      skills_.Add(other.skills_);
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 8: {
            code_ = (global::Msg.ResponseCode) input.ReadEnum();
            break;
          }
          case 18: {
            if (err_ == null) {
              err_ = new global::Msg.Error();
            }
            input.ReadMessage(err_);
            break;
          }
          case 26: {
            skills_.AddEntriesFrom(input, _repeated_skills_codec);
            break;
          }
        }
      }
    }

  }

  public sealed partial class SkillUpgradeRequest : pb::IMessage<SkillUpgradeRequest> {
    private static readonly pb::MessageParser<SkillUpgradeRequest> _parser = new pb::MessageParser<SkillUpgradeRequest>(() => new SkillUpgradeRequest());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<SkillUpgradeRequest> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Msg.SkillReflection.Descriptor.MessageTypes[2]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public SkillUpgradeRequest() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public SkillUpgradeRequest(SkillUpgradeRequest other) : this() {
      skillId_ = other.skillId_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public SkillUpgradeRequest Clone() {
      return new SkillUpgradeRequest(this);
    }

    /// <summary>Field number for the "skillId" field.</summary>
    public const int SkillIdFieldNumber = 1;
    private string skillId_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string SkillId {
      get { return skillId_; }
      set {
        skillId_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as SkillUpgradeRequest);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(SkillUpgradeRequest other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (SkillId != other.SkillId) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (SkillId.Length != 0) hash ^= SkillId.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (SkillId.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(SkillId);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (SkillId.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(SkillId);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(SkillUpgradeRequest other) {
      if (other == null) {
        return;
      }
      if (other.SkillId.Length != 0) {
        SkillId = other.SkillId;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            SkillId = input.ReadString();
            break;
          }
        }
      }
    }

  }

  public sealed partial class SkillUpgradeResponse : pb::IMessage<SkillUpgradeResponse> {
    private static readonly pb::MessageParser<SkillUpgradeResponse> _parser = new pb::MessageParser<SkillUpgradeResponse>(() => new SkillUpgradeResponse());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<SkillUpgradeResponse> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Msg.SkillReflection.Descriptor.MessageTypes[3]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public SkillUpgradeResponse() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public SkillUpgradeResponse(SkillUpgradeResponse other) : this() {
      code_ = other.code_;
      err_ = other.err_ != null ? other.err_.Clone() : null;
      skill_ = other.skill_ != null ? other.skill_.Clone() : null;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public SkillUpgradeResponse Clone() {
      return new SkillUpgradeResponse(this);
    }

    /// <summary>Field number for the "code" field.</summary>
    public const int CodeFieldNumber = 1;
    private global::Msg.ResponseCode code_ = 0;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Msg.ResponseCode Code {
      get { return code_; }
      set {
        code_ = value;
      }
    }

    /// <summary>Field number for the "err" field.</summary>
    public const int ErrFieldNumber = 2;
    private global::Msg.Error err_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Msg.Error Err {
      get { return err_; }
      set {
        err_ = value;
      }
    }

    /// <summary>Field number for the "skill" field.</summary>
    public const int SkillFieldNumber = 3;
    private global::Msg.Skill skill_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Msg.Skill Skill {
      get { return skill_; }
      set {
        skill_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as SkillUpgradeResponse);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(SkillUpgradeResponse other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Code != other.Code) return false;
      if (!object.Equals(Err, other.Err)) return false;
      if (!object.Equals(Skill, other.Skill)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Code != 0) hash ^= Code.GetHashCode();
      if (err_ != null) hash ^= Err.GetHashCode();
      if (skill_ != null) hash ^= Skill.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (Code != 0) {
        output.WriteRawTag(8);
        output.WriteEnum((int) Code);
      }
      if (err_ != null) {
        output.WriteRawTag(18);
        output.WriteMessage(Err);
      }
      if (skill_ != null) {
        output.WriteRawTag(26);
        output.WriteMessage(Skill);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Code != 0) {
        size += 1 + pb::CodedOutputStream.ComputeEnumSize((int) Code);
      }
      if (err_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Err);
      }
      if (skill_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Skill);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(SkillUpgradeResponse other) {
      if (other == null) {
        return;
      }
      if (other.Code != 0) {
        Code = other.Code;
      }
      if (other.err_ != null) {
        if (err_ == null) {
          err_ = new global::Msg.Error();
        }
        Err.MergeFrom(other.Err);
      }
      if (other.skill_ != null) {
        if (skill_ == null) {
          skill_ = new global::Msg.Skill();
        }
        Skill.MergeFrom(other.Skill);
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 8: {
            code_ = (global::Msg.ResponseCode) input.ReadEnum();
            break;
          }
          case 18: {
            if (err_ == null) {
              err_ = new global::Msg.Error();
            }
            input.ReadMessage(err_);
            break;
          }
          case 26: {
            if (skill_ == null) {
              skill_ = new global::Msg.Skill();
            }
            input.ReadMessage(skill_);
            break;
          }
        }
      }
    }

  }

  public sealed partial class Skill : pb::IMessage<Skill> {
    private static readonly pb::MessageParser<Skill> _parser = new pb::MessageParser<Skill>(() => new Skill());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<Skill> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Msg.SkillReflection.Descriptor.MessageTypes[4]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Skill() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Skill(Skill other) : this() {
      id_ = other.id_;
      name_ = other.name_;
      level_ = other.level_;
      type_ = other.type_;
      desc_ = other.desc_;
      isOpen_ = other.isOpen_;
      skillId_ = other.skillId_;
      heroId_ = other.heroId_;
      levelDesc_ = other.levelDesc_.Clone();
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Skill Clone() {
      return new Skill(this);
    }

    /// <summary>Field number for the "Id" field.</summary>
    public const int IdFieldNumber = 1;
    private int id_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int Id {
      get { return id_; }
      set {
        id_ = value;
      }
    }

    /// <summary>Field number for the "Name" field.</summary>
    public const int NameFieldNumber = 2;
    private string name_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Name {
      get { return name_; }
      set {
        name_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "Level" field.</summary>
    public const int LevelFieldNumber = 3;
    private int level_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int Level {
      get { return level_; }
      set {
        level_ = value;
      }
    }

    /// <summary>Field number for the "Type" field.</summary>
    public const int TypeFieldNumber = 4;
    private int type_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int Type {
      get { return type_; }
      set {
        type_ = value;
      }
    }

    /// <summary>Field number for the "Desc" field.</summary>
    public const int DescFieldNumber = 5;
    private string desc_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Desc {
      get { return desc_; }
      set {
        desc_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "IsOpen" field.</summary>
    public const int IsOpenFieldNumber = 6;
    private bool isOpen_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool IsOpen {
      get { return isOpen_; }
      set {
        isOpen_ = value;
      }
    }

    /// <summary>Field number for the "SkillId" field.</summary>
    public const int SkillIdFieldNumber = 7;
    private string skillId_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string SkillId {
      get { return skillId_; }
      set {
        skillId_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "HeroId" field.</summary>
    public const int HeroIdFieldNumber = 8;
    private string heroId_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string HeroId {
      get { return heroId_; }
      set {
        heroId_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "LevelDesc" field.</summary>
    public const int LevelDescFieldNumber = 9;
    private static readonly pb::FieldCodec<string> _repeated_levelDesc_codec
        = pb::FieldCodec.ForString(74);
    private readonly pbc::RepeatedField<string> levelDesc_ = new pbc::RepeatedField<string>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<string> LevelDesc {
      get { return levelDesc_; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as Skill);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(Skill other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Id != other.Id) return false;
      if (Name != other.Name) return false;
      if (Level != other.Level) return false;
      if (Type != other.Type) return false;
      if (Desc != other.Desc) return false;
      if (IsOpen != other.IsOpen) return false;
      if (SkillId != other.SkillId) return false;
      if (HeroId != other.HeroId) return false;
      if(!levelDesc_.Equals(other.levelDesc_)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Id != 0) hash ^= Id.GetHashCode();
      if (Name.Length != 0) hash ^= Name.GetHashCode();
      if (Level != 0) hash ^= Level.GetHashCode();
      if (Type != 0) hash ^= Type.GetHashCode();
      if (Desc.Length != 0) hash ^= Desc.GetHashCode();
      if (IsOpen != false) hash ^= IsOpen.GetHashCode();
      if (SkillId.Length != 0) hash ^= SkillId.GetHashCode();
      if (HeroId.Length != 0) hash ^= HeroId.GetHashCode();
      hash ^= levelDesc_.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (Id != 0) {
        output.WriteRawTag(8);
        output.WriteInt32(Id);
      }
      if (Name.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(Name);
      }
      if (Level != 0) {
        output.WriteRawTag(24);
        output.WriteInt32(Level);
      }
      if (Type != 0) {
        output.WriteRawTag(32);
        output.WriteInt32(Type);
      }
      if (Desc.Length != 0) {
        output.WriteRawTag(42);
        output.WriteString(Desc);
      }
      if (IsOpen != false) {
        output.WriteRawTag(48);
        output.WriteBool(IsOpen);
      }
      if (SkillId.Length != 0) {
        output.WriteRawTag(58);
        output.WriteString(SkillId);
      }
      if (HeroId.Length != 0) {
        output.WriteRawTag(66);
        output.WriteString(HeroId);
      }
      levelDesc_.WriteTo(output, _repeated_levelDesc_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Id != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(Id);
      }
      if (Name.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Name);
      }
      if (Level != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(Level);
      }
      if (Type != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(Type);
      }
      if (Desc.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Desc);
      }
      if (IsOpen != false) {
        size += 1 + 1;
      }
      if (SkillId.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(SkillId);
      }
      if (HeroId.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(HeroId);
      }
      size += levelDesc_.CalculateSize(_repeated_levelDesc_codec);
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(Skill other) {
      if (other == null) {
        return;
      }
      if (other.Id != 0) {
        Id = other.Id;
      }
      if (other.Name.Length != 0) {
        Name = other.Name;
      }
      if (other.Level != 0) {
        Level = other.Level;
      }
      if (other.Type != 0) {
        Type = other.Type;
      }
      if (other.Desc.Length != 0) {
        Desc = other.Desc;
      }
      if (other.IsOpen != false) {
        IsOpen = other.IsOpen;
      }
      if (other.SkillId.Length != 0) {
        SkillId = other.SkillId;
      }
      if (other.HeroId.Length != 0) {
        HeroId = other.HeroId;
      }
      levelDesc_.Add(other.levelDesc_);
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 8: {
            Id = input.ReadInt32();
            break;
          }
          case 18: {
            Name = input.ReadString();
            break;
          }
          case 24: {
            Level = input.ReadInt32();
            break;
          }
          case 32: {
            Type = input.ReadInt32();
            break;
          }
          case 42: {
            Desc = input.ReadString();
            break;
          }
          case 48: {
            IsOpen = input.ReadBool();
            break;
          }
          case 58: {
            SkillId = input.ReadString();
            break;
          }
          case 66: {
            HeroId = input.ReadString();
            break;
          }
          case 74: {
            levelDesc_.AddEntriesFrom(input, _repeated_levelDesc_codec);
            break;
          }
        }
      }
    }

  }

  #endregion

}

#endregion Designer generated code
