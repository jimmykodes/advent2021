package main

func p1(d *data) int {
	var winners []*board
	for {
		_, winners = d.DoRound()
		if len(winners) > 0 {
			return winners[0].score()
		}
	}
}