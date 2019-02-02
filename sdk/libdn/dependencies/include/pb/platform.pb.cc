// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: pb/platform.proto

#include "pb/platform.pb.h"

#include <algorithm>

#include <google/protobuf/stubs/common.h>
#include <google/protobuf/stubs/port.h>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/wire_format_lite_inl.h>
#include <google/protobuf/descriptor.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/reflection_ops.h>
#include <google/protobuf/wire_format.h>
// This is a temporary google only hack
#ifdef GOOGLE_PROTOBUF_ENFORCE_UNIQUENESS
#include "third_party/protobuf/version.h"
#endif
// @@protoc_insertion_point(includes)

namespace pb {
class emptyDefaultTypeInternal {
 public:
  ::google::protobuf::internal::ExplicitlyConstructed<empty>
      _instance;
} _empty_default_instance_;
class RPCHealthRequestDefaultTypeInternal {
 public:
  ::google::protobuf::internal::ExplicitlyConstructed<RPCHealthRequest>
      _instance;
} _RPCHealthRequest_default_instance_;
class RPCHealthReplyDefaultTypeInternal {
 public:
  ::google::protobuf::internal::ExplicitlyConstructed<RPCHealthReply>
      _instance;
} _RPCHealthReply_default_instance_;
}  // namespace pb
namespace protobuf_pb_2fplatform_2eproto {
static void InitDefaultsempty() {
  GOOGLE_PROTOBUF_VERIFY_VERSION;

  {
    void* ptr = &::pb::_empty_default_instance_;
    new (ptr) ::pb::empty();
    ::google::protobuf::internal::OnShutdownDestroyMessage(ptr);
  }
  ::pb::empty::InitAsDefaultInstance();
}

::google::protobuf::internal::SCCInfo<0> scc_info_empty =
    {{ATOMIC_VAR_INIT(::google::protobuf::internal::SCCInfoBase::kUninitialized), 0, InitDefaultsempty}, {}};

static void InitDefaultsRPCHealthRequest() {
  GOOGLE_PROTOBUF_VERIFY_VERSION;

  {
    void* ptr = &::pb::_RPCHealthRequest_default_instance_;
    new (ptr) ::pb::RPCHealthRequest();
    ::google::protobuf::internal::OnShutdownDestroyMessage(ptr);
  }
  ::pb::RPCHealthRequest::InitAsDefaultInstance();
}

::google::protobuf::internal::SCCInfo<0> scc_info_RPCHealthRequest =
    {{ATOMIC_VAR_INIT(::google::protobuf::internal::SCCInfoBase::kUninitialized), 0, InitDefaultsRPCHealthRequest}, {}};

static void InitDefaultsRPCHealthReply() {
  GOOGLE_PROTOBUF_VERIFY_VERSION;

  {
    void* ptr = &::pb::_RPCHealthReply_default_instance_;
    new (ptr) ::pb::RPCHealthReply();
    ::google::protobuf::internal::OnShutdownDestroyMessage(ptr);
  }
  ::pb::RPCHealthReply::InitAsDefaultInstance();
}

::google::protobuf::internal::SCCInfo<0> scc_info_RPCHealthReply =
    {{ATOMIC_VAR_INIT(::google::protobuf::internal::SCCInfoBase::kUninitialized), 0, InitDefaultsRPCHealthReply}, {}};

void InitDefaults() {
  ::google::protobuf::internal::InitSCC(&scc_info_empty.base);
  ::google::protobuf::internal::InitSCC(&scc_info_RPCHealthRequest.base);
  ::google::protobuf::internal::InitSCC(&scc_info_RPCHealthReply.base);
}

::google::protobuf::Metadata file_level_metadata[3];

const ::google::protobuf::uint32 TableStruct::offsets[] GOOGLE_PROTOBUF_ATTRIBUTE_SECTION_VARIABLE(protodesc_cold) = {
  ~0u,  // no _has_bits_
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::pb::empty, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _has_bits_
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::pb::RPCHealthRequest, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::pb::RPCHealthRequest, waitforminconnections_),
  ~0u,  // no _has_bits_
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::pb::RPCHealthReply, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::pb::RPCHealthReply, ready_),
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::pb::RPCHealthReply, message_),
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::pb::RPCHealthReply, basepath_),
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::pb::RPCHealthReply, numconnections_),
};
static const ::google::protobuf::internal::MigrationSchema schemas[] GOOGLE_PROTOBUF_ATTRIBUTE_SECTION_VARIABLE(protodesc_cold) = {
  { 0, -1, sizeof(::pb::empty)},
  { 5, -1, sizeof(::pb::RPCHealthRequest)},
  { 11, -1, sizeof(::pb::RPCHealthReply)},
};

static ::google::protobuf::Message const * const file_default_instances[] = {
  reinterpret_cast<const ::google::protobuf::Message*>(&::pb::_empty_default_instance_),
  reinterpret_cast<const ::google::protobuf::Message*>(&::pb::_RPCHealthRequest_default_instance_),
  reinterpret_cast<const ::google::protobuf::Message*>(&::pb::_RPCHealthReply_default_instance_),
};

void protobuf_AssignDescriptors() {
  AddDescriptors();
  AssignDescriptors(
      "pb/platform.proto", schemas, file_default_instances, TableStruct::offsets,
      file_level_metadata, NULL, NULL);
}

