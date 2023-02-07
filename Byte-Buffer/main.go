package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"os"
)

func main() {
	// struktur
	// idx
	// 0-3 'IMGD'
	// 4-7 width (uint32), BE (Big Endian)
	// 8-11 height (uint32), BE (Big Endian)
	// 12 len (byte)
	// 8..x fname

	fname := "photo.imgd"
	hdr := []byte{'I', 'M', 'G', 'D'}
	fname_len := byte(len(fname)) // 10 -> 0x0A

	width := uint32(320)  // 00 00 01 40
	height := uint32(640) // 00 00 02 80

	dumper := hex.Dumper(os.Stdout)
	defer dumper.Close()

	buf := new(bytes.Buffer)

	//write here
	buf.Write(hdr)

	binary.Write(buf, binary.BigEndian, width)
	binary.Write(buf, binary.BigEndian, height)

	buf.WriteByte(fname_len)
	buf.WriteString(fname)

	// Dump ke stdout
	buf.WriteTo(dumper)
}
