package msf

func (msf *Metasploit) SessionList() (SessionListRes, error) {
	req := &sessionListReq{
		Method: SessionList,
		Token:  msf.token,
	}
	res := make(SessionListRes)

	if err := msf.send(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (msf *Metasploit) SessionStop(SessionID string) (bool, error) {
	req := &sessionStopReq{Method: SessionStop, Token: msf.token, SessionID: SessionID}
	var res sessionStopRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result == "success", nil
}
