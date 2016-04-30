package DB

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/larspensjo/config"
)

var (
	dbfile = flag.String("dbfile", "./config/db.ini", "DB Config file")
)

var m_db *sql.DB
var m_usr []Usr
var m_curtains []Curtains

//usr data struct
type Usr struct {
	Id      int
	Phone   string
	Pwd     string
	Mail    string
	Addr    string
	Money   float32
	Coupid0 int
	Coupid1 int
	Coupid2 int
	Lock    int
	Guid    string
}

// curtains data struct
type Curtains struct {
	Id      int
	Usid    string
	Cuid    string
	Inadd   string
	Intel   string
	Intime  float64
	Inprice float32
	Inwant  string
	Instate int
	Inarea  int
	Wid     int
	Instus  string
	Ccate   int
	Cunu    int
	Scope   int
}

//------------------sub operate function-------------------------------------
// func GetUsrByUsrname(usrname string, pas string) int {

// }

//------------------sub operate function-------------------------------------

func ReadCurtains() error {
	m_curtains = make([]Curtains, 0)
	fmt.Println("ReadCurtains Begin")
	rows, err := m_db.Query("select * from curtains where Id>?", -1)
	defer rows.Close()
	if err != nil {
		return err
	}
	myusr := new(Curtains)
	for rows.Next() {
		fmt.Println("rows -----------")
		rows.Columns()
		err = rows.Scan(&myusr.Id, &myusr.Usid, &myusr.Cuid, &myusr.Inadd, &myusr.Intel, &myusr.Intime, &myusr.Inprice, &myusr.Inwant,
			&myusr.Instate, &myusr.Inarea, &myusr.Wid, &myusr.Instus, &myusr.Ccate, &myusr.Cunu, &myusr.Scope)
		if err != nil {
			fmt.Println("rows.Scan error =", err)
			return err
		}
		m_curtains = append(m_curtains, *myusr)
	}
	for i := 0; i < len(m_curtains); i++ {
		fmt.Println("m_curtains = ", m_curtains[i].Id)
	}
	return nil
}

func ReadUsr() error {
	m_usr = make([]Usr, 0)
	fmt.Println("ReadUsr Begin")
	rows, err := m_db.Query("select * from usr where Id>?", -1)
	defer rows.Close()
	if err != nil {
		return err
	}
	myusr := new(Usr)
	for rows.Next() {
		fmt.Println("rows -----------")
		rows.Columns()
		err = rows.Scan(&myusr.Id, &myusr.Phone, &myusr.Pwd, &myusr.Mail, &myusr.Addr, &myusr.Money, &myusr.Coupid0, &myusr.Coupid1, &myusr.Coupid2, &myusr.Lock)
		if err != nil {
			fmt.Println("rows.Scan error =", err)
			return err
		}
		m_usr = append(m_usr, *myusr)
	}
	return nil
}

func ConnectDB() error {
	//var myerr error
	fmt.Println("connect DB ..... begin")
	ip, port, usr, pas, dbname, err := readDBConfig()
	connectstring := usr + ":" + pas + "@tcp(" + ip + ":" + port + ")" + "/" + dbname + "?charset=utf8"
	fmt.Println("connectstring = ", connectstring)
	m_db, err = sql.Open("mysql", "luliyuan:123456@tcp(192.168.1.252:3306)/anzhuang")
	if err != nil {
		fmt.Println("sql.Open = ", err.Error())
		return err
	}
	m_db.SetMaxOpenConns(2000)
	m_db.SetMaxIdleConns(1000)
	err = m_db.Ping()
	if err != nil {
		fmt.Println("m_db.Ping = ", err)
	}
	//read usr table
	ReadUsr()
	ReadCurtains()
	return nil
}

func readDBConfig() (string, string, string, string, string, error) {
	var TOPIC = make(map[string]string)
	//server config
	cfg, err := config.ReadDefault(*dbfile)
	if err != nil {
		fmt.Println("config.ReadDefault error")
	}
	//
	//Initialized topic from the configuration
	if cfg.HasSection("db") {
		section, err := cfg.SectionOptions("db")
		if err == nil {
			for _, v := range section {
				options, err := cfg.String("db", v)
				if err == nil {
					TOPIC[v] = options
					fmt.Println("DB TOPIC[v] = ", TOPIC[v])
				}
			}
		}
	}

	if len(TOPIC) != 5 {
		return "", "", "", "", "", errors.New("par wrong")
	}
	fmt.Println("DB len(TOPIC) = ", len(TOPIC))
	return TOPIC["ip"], TOPIC["port"], TOPIC["usr"], TOPIC["pas"], TOPIC["dbname"], nil
}