void protobuf_AssignDescriptorsOnce() {
  static ::google::protobuf::internal::once_flag once;
  ::google::protobuf::internal::call_once(once, protobuf_AssignDescriptors);
}

void protobuf_RegisterTypes(const ::std::string&) GOOGLE_PROTOBUF_ATTRIBUTE_COLD;
void protobuf_RegisterTypes(const ::std::string&) {
  protobuf_AssignDescriptorsOnce();
  ::google::protobuf::internal::RegisterAllTypes(file_level_metadata, 3);
}

void AddDescriptorsImpl() {
  InitDefaults();
  static const char descriptor[] GOOGLE_PROTOBUF_ATTRIBUTE_SECTION_VARIABLE(protodesc_cold) = {
      "\n\021pb/platform.proto\022\002pb\032\024pb/matchmaking."
      "proto\032\024pb/addressbook.proto\032\020pb/storage."
      "proto\032\022pb/messaging.proto\032\022pb/publisher."
      "proto\032\034google/api/annotations.proto\"\007\n\005e"
      "mpty\"1\n\020RPCHealthRequest\022\035\n\025waitForMinCo"
      "nnections\030\001 \001(\010\"Z\n\016RPCHealthReply\022\r\n\005rea"
      "dy\030\001 \001(\010\022\017\n\007message\030\002 \001(\t\022\020\n\010basePath\030\003 "
      "\001(\t\022\026\n\016numConnections\030\004 \001(\r2\252\014\n\rDecentra"
      "lizer\022a\n\tGetHealth\022\024.pb.RPCHealthRequest"
      "\032\022.pb.RPCHealthReply\"*\202\323\344\223\002$\022\"/v1/health"
      "/{waitForMinConnections}\022_\n\rUpsertSessio"
      "n\022\033.pb.RPCUpsertSessionRequest\032\034.pb.RPCU"
      "psertSessionResponse\"\023\202\323\344\223\002\r\"\013/v1/sessio"
      "n\022L\n\rDeleteSession\022\033.pb.RPCDeleteSession"
      "Request\032\034.pb.RPCDeleteSessionResponse\"\000\022"
      "\207\001\n\026GetSessionIdsByDetails\022$.pb.RPCGetSe"
      "ssionIdsByDetailsRequest\032\034.pb.RPCGetSess"
      "ionIdsResponse\")\202\323\344\223\002#\022!/v1/sessions/{ty"
      "pe}/{key}/{value}\022^\n\026GetSessionIdsByPeer"
      "Ids\022$.pb.RPCGetSessionIdsByPeerIdsReques"
      "t\032\034.pb.RPCGetSessionIdsResponse\"\000\022b\n\nGet"
      "Session\022\030.pb.RPCGetSessionRequest\032\031.pb.R"
      "PCGetSessionResponse\"\037\202\323\344\223\002\031\022\027/v1/sessio"
      "n/{sessionId}\022C\n\nUpsertPeer\022\030.pb.RPCUpse"
      "rtPeerRequest\032\031.pb.RPCUpsertPeerResponse"
      "\"\000\022b\n\nGetPeerIds\022\030.pb.RPCGetPeerIdsReque"
      "st\032\031.pb.RPCGetPeerIdsResponse\"\037\202\323\344\223\002\031\022\027/"
      "v1/peers/{key}/{value}\022W\n\007GetPeer\022\025.pb.R"
      "PCGetPeerRequest\032\026.pb.RPCGetPeerResponse"
      "\"\035\202\323\344\223\002\027\022\025/v1/peer/{pId}/{dnId}\022L\n\rWrite"
      "PeerFile\022\033.pb.RPCWritePeerFileRequest\032\034."
      "pb.RPCWritePeerFileResponse\"\000\022F\n\013GetPeer"
      "File\022\031.pb.RPCGetPeerFileRequest\032\032.pb.RPC"
      "GetPeerFileResponse\"\000\0226\n\021SendDirectMessa"
      "ge\022\024.pb.RPCDirectMessage\032\t.pb.empty\"\000\022T\n"
      "\024ReceiveDirectMessage\022\".pb.RPCReceiveDir"
      "ectMessageRequest\032\024.pb.RPCDirectMessage\""
      "\0000\001\022J\n\027readPublisherDefinition\022\".pb.load"
      "PublisherDefinitionRequest\032\t.pb.empty\"\000\022"
      "g\n\026publishPublisherUpdate\022$.pb.RPCPublis"
      "hPublisherUpdateRequest\032%.pb.RPCPublishP"
      "ublisherUpdateResponse\"\000\022k\n\026GetPublisher"
      "Definition\022!.pb.GetPublisherDefinitionRe"
      "quest\032\027.pb.PublisherDefinition\"\025\202\323\344\223\002\017\022\r"
      "/v1/publisher\022q\n\020GetPublisherFile\022\036.pb.R"
      "PCGetPublisherFileRequest\032\037.pb.RPCGetPub"
      "lisherFileResponse\"\034\202\323\344\223\002\026\022\024/v1/publishe"
      "r/{name}b\006proto3"
  };
  ::google::protobuf::DescriptorPool::InternalAddGeneratedFile(
      descriptor, 1896);
  ::google::protobuf::MessageFactory::InternalRegisterGeneratedFile(
    "pb/platform.proto", &protobuf_RegisterTypes);
  ::protobuf_pb_2fmatchmaking_2eproto::AddDescriptors();
  ::protobuf_pb_2faddressbook_2eproto::AddDescriptors();
  ::protobuf_pb_2fstorage_2eproto::AddDescriptors();
  ::protobuf_pb_2fmessaging_2eproto::AddDescriptors();
  ::protobuf_pb_2fpublisher_2eproto::AddDescriptors();
  ::protobuf_google_2fapi_2fannotations_2eproto::AddDescriptors();
}

