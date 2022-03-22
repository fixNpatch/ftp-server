package ftp

import (
	"log"
)

func (c *Conn) pwd(args []string) {


	//if c.workDir < c.rootDir {
	//	c.workDir = c.rootDir
	//}

	log.Println(c.rootDir)
	log.Println(c.workDir)

	c.respond(status257, c.workDir)
}
