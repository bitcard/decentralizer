// This file has been automatically generated.

#pragma once
using namespace pb;

// RPC message base class
class INPRPCMessage
{
public:
	virtual void Deserialize(const uint8_t* message, size_t length) = 0;
	virtual uint8_t* Serialize(size_t* length, uint32_t id) = 0;
	
	virtual int GetType() = 0;
	
	virtual void Free() = 0;
	virtual void FreePayload() = 0;
};

// RPC message parsing callbacks
typedef INPRPCMessage* (* CreateMessageCB)();

struct rpc_message_type_s
{
	uint32_t type;
	CreateMessageCB handler;
};

extern rpc_message_type_s g_rpcMessageTypes[];
#define NUM_RPC_MESSAGE_TYPES 2 

// message class definitions
class RPCHealthRequest : public INPRPCMessage
{
private:
	NPRPCBuffer<HealthRequest> _buffer;
	
	uint8_t* _payload;
public:
	RPCHealthRequest()
	{
		_payload = NULL;
	}

	enum { Type = 1000 };
	
	HealthRequest* GetBuffer();

	virtual void Deserialize(const uint8_t* message, size_t length);
	virtual uint8_t* Serialize(size_t* length, uint32_t id);
	
	virtual int GetType();
	
	virtual void Free();
	virtual void FreePayload();
	static RPCHealthRequest* Create();
};
class RPCHealthReply : public INPRPCMessage
{
private:
	NPRPCBuffer<HealthReply> _buffer;
	
	uint8_t* _payload;
public:
	RPCHealthReply()
	{
		_payload = NULL;
	}

	enum { Type = 1001 };
	
	HealthReply* GetBuffer();

	virtual void Deserialize(const uint8_t* message, size_t length);
	virtual uint8_t* Serialize(size_t* length, uint32_t id);
	
	virtual int GetType();
	
	virtual void Free();
	virtual void FreePayload();
	static RPCHealthReply* Create();
};
