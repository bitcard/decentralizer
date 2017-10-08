package discovery

import (
	"github.com/anacrolix/utp"
	"github.com/iain17/decentralizer/discovery/pb"
	"net"
)

//Initiate a hand sake procedure.
//See (l *ListenerService) process(c net.Conn) error for the receiving side.
func connect(h *net.UDPAddr, ln *LocalNode) (*RemoteNode, error) {
	s, errSocket := utp.NewSocket("udp", ":0")
	if errSocket != nil {
		return nil, errSocket
	}

	conn, errDial := s.Dial(h.String())
	if errDial != nil {
		return nil, errDial
	}
	rn := NewRemoteNode(conn)
	err := rn.sendHeartBeat()
	if err != nil {
		conn.Close()
		return nil, err
	}

	rn.logger.Debug("Waiting for heartbeat...")
	err = pb.DecodeHeartBeat(rn.conn)
	if err != nil {
		conn.Close()
		return nil, err
	}
	rn.logger.Debug("Received heartbeat...")

	//They will respond by sending their peer info
	rn.logger.Debug("Waiting for peer info...")
	peerInfo, errPeerInfo := pb.DecodePeerInfo(rn.conn, string(ln.discovery.network.ExportPublicKey()))
	if errPeerInfo != nil {
		conn.Close()
		return nil, errPeerInfo
	}
	rn.logger.Debug("Received peer info...")
	rn.Info = peerInfo.Info

	//We send our peer peer info back
	rn.logger.Debug("Sending our peer info")
	ln.sendPeerInfo(rn.conn)

	rn.logger.Info("connected!")
	return rn, nil
}