package msf

func (msf *Metasploit) Login() error {
	ctx := &loginReq{
		Method:   "auth.login",
		Username: msf.user,
		Password: msf.pass,
	}

	var res loginRes
	if err := msf.send(ctx, &res); err != nil {
		return err
	}

	msf.token = res.Token
	return nil
}

func (msf *Metasploit) Logout() error {
	ctx := &logoutReq{
		Method:      "auth.logout",
		Token:       msf.token,
		LogoutToken: msf.token,
	}

	var res logoutReq
	if err := msf.send(ctx, &res); err != nil {
		return err
	}

	return nil
}
