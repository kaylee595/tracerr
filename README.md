# 相对官方版本做了一点新增
至少需要go1.18及以上版本, 因为加了泛型.

## 新增内容
- `Wrap1()`用于一些函数返回多值的场景当中, 例如下面这个例子.
```golang
func ParseTime(layout, v string) (time.Time, error) {
	return Wrap1(time.Parse(layout, v))
}
```