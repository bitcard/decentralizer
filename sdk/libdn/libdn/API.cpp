#include "stdafx.h"
#include "StdInc.h"
#include "libdn.h"

namespace libdn {
	LIBDN_API bool LIBDN_CALL DN_Init(ConnectLogCB callback) {
		if (!RPC_Init()) {
			return false;
		}
		//Authenticate_Init();
		return true;
	}

	LIBDN_API bool LIBDN_CALL DN_Shutdown() {
		//Friends_Shutdown();
		RPC_Shutdown();
		return true;
	}

	LIBDN_API bool LIBDN_CALL DN_RunFrame() {
		Async_RunCallbacks();
		return true;
	}
}
DN_state_s g_dn;

BOOL APIENTRY DllMain(HMODULE hModule,
	DWORD  ul_reason_for_call,
	LPVOID lpReserved
)
{
	switch (ul_reason_for_call)
	{
	case DLL_PROCESS_ATTACH:
	case DLL_THREAD_ATTACH:
	case DLL_THREAD_DETACH:
	case DLL_PROCESS_DETACH:
		break;
	}
	return TRUE;
}