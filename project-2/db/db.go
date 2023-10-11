package db

type Conn struct{
	db string
}

func NewConn(db string)(Conn,bool){
	if db!=""{
	return Conn{

		db:db,

	},true
	}
	return Conn{},false
}