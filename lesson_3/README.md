# Обработка ошибок в Go  


## 🔍 Основные концепции  

### 1. **Ошибки — это значения**  
В Go ошибки являются обычными значениями, а не исключениями. Они возвращаются из функций как второе значение.  

```go
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

### 2. **Проверка ошибок вручную**  
Ошибки нужно проверять сразу после вызова функции.  

```go
result, err := Divide(10, 0)
if err != nil {
    log.Fatal("Error:", err)
}
fmt.Println("Result:", result)
```

### 3. **Кастомные ошибки**  
Можно создавать свои типы ошибок, реализуя интерфейс `error`.  

```go
type MyError struct {
    Code    int
    Message string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}
```

### 4. **Обёртки ошибок (Go 1.13+)**  
Можно добавлять контекст с помощью `fmt.Errorf` и `%w`.  

```go
if err != nil {
    return fmt.Errorf("failed to process: %w", err)
}
```

- %w — указывает, что ошибка err должна быть "завёрнута" в новую ошибку.
- Это позволяет позже проверить оригинальную ошибку через errors.Is() или errors.As().


Зачем это нужно?

- Добавляет контекст к ошибке (например, "failed to process").
- Позволяет сохранить исходную ошибку для дальнейшего анализа.

Если просто передать ошибку в fmt.Errorf() без %w, то:

- Оригинальная ошибка потеряется — останется только её текст, но тип и структура исчезнут.
- errors.Is() и errors.As() не смогут проверить исходную ошибку.

### 5. **Сравнение ошибок**  
Проверка типа ошибки с помощью `errors.Is` и `errors.As`.  

```go
if errors.Is(err, os.ErrNotExist) {
    fmt.Println("File not found!")
}

var myErr *MyError
if errors.As(err, &myErr) {
    fmt.Println("Custom error:", myErr.Code)
}
```

- errors.Is() — проверка типа ошибки
Функция errors.Is(err, target) проверяет, содержится ли в цепочке ошибок err ошибка target.

Пример:

```go
if errors.Is(err, os.ErrNotExist) {
    fmt.Println("File not found!")
}
```

- Проверяет, была ли ошибка err (или любая из обёрнутых в ней ошибок) типа os.ErrNotExist.
- Работает даже с обёрнутыми ошибками (если использовался %w).

Раньше приходилось писать так:

```go
if err == os.ErrNotExist { ... }  // Не работало с обёрнутыми ошибками!
```

- errors.As() — приведение типа ошибки

Функция ```errors.As(err, &target)``` проверяет, можно ли привести ошибку ```err``` к определённому типу, и если да — сохраняет её в ```target```.

Пример:

```go
var myErr *MyError
if errors.As(err, &myErr) {
    fmt.Println("Custom error:", myErr.Code)
}
```

- Проверяет, является ли err (или любая из обёрнутых ошибок) типом *MyError.
- Если да — сохраняет её в переменную myErr (чтобы можно было обратиться к полям, например, myErr.Code).

**Что такое var myErr *MyError?**

- Это указатель на переменную типа MyError (кастомная ошибка, которую мы определили ранее).
- errors.As() требует передать указатель (&myErr), чтобы записать в него ошибку, если она подходит по типу.

Пример кастомной ошибки:

```go
type MyError struct {
    Code    int
    Message string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}
```

### 6. **Panic и Recover**  
Используются для критических ошибок (аналог исключений, но не для обычной логики).  

```go
func SafeCall() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
        }
    }()
    panic("Something went wrong!")
}
```

🌟 **Ключевой принцип в Go:** *"Не игнорируйте ошибки, проверяйте их явно!"*

### 🔗 Пример цепочки ошибок в Go (родитель и прародитель)  

Рассмотрим пример, где ошибки **вложены друг в друга** (`родитель -> дочерняя -> прародитель`), и как их правильно обрабатывать.  

#### 1. Создаём кастомные ошибки  
```go
import (
    "errors"
    "fmt"
)

// Кастомные ошибки
var (
    ErrConnectionFailed = errors.New("connection failed")
    ErrTimeout         = errors.New("timeout")
    ErrInvalidData     = errors.New("invalid data")
)
```

#### 2. Функция, которая возвращает обёрнутую ошибку  
```go
func fetchData() error {
    // Представим, что произошла ошибка соединения
    err := ErrConnectionFailed
    return fmt.Errorf("fetchData failed: %w", err)  // Оборачиваем ошибку
}