void AddDescriptors() {
  static ::google::protobuf::internal::once_flag once;
  ::google::protobuf::internal::call_once(once, AddDescriptorsImpl);
}
// Force AddDescriptors() to be called at dynamic initialization time.
struct StaticDescriptorInitializer {
  StaticDescriptorInitializer() {
    AddDescriptors();
  }
} static_descriptor_initializer;
}  // namespace protobuf_pb_2fplatform_2eproto
namespace pb {

// ===================================================================

void empty::InitAsDefaultInstance() {
}
#if !defined(_MSC_VER) || _MSC_VER >= 1900
#endif  // !defined(_MSC_VER) || _MSC_VER >= 1900

empty::empty()
  : ::google::protobuf::Message(), _internal_metadata_(NULL) {
  ::google::protobuf::internal::InitSCC(
      &protobuf_pb_2fplatform_2eproto::scc_info_empty.base);
  SharedCtor();
  // @@protoc_insertion_point(constructor:pb.empty)
}
empty::empty(const empty& from)
  : ::google::protobuf::Message(),
      _internal_metadata_(NULL) {
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  // @@protoc_insertion_point(copy_constructor:pb.empty)
}

void empty::SharedCtor() {
}

empty::~empty() {
  // @@protoc_insertion_point(destructor:pb.empty)
  SharedDtor();
}

void empty::SharedDtor() {
}

void empty::SetCachedSize(int size) const {
  _cached_size_.Set(size);
}
const ::google::protobuf::Descriptor* empty::descriptor() {
  ::protobuf_pb_2fplatform_2eproto::protobuf_AssignDescriptorsOnce();
  return ::protobuf_pb_2fplatform_2eproto::file_level_metadata[kIndexInFileMessages].descriptor;
}

const empty& empty::default_instance() {
  ::google::protobuf::internal::InitSCC(&protobuf_pb_2fplatform_2eproto::scc_info_empty.base);
  return *internal_default_instance();
}


void empty::Clear() {
// @@protoc_insertion_point(message_clear_start:pb.empty)
  ::google::protobuf::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _internal_metadata_.Clear();
}

bool empty::MergePartialFromCodedStream(
    ::google::protobuf::io::CodedInputStream* input) {
#define DO_(EXPRESSION) if (!GOOGLE_PREDICT_TRUE(EXPRESSION)) goto failure
  ::google::protobuf::uint32 tag;
  // @@protoc_insertion_point(parse_start:pb.empty)
  for (;;) {
    ::std::pair<::google::protobuf::uint32, bool> p = input->ReadTagWithCutoffNoLastTag(127u);
    tag = p.first;
    if (!p.second) goto handle_unusual;
  handle_unusual:
    if (tag == 0) {
      goto success;
    }
    DO_(::google::protobuf::internal::WireFormat::SkipField(
          input, tag, _internal_metadata_.mutable_unknown_fields()));
  }
success:
  // @@protoc_insertion_point(parse_success:pb.empty)
  return true;
failure:
  // @@protoc_insertion_point(parse_failure:pb.empty)
  return false;
#undef DO_
}

void empty::SerializeWithCachedSizes(
    ::google::protobuf::io::CodedOutputStream* output) const {
  // @@protoc_insertion_point(serialize_start:pb.empty)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  if ((_internal_metadata_.have_unknown_fields() &&  ::google::protobuf::internal::GetProto3PreserveUnknownsDefault())) {
    ::google::protobuf::internal::WireFormat::SerializeUnknownFields(
        (::google::protobuf::internal::GetProto3PreserveUnknownsDefault()   ? _internal_metadata_.unknown_fields()   : _internal_metadata_.default_instance()), output);
  }
  // @@protoc_insertion_point(serialize_end:pb.empty)
}

::google::protobuf::uint8* empty::InternalSerializeWithCachedSizesToArray(
    bool deterministic, ::google::protobuf::uint8* target) const {
  (void)deterministic; // Unused
  // @@protoc_insertion_point(serialize_to_array_start:pb.empty)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  if ((_internal_metadata_.have_unknown_fields() &&  ::google::protobuf::internal::GetProto3PreserveUnknownsDefault())) {
    target = ::google::protobuf::internal::WireFormat::SerializeUnknownFieldsToArray(
        (::google::protobuf::internal::GetProto3PreserveUnknownsDefault()   ? _internal_metadata_.unknown_fields()   : _internal_metadata_.default_instance()), target);
  }
  // @@protoc_insertion_point(serialize_to_array_end:pb.empty)
  return target;
}

