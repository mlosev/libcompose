package config

var schemaDataV1 = `{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "id": "config_schema_v1.json",

  "type": "object",

  "patternProperties": {
    "^[a-zA-Z0-9._-]+$": {
      "$ref": "#/definitions/service"
    }
  },

  "additionalProperties": false,

  "definitions": {
    "service": {
      "id": "#/definitions/service",
      "type": "object",

      "properties": {
        "build": {"type": "string"},
        "cap_add": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "cap_drop": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "cgroup_parent": {"type": "string"},
        "command": {
          "oneOf": [
            {"type": "string"},
            {"type": "array", "items": {"type": "string"}}
          ]
        },
        "container_name": {"type": "string"},
        "cpu_shares": {"type": ["number", "string"]},
        "cpu_quota": {"type": ["number", "string"]},
        "cpuset": {"type": "string"},
        "devices": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "dns": {"$ref": "#/definitions/string_or_list"},
        "dns_search": {"$ref": "#/definitions/string_or_list"},
        "dockerfile": {"type": "string"},
        "domainname": {"type": "string"},
        "entrypoint": {
          "oneOf": [
            {"type": "string"},
            {"type": "array", "items": {"type": "string"}}
          ]
        },
        "env_file": {"$ref": "#/definitions/string_or_list"},
        "environment": {"$ref": "#/definitions/list_or_dict"},

        "expose": {
          "type": "array",
          "items": {
            "type": ["string", "number"],
            "format": "expose"
          },
          "uniqueItems": true
        },

        "extends": {
          "oneOf": [
            {
              "type": "string"
            },
            {
              "type": "object",

              "properties": {
                "service": {"type": "string"},
                "file": {"type": "string"}
              },
              "required": ["service"],
              "additionalProperties": false
            }
          ]
        },

        "extra_hosts": {"$ref": "#/definitions/list_or_dict"},
        "external_links": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "hostname": {"type": "string"},
        "image": {"type": "string"},
        "ipc": {"type": "string"},
        "labels": {"$ref": "#/definitions/list_or_dict"},
        "links": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "log_driver": {"type": "string"},
        "log_opt": {"type": "object"},
        "mac_address": {"type": "string"},
        "mem_limit": {"type": ["number", "string"]},
        "mem_reservation": {"type": ["number", "string"]},
        "memswap_limit": {"type": ["number", "string"]},
        "mem_swappiness": {"type": "integer"},
        "net": {"type": "string"},
        "pid": {"type": ["string", "null"]},

        "ports": {
          "type": "array",
          "items": {
            "type": ["string", "number"],
            "format": "ports"
          },
          "uniqueItems": true
        },

        "privileged": {"type": "boolean"},
        "read_only": {"type": "boolean"},
        "restart": {"type": "string"},
        "security_opt": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "shm_size": {"type": ["number", "string"]},
        "stdin_open": {"type": "boolean"},
        "stop_signal": {"type": "string"},
        "tty": {"type": "boolean"},
        "ulimits": {
          "type": "object",
          "patternProperties": {
            "^[a-z]+$": {
              "oneOf": [
                {"type": "integer"},
                {
                  "type":"object",
                  "properties": {
                    "hard": {"type": "integer"},
                    "soft": {"type": "integer"}
                  },
                  "required": ["soft", "hard"],
                  "additionalProperties": false
                }
              ]
            }
          }
        },
        "user": {"type": "string"},
        "volumes": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "volume_driver": {"type": "string"},
        "volumes_from": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "working_dir": {"type": "string"}
      },

      "dependencies": {
        "memswap_limit": ["mem_limit"]
      },
      "additionalProperties": false
    },

    "string_or_list": {
      "oneOf": [
        {"type": "string"},
        {"$ref": "#/definitions/list_of_strings"}
      ]
    },

    "list_of_strings": {
      "type": "array",
      "items": {"type": "string"},
      "uniqueItems": true
    },

    "list_or_dict": {
      "oneOf": [
        {
          "type": "object",
          "patternProperties": {
            ".+": {
              "type": ["string", "number", "null"]
            }
          },
          "additionalProperties": false
        },
        {"type": "array", "items": {"type": "string"}, "uniqueItems": true}
      ]
    },

    "constraints": {
      "service": {
        "id": "#/definitions/constraints/service",
        "anyOf": [
          {
            "required": ["build"],
            "not": {"required": ["image"]}
          },
          {
            "required": ["image"],
            "not": {"anyOf": [
              {"required": ["build"]},
              {"required": ["dockerfile"]}
            ]}
          }
        ]
      }
    }
  }
}
`

