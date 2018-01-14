#pragma once

#ifndef _STDINC
#define _STDINC

#define _CRT_SECURE_NO_WARNINGS

// Windows headers
#define WIN32_LEAN_AND_MEAN
#include <windows.h>

//GRPC + Protobuf
#include <grpc++/grpc++.h>

#include "addressbook.pb.h"
#include "addressbook.grpc.pb.h"

#include "matchmaking.pb.h"
#include "matchmaking.grpc.pb.h"

#include "messaging.pb.h"
#include "messaging.grpc.pb.h"

#include "storage.pb.h"
#include "storage.grpc.pb.h"

#include "platform.pb.h"
#include "platform.grpc.pb.h"

// C/C++ headers
#include <string>
#include <vector>
#include <queue>
#include <mutex>

// app specific headers
#include "libdn.h"
#include "Utils.h"
#include "Adna.h"
#include "RPC.h"
#include "Conversions.h"
#include "Promise.h"
#include "DecentralizerClient.h"
#include "fmt/format.h"

extern const char * VERSION;

// global state
extern struct DN_state_s {
	std::mutex AdnaMutex;
	std::mutex selfMutex;
	std::mutex mutex;
	std::mutex healthMutex;
	libdn::DecentralizerClient* client;
	bool initialized = false;
	const char* host;
	int port;
	libdn::LogCB g_logCB;
	::google::protobuf::RepeatedField<::google::protobuf::uint64> sessions;
	::google::protobuf::RepeatedPtrField< ::std::string> peers;
	libdn::Peer self;
} context;

#endif