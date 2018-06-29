package variant

import "reflect"

type InputConfig struct {
	Name          string                 `yaml:"name,omitempty"`
	Description   string                 `yaml:"description,omitempty"`
	ArgumentIndex *int                   `yaml:"argument-index,omitempty"`
	Type          string                 `yaml:"type,omitempty"`
	Default       interface{}            `yaml:"default,omitempty"`
	Remainings    map[string]interface{} `yaml:",inline"`
}

func (c *InputConfig) Required() bool {
	return c.Default == nil
}

func (c *InputConfig) DefaultAsString() string {
	return getOrDefault(c.Default, reflect.String, "").(string)
}

func (c *InputConfig) DefaultAsBool() bool {
	return getOrDefault(c.Default, reflect.Bool, false).(bool)
}

func (c *InputConfig) DefaultAsInt() int {
	// Dirty work-around to avoid the conflicting two errors:
	// - panic: interface conversion: interface {} is int64, not int
	// - panic: interface conversion: interface {} is int, not int64
	var v int
	v64, is64b := getOrDefault(c.Default, reflect.Int, 0).(int64)
	if is64b {
		v = int(v64)
	} else {
		v = getOrDefault(c.Default, reflect.Int, 0).(int)
	}
	return v
}

func (c *InputConfig) TypeName() string {
	var tpe string
	if c.Type == "" {
		tpe = "string"
	} else {
		tpe = c.Type
	}
	return tpe
}

func (c *InputConfig) JSONSchema() map[string]interface{} {
	jsonschema := map[string]interface{}{}
	for k, v := range c.Remainings {
		jsonschema[k] = v
	}
	jsonschema["type"] = c.TypeName()
	return jsonschema
}

type ParameterConfig struct {
	Name        string                 `yaml:"name,omitempty"`
	Description string                 `yaml:"description,omitempty"`
	Type        string                 `yaml:"type,omitempty"`
	Default     interface{}            `yaml:"default,omitempty"`
	Required    bool                   `yaml:"required,omitempty"`
	Remainings  map[string]interface{} `yaml:",inline"`
}

type OptionConfig struct {
	Name        string                 `yaml:"name,omitempty"`
	Description string                 `yaml:"description,omitempty"`
	Type        string                 `yaml:"type,omitempty"`
	Default     interface{}            `yaml:"default,omitempty"`
	Required    bool                   `yaml:"required,omitempty"`
	Remainings  map[string]interface{} `yaml:",inline"`
}