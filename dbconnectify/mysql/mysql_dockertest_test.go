package mysql

import (
	"testing"
)

func TestMysqlDockerTest(t *testing.T) {
	// dockerTest := ginfra.InitDockerTest()
	// defer dockerTest.CleanUp()
	//
	// mysqlDockerTestConf := MysqlDockerTestConf{}
	//
	// var db *sqlx.DB
	// var err error
	//
	// dockerTest.NewContainer(mysqlDockerTestConf.ImageVersion(dockerTest, true, ""), func(res *dockertest.Resource) error {
	// 	time.Sleep(10 * time.Second)
	//
	// 	db, err = mysqlDockerTestConf.ConnectSqlx(res)
	// 	gcommon.PanicIfError(err)
	// 	return nil
	// })
	//
	// sqlxCommander := gdb.NewSqlx(db)
	// sqlxTx := gdb.NewTxSqlx(db)
	//
	// _ = sqlxTx.DoTransaction(context.Background(), nil, func(c context.Context) (bool, error) {
	// 	_, err := sqlxCommander.Commander.ExecContext(c,
	// 		"CREATE TABLE IF NOT EXISTS users (id serial PRIMARY KEY, username VARCHAR ( 50 ) NOT NULL, password VARCHAR ( 50 ) NOT NULL, email VARCHAR ( 255 ) NOT NULL, created_on TIMESTAMP NOT NULL, last_login TIMESTAMP);")
	// 	gcommon.PanicIfError(err)
	// 	return false, nil
	// })
	//
	// wg := sync.WaitGroup{}
	//
	// for i := 0; i < 100; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		_ = sqlxTx.DoTransaction(context.Background(), nil, func(c context.Context) (bool, error) {
	// 			_, err := sqlxCommander.Commander.ExecContext(c, "INSERT INTO users (username, password, email, created_on, last_login) VALUES ('test', 'test', '', NOW(), NOW());")
	// 			gcommon.PanicIfError(err)
	// 			return true, errors.New("asd")
	// 		})
	// 		wg.Done()
	// 	}()
	// }
	//
	// wg.Wait()
	//
	// rows, err := sqlxCommander.Commander.QueryxContext(context.Background(), "SELECT * FROM users;")
	// fmt.Println(rows)
	// gcommon.PanicIfError(err)
	//
	// rows, err = sqlxCommander.Commander.QueryxContext(context.Background(), "SELECT COUNT(*) FROM users;")
	// gcommon.PanicIfError(err)
	// var count int
	//
	// for rows.Next() {
	// 	err = rows.Scan(&count)
	// 	gcommon.PanicIfError(err)
	//
	// 	break
	// }
	// fmt.Println(count)

}