size_t empty::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:pb.empty)
  size_t total_size = 0;

  if ((_internal_metadata_.have_unknown_fields() &&  ::google::protobuf::internal::GetProto3PreserveUnknownsDefault())) {
    total_size +=
      ::google::protobuf::internal::WireFormat::ComputeUnknownFieldsSize(
        (::google::protobuf::internal::GetProto3PreserveUnknownsDefault()   ? _internal_metadata_.unknown_fields()   : _internal_metadata_.default_instance()));
  }
  int cached_size = ::google::protobuf::internal::ToCachedSize(total_size);
  SetCachedSize(cached_size);
  return total_size;
}

void empty::MergeFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_merge_from_start:pb.empty)
  GOOGLE_DCHECK_NE(&from, this);
  const empty* source =
      ::google::protobuf::internal::DynamicCastToGenerated<const empty>(
          &from);
  if (source == NULL) {
  // @@protoc_insertion_point(generalized_merge_from_cast_fail:pb.empty)
    ::google::protobuf::internal::ReflectionOps::Merge(from, this);
  } else {
  // @@protoc_insertion_point(generalized_merge_from_cast_success:pb.empty)
    MergeFrom(*source);
  }
}

void empty::MergeFrom(const empty& from) {
// @@protoc_insertion_point(class_specific_merge_from_start:pb.empty)
  GOOGLE_DCHECK_NE(&from, this);
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

}

void empty::CopyFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_copy_from_start:pb.empty)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

void empty::CopyFrom(const empty& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:pb.empty)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool empty::IsInitialized() const {
  return true;
}

void empty::Swap(empty* other) {
  if (other == this) return;
  InternalSwap(other);
}
void empty::InternalSwap(empty* other) {
  using std::swap;
  _internal_metadata_.Swap(&other->_internal_metadata_);
}

::google::protobuf::Metadata empty::GetMetadata() const {
  protobuf_pb_2fplatform_2eproto::protobuf_AssignDescriptorsOnce();
  return ::protobuf_pb_2fplatform_2eproto::file_level_metadata[kIndexInFileMessages];
}


// ===================================================================

void RPCHealthRequest::InitAsDefaultInstance() {
}
#if !defined(_MSC_VER) || _MSC_VER >= 1900
const int RPCHealthRequest::kWaitForMinConnectionsFieldNumber;
#endif  // !defined(_MSC_VER) || _MSC_VER >= 1900

RPCHealthRequest::RPCHealthRequest()
  : ::google::protobuf::Message(), _internal_metadata_(NULL) {
  ::google::protobuf::internal::InitSCC(
      &protobuf_pb_2fplatform_2eproto::scc_info_RPCHealthRequest.base);
  SharedCtor();
  // @@protoc_insertion_point(constructor:pb.RPCHealthRequest)
}
RPCHealthRequest::RPCHealthRequest(const RPCHealthRequest& from)
  : ::google::protobuf::Message(),
      _internal_metadata_(NULL) {
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  waitforminconnections_ = from.waitforminconnections_;
  // @@protoc_insertion_point(copy_constructor:pb.RPCHealthRequest)
}

void RPCHealthRequest::SharedCtor() {
  waitforminconnections_ = false;
}

RPCHealthRequest::~RPCHealthRequest() {
  // @@protoc_insertion_point(destructor:pb.RPCHealthRequest)
  SharedDtor();
}

void RPCHealthRequest::SharedDtor() {
}

void RPCHealthRequest::SetCachedSize(int size) const {
  _cached_size_.Set(size);
}
const ::google::protobuf::Descriptor* RPCHealthRequest::descriptor() {
  ::protobuf_pb_2fplatform_2eproto::protobuf_AssignDescriptorsOnce();
  return ::protobuf_pb_2fplatform_2eproto::file_level_metadata[kIndexInFileMessages].descriptor;
}

const RPCHealthRequest& RPCHealthRequest::default_instance() {
  ::google::protobuf::internal::InitSCC(&protobuf_pb_2fplatform_2eproto::scc_info_RPCHealthRequest.base);
  return *internal_default_instance();
}


void RPCHealthRequest::Clear() {
// @@protoc_insertion_point(message_clear_start:pb.RPCHealthRequest)
  ::google::protobuf::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  waitforminconnections_ = false;
  _internal_metadata_.Clear();
}

bool RPCHealthRequest::MergePartialFromCodedStream(
    ::google::protobuf::io::CodedInputStream* input) {
#define DO_(EXPRESSION) if (!GOOGLE_PREDICT_TRUE(EXPRESSION)) goto failure
  ::google::protobuf::uint32 tag;
  // @@protoc_insertion_point(parse_start:pb.RPCHealthRequest)
  for (;;) {
    ::std::pair<::google::protobuf::uint32, bool> p = input->ReadTagWithCutoffNoLastTag(127u);
    tag = p.first;
    if (!p.second) goto handle_unusual;
    switch (::google::protobuf::internal::WireFormatLite::GetTagFieldNumber(tag)) {
      // bool waitForMinConnections = 1;
      case 1: {
        if (static_cast< ::google::protobuf::uint8>(tag) ==
            static_cast< ::google::protobuf::uint8>(8u /* 8 & 0xFF */)) {

          DO_((::google::protobuf::internal::WireFormatLite::ReadPrimitive<
                   bool, ::google::protobuf::internal::WireFormatLite::TYPE_BOOL>(
                 input, &waitforminconnections_)));
        } else {
          goto handle_unusual;
        }
        break;
      }

      default: {
      handle_unusual:
        if (tag == 0) {
          goto success;
        }
        DO_(::google::protobuf::internal::WireFormat::SkipField(
              input, tag, _internal_metadata_.mutable_unknown_fields()));
        break;
      }
    }
  }
success:
  // @@protoc_insertion_point(parse_success:pb.RPCHealthRequest)
  return true;
failure:
  // @@protoc_insertion_point(parse_failure:pb.RPCHealthRequest)
  return false;
#undef DO_
}

