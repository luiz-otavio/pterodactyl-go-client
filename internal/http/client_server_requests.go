package http

type ServerDetailsSchema struct {
	PteroRequestSchema
	PteroMetaSchema

	Attributes struct {
		Identifier  string `json:"identifier"`
		UUID        string `json:"uuid"`
		Name        string `json:"name"`
		Node        string `json:"node"`
		SFTPDetails struct {
			IP   string `json:"ip"`
			Port int    `json:"port"`
		} `json:"sftp_details"`
		Description string `json:"description"`
		Limits      struct {
			Memory int `json:"memory"`
			Swap   int `json:"swap"`
			Disk   int `json:"disk"`
			Io     int `json:"io"`
			CPU    int `json:"cpu"`
		} `json:"limits"`
		FeatureLimits struct {
			Databases   int `json:"databases"`
			Allocations int `json:"allocations"`
		} `json:"feature_limits"`
		IsSuspended  bool `json:"is_suspended"`
		IsInstalling bool `json:"is_installing"`
		Relations    struct {
			Allocations struct {
				PteroRequestSchema
				Data []AllocationSchema `json:"data"`
			}
		} `json:"relationships"`
	} `json:"attributes"`

	Meta struct {
		ServerOwner bool     `json:"server_owner"`
		Permissions []string `json:"user_permissions"`
	} `json:"meta"`
}

type ServerWebsocketSchema struct {
	Data struct {
		Token  string `json:"token"`
		Socket string `json:"socket"`
	} `json:"data"`
}

type ServerResourcesUsageSchema struct {
	PteroRequestSchema

	Attributes struct {
		State     string `json:"current_state"`
		Suspended bool   `json:"is_suspended"`
		Limits    struct {
			Memory  int `json:"memory"`
			CPU     int `json:"cpu"`
			Disk    int `json:"disk"`
			IO      int `json:"io"`
			Network int `json:"network"`
		} `json:"limits"`
	} `json:"attributes"`
}

type ServerCommandSchema struct {
	Command string `json:"command"`
}

type ServerPowerActionSchema struct {
	Signal string `json:"signal"`
}

type ServerListDatabasesSchema struct {
	PteroRequestSchema
	Data []ServerDatabaseSchema `json:"data"`
}

type ServerCreateDatabaseSchema struct {
	Database string `json:"database"`
	Remote   string `json:"remote"`
}

type ServerDatabaseSchema struct {
	PteroRequestSchema
	Attributes struct {
		Id   string `json:"id"`
		Host struct {
			Address string `json:"address"`
			Port    int    `json:"port"`
		}
		Name            string `json:"name"`
		Username        string `json:"username"`
		ConnectionsFrom string `json:"connections_from"`
		MaxConnections  int    `json:"max_connections"`
	} `json:"attributes"`

	Relations struct {
		Password struct {
			PteroRequestSchema
			Attributes struct {
				Password string `json:"password"`
			} `json:"attributes"`
		} `json:"password"`
	} `json:"relationships"`
}

type CronSchema struct {
	DayOfWeek  string `json:"day_of_week"`
	DayOfMonth string `json:"day_of_month"`
	Hour       string `json:"hour"`
	Minute     string `json:"minute"`
}

type ServerScheduleTaskSchema struct {
	PteroRequestSchema
	Attributes struct {
		Id         int    `json:"id"`
		SequenceId int    `json:"sequence_id"`
		Action     string `json:"action"`
		Payload    string `json:"payload"`
		TimeOffset int    `json:"time_offset"`
		IsQueued   bool   `json:"is_queued"`
		CreatedAt  string `json:"created_at"`
		UpdatedAt  string `json:"updated_at"`
	} `json:"attributes"`
}

type ServerScheduleSchema struct {
	PteroRequestSchema
	Attributes struct {
		Id   int        `json:"id"`
		Name string     `json:"name"`
		Cron CronSchema `json:"cron"`
	} `json:"attributes"`

	Active     bool   `json:"is_active"`
	Processing bool   `json:"is_processing"`
	LastRunAt  string `json:"last_run_at"`
	NextRunAt  string `json:"next_run_at"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	Relations  struct {
		Tasks struct {
			PteroRequestSchema
			Data []ServerScheduleTaskSchema `json:"data"`
		} `json:"tasks"`
	} `json:"relationships"`
}

type ServerAllocationsSchema struct {
	PteroRequestSchema
	Data []AllocationSchema `json:"data"`
}

type AllocationSchema struct {
	PteroRequestSchema
	Attributes struct {
		Id        int    `json:"id"`
		IP        string `json:"ip"`
		Alias     string `json:"ip_alias"`
		Port      int    `json:"port"`
		Noted     string `json:"notes"`
		IsDefault bool   `json:"is_default"`
	} `json:"attributes"`
}

type ServerSubuserSchema struct {
	PteroRequestSchema
	Attributes struct {
		UUID        string   `json:"uuid"`
		Username    string   `json:"username"`
		Email       string   `json:"email"`
		Image       string   `json:"image"`
		TwoFactor   bool     `json:"2fa_enblaed"`
		CreatedAt   string   `json:"created_at"`
		Permissions []string `json:"permissions"`
	} `json:"attributes"`
}

type ServerSubusersSchema struct {
	PteroRequestSchema
	Data []ServerSubuserSchema `json:"data"`
}

type ServerVariableSchema struct {
	PteroRequestSchema
	Attributes struct {
		Name         string `json:"name"`
		Description  string `json:"description"`
		EnvVariable  string `json:"env_variable"`
		IsEditable   bool   `json:"is_editable"`
		DefaultValue string `json:"default_value"`
		ServerValue  string `json:"server_value"`
		Rules        string `json:"rules"`
	} `json:"attributes"`
}

type ServerVariablesSchema struct {
	PteroRequestSchema
	Data []ServerVariableSchema `json:"data"`
}

type ClientPermissionSchema struct {
	Description string            `json:"description"`
	Keys        map[string]string `json:"keys"`
}

type ClientPermissionsSchema struct {
	PteroRequestSchema
	Permissions map[string]ClientPermissionSchema `json:"data"`
}
