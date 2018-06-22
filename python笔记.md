# Python笔记  

  
###  2.1.1. Argument Passing 
`python3 -c 'import sys;print(sys.argv[0]);print(sys.argv[1]),print(sys.argv[2])' 1 2`
输出为:   

```python 
-c
1
2  
```  
可见argv[0]:-c argv[1]:1 argv[2]:2....
总之从argv[1]开始是传入并且由command或者module或者file处理的参数，而argv[0]则要看具体传递的形式，或者是“-c”本身或者是个文件名等等。    
### 2.2.1. Source Code Encoding 以及shell脚本
`# -*- coding: cp1252 -*-`  
  
或者
  
```  
#!/usr/bin/env python3
# -*- coding: cp1252 -*-  
```
第二个例子还说明了用python写shell脚本的方式  
### 3.1.1. Numbers 幂运算  
 
```python  
>>> 5 ** 2  # 5 squared
25
>>> 2 ** 7  # 2 to the power of 7
128  
```  
### 3.1.1. Numbers 变量'_' 
```python  
>>> tax = 12.5 / 100
>>> price = 100.50
>>> price *tax
12.5625
>>> price +_
113.0625
>>> round(_, 2)
113.06  
```
互动模式下，_代表上一条刚打印的表达式，这是个只读变量  
### 3.1.2. Strings escape以及raw string  
作为参数调用print时escape才起作用  
 
```python 
>>> '"Isn\'t," she said.'

'"Isn\'t," she said.'`

>>> print('"Isn\'t," she said.')

"Isn't," she said.

>>> s = 'First line.\nSecond line.'  # \n means newline
>>> s  # without print(), \n is included in the output

'First line.\nSecond line.'

>>> print(s)  # with print(), \n produces a new line

First line.
Second line.
```
raw string  

```python
>>> print(r'C:\some\name')  # note the r before the quote
C:\some\name
```  
### 3.1.2. Strings 多行输入以及\的作用  
一个string分成多行输入时，\可以取消编辑时输入的换行符  

```python
>>> print("""\
... Usage: thingy [OPTIONS]
...      -h                        Display this usage message
...      -H hostname               Hostname to connect to
... "”")
```
```python
Usage: thingy [OPTIONS]
     -h                        Display this usage message
     -H hostname               Hostname to connect to
```  
```python
>>> print("""
... Usage: thingy [OPTIONS]
...      -h                        Display this usage message
...      -H hostname               Hostname to connect to
... "”")
```
```python

Usage: thingy [OPTIONS]
     -h                        Display this usage message
     -H hostname               Hostname to connect to
