package cmd

func main() {
	// database connect
	var cfg model.Config
	loadConfig(&cfg)
	var sql = new(db.SQL)
	sql.Connect(cfg)
	defer sql.Close()
}
