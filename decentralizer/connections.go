package decentralizer

import (
	"github.com/ccding/go-stun/stun"
	"net"
	"errors"
	"github.com/iain17/decentralizer/decentralizer/upnp"
	"fmt"
	logger "github.com/Sirupsen/logrus"
)

//Returns a forwarded udp connection.
func getUdpConn() (*net.UDPConn, *stun.Host, error) {
	conn, err := net.ListenUDP("udp", nil)
	if err != nil {
		return nil, nil, err
	}
	port := conn.LocalAddr().(*net.UDPAddr).Port
	err = upnp.Open(port, port, "udp")
	if err != nil {
		logger.Warning(err)
	}
	nat, host, err := stun.NewClientWithConnection(conn).Discover()
	logger.Infof("Nat type is %s, %s", nat.String(), conn.LocalAddr().(*net.UDPAddr).String())
	if nat != stun.NATFull && nat != stun.NATNone {
		logger.Warn(errors.New("Could not forward UDP connection"+nat.String()))
	}
	conn.Close()
	conn, err = net.ListenUDP("udp", conn.LocalAddr().(*net.UDPAddr))

	return conn, host, nil
}

//TODO: Reachable check
func getTcpConn() (net.Listener, error) {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		return nil, err
	}
	//Forward the port
	port := lis.Addr().(*net.TCPAddr).Port
	err = upnp.Open(port, port, "tcp")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not forward TCP RPC server. %v", err))
	}
	return lis, nil
}