void RPCHealthRequest::SerializeWithCachedSizes(
    ::google::protobuf::io::CodedOutputStream* output) const {
  // @@protoc_insertion_point(serialize_start:pb.RPCHealthRequest)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // bool waitForMinConnections = 1;
  if (this->waitforminconnections() != 0) {
    ::google::protobuf::internal::WireFormatLite::WriteBool(1, this->waitforminconnections(), output);
  }

  if ((_internal_metadata_.have_unknown_fields() &&  ::google::protobuf::internal::GetProto3PreserveUnknownsDefault())) {
    ::google::protobuf::internal::WireFormat::SerializeUnknownFields(
        (::google::protobuf::internal::GetProto3PreserveUnknownsDefault()   ? _internal_metadata_.unknown_fields()   : _internal_metadata_.default_instance()), output);
  }
  // @@protoc_insertion_point(serialize_end:pb.RPCHealthRequest)
}

::google::protobuf::uint8* RPCHealthRequest::InternalSerializeWithCachedSizesToArray(
    bool deterministic, ::google::protobuf::uint8* target) const {
  (void)deterministic; // Unused
  // @@protoc_insertion_point(serialize_to_array_start:pb.RPCHealthRequest)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // bool waitForMinConnections = 1;
  if (this->waitforminconnections() != 0) {
    target = ::google::protobuf::internal::WireFormatLite::WriteBoolToArray(1, this->waitforminconnections(), target);
  }

  if ((_internal_metadata_.have_unknown_fields() &&  ::google::protobuf::internal::GetProto3PreserveUnknownsDefault())) {
    target = ::google::protobuf::internal::WireFormat::SerializeUnknownFieldsToArray(
        (::google::protobuf::internal::GetProto3PreserveUnknownsDefault()   ? _internal_metadata_.unknown_fields()   : _internal_metadata_.default_instance()), target);
  }
  // @@protoc_insertion_point(serialize_to_array_end:pb.RPCHealthRequest)
  return target;
}

size_t RPCHealthRequest::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:pb.RPCHealthRequest)
  size_t total_size = 0;

  if ((_internal_metadata_.have_unknown_fields() &&  ::google::protobuf::internal::GetProto3PreserveUnknownsDefault())) {
    total_size +=
      ::google::protobuf::internal::WireFormat::ComputeUnknownFieldsSize(
        (::google::protobuf::internal::GetProto3PreserveUnknownsDefault()   ? _internal_metadata_.unknown_fields()   : _internal_metadata_.default_instance()));
  }
  // bool waitForMinConnections = 1;
  if (this->waitforminconnections() != 0) {
    total_size += 1 + 1;
  }

  int cached_size = ::google::protobuf::internal::ToCachedSize(total_size);
  SetCachedSize(cached_size);
  return total_size;
}

void RPCHealthRequest::MergeFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_merge_from_start:pb.RPCHealthRequest)
  GOOGLE_DCHECK_NE(&from, this);
  const RPCHealthRequest* source =
      ::google::protobuf::internal::DynamicCastToGenerated<const RPCHealthRequest>(
          &from);
  if (source == NULL) {
  // @@protoc_insertion_point(generalized_merge_from_cast_fail:pb.RPCHealthRequest)
    ::google::protobuf::internal::ReflectionOps::Merge(from, this);
  } else {
  // @@protoc_insertion_point(generalized_merge_from_cast_success:pb.RPCHealthRequest)
    MergeFrom(*source);
  }
}

void RPCHealthRequest::MergeFrom(const RPCHealthRequest& from) {
// @@protoc_insertion_point(class_specific_merge_from_start:pb.RPCHealthRequest)
  GOOGLE_DCHECK_NE(&from, this);
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  if (from.waitforminconnections() != 0) {
    set_waitforminconnections(from.waitforminconnections());
  }
}

void RPCHealthRequest::CopyFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_copy_from_start:pb.RPCHealthRequest)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

void RPCHealthRequest::CopyFrom(const RPCHealthRequest& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:pb.RPCHealthRequest)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool RPCHealthRequest::IsInitialized() const {
  return true;
}

void RPCHealthRequest::Swap(RPCHealthRequest* other) {
  if (other == this) return;
  InternalSwap(other);
}
void RPCHealthRequest::InternalSwap(RPCHealthRequest* other) {
  using std::swap;
  swap(waitforminconnections_, other->waitforminconnections_);
  _internal_metadata_.Swap(&other->_internal_metadata_);
}

