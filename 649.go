func predictPartyVictory(senate string) string {
	rSenate := []int{}
	dSenate := []int{}
	for i, v := range senate {
		if v == 'R' {
			rSenate = append(rSenate, i)
		} else {
			dSenate = append(dSenate, i)
		}
	}
	for len(rSenate) > 0 && len(dSenate) > 0 {
		rv := rSenate[0]
		dv := dSenate[0]
		dSenate = dSenate[1:]
		rSenate = rSenate[1:]
		if rv < dv {
			rSenate = append(rSenate, rv+len(senate))
		} else {
			dSenate = append(dSenate, dv+len(senate))
		}
	}
	if len(rSenate) > 0 {
		return "Radiant"
	}
	return "Dire"
}