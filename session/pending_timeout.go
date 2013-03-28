package session

type pendingTimeout struct {
	inSession
}

func (currentState pendingTimeout) Timeout(session *session, event event) (nextState state) {
	switch event {
	case peerTimeout:
		session.log.OnEvent("Session Timeout")
		return latentState{}
	}

	return currentState
}
