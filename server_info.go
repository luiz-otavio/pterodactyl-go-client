package serverinfo

import "strconv"

type ServerInfo struct {
	environment Environment
	info        map[string]string
}

type Environment map[string]string
type Limit map[string]int

func (info ServerInfo) Name(name string) ServerInfo {
	return *info.set("name", name)
}

func (info ServerInfo) User(id int) ServerInfo {
	return *info.set("user", strconv.Itoa(id))
}

func (info ServerInfo) Egg(id int) ServerInfo {
	return *info.set("egg", strconv.Itoa(id))
}

func (info ServerInfo) Image(name string) ServerInfo {
	return *info.set("docker_image", name)
}

func (info ServerInfo) Startup(name string) ServerInfo {
	return *info.set("startup", name)
}

func (environment Environment) SetEnv(name string, value string) Environment {
	environment[name] = value

	return environment
}

func (info *ServerInfo) set(key string, value string) *ServerInfo {
	info.info[key] = value

	return info
}
