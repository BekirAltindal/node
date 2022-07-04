/** Dump fields and values in a socket to plot them with villasDump.py.
 *
 * @author Manuel Pitz <manuel.pitz@eonerc.rwth-aachen.de>
 * @copyright 2014-2022, Institute for Automation of Complex Power Systems, EONERC
 * @license Apache 2.0
 *********************************************************************************/

#include <sys/socket.h>
#include <sys/types.h>
#include <sys/un.h>
#include <unistd.h>

#include <sstream>
#include <cstring>

#include <villas/dumper.hpp>

using namespace villas;
using namespace villas::node;

Dumper::Dumper(const std::string &socketNameIn) :
	socketName(socketNameIn),
	supressRepeatedWarning(true),
	warningCounter(0),
	logger(logging.get("dumper"))
{
	openSocket();
}

Dumper::~Dumper() {
	closeSocket();
}

int Dumper::openSocket()
{
	socketFd = socket(AF_LOCAL, SOCK_STREAM, 0);
	if (socketFd < 0) {
		logger->info("Error creating socket {}", socketName);
		return -1;
	}

	sockaddr_un socketaddrUn;
	socketaddrUn.sun_family = AF_UNIX;
	strcpy(socketaddrUn.sun_path, socketName.c_str());

	int ret = connect(socketFd, (struct sockaddr *) &socketaddrUn, sizeof(socketaddrUn));
	if (!ret)
		return ret;

	return 0;
}

int Dumper::closeSocket()
{
	int ret = close(socketFd);
	if (!ret)
		return ret;

	return 0;
}

void Dumper::writeDataBinary(unsigned len, double *yData, double *xData){

	if (warningCounter > 10)
		return;

	if (yData == nullptr)
		return;

	unsigned dataLen = len * sizeof(yData[0]);
	ssize_t bytesWritten = write(socketFd, &dataLen, sizeof(dataLen));
	if ((size_t) bytesWritten != sizeof(len)) {
		logger->warn("Could not send all content (Len) to socket {}", socketName);
		warningCounter++;
	}

	bytesWritten = write(socketFd, "d000", 4);
	if (bytesWritten != 4) {
		logger->warn("Could not send all content (Type) to socket {}", socketName);
		warningCounter++;
	}

	bytesWritten = write(socketFd, yData, dataLen );
	if (bytesWritten != (ssize_t) dataLen && (!supressRepeatedWarning || warningCounter <1 )) {
		logger->warn("Could not send all content (Data) to socket {}", socketName);
		warningCounter++;
	}
}

void Dumper::writeDataCSV(unsigned len, double *yData, double *xData)
{
	for (unsigned i = 0; i < len; i++) {
		std::stringstream ss;

		ss << yData[i];

		if (xData != nullptr)
			ss << ";" << xData[i];

		ss << std::endl;

		auto str = ss.str();
		auto bytesWritten =  write(socketFd, str.c_str(), str.length());
		if ((size_t) bytesWritten != str.length() && (!supressRepeatedWarning || warningCounter < 1)) {
			logger->warn("Could not send all content to socket {}", socketName);
			warningCounter++;
		}
	}
}
