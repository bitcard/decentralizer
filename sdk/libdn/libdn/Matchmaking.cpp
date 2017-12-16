#include "StdInc.h"

namespace libdn {
	LIBDN_API Promise<UpsertSessionResult>* LIBDN_CALL UpsertSession(libdn::Session * session) {
		auto result = new Promise<UpsertSessionResult>([session](Promise<UpsertSessionResult>* promise) {
			// Data we are sending to the server.
			pb::RPCUpsertSessionRequest request;
			request.set_allocated_session(DNSessionToPBSession(session));

			// Container for the data we expect from the server.
			pb::RPCUpsertSessionResponse reply;

			// Context for the client. It could be used to convey extra information to
			// the server and/or tweak certain RPC behaviors.
			grpc::ClientContext ctx;
			grpc::Status status = context.client->stub_->UpsertSession(&ctx, request, &reply);

			UpsertSessionResult result;
			if (status.ok()) {
				result.sessionId = reply.sessionid();
			} else {
				promise->reject(va("[Could not upsert session] %i: %s", status.error_code(), status.error_message().c_str()));
			}
			return result;
		});
		return result;
	}

	LIBDN_API Promise<bool>* LIBDN_CALL DeleteSession(DNSID sid) {
		auto result = new Promise<bool>([sid](Promise<bool>* promise) {
			//build request.
			pb::RPCDeleteSessionRequest request;
			request.set_sessionid(sid);

			// Container for the data we expect from the server.
			pb::RPCDeleteSessionResponse reply;

			// Context for the client. It could be used to convey extra information to
			// the server and/or tweak certain RPC behaviors.
			grpc::ClientContext ctx;
			grpc::Status status = context.client->stub_->DeleteSession(&ctx, request, &reply);

			if (status.ok()) {
				return true;
			} else {
				promise->reject(va("[Could not delete session] %i: %s", status.error_code(), status.error_message().c_str()));
			}

			return false;
		});
		return result;
	}

	LIBDN_API Promise<int>* LIBDN_CALL GetNumSessions(uint32_t type, const char* key, const char* value) {
		auto result = new Promise<int>([type, key, value](Promise<int>* promise) {
			//build request.
			pb::RPCGetSessionIdsRequest request;
			request.set_type(type);
			request.set_key(key);
			request.set_value(value);

			// Container for the data we expect from the server.
			pb::RPCGetSessionIdsResponse reply;

			// Context for the client. It could be used to convey extra information to
			// the server and/or tweak certain RPC behaviors.
			grpc::ClientContext ctx;
			grpc::Status status = context.client->stub_->GetSessionIds(&ctx, request, &reply);

			if (status.ok()) {
				context.sessions = reply.sessionids();
				int size = context.sessions.size();
				return size;
			} else {
				promise->reject(va("[Could not get session ids] %i: %s", status.error_code(), status.error_message().c_str()));
			}

			return 0;
		});
		return result;
	}

	LIBDN_API Promise<libdn::Session*>* LIBDN_CALL GetSessionBySessionId(DNSID sessionId) {
		auto result = new Promise<libdn::Session*>([sessionId](Promise<libdn::Session*>* promise) {
			//build request.
			pb::RPCGetSessionRequest request;
			request.set_sessionid(sessionId);

			// Container for the data we expect from the server.
			pb::RPCGetSessionResponse reply;

			// Context for the client. It could be used to convey extra information to
			// the server and/or tweak certain RPC behaviors.
			grpc::ClientContext ctx;
			grpc::Status status = context.client->stub_->GetSession(&ctx, request, &reply);

			if (!status.ok()) {
				promise->reject(va("[Could not get session] %i: %s", status.error_code(), status.error_message().c_str()));
			}
			auto wtf = reply.session();
			return PBSessionToDNSession(&wtf);
		});
		return result;
	}


	LIBDN_API Session* LIBDN_CALL GetSessionByIndex(int index) {
		if (index > MAX_SESSIONS || index > context.sessions.size() - 1) {
			return NULL;
		}
		auto req = GetSessionBySessionId(context.sessions.Get(index));
		req->fail([](std::string reason) {
			Log_Print(reason.c_str());
		});
		if (req->wait()) {
			return req->get();
		}
		return NULL;
	}

}