::google::protobuf::Metadata RPCHealthRequest::GetMetadata() const {
  protobuf_pb_2fplatform_2eproto::protobuf_AssignDescriptorsOnce();
  return ::protobuf_pb_2fplatform_2eproto::file_level_metadata[kIndexInFileMessages];
}


// ===================================================================

void RPCHealthReply::InitAsDefaultInstance() {
}
#if !defined(_MSC_VER) || _MSC_VER >= 1900
const int RPCHealthReply::kReadyFieldNumber;
const int RPCHealthReply::kMessageFieldNumber;
const int RPCHealthReply::kBasePathFieldNumber;
const int RPCHealthReply::kNumConnectionsFieldNumber;
#endif  // !defined(_MSC_VER) || _MSC_VER >= 1900

RPCHealthReply::RPCHealthReply()
  : ::google::protobuf::Message(), _internal_metadata_(NULL) {
  ::google::protobuf::internal::InitSCC(
      &protobuf_pb_2fplatform_2eproto::scc_info_RPCHealthReply.base);
  SharedCtor();
  // @@protoc_insertion_point(constructor:pb.RPCHealthReply)
}
RPCHealthReply::RPCHealthReply(const RPCHealthReply& from)
  : ::google::protobuf::Message(),
      _internal_metadata_(NULL) {
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  message_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  if (from.message().size() > 0) {
    message_.AssignWithDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), from.message_);
  }
  basepath_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  if (from.basepath().size() > 0) {
    basepath_.AssignWithDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), from.basepath_);
  }
  ::memcpy(&ready_, &from.ready_,
    static_cast<size_t>(reinterpret_cast<char*>(&numconnections_) -
    reinterpret_cast<char*>(&ready_)) + sizeof(numconnections_));
  // @@protoc_insertion_point(copy_constructor:pb.RPCHealthReply)
}

void RPCHealthReply::SharedCtor() {
  message_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  basepath_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  ::memset(&ready_, 0, static_cast<size_t>(
      reinterpret_cast<char*>(&numconnections_) -
      reinterpret_cast<char*>(&ready_)) + sizeof(numconnections_));
}

RPCHealthReply::~RPCHealthReply() {
  // @@protoc_insertion_point(destructor:pb.RPCHealthReply)
  SharedDtor();
}

void RPCHealthReply::SharedDtor() {
  message_.DestroyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  basepath_.DestroyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}

void RPCHealthReply::SetCachedSize(int size) const {
  _cached_size_.Set(size);
}
const ::google::protobuf::Descriptor* RPCHealthReply::descriptor() {
  ::protobuf_pb_2fplatform_2eproto::protobuf_AssignDescriptorsOnce();
  return ::protobuf_pb_2fplatform_2eproto::file_level_metadata[kIndexInFileMessages].descriptor;
}

const RPCHealthReply& RPCHealthReply::default_instance() {
  ::google::protobuf::internal::InitSCC(&protobuf_pb_2fplatform_2eproto::scc_info_RPCHealthReply.base);
  return *internal_default_instance();
}


void RPCHealthReply::Clear() {
// @@protoc_insertion_point(message_clear_start:pb.RPCHealthReply)
  ::google::protobuf::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  message_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  basepath_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  ::memset(&ready_, 0, static_cast<size_t>(
      reinterpret_cast<char*>(&numconnections_) -
      reinterpret_cast<char*>(&ready_)) + sizeof(numconnections_));
  _internal_metadata_.Clear();
}

bool RPCHealthReply::MergePartialFromCodedStream(
    ::google::protobuf::io::CodedInputStream* input) {
#define DO_(EXPRESSION) if (!GOOGLE_PREDICT_TRUE(EXPRESSION)) goto failure
  ::google::protobuf::uint32 tag;
  // @@protoc_insertion_point(parse_start:pb.RPCHealthReply)
  for (;;) {
    ::std::pair<::google::protobuf::uint32, bool> p = input->ReadTagWithCutoffNoLastTag(127u);
    tag = p.first;
    if (!p.second) goto handle_unusual;
    switch (::google::protobuf::internal::WireFormatLite::GetTagFieldNumber(tag)) {
      // bool ready = 1;
      case 1: {
        if (static_cast< ::google::protobuf::uint8>(tag) ==
            static_cast< ::google::protobuf::uint8>(8u /* 8 & 0xFF */)) {

          DO_((::google::protobuf::internal::WireFormatLite::ReadPrimitive<
                   bool, ::google::protobuf::internal::WireFormatLite::TYPE_BOOL>(
                 input, &ready_)));
        } else {
          goto handle_unusual;
        }
        break;
      }

      // string message = 2;
      case 2: {
        if (static_cast< ::google::protobuf::uint8>(tag) ==
            static_cast< ::google::protobuf::uint8>(18u /* 18 & 0xFF */)) {
          DO_(::google::protobuf::internal::WireFormatLite::ReadString(
                input, this->mutable_message()));
          DO_(::google::protobuf::internal::WireFormatLite::VerifyUtf8String(
            this->message().data(), static_cast<int>(this->message().length()),
            ::google::protobuf::internal::WireFormatLite::PARSE,
            "pb.RPCHealthReply.message"));
        } else {
          goto handle_unusual;
        }
        break;
      }

      // string basePath = 3;
      case 3: {
        if (static_cast< ::google::protobuf::uint8>(tag) ==
            static_cast< ::google::protobuf::uint8>(26u /* 26 & 0xFF */)) {
          DO_(::google::protobuf::internal::WireFormatLite::ReadString(
                input, this->mutable_basepath()));
          DO_(::google::protobuf::internal::WireFormatLite::VerifyUtf8String(
            this->basepath().data(), static_cast<int>(this->basepath().length()),
            ::google::protobuf::internal::WireFormatLite::PARSE,
            "pb.RPCHealthReply.basePath"));
        } else {
          goto handle_unusual;
        }
        break;
      }

      // uint32 numConnections = 4;
      case 4: {
        if (static_cast< ::google::protobuf::uint8>(tag) ==
            static_cast< ::google::protobuf::uint8>(32u /* 32 & 0xFF */)) {

          DO_((::google::protobuf::internal::WireFormatLite::ReadPrimitive<
                   ::google::protobuf::uint32, ::google::protobuf::internal::WireFormatLite::TYPE_UINT32>(
                 input, &numconnections_)));
        } else {
          goto handle_unusual;
        }
        break;
      }

      default: {
      handle_unusual:
        if (tag == 0) {
          goto success;
        }
        DO_(::google::protobuf::internal::WireFormat::SkipField(
              input, tag, _internal_metadata_.mutable_unknown_fields()));
        break;
      }
    }
  }
success:
  // @@protoc_insertion_point(parse_success:pb.RPCHealthReply)
  return true;
failure:
  // @@protoc_insertion_point(parse_failure:pb.RPCHealthReply)
  return false;
#undef DO_
}

