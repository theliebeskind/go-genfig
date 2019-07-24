// Code generated by genfig (config built by merging 'default.yml' and 'test.json') on 2019-07-24T22:16:50+02:00; DO NOT EDIT.

package config

func init() {
	Envs.Test = Config{
		Server: ConfigServer{
			Port: 1234,
			Host: "localhost",
		},
		Secrets: []string{"ChuckNorriscanwinagameofConnectFourinonlythreemoves"},
		Apis: ConfigApis{
			Google: ConfigApisGoogle{
				Uri: "google.com",
			},
		},
		LongDesc: ConfigLongDesc{
			En: "Long description",
			De: "Lange Beschreibung",
		},
		Wip:     true,
		Version: "1-test",
		Project: "genfig",
		Db: ConfigDb{
			User: "chuck",
			Pass: "norris",
			Uri:  "mongdb://${db.user}:${db.pass}@remotedb:27018/proddb",
		},
		Randomizer: ConfigRandomizer{
			Threshold: 0.12345,
		},
	}
	Envs.Test._map = map[string]interface{}{"apis": map[interface{}]interface{}{"google": map[interface{}]interface{}{"uri": "google.com"}}, "db": map[interface{}]interface{}{"pass": "norris", "uri": "mongdb://${db.user}:${db.pass}@remotedb:27018/proddb", "user": "chuck"}, "longDesc": map[interface{}]interface{}{"de": "Lange Beschreibung", "en": "Long description"}, "project": "genfig", "randomizer": map[interface{}]interface{}{"threshold": 0.12345}, "secrets": []interface{}{"ChuckNorriscanwinagameofConnectFourinonlythreemoves"}, "server": map[interface{}]interface{}{"host": "localhost", "port": 1234}, "version": "1-test", "wip": true}
}