func process() error {
    err := fetchData()
    if err != nil {
        return fmt.Errorf("process failed: %w", err)  // Ещё один уровень обёртки
    }
    return nil
}
```

#### 3. Проверяем цепочку ошибок  
```go
func main() {
    err := process()

    // Проверяем, есть ли в цепочке ErrConnectionFailed
    if errors.Is(err, ErrConnectionFailed) {
        fmt.Println("Основная ошибка:", ErrConnectionFailed) // Основная ошибка: connection failed
    }

    // Разворачиваем и печатаем всю цепочку
    fmt.Println("Полная цепочка ошибок:", err) // process failed: fetchData failed: connection failed
}
```
**Вывод:**  
```
Основная ошибка: connection failed
Полная цепочка ошибок: process failed: fetchData failed: connection failed
```

---

### 🛠 Основные способы обработки ошибок в Go  

#### 1. **Проверка ошибок (`if err != nil`)**
```go
file, err := os.Open("data.txt")
if err != nil {
    log.Fatal("Ошибка:", err)  // Логируем и завершаем программу
}
defer file.Close()  // Закрываем файл при выходе
```

#### 2. **`errors.Is()` — проверка типа ошибки**
```go
if errors.Is(err, os.ErrNotExist) {
    fmt.Println("Файл не найден!")
}
```

#### 3. **`errors.As()` — приведение к кастомному типу**
```go
var myErr *MyError
if errors.As(err, &myErr) {
    fmt.Println("Код ошибки:", myErr.Code)
}
```

#### 4. **`fmt.Errorf("%w")` — обёртка ошибок**
```go
if err != nil {
    return fmt.Errorf("не удалось прочитать файл: %w", err)
}
```

#### 5. **`panic()` и `recover()` — аварийные ситуации**
```go
func safeCall() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Программа восстановлена:", r)
        }
    }()
    panic("что-то пошло не так!")
}
```

---

### 🌀 **`defer` — отложенные вызовы**  
`defer` гарантирует, что функция выполнится **перед выходом из текущей функции**, даже если произошла ошибка.  

#### Пример с `defer` для закрытия файла:
```go
func readFile() error {
    file, err := os.Open("data.txt")
    if err != nil {
        return err
    }
    defer file.Close()  // Закроет файл при выходе из readFile()

    // Читаем файл...
    return nil
}
```

#### Пример с `defer` + `recover()`:
```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Перехвачена паника:", r)
        }
    }()

    panic("всё сломалось!")
}
```

---

### 📌 **Сделаем небольшие выводы**  
- **Цепочки ошибок** создаются через `fmt.Errorf("%w")` и проверяются через `errors.Is()`/`errors.As()`.  
- **`defer`** помогает управлять ресурсами (файлы, сетевые соединения).  
- **`panic` и `recover`** — только для критических ошибок, а не для обычной логики.  
- **Ошибки в Go — это значения**, их нужно проверять явно.  

Такой подход делает код **надёжным и предсказуемым**, но требует внимательной обработки ошибок.


# `defer` в Go первый взгляд


## 1. Простейший пример с `defer`

`defer` — это способ сказать Go: "Выполни это действие в самом конце функции".

```go
func main() {
    defer fmt.Println("Это выполнится последним")  // (1)
    fmt.Println("Это выполнится первым")           // (2)
}
```

Вывод:
```
Это выполнится первым
Это выполнится последним
```

### Где это полезно?
- Закрытие файлов
- Освобождение ресурсов
- Завершение операций

## 2. Реальный пример: работа с файлом

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Открываем файл
    file, err := os.Open("test.txt")
    if err != nil {
        fmt.Println("Не удалось открыть файл:", err)
        return
    }
    
    // Гарантируем, что файл закроется при выходе из функции
    defer file.Close()
    
    // Работаем с файлом...
    fmt.Println("Файл успешно открыт!")
    
    // Здесь можно читать файл
    // file.Read(...)
}
```

### Что важно:
1. Сначала проверяем ошибку (`if err != nil`)
2. Потом ставим `defer` для закрытия файла
3. Только потом работаем с файлом

## 3. Как работает `defer` с несколькими вызовами

```go
func main() {
    defer fmt.Println("Первый defer")  // (3)
    defer fmt.Println("Второй defer")  // (2)
    defer fmt.Println("Третий defer")  // (1)
    
    fmt.Println("Основной код")        // (0)
}
```