void RPCHealthReply::SerializeWithCachedSizes(
    ::google::protobuf::io::CodedOutputStream* output) const {
  // @@protoc_insertion_point(serialize_start:pb.RPCHealthReply)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // bool ready = 1;
  if (this->ready() != 0) {
    ::google::protobuf::internal::WireFormatLite::WriteBool(1, this->ready(), output);
  }

  // string message = 2;
  if (this->message().size() > 0) {
    ::google::protobuf::internal::WireFormatLite::VerifyUtf8String(
      this->message().data(), static_cast<int>(this->message().length()),
      ::google::protobuf::internal::WireFormatLite::SERIALIZE,
      "pb.RPCHealthReply.message");
    ::google::protobuf::internal::WireFormatLite::WriteStringMaybeAliased(
      2, this->message(), output);
  }

  // string basePath = 3;
  if (this->basepath().size() > 0) {
    ::google::protobuf::internal::WireFormatLite::VerifyUtf8String(
      this->basepath().data(), static_cast<int>(this->basepath().length()),
      ::google::protobuf::internal::WireFormatLite::SERIALIZE,
      "pb.RPCHealthReply.basePath");
    ::google::protobuf::internal::WireFormatLite::WriteStringMaybeAliased(
      3, this->basepath(), output);
  }

  // uint32 numConnections = 4;
  if (this->numconnections() != 0) {
    ::google::protobuf::internal::WireFormatLite::WriteUInt32(4, this->numconnections(), output);
  }

  if ((_internal_metadata_.have_unknown_fields() &&  ::google::protobuf::internal::GetProto3PreserveUnknownsDefault())) {
    ::google::protobuf::internal::WireFormat::SerializeUnknownFields(
        (::google::protobuf::internal::GetProto3PreserveUnknownsDefault()   ? _internal_metadata_.unknown_fields()   : _internal_metadata_.default_instance()), output);
  }
  // @@protoc_insertion_point(serialize_end:pb.RPCHealthReply)
}

::google::protobuf::uint8* RPCHealthReply::InternalSerializeWithCachedSizesToArray(
    bool deterministic, ::google::protobuf::uint8* target) const {
  (void)deterministic; // Unused
  // @@protoc_insertion_point(serialize_to_array_start:pb.RPCHealthReply)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // bool ready = 1;
  if (this->ready() != 0) {
    target = ::google::protobuf::internal::WireFormatLite::WriteBoolToArray(1, this->ready(), target);
  }

  // string message = 2;
  if (this->message().size() > 0) {
    ::google::protobuf::internal::WireFormatLite::VerifyUtf8String(
      this->message().data(), static_cast<int>(this->message().length()),
      ::google::protobuf::internal::WireFormatLite::SERIALIZE,
      "pb.RPCHealthReply.message");
    target =
      ::google::protobuf::internal::WireFormatLite::WriteStringToArray(
        2, this->message(), target);
  }

  // string basePath = 3;
  if (this->basepath().size() > 0) {
    ::google::protobuf::internal::WireFormatLite::VerifyUtf8String(
      this->basepath().data(), static_cast<int>(this->basepath().length()),
      ::google::protobuf::internal::WireFormatLite::SERIALIZE,
      "pb.RPCHealthReply.basePath");
    target =
      ::google::protobuf::internal::WireFormatLite::WriteStringToArray(
        3, this->basepath(), target);
  }

  // uint32 numConnections = 4;
  if (this->numconnections() != 0) {
    target = ::google::protobuf::internal::WireFormatLite::WriteUInt32ToArray(4, this->numconnections(), target);
  }

  if ((_internal_metadata_.have_unknown_fields() &&  ::google::protobuf::internal::GetProto3PreserveUnknownsDefault())) {
    target = ::google::protobuf::internal::WireFormat::SerializeUnknownFieldsToArray(
        (::google::protobuf::internal::GetProto3PreserveUnknownsDefault()   ? _internal_metadata_.unknown_fields()   : _internal_metadata_.default_instance()), target);
  }
  // @@protoc_insertion_point(serialize_to_array_end:pb.RPCHealthReply)
  return target;
}

