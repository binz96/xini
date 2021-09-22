package xini

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Section struct {
	Name string
	kvs  map[string]string
}

func (s *Section) Key(keyName string) (value string) {
	return s.kvs[keyName]
}

type Config struct {
	Sections []Section
}

func (c *Config) Section(secName string) *Section {
	for i := 0; i < len(c.Sections); i++ {
		if c.Sections[i].Name == secName {
			return &c.Sections[i]
		}
	}
	return nil
}

func Load(filename string) (*Config, error) {
	sec := Section{Name: "", kvs: make(map[string]string)}
	cfg := &Config{Sections: make([]Section, 0)}
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(fileBytes), "\n")
	for lineNum, line := range lines {
		if lineNum == len(lines)-1 {
			cfg.Sections = append(cfg.Sections, sec)
		}
		line = strings.TrimSpace(line)
		// 忽略空行
		if len(line) == 0 {
			continue
		}
		// 忽略注释
		if line[0] == ';' || line[0] == '#' {
			continue
		}

		// section
		if line[0] == '[' {
			cfg.Sections = append(cfg.Sections, sec)
			// 非法section
			if line[len(line)-1] != ']' || len(line) < 3 {
				return nil, fmt.Errorf("section syntax error in line %d", lineNum+1)
			}
			secName := strings.TrimSpace(line[1 : len(line)-1])
			if len(secName) == 0 {
				return nil, fmt.Errorf("section syntax error in line %d", lineNum+1)
			}

			sec = Section{
				Name: secName,
				kvs:  make(map[string]string),
			}
			continue
		}

		// kv
		eq := strings.Index(line, "=")
		if eq == -1 {
			return nil, fmt.Errorf("syntax error in line %d", lineNum+1)
		}
		key := strings.TrimSpace(line[0:eq])
		if len(key) == 0 {
			return nil, fmt.Errorf("syntax error in line %d", lineNum+1)
		}
		value := ""
		if eq < len(line) {
			value = strings.TrimSpace(line[eq+1:])
		}
		sec.kvs[key] = value
	}
	return cfg, nil
}
