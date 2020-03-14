package singleton

type Config interface {
	GetProperty(string) string
	SetProperty(string, string)
}

var config Config = configInstance{map[string]string {
		"p1": "v1",
		"p2": "v2",
	},
}

type configInstance struct {
	properties map[string]string
}

func (c configInstance) GetProperty(key string)  string {
	return c.properties[key]
}

func (c configInstance) SetProperty(key string, value string) {
	c.properties[key] = value
}

func GetInstance() Config {
	return config
}
