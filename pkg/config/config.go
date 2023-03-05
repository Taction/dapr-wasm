/**
 * @Author: zhangchao
 * @Description:
 * @Date: 2023/3/2 8:41 PM
 */
package config

const (
	TriggerTypeHTTP    = "http"
	TriggerTypeBinding = "binding"
	TriggerTypePubSub  = "pubsub"
)

type Config struct {
	Version      string `yaml:"version"`
	TLS          `yaml:"tls"`
	Applications []Application `yaml:"applications"`
}

type TLS struct {
	Enabled bool `yaml:"enabled"`
}

type Application struct {
	Trigger `yaml:"trigger,omitempty"`
	Source  `yaml:"source,omitempty"`
}

type Trigger struct {
	Type     string         `yaml:"type"`
	Name     string         `yaml:"name"`
	Metadata []MetadataItem `yaml:"metadata" json:"metadata"`
}

type Source struct {
	Type string `yaml:"type"`
	Path string `yaml:"path"`
}

// MetadataItem is a name/value pair for a metadata.
type MetadataItem struct {
	Name string `json:"name"`
	//+optional
	Value string `json:"value,omitempty"`
}
