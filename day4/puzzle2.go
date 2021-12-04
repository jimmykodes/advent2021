package main

func p2(d *data) int {
	for {
		wIndexes, ws := d.DoRound()
		if len(ws) > 0 {
			if len(d.boards) == 1 {
				return ws[0].score()
			}
			for i := len(wIndexes) - 1; i >= 0; i-- {
				// iterate backwards
				d.removeBoard(wIndexes[i])
			}
		}
	}
}
