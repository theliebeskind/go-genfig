// Code generated by genfig (config built from 'default.yml') on 2019-07-01T12:51:25+02:00; DO NOT EDIT.

package config

var Default = Config{
  Db: ConfigDb{
    User: "",
    Pass: "",
    Uri: "mongdb://localhos:27017/db",
  },
  Secrets: []string{""},
  Randomizer: ConfigRandomizer{
    Threshold: 0.75,
  },
  Apis: ConfigApis{
    Google: ConfigApisGoogle{
      Uri: "google.com",
    },
  },
  Version: "0.1.0",
  Project: "genfig",
  Server: ConfigServer{
    Port: 1234,
    Host: "localhost",
  },
}