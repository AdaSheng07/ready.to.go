## Map

`Map`是一个`key-value`组合的结构体，一种无序的键值对的集合，常用实现方式是二叉树和哈希表（散列表）。

`Map`通过`key`来快速检索数据，`key`的作用类似于数组与切片中的索引，指向数据的值`value`。`key`总是唯一的，相同的`key`拿到相同的`value`，更新`value`时，会覆盖相同`key`的原有`value`值。

### 🔸 使用`Map`的优势：快速查找，从`key`定位到`value`

当数据量很大时，e.g. >10000，仍然用数组或者切片来存储，会出现什么问题？

如果我们需要查找数据集合中的某一组数据，需要`for-loop`遍历整个数组/切片，再比较值是否相等来锁定目标数据，时间复杂度是`O(n)`，很大。

如果在总共` 10000 `人中找第` 5000 `个人` 2000 `次，一共在切片上比较`5000 * 2000`次，效率太低。

`Map`是为了解决这样的问题而存在的，只要数据组中的每一组数据都是是唯一的，我们就可以用`key-value`构造`Map`来进行存储。

![image](https://github.com/AdaSheng07/ready.to.go/blob/6e686eaf7221d395a294a18280410052cbb455a9/pics/map.png)

### 🔸 `Map`的定义与初始化 [link](https://github.com/AdaSheng07/ready.to.go/blob/0bc031aec6339e4f13d7ab4705546030a8ec0dc6/chapter1/008.map1/main.go)

定义`map`时注意`key`与`value`的类型都需要声明。
`map`定义可以是符合类型，它的`key`与`value`都可以是任意类型，e.g. `int`,`float64`, `array`, `slice`, etc.

主要有两种定义方式：
```
  1. use keyword `map` to declare and initialize
  var map_variable_name map[key_data_type]value_data_type
  map_variable_name := map[key_data_type]value_data_type{}
  map_variable_name := map[key_data_type]value_data_type{key1: value1, key2: value2, key3: value3, ...}
  
  2. use keyword `map` to declare, then use built-in function make() to initialize
  var map_variable_name map[key_data_type]value_data_type
  map_variable_name = make(map[key_data_type]value_data_type, length)
  
  3. use built-in function make() to declare, then add in key-value elements
  map_variable_name := make(map[key_data_type]value_data_type, length)
  map_variable_name[key1] = value1
  map_variable_name[key2] = value2
  ...
 ```
`Map`定义的`key`与`value`类型可以**嵌套使用**，但要注意多层嵌套后代码意义是否会混淆，影响可读性，如：
```
  map_variable_name := map[string]map[string]map[int]float64{}
```

### 🔸 对`map`的操作 [link](https://github.com/AdaSheng07/ready.to.go/blob/0bc031aec6339e4f13d7ab4705546030a8ec0dc6/chapter1/008.map2/main.go)

如果不初始化/实例化`map`，就会默认初始化为`nil map`。`Map`不用实例化就可以读取和删除，但`nil map`**不能**用来写入/存放键值对：
```
  panic: assignment to entry in nil map
```

**`Map`的增删改查**

`Map`属于`Golang`的变量范畴，也是强类型的。当定义好`map`后，它只能容纳对应类型的数据。

>【回顾】如何在切片中间插入一个元素?
>
> 备份切片扩容 >> 拷贝备份 >> 用`append`插入值 >> 用`append`结合备份与追加元素后的新切片

此操作的**风险高，代价大**，而`Map`的增删改查是极其方便的：

```
  1. add/change key-value
  map_variable_name[key_name] = key_value
  * If this key does not exist in map, add this key-value pair in map; 
    else, change its value in the map.
  
  2. delete key-value
  delete(map_variable_name, key)
  * We can delete the same key repeatedly from map.
  
  3. lookup key-value：take key as index of map
  map_variable_name[key]
```
【注意】如果此时`map`中不存在这个`key`，会自动加入此`key`，但返回的`value`是**假值**。

> 如何判断验证真假值（`key`在`map`中是否存在）呢？
> ```
> value, ok := map_variable_name[key]
> * If ok is true, value is a true value;
>   else if ok is false, value is a false value.
> ```


**`Map`的遍历**
```
  for key, value := range map_variable_name {
      fmt.Printf("%v\t%v\n", key, value)
  }
  
  for key := range map_variable_name{
      fmt.Println(map_variable_name[key])
  }
```

**`Map`的合并**
```
  map_variable_name1, map_variable_name2 := map[key_type]value_type, map[key_type]value_type
  map_variable_name1[key1] = value1
  map_variable_name2[key2] = value2
  for k, v := range map_variable_name1{
      map_variable_name2[k] = v
  }
```