Вывод:
```
Основной код
Третий defer
Второй defer
Первый defer
```

### Правило:
- `defer` работает по принципу стека (последний добавленный — первый выполненный)
- Это называется LIFO (Last In, First Out)
- Отложенные вызовы выполняются перед возвратом из функции
- Порядок выполнения — последний добавленный, первый выполненный (LIFO)
- Аргументы функции вычисляются в момент объявления defer

## 4. Частые ошибки новичков

### Ошибка 1: `defer` в цикле
```go
for i := 0; i < 5; i++ {
    file, _ := os.Open("file.txt")
    defer file.Close()  // Плохо! Все 5 defer'ов накопятся
}
```
Исправление:
```go
for i := 0; i < 5; i++ {
    func() {
        file, _ := os.Open("file.txt")
        defer file.Close()  // Хорошо — закроется в конце каждой итерации
    }()
}
```

### Ошибка 2: Игнорирование ошибок
```go
file, _ := os.Open("file.txt")  // Плохо! Игнорируем ошибку
```
Всегда проверяй ошибки!

## 6. Практические советы

1. Всегда проверяй ошибки (`if err != nil`)
2. Используйте `defer` для очистки ресурсов
3. Помните, что `defer` выполняется перед return
4. Не злоупотребляйте `defer` в циклах
5. Для сложных случаев используйте `defer` с анонимными функциями

## 7. Итоговый пример

```go
func copyFile(src, dst string) error {
    // Открываем исходный файл
    srcFile, err := os.Open(src)
    if err != nil {
        return fmt.Errorf("не удалось открыть файл %s: %v", src, err)
    }
    defer srcFile.Close()
    
    // Создаём новый файл
    dstFile, err := os.Create(dst)
    if err != nil {
        return fmt.Errorf("не удалось создать файл %s: %v", dst, err)
    }
    defer dstFile.Close()
    
    // Копируем содержимое
    _, err = io.Copy(dstFile, srcFile)
    if err != nil {
        return fmt.Errorf("ошибка копирования: %v", err)
    }
    
    return nil
}
```

Этот код:
1. Открывает файлы с проверкой ошибок
2. Гарантирует их закрытие через `defer`
3. Возвращает понятные сообщения об ошибках
4. Корректно обрабатывает все случаи


# Функции в Go

## 1. Основы функций

Функция в Go объявляется с помощью ключевого слова `func`:

```go
func sayHello() {
    fmt.Println("Привет!")
}
```

Вызов функции:
```go
sayHello() // Выведет: Привет!
```

## 2. Функции с параметрами

```go
func greet(name string) {
    fmt.Println("Привет,", name)
}

// Вызов
greet("Анна") // Выведет: Привет, Анна
```

## 3. Возвращаемые значения

Функции могут возвращать значения:

```go
func sum(a, b int) int {
    return a + b
}

result := sum(3, 5) // result = 8
```

### Возврат нескольких значений

В Go функции могут возвращать несколько значений:

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("деление на ноль")
    }
    return a / b, nil
}

result, err := divide(10, 2)
if err != nil {
    fmt.Println("Ошибка:", err)
} else {
    fmt.Println("Результат:", result)
}
```

## 4. Именованные возвращаемые значения

Можно задать имена возвращаемым значениям:

```go
func calc(a, b int) (sum int, diff int) {
    sum = a + b
    diff = a - b
    return // автоматически вернет sum и diff
}

s, d := calc(10, 5) // s=15, d=5
```

## 5. Анонимные функции (лямбда-функции)

Анонимные функции - это функции без имени:

```go
func main() {
    // Объявление и вызов анонимной функции
    func() {
        fmt.Println("Я анонимная функция!")
    }()
    
    // Сохранение анонимной функции в переменную
    greet := func(name string) {
        fmt.Println("Привет,", name)
    }
    
    greet("Мир") // Вызов через переменную
}
```

## 6. Функции как значения

Функции можно передавать как аргументы и возвращать из других функций:

```go
func applyOperation(a, b int, operation func(int, int) int) int {
    return operation(a, b)
}

func add(x, y int) int {
    return x + y
}

func multiply(x, y int) int {
    return x * y
}

