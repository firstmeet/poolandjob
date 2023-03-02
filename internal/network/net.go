package network

import (
	"awesomeProject/internal"
	"encoding/binary"
	"encoding/json"
	"io"
	"net"
)

func ReadPackage(conn net.Conn) (*internal.Package, error) {
	buf := make([]byte, internal.DefaultHeaderSize)
	n, err := io.ReadFull(conn, buf)
	if err != nil {
		return nil, err
	}
	length := binary.BigEndian.Uint32(buf[:n])
	buf = make([]byte, length)
	n, err = io.ReadFull(conn, buf)
	if err != nil {
		return nil, err
	}
	pkg := &internal.Package{}
	err = json.Unmarshal(buf[:n], pkg)
	if err != nil {
		return nil, err
	}
	return pkg, nil
}
func SendPackage(conn net.Conn, pkg *internal.Package) error {
	buf, err := json.Marshal(pkg)
	if err != nil {
		return err
	}
	length := uint32(len(buf))
	buf2 := make([]byte, internal.DefaultHeaderSize)
	binary.BigEndian.PutUint32(buf2, length)
	buf2 = append(buf2, buf...)
	_, err = conn.Write(buf2)
	if err != nil {
		return err
	}
	return nil
}
