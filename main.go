package main

import "test/ftp"

// Time format used by the MDTM and MFMT commands
const timeFormat = "20060102150405"

func main()  {

	Server := new(ftp.Server)
	Server.Listen(ftp.Addr{URI:"127.0.0.1:21", Network:"tcp"})


}