func main() {
    result := applyOperation(3, 4, add)      // 7
    result2 := applyOperation(3, 4, multiply) // 12
}
```

## 7. Возврат функций из функций

Функции могут возвращать другие функции:

```go
func makeGreeter(prefix string) func(string) {
    return func(name string) {
        fmt.Println(prefix, name)
    }
}

func main() {
    russianGreet := makeGreeter("Привет")
    englishGreet := makeGreeter("Hello")
    
    russianGreet("Иван") // Выведет: Привет Иван
    englishGreet("John") // Выведет: Hello John
}
```

## 8. Замыкания (closures)

Анонимные функции могут захватывать переменные из окружающего контекста:

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    myCounter := counter()
    fmt.Println(myCounter()) // 1
    fmt.Println(myCounter()) // 2
    fmt.Println(myCounter()) // 3
}
```

## 9. Вариативные функции

Функции могут принимать переменное количество аргументов:

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3))       // 6
    fmt.Println(sum(1, 2, 3, 4, 5)) // 15
    
    nums := []int{10, 20, 30}
    fmt.Println(sum(nums...))       // 60
}
```

## 10. Рекурсивные функции

Функции могут вызывать сами себя:

```go
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    fmt.Println(factorial(5)) // 120 (5! = 5*4*3*2*1)
}
```

## Итоговый пример

```go
package main

import "fmt"

// Обычная функция
func add(a, b int) int {
    return a + b
}

// Функция, возвращающая функцию
func makeMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

func main() {
    // Использование обычной функции
    sum := add(3, 5)
    fmt.Println("3 + 5 =", sum)
    
    // Создание и использование функции-умножителя
    double := makeMultiplier(2)
    triple := makeMultiplier(3)
    
    fmt.Println("Двойное от 5:", double(5))   // 10
    fmt.Println("Тройное от 5:", triple(5))   // 15
    
    // Анонимная функция
    square := func(x int) int {
        return x * x
    }
    fmt.Println("Квадрат 4:", square(4))     // 16
    
    // Рекурсия
    fmt.Println("Факториал 6:", factorial(6)) // 720
}
```

### Ключевые моменты про функции:
1. Функции объявляются с `func`
2. Могут принимать параметры и возвращать значения
3. Могут возвращать несколько значений
4. Функции - это объекты первого класса (можно передавать, возвращать)
5. Анонимные функции полезны для коротких операций
6. Замыкания позволяют сохранять состояние между вызовами

# Рассмотрим подробнее анонимные функции, рекурсию и функции как объекты

## Анонимные функции (лямбда-функции)

Анонимные функции — это функции без имени, которые можно объявлять прямо в коде.

### Базовый пример:
```go
func main() {
    // Объявление и немедленный вызов
    func() {
        fmt.Println("Я анонимная функция!")
    }()  // Круглые скобки в конце означают вызов
    
    // Сохранение в переменную
    greet := func(name string) {
        fmt.Println("Привет,", name)
    }
    
    greet("Мир")  // Вызов через переменную
}
```

### Особенности:
1. Могут быть вызваны сразу после объявления
2. Могут быть присвоены переменным
3. Могут принимать параметры и возвращать значения

### Пример с параметрами и возвратом:
```go
func main() {
    sum := func(a, b int) int {
        return a + b
    }
    
    result := sum(3, 5)
    fmt.Println("Сумма:", result)  // Выведет: Сумма: 8
}
```

## Рекурсия в Go

Рекурсия — это когда функция вызывает саму себя.

### Классический пример — факториал:
```go
func factorial(n int) int {
    // Базовый случай (остановка рекурсии)
    if n == 0 {
        return 1
    }
    // Рекурсивный случай
    return n * factorial(n-1)
}

func main() {
    fmt.Println(factorial(5))  // 120 (5! = 5*4*3*2*1)
}
```

### Важные моменты о рекурсии:
1. Всегда должен быть **базовый случай** (условие выхода)
2. Каждый рекурсивный вызов должен **упрощать задачу**
3. В Go нет оптимизации хвостовой рекурсии (в отличие от некоторых других языков)

### Пример с числами Фибоначчи:
```go
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}
```

## Функции как объекты первого класса

В Go функции являются **объектами первого класса**, что означает:

1. Функции можно присваивать переменным
2. Функции можно передавать как аргументы другим функциям
3. Функции можно возвращать из других функций

### 1. Функция как переменная:
```go
func main() {
    // Присваиваем функцию переменной
    var greetFunc func(string) = func(name string) {
        fmt.Println("Hello,", name)
    }
    
    greetFunc("John")  // Hello, John
}
```

### 2. Функция как аргумент:
```go
func apply(numbers []int, operation func(int) int) []int {
    result := make([]int, len(numbers))
    for i, v := range numbers {
        result[i] = operation(v)
    }
    return result
}

