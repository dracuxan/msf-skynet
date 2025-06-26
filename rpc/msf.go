package rpc

type sessionListReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type SessionListReq struct {
	ID          uint32 `msgpack:",omitempty"`
	Type        string `msgpack:"type,omitempty"`
	TunnelLocal string `msgpack:"tunnel_local,omitempty"`
	TunnelPeer  string `msgpack:"tunnel_peer,omitempty"`
	ViaExploit  string `msgpack:"via_exploit,omitempty"`
	ViaPayload  string `msgpack:"via_payload,omitempty"`
	Description string `msgpack:"desc,omitempty"`
	Info        string `msgpack:"info,omitempty"`
	Workspace   string `msgpack:"workspace,omitempty"`
	SessionHost string `msgpack:"session_host,omitempty"`
	SessionPort int    `msgpack:"session_port,omitempty"`
	Username    string `msgpack:"username,omitempty"`
	UUID        string `msgpack:"uuid,omitempty"`
	ExploitUUID string `msgpack:"exploit_uuid,omitempty"`
}

type loginReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Username string
	Password string
}

type loginRes struct {
	Result       string `msgpack:"result"`
	Token        string `msgpack:"token"`
	Error        bool   `msgpack:"error"`
	ErrorClass   string `msgpack:"error_class,omitempty"`
	ErrorMessage string `msgpack:"error_message,omitempty"`
}

type logoutReq struct {
	_msgpack    struct{} `msgpack:",asArray"`
	Method      string
	Token       string
	LogoutToken string
}

type logoutRes struct {
	Result string `msgpack:"result"`
}

type Metasploit struct {
	host  string
	user  string
	pass  string
	token string
}

func New(host, user, pass string) *Metasploit {
	msf := &Metasploit{
		host: host,
		user: user,
		pass: pass,
	}

	return msf
}
