package Hedayat

type Config struct {
	Kafka Kafka `json:"kafka" yaml:"kafka" mapstructure:"kafka" koanf:"kafka"`
	Avro  Avro  `json:"avro"  yaml:"avro"  mapstructure:"avro"  koanf:"avro"`
}

type Kafka struct {
	Brokers                   []string `json:"brokers"                      yaml:"brokers"                      mapstructure:"brokers"                      koanf:"brokers"`
	SASLMechanisms            string   `json:"sasl_mechanisms"              yaml:"sasl_mechanisms"              mapstructure:"sasl_mechanisms"              koanf:"sasl_mechanisms"`
	SASLUsername              string   `json:"sasl_username"                yaml:"sasl_username"                mapstructure:"sasl_username"                koanf:"sasl_username"`
	SASLPassword              string   `json:"sasl_password"                yaml:"sasl_password"                mapstructure:"sasl_password"                koanf:"sasl_password"`
	SecurityProtocol          string   `json:"security_protocol"            yaml:"security_protocol"            mapstructure:"security_protocol"            koanf:"security_protocol"`
	ClientID                  string   `json:"client_id"                    yaml:"client_id"                    mapstructure:"client_id"                    koanf:"client_id"`
	Acks                      string   `json:"acks"                         yaml:"acks"                         mapstructure:"acks"                         koanf:"acks"`
	MessageTimeoutMs          uint32   `json:"message_timeout_ms"           yaml:"message_timeout_ms"           mapstructure:"message_timeout_ms"           koanf:"message_timeout_ms"`
	Retries                   uint     `json:"retries"                      yaml:"retries"                      mapstructure:"retries"                      koanf:"retries"`
	RetryBackoffMs            uint32   `json:"retry_backoff_ms"             yaml:"retry_backoff_ms"             mapstructure:"retry_backoff_ms"             koanf:"retry_backoff_ms"`
	QueueBufferingMaxMessages uint32   `json:"queue_buffering_max_messages" yaml:"queue_buffering_max_messages" mapstructure:"queue_buffering_max_messages" koanf:"queue_buffering_max_messages"`
	BatchNumMessages          uint32   `json:"batch_num_messages"           yaml:"batch_num_messages"           mapstructure:"batch_num_messages"           koanf:"batch_num_messages"`
}

type Avro struct {
	SchemaRegistryURL          string `json:"schema_registry_url"           yaml:"schema_registry_url"           mapstructure:"schema_registry_url"           koanf:"schema_registry_url"`
	BasicAuthCredentialsSource string `json:"basic_auth_credentials_source" yaml:"basic_auth_credentials_source" mapstructure:"basic_auth_credentials_source" koanf:"basic_auth_credentials_source"`
	MaxRetries                 uint   `json:"max_retries"                   yaml:"max_retries"                   mapstructure:"max_retries"                   koanf:"max_retries"`
	RetriesMaxWaitMs           uint32 `json:"retries_max_wait_ms"           yaml:"retries_max_wait_ms"           mapstructure:"retries_max_wait_ms"           koanf:"retries_max_wait_ms"`
}
