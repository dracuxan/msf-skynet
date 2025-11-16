package msf

func (msf *Metasploit) SessionList() (SessionListRes, error) {
	req := &sessionListReq{
		Method: "session.list",
		Token:  msf.token,
	}
	res := make(SessionListRes)

	if err := msf.send(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// func (msf *Metasploit) SessionStop(sessionID uint32) error {
// }
