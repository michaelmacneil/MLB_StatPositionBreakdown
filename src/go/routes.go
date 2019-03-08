package main

func (s *server) routes() {
  // Setup endpoints
	s.router.HandleFunc("/getStatisticsByPositionForPlayer", s.getStatisticsByPositionForPlayerRequest())
	s.router.HandleFunc("/getMLBTeams", s.getMLBTeamsRequest())
	s.router.HandleFunc("/getPlayersFromTeam", s.getPlayersFromTeamRequest())
}
