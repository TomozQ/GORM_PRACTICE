# GORM

## db.Find(&モデル配列)
モデル構造体の配列をポインタで渡す。
検索されたレコードが指定のモデルの構造体に変換され、配列にまとめられて渡される。
→Postのリストを取得... 
```
var pl []my.Post
db.Find(&pl)
```

一つのレコードのみ取り出したい
```
var pt my.Post
db.First(&pt)
```

### 検索条件の設定
## db.Where( 条件, 値1, ......)
第一引数に条件となる指揮などをstringにまとめたものを指定
→db.Where("account = ?", ac).First(&user)
?にacが代入される。

### 並べ替え
## db.Order()

### オフセットと最大取得数
## db.Offset( 開始数 ).Limit( 最大数 )

### 関連モデルを取り出す
## db.Model( &モデル ).Related( &モデル )
```
var grp my.Group
var pts []my.Post
db.where("id = ?", gid).First(&grp)
db.Order("created_at desc").Model(&grp).Related(&pts)
```
id=gidのGroupに関連づけられている（つまりそのGroupに投稿されている）Postを全て取り出すことができる。

Table, Select, JoinsによるJOINの実行
db.Joins( 連結の設定 )

## db.Table( テーブル ).Select( 項目の設定 )
Table→検索するテーブルの名前をstringで指定
Select→値を取り出す項目名をstringで全て指定。Joinsで連結するテーブルにある項目も用意する必要がある。

ex)
Postに紐付けられたCommentを取得する
```
var pst my.Post
var cmts []my.CommentJoin

db.Where("id = ?", pid).First(&pst)
db.Table("comments").Select("comments.*, users.id, users.name).Joins("join users on users.id + comments.user_id").Where("comments.post_id = ?", pid).Order("created_at desc").Find(&cmts)
```

CommentJoinは？
```
type CommentJoin struct {
  Comment
  User
  Post
}
```
CommentにUserやPostの情報を追加したもの<br>
### 「○○Join」という名前のモデルは、複数のモデルをひとまとめにしたもの

```
.Table("comments").Select("comments.*, users.id, users.name)
```
取り出すのはCommentsテーブル
取り出す項目→comments.*, users.id, users.name

### レコードの新規作成
## db.Create( &モデル )
```
// Post構造体の値を作成
pt := my.Post{
  UserId: int(user.Model.ID),
  Address: ad,
  Message: rq.PostFormValue("message),
}

// ポインタを引数に指定してCreate
db.Create(&pt)
```
