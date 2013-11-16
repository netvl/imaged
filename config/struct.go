package config

type ValidationError string

func (e ValidationError) Error() string {
    return "Config validation error: " + string(e)
}

type ConfigWrapper struct {
    *Config
    Overrides Overrides
}

type Overrides struct {
    Paths Paths
}

type Config struct {
    Paths     Paths
    Interface Interface
}

type Paths struct {
    StageDir   string `toml:"stage-dir"`
    StorageDir string `toml:"storage-dir"`
}

type Interface struct {
    Socket  Socket
    Network Network
}

type Socket struct {
    Enabled bool
    Path    string
}

type Network struct {
    Enabled bool
    Host    string
    Port    int
}
