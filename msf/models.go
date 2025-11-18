package msf

import "bytes"

// login req stuff
type loginReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Username string
	Password string
}

// login res stuff
type loginRes struct {
	Result       string `msgpack:"result"`
	Token        string `msgpack:"token"`
	Error        bool   `msgpack:"error"`
	ErrorClass   string `msgpack:"error_class"`
	ErrorMessage string `msgpack:"error_message"`
}

// logout req stuff
type logoutReq struct {
	_msgpack struct{} `msgpack:",asArray"`

	Method      string
	Token       string
	LogoutToken string
}

// logout res stuff
// NOTE: currently unused
type logoutRes struct {
	Result string `msgpack:"result"`
}

// session req stuff
type sessionListReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

// session res stuff
type (
	SessionListRes map[uint32]sessionInfo

	sessionInfo struct {
		ID          uint32 `msgpack:",omitempty"`
		Type        string `msgpack:"type"`
		TunnelLocal string `msgpack:"tunnel_local"`
		TunnelPeer  string `msgpack:"tunnel_peer"`
		ViaExploit  string `msgpack:"via_exploit"`
		ViaPayload  string `msgpack:"via_payload"`
		Description string `msgpack:"desc"`
		Info        string `msgpack:"info"`
		Workspace   string `msgpack:"workspace"`
		SessionHost string `msgpack:"session_host"`
		SessionPort int    `msgpack:"session_port"`
		Username    string `msgpack:"username"`
		UUID        string `msgpack:"uuid"`
		ExploitUUID string `msgpack:"exploit_uuid"`
	}
)

// session stop req
type sessionStopReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID string
}

type sessionStopRes struct {
	Result string `msgpack:"result"`
}

type ErrorRes struct {
	Error        bool   `msgpack:"error,omitempty"`
	ErrorClass   string `msgpack:"error_class,omitempty"`
	ErrorMessage string `msgpack:"error_message,omitempty"`
}

type Metasploit struct {
	host  string
	user  string
	pass  string
	token string
	buf   *bytes.Buffer
}

func New(host, user, pass string) (*Metasploit, error) {
	msf := &Metasploit{
		host: host,
		user: user,
		pass: pass,
		buf:  new(bytes.Buffer),
	}

	if err := msf.Login(); err != nil {
		return nil, err
	}

	return msf, nil
}
