// Code generated by genfig (config built from 'development.yaml') on 2019-07-01T12:51:25+02:00; DO NOT EDIT.

package config

var Development = Config{
  Version: "1",
  Server: ConfigServer{
    Port: 1234,
    Host: "dev.mydomain.com",
  },
  Db: ConfigDb{
    Uri: "mongdb://${db.user}:${db.pass}@remotedb:27018/dev-db",
  },
}