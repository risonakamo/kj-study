package kj_study

import "kj-study/lib/utils"

// run time configuration of kj study exe
type KjStudyConfig struct {
	// name of a dir that exists in data/split-data. will be used
	// as target data set.
	DataDir string `yaml:"dataDir"`

	// when in a session, number of sentences that will be presented for
	// each word. does not affect the data (which has many words per sentence),
	// but will change size of each session
	SentenceConfig struct {
		Min int `yaml:"min"`
		Max int `yaml:"max"`
	} `yaml:"sentences"`

    // port to host on
    Port int `yaml:"port"`
}

// read kj study config from path
func ReadKjStudyConfig(path string) KjStudyConfig {
    var config KjStudyConfig
    var e error
    config,e=utils.ReadYaml[KjStudyConfig](path)

    if e!=nil {
        panic(e)
    }

    return config
}