size_t RPCHealthReply::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:pb.RPCHealthReply)
  size_t total_size = 0;

  if ((_internal_metadata_.have_unknown_fields() &&  ::google::protobuf::internal::GetProto3PreserveUnknownsDefault())) {
    total_size +=
      ::google::protobuf::internal::WireFormat::ComputeUnknownFieldsSize(
        (::google::protobuf::internal::GetProto3PreserveUnknownsDefault()   ? _internal_metadata_.unknown_fields()   : _internal_metadata_.default_instance()));
  }
  // string message = 2;
  if (this->message().size() > 0) {
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::StringSize(
        this->message());
  }

  // string basePath = 3;
  if (this->basepath().size() > 0) {
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::StringSize(
        this->basepath());
  }

  // bool ready = 1;
  if (this->ready() != 0) {
    total_size += 1 + 1;
  }

  // uint32 numConnections = 4;
  if (this->numconnections() != 0) {
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::UInt32Size(
        this->numconnections());
  }

  int cached_size = ::google::protobuf::internal::ToCachedSize(total_size);
  SetCachedSize(cached_size);
  return total_size;
}

void RPCHealthReply::MergeFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_merge_from_start:pb.RPCHealthReply)
  GOOGLE_DCHECK_NE(&from, this);
  const RPCHealthReply* source =
      ::google::protobuf::internal::DynamicCastToGenerated<const RPCHealthReply>(
          &from);
  if (source == NULL) {
  // @@protoc_insertion_point(generalized_merge_from_cast_fail:pb.RPCHealthReply)
    ::google::protobuf::internal::ReflectionOps::Merge(from, this);
  } else {
  // @@protoc_insertion_point(generalized_merge_from_cast_success:pb.RPCHealthReply)
    MergeFrom(*source);
  }
}

void RPCHealthReply::MergeFrom(const RPCHealthReply& from) {
// @@protoc_insertion_point(class_specific_merge_from_start:pb.RPCHealthReply)
  GOOGLE_DCHECK_NE(&from, this);
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  if (from.message().size() > 0) {

    message_.AssignWithDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), from.message_);
  }
  if (from.basepath().size() > 0) {

    basepath_.AssignWithDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), from.basepath_);
  }
  if (from.ready() != 0) {
    set_ready(from.ready());
  }
  if (from.numconnections() != 0) {
    set_numconnections(from.numconnections());
  }
}

void RPCHealthReply::CopyFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_copy_from_start:pb.RPCHealthReply)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

void RPCHealthReply::CopyFrom(const RPCHealthReply& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:pb.RPCHealthReply)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool RPCHealthReply::IsInitialized() const {
  return true;
}

void RPCHealthReply::Swap(RPCHealthReply* other) {
  if (other == this) return;
  InternalSwap(other);
}
void RPCHealthReply::InternalSwap(RPCHealthReply* other) {
  using std::swap;
  message_.Swap(&other->message_, &::google::protobuf::internal::GetEmptyStringAlreadyInited(),
    GetArenaNoVirtual());
  basepath_.Swap(&other->basepath_, &::google::protobuf::internal::GetEmptyStringAlreadyInited(),
    GetArenaNoVirtual());
  swap(ready_, other->ready_);
  swap(numconnections_, other->numconnections_);
  _internal_metadata_.Swap(&other->_internal_metadata_);
}

::google::protobuf::Metadata RPCHealthReply::GetMetadata() const {
  protobuf_pb_2fplatform_2eproto::protobuf_AssignDescriptorsOnce();
  return ::protobuf_pb_2fplatform_2eproto::file_level_metadata[kIndexInFileMessages];
}


// @@protoc_insertion_point(namespace_scope)
}  // namespace pb
namespace google {
namespace protobuf {
template<> GOOGLE_PROTOBUF_ATTRIBUTE_NOINLINE ::pb::empty* Arena::CreateMaybeMessage< ::pb::empty >(Arena* arena) {
  return Arena::CreateInternal< ::pb::empty >(arena);
}
template<> GOOGLE_PROTOBUF_ATTRIBUTE_NOINLINE ::pb::RPCHealthRequest* Arena::CreateMaybeMessage< ::pb::RPCHealthRequest >(Arena* arena) {
  return Arena::CreateInternal< ::pb::RPCHealthRequest >(arena);
}
template<> GOOGLE_PROTOBUF_ATTRIBUTE_NOINLINE ::pb::RPCHealthReply* Arena::CreateMaybeMessage< ::pb::RPCHealthReply >(Arena* arena) {
  return Arena::CreateInternal< ::pb::RPCHealthReply >(arena);
}
}  // namespace protobuf
}  // namespace google

// @@protoc_insertion_point(global_scope)
