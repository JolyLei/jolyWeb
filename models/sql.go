package main
import(
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"strings"
	"fmt"
)
/*
CREATE TABLE `book` (
`idbook` int(11) NOT NULL,
`name` varchar(45) DEFAULT NULL,
`language` varchar(45) DEFAULT NULL,
`describe` varchar(45) DEFAULT NULL,
PRIMARY KEY (`idbook`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
*/

//数据库的配置
const(
	userName = "root"
	password = "didisql"
	ip = "127.0.0.1"
	port = "3306"
	dbName = "sql_lesson"
)


//Db数据库连接池
var DB *sql.DB

func InitDB(){
	//sql.Open("mysql", "root:didisql@tcp(127.0.0.1:3306)/test?parseTime=true")
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"},"")

	//打开数据库
	DB, _ = sql.Open("mysql", path)
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil{
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connect success")

}

func Insert() (bool){
	//开启事务

	tx, err := DB.Begin()
	if err != nil{
		fmt.Println("tx fail")
		return false
	}
	//准备sql语句
	stmt, err := tx.Prepare(`insert into book (language,name,idbook) values(?, ?, ?)`)
	if err != nil{
		fmt.Println("prepare fail")
		return false
	}

	//将参数传递到sql语句并且执行
	res, err := stmt.Exec("c", "kafak实战", 8)

	if err != nil{
		fmt.Println("exec fail")
		return false
	}
	//将事务提交
	tx.Commit()
	//获得上一个自增的id
	fmt.Println(res.LastInsertId())
	return true


/*
	res, err := DB.Exec("delete from book where name = 'kafak实战'")
	if err != nil{
		fmt.Println(err.Error())
	}
	if res != nil{
		fmt.Println(res.LastInsertId())
	}
	return true
	*/
}

func Delete() (bool) {

	//开启事务
	tx, err := DB.Begin()
	if err != nil{
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("DELETE FROM book WHERE idbook = ?")
	if err != nil{
		fmt.Println("Prepare fail")
		return false
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec(2)
	if err != nil{
		fmt.Println("Exec fail")
		return false
	}
	//提交事务
	tx.Commit()
	//获得上一个insert的id
	fmt.Println(res.LastInsertId())
	return true
}

func Update() (bool) {
	//开启事务
	tx, err := DB.Begin()
	if err != nil{
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("update book from book where name = `kafak`")
	if err != nil{
		fmt.Println("Prepare fail")
		return false
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec(`Java`, `from change`)
	if err != nil{
		fmt.Println("Exec fail")
		return false
	}
	//提交事务
	tx.Commit()
	fmt.Println(res.LastInsertId())
	return true
}
func main(){
	InitDB()
	if err := Update(); !err{
		fmt.Println("插入出错")
	}
	DB.Close()
}