```  
```python
>>> print("""\
... Usage: thingy [OPTIONS]\
...      -h                        Display this usage message\
...      -H hostname               Hostname to connect to\
...  "””)
```
```python
Usage: thingy [OPTIONS]     -h                        Display this usage message     -H hostname               Hostname to connect to
```  
### 3.1.2. Strings index以及slice  
```python
>>> text='python'
>>> text[3]
‘h'  
```
是可以的  

`‘python'[0]`  

是不允许的  

```python  
|P |y |t |h |o |n |
 0  1  2  3  4  5  6
-6 -5 -4 -3 -2 -1
```  
所以  

```python  
text[0]==text[-6]==‘P’  
text[-1]=text[5]=’n’
```  

因为-0==0，负索引从-1开始，也就是说最后一个字符索引是-1,第一个字符是-len(string)

slice的索引范围是个半开区间，也就是[b,e)这样的区间`text[-2:]==‘on’==text[4:]`  

```python  
>>> text[4:30]
'on'
>>> text[6:]
''
```
### 3.1.2. Strings 连接  

```python
>>> 'un'*3
'ununun'
>>> 3*'un'
'ununun'  
```
下面连接方式只适用于字符串文本，不能用于连接变量和表达式，要连接变量或者表达式时要用+号

```python
>>> 'py''thon'
'python'
>>> text=('py'
... 'thon')
>>> text
'python'
```  
### 3.1.3. Lists  
list是mutable的，string是immutable的  

```python
>>> cubes =[1, 8, 27, 65, 125]  # something's wrong here
>>> cubes[3] = 64  # replace the wrong value
>>> cubes
[1, 8, 27, 64, 125]  
```
还可以append  

```python
>>> cubes.append(216)  # add the cube of 6
>>> cubes.append(7 ** 3)  # and the cube of 7
>>> cubes
[1, 8, 27, 64, 125, 216, 343]  
```  
*******  
```python
>>> letters =['a', 'b', 'c', 'd', 'e', 'f', 'g']
>>> letters
['a', 'b', 'c', 'd', 'e', 'f', 'g']  
```
*******
可以replace一些值  

```python
>>> letters[2:5] =['C', 'D', 'E']
>>> letters
['a', 'b', 'C', 'D', 'E', 'f', 'g']  
```  

可以remove一些值  

```python
>>> letters[2:5] =[]
>>> letters
['a', 'b', 'f', 'g']  
```  

可以清空一个list  

```python
>>> letters[:] =[]
>>> letters
[]  
```
可以嵌套  

```python
>>> a =['a', 'b', 'c']
>>> n =[1, 2, 3]
>>> x =[a,n]
>>> x
[['a', 'b', 'c'], [1, 2, 3]]
>>> x[0]
['a', 'b', 'c']
>>> x[0][1]
'b'   
```  
### 4.2. for Statements  
```python
>>> for w in words:
...     print(w, len(w))
...  
cat 3
window 6
defenestrate 12
```
如果要在迭代中修改words，则需要：  

```python
>>> for w in words[:]:  # Loop over a slice copy of the entire list.
...     if len(w)>6:
...             words.insert(0, w)
... 
>>> words
['defenestrate', 'cat', 'window', 'defenestrate']
```
words[:]是words的一个copy，而不是words本身，所以在words中insert并不会改变这个copy，否则：  

```python
>>> for w in words:  # Loop over a slice copy of the entire list.
...     if len(w)>6:
...             words.insert(0, w)
```
会无限循环下去，不停的在前面插入‘defenestrate'
### 4.3. The range() Function  
```python
>>> print(range(10))
range(0, 10)
>>> x=range(10)
>>> print(x)
range(0, 10
```
可见range返回一个可迭代的对象，但不是list本身：
可以用for循环迭代，也可以使用list()函数生成一个list  

```python
>>> list(range(5))
[0, 1, 2, 3, 4]
```
### 4.4. break and continue Statements, and else Clauses on Loops  

```python
>>> forn in range(2, 10):
…     for x in range(2,n):
…         if n % x == 0:
...             print(n, 'equals',x, '*',n//x)
...             break
...       else:
...         # loop fell through without finding a factor
...             print(n, 'is a prime number')
...
2 is a prime number
3 is a prime number
4 equals 2 * 2
5 is a prime number
6 equals 2 * 3
7 is a prime number
8 equals 2 * 4
9 equals 3 * 3
```
break和continue与c语言或者其他什么语言的语义是一样的，而这里的else会在内层循环正常迭代结束之后执行，但是break出来之后不会执行else  
### 4.6. Defining Functions None以及函数对象  
任何一个函数都会有一个返回值，即使没有return语句，也会返回一个None。  
函数是一个对象，所以可以  

```python
f=func
f(100)
```
也可以  

```python
f100=func(100)
f100
```  
### 4.7.1. Default Argument Values 及in和is None  

```python
def ask_ok(prompt,retries=4,reminder='Please try again!'):
    while True:
        ok = input(prompt)
        if ok in('y', 'ye', 'yes'):
            return True
        if ok in ('n', 'no', 'nop', 'nope'):
            return False
        retries =retries - 1
        if retries < 0:
            raise ValueError('invalid user response')
        print(reminder)  

```
有三种调用方式  

1. `ask_ok('Do you really want to quit?')`
2. `ask_ok('OK to overwrite the file?', 2)  #retries=2`
3. `ask_ok('OK to overwrite the file?', 2, 'Come on, only yesor no!')`

in用来检查一个序列中是否包含某个值

缺省参数在函数定义的地方计算，只计算一次，而不是在调用时计算  

```python
>>> i=5
>>> def f(arg=i):
...     print(arg)
... 
>>> i=6
>>> f()
5  
```
在f的定义处，arg已经被初始化为5，之后与i已经没有关系了。
而  

```python
>>> def f(a, L=[]):
...     L.append(a)
...     return L
... 
>>> print(f(1))
[1]
>>> print(f(2))
[1, 2]
>>> print(f(3))
[1, 2, 3]  
```
在定义处初始化的这个list是共享的，这对于mutable的对象都是如此，比如dictionary和一些class。这也说明每次调用函数时，最初指向的都是同一个函数对象，而不是马上产生一个函数副本。  

如果不希望被共享，可以：  

```python
def f(a,L=None):
    if L is None:
        L =[]
        L.append(a)
    return  L  
```
另外还可以看到None是用is来比较，而不是==
### 4.7.2. Keyword Arguments  
keyword实参[^参数]指在调用函数时使用arg=value的形式传入的参数，而在函数定义形式参数列表中arg=value形式的参数arg称为可选参数，value称为缺省值  

~~~python
def parrot(voltage, state='a stiff', action='voom', type='Norwegian Blue'):
    print("-- This parrot wouldn't", action, end=' ')
    print("if you put", voltage, "volts through it.")
    print("-- Lovely plumage, the", type)
    print("-- It's", state, "!")  
~~~  
这个函数接受一个required实参(voltage),三个可选实参(state,action,type)。keyword参数出现的顺序任意的，毕竟已经提供参数名字了，还要求顺序对程序员来说显得略多余。  
下面这样调用都是非法的：  

~~~python
parrot()                     # 缺少参数
parrot(voltage=5.0, 'dead')  # 在keyword参数之后不允许再出现非keyword参数
parrot(110, voltage=220)     # voltage重复了
parrot(actor='John Cleese')  # actor未知参数
~~~  

[^参数]: 这里用参数还是实参犹豫了一阵子，虽然parameter和argument意义是不同的，前者指函数定义处的参数，也就是形式参数或者形参，后者指函数调用时的参数，也就是实际参数或者实参，不过在不会混淆的情况下还是会使用参数  
***
~~~python
def cheeseshop(kind, *arguments, **keywords):
    print("-- Do you have any", kind, "?")
    print("-- I'm sorry, we're all out of", kind)
    for arg in arguments:
        print(arg)
    print("-" * 40)
    for kw in keywords:
        print(kw, ":", keywords[kw])
~~~  
复习一下`print("-" * 40)`这种字符串生成方式，以及注意一下arguments和keywords在函数中的使用方式。  
这个函数可以这样调用  

~~~python
cheeseshop("Limburger", "It's very runny, sir.",
           "It's really very, VERY runny, sir.",
           shopkeeper="Michael Palin",
           client="John Cleese",
           sketch="Cheese Shop Sketch")  
~~~  
`*arguments`是个元组:  
`("It's very runny, sir.","It's really very, VERY runny, sir.")`
`**keywords"`则是个字典:  
`{"shopkeeper":"Michael Palin","client":"John Cleese",      "sketch":"Cheese Shop Sketch"}`  
***_注意并且理解参数列表中的顺序。_***
  
### 4.7.4. Unpacking Argument Lists `*`和`**`  

~~~python
>>> list(range(3, 6))            # normal call with separate arguments
[3, 4, 5]
>>> args = [3, 6]
>>> list(range(*args))            # call with arguments unpacked from a list
[3, 4, 5]  
~~~
args是个list,用元组也是可以的，但是args必须加上`*`号  

~~~python
>>> def parrot(voltage, state='a stiff', action='voom'):
...     print("-- This parrot wouldn't", action, end=' ')
...     print("if you put", voltage, "volts through it.", end=' ')
...     print("E's", state, "!")
...
>>> d = {"voltage": "four million", "state": "bleedin' demised", "action": "VOOM"}
>>> parrot(**d)
-- This parrot wouldn't VOOM if you put four million volts through it. E's bleedin' demised !  
~~~  
同样的，d也必须加上`**`号
### 4.7.5. Lambda Expressions  

```python
>>> def make_incrementor(n):
...     return lambda x: x + n
...
>>> f = make_incrementor(42)
>>> f(0)
42
>>> f(1)
43
```  
第4行生成了一个函数对象**f**:x->x+42[^函数]，接下来第5和第7行是对这个函数**f**的调用，而不是对函数make_incrementor的调用。  

[^函数]: 好吧，这个函数的写法是数学意义上的函数的写法，理解就行  

```python
>>> pairs = [(1, 'one'), (2, 'two'), (3, 'three'), (4, 'four')]
>>> pairs.sort(key=lambda pair: pair[1])
>>> pairs
[(4, 'four'), (1, 'one'), (3, 'three'), (2, 'two')]
```  

lamda表达式作为参数传入了sort函数，需要注意的是必须使用keyword实参的形式传入。  
**注意**：这里只是说明了在函数的语境中lamda表达式的使用：作为返回值或者参数。  
不过我们可以看另外一个例子，可能更好理解一些。

```python
>>> def printPairs(key):
...     pairs = [(1, 'one'), (2, 'two'), (3, 'three'), (4, 'four')]
...     for tup in pairs:
...             print(key(tup))

>>> printPairs(key=lambda pair: pair[1])
one
two
three
four
>>> printPairs(key=lambda pair: pair[0])
1
2
3
4
```  
### 4.7.6. Documentation Strings:`__doc__`  

```python
>>> def my_function():
...     """Do nothing, but document it.
...
...     No, really, it doesn't do anything.
...     """
...     pass
...
>>> print(my_function.__doc__)
Do nothing, but document it.

    No, really, it doesn't do anything.  

>>> def my_function():
...     """\
...     Do nothing, but document it.
...
...     No, really, it doesn't do anything.
...     """
...     pass
... 
>>> print(my_function.__doc__)
	Do nothing, but document it.
	
	No, really, it doesn't do anything.

```  
第一个例子是惯常的函数文档的写法，第二个只是复习一下相关内容。

### 4.7.7. Function Annotations: `__annotations__`  
```python
>>> def f(ham: str, eggs: str = 'eggs') -> str:
...     print("Annotations:", f.__annotations__)
...     print("Arguments:", ham, eggs)
...     return ham + ' and ' + eggs
...
>>> f('spam')
Annotations: {'ham': <class 'str'>, 'return': <class 'str'>, 'eggs': <class 'str'>}
Arguments: spam eggs
'spam and eggs'

>>> def f(ham, eggs= 'eggs'):
...     print("Annotations:", f.__annotations__)
...     print("Arguments:", ham, eggs)
...     return ham + ' and ' + eggs
... 
>>> f('spam')
Annotations: {}
Arguments: spam eggs
'spam and eggs'
```
似乎不需要解释什么了。
### 4.8. Intermezzo: Coding Style  
***`CamelCase` for classes and `lower_case_with_underscores` for functions and methods***
### 5.1. More on Lists  
所有list的方法都在这了：  

- `list.append(x)`  
* `list.extend(iterable)`
- `list.insert(i, x)`
- `list.remove(x)`
- `list.pop([i])`
- `list.clear()`
- `list.index(x[, start[, end]])`
- `list.count(x)`
- `list.sort(key=None, reverse=False)`
- `list.reverse()`
- `list.copy()`  

```python
>>> fruits = ['orange', 'apple', 'pear', 'banana', 'kiwi', 'apple', 'banana']
>>> fruits.count('apple')
2
>>> fruits.count('tangerine')
0
>>> fruits.index('banana')
3
>>> fruits.index('banana', 4)  # Find next banana starting a position 4
6
>>> fruits.reverse()
>>> fruits
['banana', 'apple', 'kiwi', 'banana', 'pear', 'apple', 'orange']
>>> fruits.append('grape')
>>> fruits
['banana', 'apple', 'kiwi', 'banana', 'pear', 'apple', 'orange', 'grape']
>>> fruits.sort()
>>> fruits
['apple', 'apple', 'banana', 'banana', 'grape', 'kiwi', 'orange', 'pear']
>>> fruits.pop()
'pear'
```
`insert`, `remove` 和 `sort`返回`None`
### 5.1.1. Using Lists as Stacks

```python
>>> stack = [3, 4, 5]
>>> stack.append(6)
>>> stack.append(7)
>>> stack
[3, 4, 5, 6, 7]
>>> stack.pop()
7
>>> stack
[3, 4, 5, 6]
>>> stack.pop()
6
>>> stack.pop()
5
>>> stack
[3, 4]
```
### 5.1.2. Using Lists as Queues
由于list存储方式的原因，直接使用list做为队列效率并不高，可以用 `collections.deque`:

```python
>>> from collections import deque
>>> queue = deque(["Eric", "John", "Michael"])
>>> queue.append("Terry")           # Terry arrives
>>> queue.append("Graham")          # Graham arrives
>>> queue.popleft()                 # The first to arrive now leaves
'Eric'
>>> queue.popleft()                 # The second to arrive now leaves
'John'
>>> queue                           # Remaining queue in order of arrival
deque(['Michael', 'Terry', 'Graham'])
```
### 5.1.3. List Comprehensions
不确定Comprehensions该怎么翻译，看例子：

```
[x**2 for x in range(10)]

[(x, y) for x in [1,2,3] for y in [3,1,4] if x != y]
```
一个list comprehension包括一对方括号`[]`,在方括号里边有一个表达式，比如上面两个例子里的`x**2`和`(x,y)`, 这个表达式后面跟着一个`for`语句, 然后是0个或者多个`for`或者`if`语句. 
上面两个例子分别等价于：

```python
>>> squares = []
>>> for x in range(10):
...     squares.append(x**2)
...
>>> squares
[0, 1, 4, 9, 16, 25, 36, 49, 64, 81]
```
和

```python
>>> combs = []
>>> for x in [1,2,3]:
...     for y in [3,1,4]:
...         if x != y:
...             combs.append((x, y))
...
>>> combs
[(1, 3), (1, 4), (2, 3), (2, 1), (2, 4), (3, 1), (3, 4)]
```
对于第一个例子还有等价的做法，不过这是另外一个问题了。

```python
squares = list(map(lambda x: x**2, range(10)))
```
更多的例子：

```python
>>> vec = [-4, -2, 0, 2, 4]
>>> # create a new list with the values doubled

>>> [x*2 for x in vec]
[-8, -4, 0, 4, 8]

>>> # filter the list to exclude negative numbers
>>> [x for x in vec if x >= 0]
[0, 2, 4]

>>> # apply a function to all the elements
>>> [abs(x) for x in vec]
[4, 2, 0, 2, 4]

>>> # call a method on each element
>>> freshfruit = ['  banana', '  loganberry ', 'passion fruit  ']
>>> [weapon.strip() for weapon in freshfruit]
['banana', 'loganberry', 'passion fruit']

>>> # create a list of 2-tuples like (number, square)
>>> [(x, x**2) for x in range(6)]
[(0, 0), (1, 1), (2, 4), (3, 9), (4, 16), (5, 25)]

>>> # the tuple must be parenthesized, otherwise an error is raised
>>> [x, x**2 for x in range(6)]
  File "<stdin>", line 1, in <module>
    [x, x**2 for x in range(6)]
               ^
SyntaxError: invalid syntax

>>> # flatten a list using a listcomp with two 'for'
>>> vec = [[1,2,3], [4,5,6], [7,8,9]]
>>> [num for elem in vec for num in elem]
[1, 2, 3, 4, 5, 6, 7, 8, 9]

>>> from math import pi
>>> [str(round(pi, i)) for i in range(1, 6)]
['3.1', '3.14', '3.142', '3.1416', '3.14159']
```
***注意第21和第24行两个例子的区别。***

### Nested List Comprehensions

```python
>>> matrix = [
...     [1, 2, 3, 4],
...     [5, 6, 7, 8],
...     [9, 10, 11, 12],
... ]
```
我们要把这个矩阵转置，不过为了好理解，首先看这个例子:

```python
>>> [row[i] for row in matrix for i in range(4)]
[1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12]
```
比较下：

```python
>>> [[row[i] for row in matrix] for i in range(4)]
[[1, 5, 9], [2, 6, 10], [3, 7, 11], [4, 8, 12]
```
第一个例子等价于：

```python
>>> transposed = []
>>> for row in matrix:
...     for i in range(4):
...         transposed.append(row[i])
...
```
相应的第二个例子就等价于:

```python
>>> transposed = []
>>> for i in range(4):
...     # the following 3 lines implement the nested listcomp
...     transposed_row = []
...     for row in matrix:
...         transposed_row.append(row[i])
...     transposed.append(transposed_row)
...
>>> transposed
[[1, 5, 9], [2, 6, 10], [3, 7, 11], [4, 8, 12]]
```
也就是

```python
>>> transposed = []
>>> for i in range(4):
...     transposed.append([row[i] for row in matrix])
...
>>> transposed
[[1, 5, 9], [2, 6, 10], [3, 7, 11], [4, 8, 12]]
```
当然实现矩阵转置可以调用内置函数`zip`:

```python
>>> list(zip(*matrix))
[(1, 5, 9), (2, 6, 10), (3, 7, 11), (4, 8, 12)]
```
