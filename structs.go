package provider

import (
)


type (
	Credentials map[string]string
      // {
      //   id:                'aws',
      //   name:              'Amazon Web Services',
      //   server_nick_name:  'EC2 Instance',
      //   default_region:    'us-east-1',
      //   default_plan:      'general_purpose',
      //   default_size:      't2.nano',
      //   can_reboot:        true,
      //   can_rename:        true,
      //   ssh_user:          'ubuntu',
      //   external_iface:    nil,
      //   internal_iface:    'eth0',
      //   credential_fields: [
      //     { key: :access_key_id, label: 'Access Key ID' },
      //     { key: :secret_access_key, label: 'Secret Access Key' }
      //   ],
      //   instructions:      instructions,
      //   bootstrap_script:  'https://s3.amazonaws.com/tools.nanobox.io/bootstrap/ubuntu.sh'
      // }
	Metadata struct {
		ID string `json:"id"`
		Name string `json:"name"`
		NickName string `json:"server_nick_name"`
		DefaultRegion string `json:"default_region"`
		DefaultPlan string `json:"default_plan"`
		DefaultSize string `json:"default_size"`
		Rebootable bool `json:"can_reboot"`
		Renamable bool `json:"can_rename"`
		SSHAuthMethod string `json:"ssh_auth_method"` // key or password
		SSHKeyMethod string `json:"ssh_key_method"` // reference or object
		SSHUser string `json:"ssh_user"`
		ExternalInterface string `json:"external_iface"`
		InternalInterface string `json:"internal_iface"`
		CredentialFields []CredentialField `json:"credential_fields"`
		Instructions string `json:"instructions"`
		BootstrapScript string `json:"bootstrap_script"`
	}

	CredentialField struct {
		Key string `json:"key"`
		Label string `json:"label"`
	}

	// [
	//   {
	//     "id":    "sfo1",
	//     "name":  "San Francisco 1",
	//     "plans": [
	//       {
	//         "id": "standard",
	//         "name": "Standard Configuration",
	//         "specs": [
	//           {"id": "512mb", "ram": 512, "cpu": 1, "disk": 20, "transfer": 1.0, "dollars_per_hr": 0.00744, "dollars_per_mo": 5.0},
	//           {"id": "1gb", "ram": 1024, "cpu": 1, "disk": 30, "transfer": 2.0, "dollars_per_hr": 0.01488, "dollars_per_mo": 10.0}
	//         ]
	//       }
	//     ]
	//   }
	// ]

	ServerOption struct {
		ID string `json:"id"`
		Name string `json:"name"`
		Plans []ServerPlan `json:"plans"`
	}

	ServerPlan struct {
		ID string `json:"id"`
		Name string `json:"name"`
		Specs []map[string]interface{} `json:"specs"`
	}

	// {
	//   "name": "nanobox-provider-account-ID",
	//   "key":  "CONTENTS OF PUBLIC KEY"
	// }
	KeyOrder struct {
		Name string `json:"name"`
		Key string `json:"key"`
	}

	// id: fingerprint or key identifier to use when ordering servers
	// name: the user-friendly name of the key.
	// public_key: "CONTENTS OF PUBLIC KEY"
	Key struct {
		ID string `json:"id"`
		Name string `json:"name"`
		PublicKey string `json:"public_key"`		
	}

	// {
	//   "name":    "nanobox.io-cool-app-do.1.1",
	//   "region":  "sfo1",
	//   "size":    "a1",
	//   "ssh_key": "12345"
	// }
	ServerOrder struct {
		Name string `json:"name"`
		Region string `json:"region"`
		Size string `json:"size"`
		SSHKey string `json:"ssh_key"`
		
	}

	// id: the server id
	// status: the status or availability of the server. (active indicates server is ready)
	// name: name of the server
	// external_ip: external or public IP of the server
	// internal_ip: internal or private IP of the server
	// password: the ssh password to use (if sshauthmethod is password)
	Server struct {
		ID string `json:"id"`
		Status string `json:"status"`
		Name string `json:"name"`
		ExternalIP string `json:"external_ip"`
		InternalIP string `json:"internal_ip"`
		Password string `json:"password,omitempty"`
	}
)

