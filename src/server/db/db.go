package db

type DBHandler struct {
	DB     *sql.DB
	Stream chan *model.User
}

func main() {}