var servicesSchemaDataV2 = `{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "id": "config_schema_v2.2.json",
  "type": "object",

  "patternProperties": {
    "^[a-zA-Z0-9._-]+$": {
      "$ref": "#/definitions/service"
    }
  },

  "additionalProperties": false,

  "definitions": {

    "service": {
      "id": "#/definitions/service",
      "type": "object",

      "properties": {
        "blkio_config": {
          "type": "object",
          "properties": {
            "device_read_bps": {
              "type": "array",
              "items": {"$ref": "#/definitions/blkio_limit"}
            },
            "device_read_iops": {
              "type": "array",
              "items": {"$ref": "#/definitions/blkio_limit"}
            },
            "device_write_bps": {
              "type": "array",
              "items": {"$ref": "#/definitions/blkio_limit"}
            },
            "device_write_iops": {
              "type": "array",
              "items": {"$ref": "#/definitions/blkio_limit"}
            },
            "weight": {"type": "integer"},
            "weight_device": {
              "type": "array",
              "items": {"$ref": "#/definitions/blkio_weight"}
            }
          },
          "additionalProperties": false
        },

        "build": {
          "oneOf": [
            {"type": "string"},
            {
              "type": "object",
              "properties": {
                "context": {"type": "string"},
                "dockerfile": {"type": "string"},
                "args": {"$ref": "#/definitions/list_or_dict"},
                "labels": {"$ref": "#/definitions/list_or_dict"},
                "cache_from": {"$ref": "#/definitions/list_of_strings"},
                "network": {"type": "string"}
              },
              "additionalProperties": false
            }
          ]
        },
        "cap_add": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "cap_drop": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "cgroup_parent": {"type": "string"},
        "command": {
          "oneOf": [
            {"type": "string"},
            {"type": "array", "items": {"type": "string"}}
          ]
        },
        "container_name": {"type": "string"},
        "cpu_count": {"type": "integer", "minimum": 0},
        "cpu_percent": {"type": "integer", "minimum": 0, "maximum": 100},
        "cpu_shares": {"type": ["number", "string"]},
        "cpu_quota": {"type": ["number", "string"]},
        "cpus": {"type": "number", "minimum": 0},
        "cpuset": {"type": "string"},
        "depends_on": {
          "oneOf": [
            {"$ref": "#/definitions/list_of_strings"},
            {
              "type": "object",
              "additionalProperties": false,
              "patternProperties": {
                "^[a-zA-Z0-9._-]+$": {
                  "type": "object",
                  "additionalProperties": false,
                  "properties": {
                    "condition": {
                      "type": "string",
                      "enum": ["service_started", "service_healthy"]
                    }
                  },
                  "required": ["condition"]
                }
              }
            }
          ]
        },
        "devices": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "dns_opt": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "uniqueItems": true
        },
        "dns": {"$ref": "#/definitions/string_or_list"},
        "dns_search": {"$ref": "#/definitions/string_or_list"},
        "domainname": {"type": "string"},
        "entrypoint": {
          "oneOf": [
            {"type": "string"},
            {"type": "array", "items": {"type": "string"}}
          ]
        },
        "env_file": {"$ref": "#/definitions/string_or_list"},
        "environment": {"$ref": "#/definitions/list_or_dict"},

        "expose": {
          "type": "array",
          "items": {
            "type": ["string", "number"],
            "format": "expose"
          },
          "uniqueItems": true
        },

        "extends": {
          "oneOf": [
            {
              "type": "string"
            },
            {
              "type": "object",

              "properties": {
                "service": {"type": "string"},
                "file": {"type": "string"}
              },
              "required": ["service"],
              "additionalProperties": false
            }
          ]
        },

        "external_links": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "extra_hosts": {"$ref": "#/definitions/list_or_dict"},
        "healthcheck": {"$ref": "#/definitions/healthcheck"},
        "hostname": {"type": "string"},
        "image": {"type": "string"},
        "init": {"type": ["boolean", "string"]},
        "ipc": {"type": "string"},
        "isolation": {"type": "string"},
        "labels": {"$ref": "#/definitions/list_or_dict"},
        "links": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},

        "logging": {
            "type": "object",

            "properties": {
                "driver": {"type": "string"},
                "options": {"type": "object"}
            },
            "additionalProperties": false
        },

        "mac_address": {"type": "string"},
        "mem_limit": {"type": ["number", "string"]},
        "mem_reservation": {"type": ["string", "integer"]},
        "mem_swappiness": {"type": "integer"},
        "memswap_limit": {"type": ["number", "string"]},
        "network_mode": {"type": "string"},

        "networks": {
          "oneOf": [
            {"$ref": "#/definitions/list_of_strings"},
            {
              "type": "object",
              "patternProperties": {
                "^[a-zA-Z0-9._-]+$": {
                  "oneOf": [
                    {
                      "type": "object",
                      "properties": {
                        "aliases": {"$ref": "#/definitions/list_of_strings"},
                        "ipv4_address": {"type": "string"},
                        "ipv6_address": {"type": "string"},
                        "link_local_ips": {"$ref": "#/definitions/list_of_strings"}
                      },
                      "additionalProperties": false
                    },
                    {"type": "null"}
                  ]
                }
              },
              "additionalProperties": false
            }
          ]
        },
        "oom_kill_disable": {"type": "boolean"},
        "oom_score_adj": {"type": "integer", "minimum": -1000, "maximum": 1000},
        "group_add": {
            "type": "array",
            "items": {
                "type": ["string", "number"]
            },
            "uniqueItems": true
        },
        "pid": {"type": ["string", "null"]},

        "ports": {
          "type": "array",
          "items": {
            "type": ["string", "number"],
            "format": "ports"
          },
          "uniqueItems": true
        },

        "privileged": {"type": "boolean"},
        "read_only": {"type": "boolean"},
        "restart": {"type": "string"},
        "scale": {"type": "integer"},
        "security_opt": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "shm_size": {"type": ["number", "string"]},
        "sysctls": {"$ref": "#/definitions/list_or_dict"},
        "pids_limit": {"type": ["number", "string"]},
        "stdin_open": {"type": "boolean"},
        "stop_grace_period": {"type": "string", "format": "duration"},
        "stop_signal": {"type": "string"},
        "storage_opt": {"type": "object"},
        "tmpfs": {"$ref": "#/definitions/string_or_list"},
        "tty": {"type": "boolean"},
        "ulimits": {
          "type": "object",
          "patternProperties": {
            "^[a-z]+$": {
              "oneOf": [
                {"type": "integer"},
                {
                  "type":"object",
                  "properties": {
                    "hard": {"type": "integer"},
                    "soft": {"type": "integer"}
                  },
                  "required": ["soft", "hard"],
                  "additionalProperties": false
                }
              ]
            }
          }
        },
        "user": {"type": "string"},
        "userns_mode": {"type": "string"},
        "volumes": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "volume_driver": {"type": "string"},
        "volumes_from": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "working_dir": {"type": "string"}
      },

      "dependencies": {
        "memswap_limit": ["mem_limit"]
      },
      "additionalProperties": false
    },

    "healthcheck": {
      "id": "#/definitions/healthcheck",
      "type": "object",
      "additionalProperties": false,
      "properties": {
        "disable": {"type": "boolean"},
        "interval": {"type": "string"},
        "retries": {"type": "number"},
        "test": {
          "oneOf": [
            {"type": "string"},
            {"type": "array", "items": {"type": "string"}}
          ]
        },
        "timeout": {"type": "string"}
      }
    },

    "network": {
      "id": "#/definitions/network",
      "type": "object",
      "properties": {
        "driver": {"type": "string"},
        "driver_opts": {
          "type": "object",
          "patternProperties": {
            "^.+$": {"type": ["string", "number"]}
          }
        },
        "ipam": {
            "type": "object",
            "properties": {
                "driver": {"type": "string"},
                "config": {
                    "type": "array"
                },
                "options": {
                  "type": "object",
                  "patternProperties": {
                    "^.+$": {"type": "string"}
                  },
                  "additionalProperties": false
                }
            },
            "additionalProperties": false
        },
        "external": {
          "type": ["boolean", "object"],
          "properties": {
            "name": {"type": "string"}
          },
          "additionalProperties": false
        },
        "internal": {"type": "boolean"},
        "enable_ipv6": {"type": "boolean"},
        "labels": {"$ref": "#/definitions/list_or_dict"}
      },
      "additionalProperties": false
    },

    "volume": {
      "id": "#/definitions/volume",
      "type": ["object", "null"],
      "properties": {
        "driver": {"type": "string"},
        "driver_opts": {
          "type": "object",
          "patternProperties": {
            "^.+$": {"type": ["string", "number"]}
          }
        },
        "external": {
          "type": ["boolean", "object"],
          "properties": {
            "name": {"type": "string"}
          },
          "additionalProperties": false
        },
        "labels": {"$ref": "#/definitions/list_or_dict"},
        "name": {"type": "string"}
      },
      "additionalProperties": false
    },

    "string_or_list": {
      "oneOf": [
        {"type": "string"},
        {"$ref": "#/definitions/list_of_strings"}
      ]
    },

    "list_of_strings": {
      "type": "array",
      "items": {"type": "string"},
      "uniqueItems": true
    },

    "list_or_dict": {
      "oneOf": [
        {
          "type": "object",
          "patternProperties": {
            ".+": {
              "type": ["string", "number", "null"]
            }
          },
          "additionalProperties": false
        },
        {"type": "array", "items": {"type": "string"}, "uniqueItems": true}
      ]
    },

    "blkio_limit": {
      "type": "object",
      "properties": {
        "path": {"type": "string"},
        "rate": {"type": ["integer", "string"]}
      },
      "additionalProperties": false
    },
    "blkio_weight": {
      "type": "object",
      "properties": {
        "path": {"type": "string"},
        "weight": {"type": "integer"}
      },
      "additionalProperties": false
    },

    "constraints": {
      "service": {
        "id": "#/definitions/constraints/service",
        "anyOf": [
          {"required": ["build"]},
          {"required": ["image"]}
        ],
        "properties": {
          "build": {
            "required": ["context"]
          }
        }
      }
    }
  }
}
`
