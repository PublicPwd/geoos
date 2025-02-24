# Geoos
我们的组织`spatial-go`正式成立，这是我们的第一个开源项目`Geoos`,`Geoos`提供有关空间数据和几何算法,使用`Go`语言包装实现。
欢迎大家使用并提出宝贵意见！

## 内容列表
  - [内容列表](#内容列表)
  - [目录结构](#目录结构)
  - [使用说明](#使用说明)
  - [维护者](#维护者)
  - [如何贡献](#如何贡献)
  - [使用许可](#使用许可)



## 目录结构
1. `algorithm` 是对外暴露的空间运算方法定义。
2. `strategy.go` 定义了空间运算底层算法的选择实现。

## 使用说明
以计算面积`Area`为例。
```
package main

import (
	"bytes"
	"fmt"

	"github.com/spatial-go/geoos/geoencoding"
	"github.com/spatial-go/geoos/planar"
)

func main() {
	// First, choose the default algorithm.
	strategy := planar.NormalStrategy()
	// Secondly, manufacturing test data and convert it to geometry
	const polygon = `POLYGON((-1 -1, 1 -1, 1 1, -1 1, -1 -1))`
	// geometry, _ := wkt.UnmarshalString(polygon)

	buf0 := new(bytes.Buffer)
	buf0.Write([]byte(polygon))
	geometry, _ := geoencoding.Read(buf0, geoencoding.WKT)

	// Last， call the Area () method and get result.
	area, e := strategy.Area(geometry)
	if e != nil {
		fmt.Printf(e.Error())
	}
	fmt.Printf("%f", area)
	// get result 4.0
}
```
Example: geoencoding
[example_encoding.go](https://github.com/spatial-go/geoos/example/example_encoding.go)

## 维护者

[@spatial-go](https://github.com/spatial-go)。


## 如何贡献

我们也将秉承“开放、共创、共赢”的目标理念在空间计算领域贡献自己的一份力量。

非常欢迎你的加入！[提一个 Issue](https://github.com/spatial-go/geoos/issues/new)

联系邮箱： [geoos@changjing.ai](mailto:geoos@changjing.ai)

## 使用许可

[LGPL-2.1 ](LICENSE)