func main() {
    numbers := []int{1, 2, 3}
    
    // Передаем анонимную функцию как аргумент
    squared := apply(numbers, func(x int) int {
        return x * x
    })
    
    fmt.Println(squared)  // [1 4 9]
}
```

### 3. Функция как возвращаемое значение:
```go
func makeMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

func main() {
    double := makeMultiplier(2)
    triple := makeMultiplier(3)
    
    fmt.Println(double(5))  // 10
    fmt.Println(triple(5))  // 15
}
```

## Замыкания (closures)

Замыкания — это функции, которые запоминают окружение, в котором они были созданы.

### Пример счетчика:
```go
func counter() func() int {
    count := 0  // Переменная сохраняется между вызовами
    return func() int {
        count++
        return count
    }
}

func main() {
    myCounter := counter()
    fmt.Println(myCounter())  // 1
    fmt.Println(myCounter())  // 2
    fmt.Println(myCounter())  // 3
    
    anotherCounter := counter()
    fmt.Println(anotherCounter())  // 1 (новая независимая копия)
}
```

### Практическое применение замыканий:
1. Создание генераторов
2. Реализация приватных переменных
3. Мемоизация (кэширование результатов)

## Отличия от других языков

1. **Более явная работа с функциями**:
   - В Python/JavaScript функции тоже объекты первого класса, но в Go это более явно
   - Нет "поднятия" (hoisting) как в JavaScript

2. **Строгая типизация**:
   ```go
   var funcVar func(int) int  // Явное объявление типа функции
   ```

3. **Нет перегрузки функций**:
   - В Go нельзя создать две функции с одним именем, но разными параметрами

4. **Простота и предсказуемость**:
   - Поведение функций всегда четко определено
   - Нет сложных правил контекста как в JavaScript



## Вывод

В Go функции — это полноценные объекты, которые можно:
- Создавать без имени (анонимные)
- Присваивать переменным
- Передавать как аргументы
- Возвращать из других функций
- Использовать в рекурсии
- Создавать замыкания

Это делает Go мощным языком с функциональными элементами, сохраняя при этом простоту и понятность.

# Попрактикуемся

## 1. Рекурсия: Сумма чисел от 1 до N
**Задача:** Напишите рекурсивную функцию `sumToN(n int) int`, которая возвращает сумму чисел от 1 до n.

```go
// Ввод:
sumToN(5)
// Ожидаемый вывод:
15
```

<details>
<summary>🔍 Решение</summary>

```go
func sumToN(n int) int {
    if n == 1 {
        return 1
    }
    return n + sumToN(n-1)
}
```
</details>

---

## 2. Defer: Простой логгер
**Задача:** Напишите функцию `logMessage(msg string)`, которая выводит сообщение, а через defer добавляет "[END]".

```go
// Ввод:
logMessage("Старт программы")
// Ожидаемый вывод:
Старт программы
[END]
```

<details>
<summary>🔍 Решение</summary>

```go
func logMessage(msg string) {
    defer fmt.Println("[END]")
    fmt.Println(msg)
}
```
</details>

---

## 3. Анонимная функция: Квадраты чисел
**Задача:** Создайте анонимную функцию, которая возводит число в квадрат и сразу вызовите её.

```go
// Ввод:
func() {
    x := 5
    // Ваш код здесь
}()
// Ожидаемый вывод:
25
```

<details>
<summary>🔍 Решение</summary>

```go
func() {
    x := 5
    square := func(n int) int { return n * n }
    fmt.Println(square(x))
}()
```
</details>

---

## 4. Обработка ошибок: Делитель
**Задача:** Напишите функцию `divide(a, b int) (int, error)`, которая возвращает ошибку при делении на ноль.

```go
// Ввод:
res, err := divide(10, 0)
// Ожидаемый вывод:
0 ошибка: деление на ноль
```

<details>
<summary>🔍 Решение</summary>

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("ошибка: деление на ноль")
    }
    return a / b, nil
}
```
</details>

---

## 5. Функция как значение: Операции
**Задача:** Создайте переменную `operation`, которая может хранить функцию для сложения или вычитания.

