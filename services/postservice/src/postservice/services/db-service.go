package services

import (
	"gopkg.in/mgo.v2"
	"net"
	"crypto/tls"
	"time"
)

type DbService struct {
	DbUser, DbName, DbPass string
	DbServers              []string
}

func (dbSettings DbService) GetCollection(collectionName string) (*mgo.Session, error, *mgo.Collection) {
	dialInfo := &mgo.DialInfo{
		Addrs:    dbSettings.DbServers,
		Username: dbSettings.DbUser,
		Password: dbSettings.DbPass,
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
		Timeout: time.Second * 10,
	}

	session, err := mgo.DialWithInfo(dialInfo)
	collection := session.DB(dbSettings.DbName).C(collectionName)
	return session, err, collection

}
