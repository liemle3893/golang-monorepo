package viper

import (
	"github.com/spf13/viper"
	"reflect"
	"strings"
)

// CreateEnvViper create a new instance of Viper that support Environment variables binding to Struct.
// Example:
//	os.Setenv("NAME", "TestViper")
//	os.Setenv("APPLICATION_ID", "TestApp")
//	os.Setenv("NESTED_ASTRING", "nested string")
//	type NestedConfig struct {
//		AString string `mapstructure:"astring"`
//	}
//	type Config struct {
//		Name          string       `mapstructure:"name"`
//		ApplicationID string       `mapstructure:"application-id"`
//		Nested        NestedConfig `mapstructure:"nested"`
//	}
//
//	confReader := CreateEnvViper(Config{})
//	c := &Config{}
//	confReader.Unmarshal(c)
//	fmt.Printf("%+v", c)
// Result will be: &{Name:TestViper ApplicationID:TestApp Nested:{AString:nested string}}
func CreateEnvViper(v *viper.Viper, ifaces ...interface{}) *viper.Viper {
	if v == nil {
		v = viper.New()
	}
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()
	for _, iface := range ifaces {
		BindEnvs(v, iface)
	}
	return v
}

// BindEnvs bind environment into struct, nested attribute was joined by "."
func BindEnvs(vp *viper.Viper, iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			BindEnvs(vp, v.Interface(), append(parts, tv)...)
		default:
			s := strings.Join(append(parts, tv), ".")
			vp.BindEnv(s)
		}
	}
}
