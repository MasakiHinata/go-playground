# Golang MySql
## データベースと接続

```go
import (
	_ "github.com/go-sql-driver/mysql" // インポート忘れ注意！
	"github.com/jinzhu/gorm"
)

const (
	DBMS     = "mysql"
	USER     = "root"
	PASS     = "mysql"
	PROTOCOL = "tcp(localhost:3306)"
	DBNAME   = "user_schema"
)

connect := fmt.Sprintf("%s:%s@%s/%s", USER, PASS, PROTOCOL, DBNAME)
db, err := gorm.Open(DBMS, connect)
```

## データ
```go
type User struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"size:32"`
	Age  int
}
```

### gorm.Model

DBのエンティティにあった要素が追加される

```go
type User struct {
	gorm.Model
	...
}
```

```go
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
```

## データ操作
### INSERT

```go
// アドレスを渡す
err := db.Create(&user).Error
```

### DELETE

```go
// 値を渡す
db.Delete(*user).Error
```

### SELECT

```go
var users []*model.User
err := db.Find(&users).Error
```

### WHERE

```go
var user model.User
err := rep.db.Where("id = ?", id).Find(&user).Error
```

### UPDATE

```go
err := db.Model(&model.User{}).Where("id = ?", id).Update("name", name).Error
```