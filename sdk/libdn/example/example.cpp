// example.cpp : Defines the entry point for the console application.
//

#include "stdafx.h"
#include "libdn.h"
#include <iostream>

void DN_LogCB(const char* message)
{
	printf("[NP] %s", message);
}

int main()
{
	printf("DN_Init()\n");
	DN_Init(DN_LogCB);
	bool status = false;
	while (!status) {
		status = DN_Connect("10.1.1.34", 3036);
	}	
	while (true)
	{
		Sleep(100);
		DN_RunFrame();
		DNSessionInfo info;
		info.name = "Big tests";
		auto request = DN_UpsertSession(&info);
		request->SetTimeoutCallback([](DNAsync<DNUpsertSessionResult>* async) {
			DNUpsertSessionResult* answer = async->GetResult();
			printf("session id: %d", answer->sessionId);
		}, 1000);

		printf("Frame\n");
	}
}