```go
// Ввод:
operation := // присвойте функцию сложения
fmt.Println(operation(2, 3))
// Ожидаемый вывод:
5
```

<details>
<summary>🔍 Решение</summary>

```go
operation := func(a, b int) int { return a + b }
fmt.Println(operation(2, 3))
```
</details>

---

## 6. Рекурсия + defer: Обратный отсчёт
**Задача:** Напишите функцию `countdown(n int)`, которая рекурсивно ведёт отсчёт с defer.

```go
// Ввод:
countdown(3)
// Ожидаемый вывод:
Старт
3
2
1
Пуск!
```

<details>
<summary>🔍 Решение</summary>

```go
func countdown(n int) {
    fmt.Println("Старт")
    defer fmt.Println("Пуск!")
    
    if n == 0 {
        return
    }
    fmt.Println(n)
    countdown(n-1)
}
```
</details>

---

## 7. Анонимная функция + defer: Счётчик вызовов
**Задача:** Создайте анонимную функцию с defer, которая считает свои вызовы.

```go
// Ввод:
count := 0
func() {
    // Ваш код с defer
    count++
}()
fmt.Println(count)
// Ожидаемый вывод после 3 вызовов:
3
```

<details>
<summary>🔍 Решение</summary>

```go
count := 0
f := func() {
    defer func() { count++ }()
    f()
}
f()
fmt.Println(count)
```
</details>

---

## 8. Простая рекурсия: Степень числа
**Задача:** Напишите рекурсивную функцию `power(base, exp int) int`, вычисляющую base^exp.

```go
// Ввод:
power(2, 3)
// Ожидаемый вывод:
8
```

<details>
<summary>🔍 Решение</summary>

```go
func power(base, exp int) int {
    if exp == 0 {
        return 1
    }
    return base * power(base, exp-1)
}
```
</details>

---

## 9. Defer: Изменение возвращаемого значения
**Задача:** Напишите функцию `withMessage()`, которая через defer добавляет к результату "[OK]".

```go
// Ввод:
func withMessage() string {
    msg := "Выполнено"
    // Ваш defer код
    return msg
}
// Ожидаемый вывод:
Выполнено [OK]
```

<details>
<summary>🔍 Решение</summary>

```go
func withMessage() (msg string) {
    msg = "Выполнено"
    defer func() { msg += " [OK]" }()
    return
}
```
</details>

---

## 10. Комплекс: Рекурсия + функции как объекты
**Задача:** Создайте функцию-генератор `sequenceGenerator`, которая возвращает следующее число последовательности (1, 3, 6, 10...).

```go
// Ввод:
gen := sequenceGenerator()
fmt.Println(gen(), gen(), gen())
// Ожидаемый вывод:
1 3 6
```

<details>
<summary>🔍 Решение</summary>

```go
func sequenceGenerator() func() int {
    n, sum := 0, 0
    return func() int {
        n++
        sum += n
        return sum
    }
}
```
</details>

---

# Про IDE и установку GO

## Компилятор vs Интерпретатор в контексте Go

**Go — компилируемый язык**, что означает:

### 🔹 Основные отличия:
| Характеристика       | Компилятор (Go)               | Интерпретатор (Python)       |
|----------------------|-------------------------------|------------------------------|
| **Процесс**          | Преобразует код в бинарник до выполнения | Выполняет код построчно |
| **Скорость**         | Быстрое выполнение            | Медленнее (анализ на лету)   |
| **Распространение**  | Один бинарный файл            | Нужен интерпретатор          |
| **Отладка**          | Сложнее (нет доступа к исходнику при ошибке) | Проще (видна строка с ошибкой) |

**Пример компиляции в Go:**
```bash
go build main.go  # Компиляция
./main           # Запуск
```

## Установка Go

