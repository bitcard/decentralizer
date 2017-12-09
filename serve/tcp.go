package serve

import (
	"net"
	"github.com/iain17/logger"
	"fmt"
	"github.com/iain17/decentralizer/pb"
	"reflect"
	"io"
)

func (s *Serve) ListenTCP(port int) {
	address := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.Error(err)
		return
	}
	defer listener.Close()
	logger.Infof("TCP API serving on %s", address)

	for {
		conn, err := listener.Accept()
		logger.Infof("New connection: %s", conn.RemoteAddr())
		if err != nil {
			logger.Warning(err)
			continue
		}
		go s.handleConnection(conn)
	}
}


func (s *Serve) handleConnection(conn net.Conn) {
	defer func() {
		if error := recover(); error != nil {
			logger.Errorf("Recover error: %s", error)
		}

		conn.Close()
	}()

	for {
		packets := make(chan *pb.RPCMessage)
		go s.handlePackets(conn, packets)
		err := pb.Decode(conn, packets)
		if err != nil {
			if err == io.EOF {
				break
			}
			logger.Warning(err)
			continue
		}
	}
}

func (s *Serve) handlePackets(conn net.Conn, packets chan *pb.RPCMessage) {
	for packet := range packets {
		logger.Infof("Received a packet with id %d and type %v", packet.Id, packet.GetMsg())

		handler := s.handlers[reflect.TypeOf(packet.GetMsg())]
		if handler != nil {
			res, err := handler(packet)
			if err != nil {
				logger.Warning(err)
				continue
			}
			if res != nil {
				res.Id = packet.Id
				logger.Infof("Writing reply back a packet with id %d and type %v", res.Id, res.GetMsg())
				err = pb.Write(conn, res)
				if err != nil {
					logger.Warning(err)
					continue
				}
			}
		} else {
			logger.Infof("No handler found for type %v", packet.GetMsg())
		}
	}
}