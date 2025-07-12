package msf

func (msf *Metasploit) SessionList() (map[uint32]SessionListRes, error) {
	req := &sessionListReq{
		Method: "session.list",
		Token:  msf.token,
	}
	res := make(map[uint32]SessionListRes)

	if err := msf.send(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
