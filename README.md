# a-tour-of-Go-plactice
[A Tour of Go](https://go-tour-jp.appspot.com/list)

Go version `1.14.3`

### メモ

- 数字区切り
```
5_000_000_000
```
- 変数
```
var n int = 100

var n int
n = 100

var n = 100

n := 100

var (
    n = 100
    m - 200
)
```
- 定数
```
const a int = 100
const b = 100
const s = "Hello,  " + "世界"

const (
    x = 100
    y = 200
)
```

- iota: 連続した定数を作るための仕組み
```
const (
    a = iota
    b
)

const (
    c = 1 << iota
    d
    e
)
fmt.Println(a, b, c, d, e)
// 0 1 1 2 4
```

- 算術演算子
```
+, -, /, *, %
```

- 代入演算子
```
=, :=, +=など, ++, --
```

- ビット演算
```
a | b   論理和
a & b   論理積
^a      否定
a ^b    排他的論理和 
a&^b    論理積の否定
1 << a  2のa乗
1 >> a  2の1/a乗
```

- 論理演算
```
a || b      または
a && b      かつ
|a          否定
```

- アドレス演算
```
&a   ポインタを取得
*p   ポインタが示す値を取得
```

- チャネル演算
```
ch<-100     チャネルへの送信
<-ch        チャネルへの受信
```