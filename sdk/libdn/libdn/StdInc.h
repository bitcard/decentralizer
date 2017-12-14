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
#include "RPC.h"
#include "Conversions.h"
#include "Promise.h"
#include "DecentralizerClient.h"

const int MAX_SESSIONS = 1024;

// global state
extern struct DN_state_s {
	libdn::DecentralizerClient* client;
	bool initialized = false;
	libdn::ConnectLogCB g_logCB;
	::google::protobuf::RepeatedField<::google::protobuf::uint64> sessions;
} context;

#endif