### 🖥️ Для Windows:
1. Скачайте установщик с [официального сайта](https://golang.org/dl/)
2. Запустите `.msi` файл и следуйте инструкциям
3. Проверьте установку:
```cmd
go version
```

### 🍏 Для Mac:
1. Вариант 1: Через Homebrew
```bash
brew install go
```

2. Вариант 2: Скачайте `.pkg` с [официального сайта](https://golang.org/dl/)
3. Проверьте установку:
```bash
go version
```

## Настройка IDE/редакторов

### 🔧 Visual Studio Code (рекомендуется)
1. Установите [VS Code](https://code.visualstudio.com/)
2. Откройте расширения (Ctrl+Shift+X) и установите:
   - Go (от Google)
   - Go Test Explorer (для тестов)

3. Настройки для Go (settings.json):
```json
{
  "go.formatTool": "goimports",
  "go.useLanguageServer": true,
  "go.lintOnSave": "package"
}
```

### 🛠️ Настройка GOPATH (важно!)
1. Создайте рабочую директорию (например `~/go`)
2. Добавьте в переменные среды:
```bash
# Mac/Linux
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# Windows (через GUI или cmd)
setx GOPATH "%USERPROFILE%\go"
setx PATH "%PATH%;%GOPATH%\bin"
```

## Альтернативные IDE

### 1. Goland (JetBrains)
- Платная (есть бесплатный пробный период)
- Самая продвинутая среда для Go
- [Скачать](https://www.jetbrains.com/go/)

### 2. LiteIDE
- Бесплатная, легковесная
- Специально разработана для Go
- [Скачать](https://github.com/visualfc/liteide)

## Полезные команды после установки

```bash
go env      # Проверить настройки окружения
go help     # Список всех команд
go get -u github.com/gin-gonic/gin  # Пример установки пакета
```

## Советы по настройке

1. **Автодополнение**: В VS Code установите `gopls`:
```bash
go install golang.org/x/tools/gopls@latest
```

2. **Форматирование**: Установите `goimports`:
```bash
go install golang.org/x/tools/cmd/goimports@latest
```

3. **Отладка**: В VS Code нажмите F5 и выберите "Go" для настройки дебаггера.

Теперь вы готовы к разработке на Go! Для проверки создайте файл `main.go`:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```
И запустите:
```bash
go run main.go
```


# Работа с Git: от установки до управления версиями

## Установка Git

### Windows
1. Скачайте установщик с [официального сайта](https://git-scm.com/downloads)
2. Запустите установщик, оставив все настройки по умолчанию
3. Проверьте установку в командной строке:
```bash
git --version
```

### Mac
1. Установите через Homebrew:
```bash
brew install git
```
2. Или скачайте с [официального сайта](https://git-scm.com/download/mac)
3. Проверьте установку:
```bash
git --version
```

## Настройка SSH-ключей

SSH-ключи нужны для безопасного подключения к GitHub/GitLab без ввода пароля.

### 1. Генерация ключей
```bash
ssh-keygen -t ed25519 -C "your_email@example.com"
```
(Нажимайте Enter на всех вопросах)

### 2. Добавление ключа в ssh-agent
```bash
eval "$(ssh-agent -s)"
ssh-add ~/.ssh/id_ed25519
```

### 3. Показ и копирование публичного ключа
```bash
cat ~/.ssh/id_ed25519.pub
```
Скопируйте вывод и добавьте в настройках GitHub/GitLab:  
Settings → SSH and GPG keys → New SSH key

### 4. Проверка подключения
```bash
ssh -T git@github.com
```

## Основные команды Git

### Инициализация репозитория
```bash
git init
```

### Проверка состояния файлов
```bash
git status
```
Показывает:
- Измененные файлы
- Новые файлы
- Файлы готовые к коммиту

### Добавление файлов в отслеживание
```bash
git add <file>      # Добавить конкретный файл
git add .           # Добавить все измененные файлы
git add -A          # Добавить все файлы (включая удаленные)
```

### Создание коммита
```bash
git commit -m "Описание изменений"
```
Фиксирует изменения в истории проекта

### Отправка изменений на сервер
```bash
git push origin main  # Отправить в ветку main
```

## Работа с историей изменений

### Просмотр истории
```bash
git log              # Полная история
git log --oneline    # Краткая история
git log --graph      # История с визуализацией веток
```

### Откат изменений

1. **Отмена незакоммиченных изменений**:
```bash
git checkout -- <file>  # Откат одного файла
git reset --hard        # Полный откат всех изменений
```

2. **Отмена коммита (еще не отправленного)**:
```bash
git reset HEAD~1        # Отменить последний коммит, сохранив изменения
git reset --hard HEAD~1 # Полное удаление последнего коммита
```

3. **Отмена уже отправленного коммита**:
```bash
git revert <commit-hash>  # Создает новый коммит, отменяющий изменения
git push origin main      # Отправляем изменения
```

## Типичный рабочий процесс

1. Получить последние изменения:
```bash
git pull origin main
```

2. Создать новую ветку для задачи:
```bash
git checkout -b feature/new-button
```

3. Работать с файлами, затем:
```bash
git add .
git commit -m "Добавлена новая кнопка"
```

4. Отправить изменения:
```bash
git push origin feature/new-button
```

5. Создать Pull Request в GitHub/GitLab

## Полезные советы

1. **Игнорирование файлов**  
   Создайте файл `.gitignore` в корне проекта с перечнем файлов/папок, которые Git должен игнорировать

2. **Просмотр изменений**  
   ```bash
   git diff            # Покажет все изменения
   git diff <file>     # Покажет изменения в конкретном файле
   ```

3. **Временное сохранение изменений**  
   ```bash
   git stash           # Сохранить изменения "в ящик"
   git stash pop       # Вернуть изменения обратно
   ```


## **1. Вопросы которые могли возникнуть по go и git**  

Когда вы устанавливаете Go, его бинарные файлы (исполняемые программы, например, `go`, `gofmt`) помещаются в определенную папку (например, `C:\Go\bin` на Windows или `/usr/local/go/bin` на macOS/Linux).  

**Зачем добавлять в `PATH`?**  
- Чтобы терминал/командная строка знал, где искать команду `go`.  
- Без этого при вводе `go version` система может сказать:  
  ```bash
  'go' is not recognized as an internal or external command...
  ```
  
### **Как добавить Go в `PATH`?**  

#### **Windows**  
1. Откройте **Системные свойства** → **Переменные среды**  
2. В **`PATH`** добавьте путь к папке `bin` (например, `C:\Go\bin`)  

#### **macOS / Linux**  
Добавьте в `~/.bashrc` или `~/.zshrc`:  
```bash
export PATH=$PATH:/usr/local/go/bin
```
Затем перезагрузите терминал:  
```bash
source ~/.bashrc  # или source ~/.zshrc
```

---

## **2. Что такое SSH-ключ простыми словами?**  

SSH-ключ — это **"электронный пропуск"** для безопасного доступа к GitHub/GitLab без ввода пароля.  

- **Приватный ключ** (`id_rsa`) — хранится у вас на компьютере (никому не отдавать!).  
- **Публичный ключ** (`id_rsa.pub`) — загружается в GitHub/GitLab.  

🔐 **Как это работает?**  
1. Вы подключаетесь к GitHub через SSH.  
2. GitHub проверяет, есть ли у вас **приватный ключ**, который соответствует **публичному**.  
3. Если ключи совпадают — доступ разрешён!  

---

## **3. Как быстро сгенерировать SSH-ключ?**  

### **1. Генерация ключа (RSA, самый популярный)**  
```bash
ssh-keygen -t rsa -b 4096 -C "ваша_почта@example.com"
```
- Нажимайте **Enter** на всех вопросах (пароль можно не ставить).  
- Ключи сохранятся в `~/.ssh/id_rsa` (приватный) и `~/.ssh/id_rsa.pub` (публичный).  

### **2. Просмотр ключей**  
```bash
ls -a ~/.ssh  # Посмотреть список файлов в папке .ssh
```
Вывод:  
```
id_rsa      # Приватный ключ (никому не показывать!)
id_rsa.pub  # Публичный ключ (можно загружать в GitHub)
```

### **3. Показать публичный ключ**  
```bash
cat ~/.ssh/id_rsa.pub
```
Вы увидите что-то вроде:  
```
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC... ваша_почта@example.com
```

### **4. Добавление ключа в GitHub/GitLab**  
1. Скопируйте вывод `cat ~/.ssh/id_rsa.pub`  
2. Перейдите в:  
   - GitHub: **Settings → SSH and GPG Keys → New SSH Key**  
   - GitLab: **Preferences → SSH Keys**  
3. Вставьте ключ и сохраните.  

### **5. Проверка подключения**  
```bash
ssh -T git@github.com
```
Если всё ок, увидите:  
```
Hi username! You've successfully authenticated...
```

---

## **Итог**  
✅ **Go в `PATH`** — чтобы команда `go` работала в любом месте терминала.  
🔑 **SSH-ключ** — безопасный вход в GitHub без пароля.  
🚀 **Быстрая генерация** → `ssh-keygen -t rsa`, `cat ~/.ssh/id_rsa.pub`.  
