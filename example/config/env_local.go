// Code generated by genfig (config built by merging 'default.yml' and '.env.local') on 2019-07-24T22:16:50+02:00; DO NOT EDIT.

package config

func init() {
	Envs.Local = Config{
		Server: ConfigServer{
			Port: 1212,
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
		Version: "0.1.0",
		Project: "genfig",
		Db: ConfigDb{
			User: "chuck",
			Pass: "norris",
			Uri:  "mongdb://localhos:27017/db",
		},
		Randomizer: ConfigRandomizer{
			Threshold: 0.12345,
		},
	}
	Envs.Local._map = map[string]interface{}{"apis": map[interface{}]interface{}{"google": map[interface{}]interface{}{"uri": "google.com"}}, "db": map[interface{}]interface{}{"pass": "norris", "uri": "mongdb://localhos:27017/db", "user": "chuck"}, "longDesc": map[interface{}]interface{}{"de": "Lange Beschreibung", "en": "Long description"}, "project": "genfig", "randomizer": map[interface{}]interface{}{"threshold": 0.12345}, "secrets": []interface{}{"ChuckNorriscanwinagameofConnectFourinonlythreemoves"}, "server": map[interface{}]interface{}{"host": "localhost", "port": 1212}, "version": "0.1.0", "wip": true}
}
