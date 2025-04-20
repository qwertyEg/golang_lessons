# **Обучение Go: Требования к коду, улучшения и Code Style**

## **1. Требования к коду на Go**
Go — язык со строгими правилами оформления. Вот основные:

### **1.1 Именование**
- **CamelCase**: Используйте для имен переменных, функций, типов.
  - Правильно: `getUser`, `MaxSize`
  - Неправильно: `get_user`, `max_size`

- **Acronyms**: Пишите аббревиатуры в верхнем регистре.
  - Правильно: `HTTPRequest`, `JSONParser`

- **Пакеты**: Короткие имена в нижнем регистре. Избегайте подчеркиваний.
  - Правильно: `utils`, `httpclient`

### **1.2 Форматирование**
- **Отступы**: Используйте табы (не пробелы).
- **Длина строки**: Старайтесь не превышать 80–100 символов.
- **Фигурные скобки**: Открывающая скобка — на той же строке.
  ```go
  // Правильно
  if x > 0 {
      fmt.Println("Hello")
  }

  // Неправильно
  if x > 0 
  {
      fmt.Println("Hello")
  }
  ```

**Совет**: Команда `gofmt` автоматически форматирует код. Запустите:
```bash
gofmt -w yourfile.go
```

---

## **2. Улучшения кода**
### **2.1 Избегайте лишних переменных**
**Плохо**:
```go
result := add(5, 3)
fmt.Println(result)
```

**Хорошо**:
```go
fmt.Println(add(5, 3))
```

### **2.2 Обработка ошибок**
Всегда проверяйте ошибки. Не игнорируйте их!

**Плохо**:
```go
data, _ := ioutil.ReadFile("file.txt")
```

**Хорошо**:
```go
data, err := ioutil.ReadFile("file.txt")
if err != nil {
    log.Fatal("Ошибка чтения файла:", err)
}
```

### **2.3 Используйте `defer` для очистки**
**Пример**:
```go
func readFile() {
    file, err := os.Open("file.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close() // Закрытие файла после завершения функции
    // ... работа с файлом ...
}
```

---

## **3. Правила оформления кода**
### **3.1 Структура файла**
1. Объявление пакета (`package main`)
2. Импорты (группируйте стандартные и сторонние)
3. Остальной код

**Пример**:
```go
package main

import (
    "fmt"
    "math/rand"

    "github.com/example/utils"
)
```

### **3.2 Комментарии**
- **Однострочные**: `// Комментарий`
- **Многострочные**: 
  ```go
  /* 
  Это многострочный 
  комментарий
  */
  ```
- **Документирующие**: Начинаются с имени сущности и пишутся перед ней.
  ```go
  // CalculateSum возвращает сумму двух чисел.
  func CalculateSum(a, b int) int {
      return a + b
  }
  ```

---

## **4. Линтеры**
Линтеры проверяют код на соответствие стандартам и потенциальные ошибки.

### **4.1 Установка**
```bash
go install golang.org/x/lint/golint@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
```

### **4.2 Пример проверки**
**Код** (`main.go`):
```go
package main

func test() { // Не экспортированная функция с именем в нижнем регистре
    fmt.Println("test")
}
```

**Запуск**:
```bash
golint main.go
```
**Вывод**:
```
main.go:3:1: exported function Test should have comment or be unexported
```

**Исправленный код**:
```go
package main

// test выводит строку "test".
func test() {
    fmt.Println("test")
}
```

---

## **5. Документирование кода**
Go автоматически генерирует документацию из комментариев. Используйте `godoc`.

### **5.1 Самодокументируемый код**
Код на Go может быть «самодокументирующимся» за счет:
1. **Понятных имен** функций, переменных и типов.
2. **Логичной структуры**.
3. **Комментариев в формате godoc**, которые автоматически генерируют документацию.

### **5.2 Пример самодокументируемого кода**
```go
// Package geometry предоставляет базовые операции с геометрическими фигурами.
package geometry

import "math"

// Point представляет точку в 2D-пространстве.
type Point struct {
    X, Y float64
}

// NewPoint создает новую точку с заданными координатами.
func NewPoint(x, y float64) Point {
    return Point{X: x, Y: y}
}

// DistanceTo вычисляет расстояние между текущей точкой и целевой.
func (p Point) DistanceTo(target Point) float64 {
    dx := p.X - target.X
    dy := p.Y - target.Y
    return math.Sqrt(dx*dx + dy*dy)
}

// Circle представляет окружность с центром и радиусом.
type Circle struct {
    Center Point
    Radius float64
}

// Area вычисляет площадь окружности.
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

// IsPointInside проверяет, находится ли точка внутри окружности.
func (c Circle) IsPointInside(p Point) bool {
    return c.Center.DistanceTo(p) <= c.Radius
}
```

### **5.3 Как это работает?**
#### **Самодокументируемые имена**
- `DistanceTo` ясно указывает, что метод вычисляет расстояние.
- `IsPointInside` прямо отвечает на вопрос «находится ли точка внутри?».

#### **Комментарии для godoc**
- Комментарии начинаются с имени сущности (`// Point представляет...`).
- Описывают **что делает** функция/тип, а не как реализовано.

### **5.4 Сгенерированная документация**
Запустите локальный сервер документации:
```bash
godoc -http=:6060
```

Перейдите по адресу `http://localhost:6060/pkg/ваш_модуль/geometry/`.

**Вы увидите:**
```
PACKAGE DOCUMENTATION

package geometry
    import "ваш_модуль/geometry"

Package geometry предоставляет базовые операции с геометрическими фигурами.

TYPES

type Circle struct {
    Center Point
    Radius float64
}
    Circle представляет окружность с центром и радиусом.

func (c Circle) Area() float64
    Area вычисляет площадь окружности.

func (c Circle) IsPointInside(p Point) bool
    IsPointInside проверяет, находится ли точка внутри окружности.

type Point struct {
    X, Y float64
}
    Point представляет точку в 2D-пространстве.

func NewPoint(x, y float64) Point
    NewPoint создает новую точку с заданными координатами.

func (p Point) DistanceTo(target Point) float64
    DistanceTo вычисляет расстояние между текущей точкой и целевой.
```

### **5.5 Почему это «самодокументируемый» код?**
1. **Имена говорят сами за себя**:
   - `circle.Area()` не требует комментария, чтобы понять, что она делает.
   - `p.DistanceTo(target)` читается как предложение.

2. **Минимум технического жаргона**:
   ```go
   // Плохой комментарий:
   // Используем теорему Пифагора
   dx := p.X - target.X
   dy := p.Y - target.Y
   return math.Sqrt(dx*dx + dy*dy)

   // Хорошо: код и так понятен, комментарий не нужен.
   ```

3. **Комментарии добавляют контекст**, а не повторяют код:
   - Вместо:
     ```go
     // Умножает радиус на радиус
     return math.Pi * c.Radius * c.Radius
     ```
   - Лучше:
     ```go
     // Area вычисляет площадь окружности.
     ```

### **5.6 Дополнительный прием: Примеры использования**
Добавьте файл `geometry_example_test.go`:
```go
package geometry_test

import (
    "fmt"
    "ваш_модуль/geometry"
)

func ExampleCircle_Area() {
    c := geometry.Circle{
        Center: geometry.NewPoint(0, 0),
        Radius: 5,
    }
    fmt.Printf("Площадь окружности: %.2f", c.Area())
    // Output: Площадь окружности: 78.54
}

---
``` 

## **6. Code Style: Теория и практика**
### **6.1 Основные принципы**
- **Простота**: Пишите код так, чтобы его было легко читать.
- **Единообразие**: Следуйте стилю проекта.
- **Эффективность**: Не оптимизируйте преждевременно.

### **6.2 Пример рефакторинга**
**Было**:
```go
func ProcessData(data []int) {
    var sum int = 0
    for i := 0; i < len(data); i++ {
        sum += data[i]
    }
    avg := sum / len(data)
    fmt.Println(avg)
}
```

**Стало** (используем range, убираем лишние переменные):
```go
// ProcessData вычисляет среднее значение среза целых чисел.

func ProcessData(data []int) {
    sum := 0
    for _, value := range data {
        sum += value
    }
    fmt.Println(sum / len(data))
}
```

---

## **7. Помни!!!!!**
- Всегда запускай `gofmt` и `golint`.
- Пиши комментарии для экспортированных сущностей.
- Следуй официальному гайду [Effective Go](https://golang.org/doc/effective_go).

**Практическое задание**:
1. Напишите функцию `Multiply(a, b int)`, которая возвращает произведение чисел.
2. Добавьте документирующий комментарий.
3. Проверьте код через `golint`.

**Пример решения**:
```go
// Multiply возвращает произведение двух целых чисел.
func Multiply(a, b int) int {
    return a * b
}
```

## **8. Тестирование в Go**

Тестирование — важная часть разработки. В Go тестирование встроено в язык через пакет `testing`, а библиотека `Testify` расширяет возможности. Разберем все по шагам.

---

### **8.1 Основы тестирования в Go**

#### **8.1.1 Как устроены тесты**
- Тесты пишутся в файлах с суффиксом `_test.go`.
- Функции тестов начинаются с `TestXxx`, где `Xxx` — имя тестируемой функции.
- Используется тип `*testing.T` для управления тестами.

**Пример:**
```go
// Файл: math.go
package math

func Add(a, b int) int {
    return a + b
}
```

```go
// Файл: math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Ожидалось 5, получено %d", result)
    }
}
```

#### **8.1.2 Запуск тестов**
```bash
go test -v ./...  # -v для подробного вывода
```

#### **8.1.3 Табличные тесты**
Позволяют тестировать множество случаев в одном тесте.

```go
func TestAdd_TableDriven(t *testing.T) {
    tests := []struct {
        name   string
        a, b   int
        want   int
    }{
        {"Позитивные числа", 2, 3, 5},
        {"Ноль", 0, 0, 0},
        {"Отрицательные числа", -1, -2, -3},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Add(tt.a, tt.b); got != tt.want {
                t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
            }
        })
    }
}
```

---

### **8.2 Пакет Testify**

**Testify** — это популярная библиотека для Go, которая расширяет возможности стандартного пакета `testing`. Она делает тесты более читаемыми, удобными и мощными. Давайте разберем её основные компоненты на примерах.

#### **8.2.1 Установка**
```bash
go get github.com/stretchr/testify
```

#### **8.2.2 Основные компоненты Testify**

##### **8.2.2.1 Пакет `assert`**
Упрощает проверки в тестах. Если проверка не проходит, тест помечается как проваленный, но продолжает выполняться.

**Пример:**
```go
import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    assert.Equal(t, 5, result, "Сумма 2+3 должна быть 5")
    
    // Проверка на наличие ошибки
    _, err := Divide(10, 0)
    assert.Error(t, err)
    
    // Проверка на равенство слайсов
    expected := []int{1, 2, 3}
    assert.Equal(t, expected, []int{1, 2, 3})
}
```

**Доступные методы:**
- `Equal(t, expected, actual)` — проверка равенства.
- `NotEqual(t, expected, actual)` — проверка неравенства.
- `True(t, condition)` — проверка, что условие истинно.
- `Nil(t, obj)` — проверка, что объект `nil`.
- `Contains(t, str, substring)` — проверка наличия подстроки.
- И многие другие: [документация](https://pkg.go.dev/github.com/stretchr/testify/assert).

##### **8.2.2.2 Пакет `require`**
Аналогичен `assert`, но при провале проверки **немедленно завершает тест**. Полезно для критических проверок.

**Пример:**
```go
import (
    "testing"
    "github.com/stretchr/testify/require"
)

func TestCriticalCheck(t *testing.T) {
    // Если файл не найден, тест остановится здесь
    data, err := os.ReadFile("config.json")
    require.NoError(t, err)
    
    // Дальнейшие проверки...
    require.Contains(t, string(data), "admin")
}
```

##### **8.2.2.3 Пакет `mock`**
Позволяет создавать **моки (имитации)** объектов для изоляции тестируемого кода от внешних зависимостей.

**Пример: Мок HTTP-клиента**
```go
import (
    "testing"
    "github.com/stretchr/testify/mock"
)

// Интерфейс HTTP-клиента
type HTTPClient interface {
    Get(url string) (string, error)
}

// Мок
type MockHTTPClient struct {
    mock.Mock
}

func (m *MockHTTPClient) Get(url string) (string, error) {
    args := m.Called(url)
    return args.String(0), args.Error(1)
}

// Тест
func TestFetchData(t *testing.T) {
    mockClient := new(MockHTTPClient)
    
    // Настраиваем мок: при вызове Get("https://api.com") вернуть "OK" и nil
    mockClient.On("Get", "https://api.com").Return("OK", nil)
    
    // Передаем мок в тестируемый код
    result, err := FetchData(mockClient, "https://api.com")
    
    assert.NoError(t, err)
    assert.Equal(t, "OK", result)
    
    // Проверяем, что метод Get был вызван с нужным аргументом
    mockClient.AssertCalled(t, "Get", "https://api.com")
}
```

##### **8.2.2.4 Пакет `suite`**
Позволяет создавать **наборы тестов (Test Suites)** с общими настройками (например, подключение к БД перед всеми тестами).

**Пример:**
```go
import (
    "testing"
    "github.com/stretchr/testify/suite"
)

// Создаем структуру для набора тестов
type MySuite struct {
    suite.Suite
    db *Database // Общие ресурсы
}

// Настройка перед всеми тестами
func (s *MySuite) SetupSuite() {
    s.db = ConnectToTestDB()
}

// Очистка после всех тестов
func (s *MySuite) TearDownSuite() {
    s.db.Close()
}

// Настройка перед каждым тестом
func (s *MySuite) SetupTest() {
    s.db.ClearTables()
}

// Тест 1
func (s *MySuite) TestAddUser() {
    err := s.db.AddUser("Alice")
    s.NoError(err)
}

// Тест 2
func (s *MySuite) TestGetUser() {
    user, err := s.db.GetUser("Alice")
    s.NoError(err)
    s.Equal("Alice", user.Name)
}

// Запуск набора тестов
func TestMySuite(t *testing.T) {
    suite.Run(t, new(MySuite))
}
```

#### **8.2.3 Преимущества Testify**
- **Удобные проверки**: Методы вроде `Equal` или `Contains` делают код тестов чище.
- **Читаемость**: Тесты выглядят как обычные утверждения на естественном языке.
- **Моки**: Легко изолировать код от внешних систем.
- **Группировка тестов**: `Test Suites` упрощают работу с общими ресурсами.

#### **8.2.4 Лучшие практики**
1. **Используйте `assert` для не критических проверок** (например, проверка значения).
2. **Используйте `require` для обязательных проверок** (например, инициализация ресурсов).
3. **Не злоупотребляйте моками**: Если можно использовать реальный объект (например, базу в памяти), делайте это.
4. **Документируйте моки**: Комментарии помогут понять, какие вызовы ожидаются.

#### **8.2.5 Пример: Полный тест с Testify**
```go
func TestUserService(t *testing.T) {
    // Создаем мок БД
    mockDB := new(MockDatabase)
    mockDB.On("GetUser", 1).Return("Alice", nil)
    mockDB.On("SaveUser", "Bob").Return(nil)

    // Создаем сервис
    service := UserService{db: mockDB}

    t.Run("GetUser", func(t *testing.T) {
        name, err := service.GetUserName(1)
        assert.NoError(t, err)
        assert.Equal(t, "Alice", name)
    })

    t.Run("SaveUser", func(t *testing.T) {
        err := service.SaveUser("Bob")
        require.NoError(t, err)
    })

    // Проверяем, что все ожидаемые методы были вызваны
    mockDB.AssertExpectations(t)
}
```

#### **8.2.6 Полезные ссылки**
- [Документация Testify](https://github.com/stretchr/testify)
- [Примеры использования моков](https://github.com/stretchr/testify#mock-package)
- [Effective Testing with Testify](https://blog.alexellis.io/golang-testing-with-testify/)

---

### **8.3 Покрытие кода**
Проверьте, какой процент кода покрыт тестами:
```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out  # Открыть в браузере
```

---

### **8.4 Лучшие практики**
1. **Имена тестов**: `TestФункция_Сценарий_Ожидание` (например, `TestAdd_NegativeNumbers_ReturnsSum`).
2. **Табличные тесты**: Для множества кейсов.
3. **Тестируйте ошибки**: Убедитесь, что функции возвращают ошибки, когда должны.
4. **Избегайте зависимостей**: Используйте моки для внешних сервисов.
5. **Покрытие**: Стремитесь к 70-80%, но не гонитесь за 100%.

---

### **8.5 Пример: Полный тест с Testify**
```go
func TestCalculateDiscount(t *testing.T) {
    t.Run("Скидка 10%", func(t *testing.T) {
        price := 100.0
        discount := 0.1
        expected := 90.0

        result, err := CalculateDiscount(price, discount)
        assert.NoError(t, err)
        assert.InDelta(t, expected, result, 0.001) // Для сравнения float
    })

    t.Run("Отрицательная скидка", func(t *testing.T) {
        _, err := CalculateDiscount(100, -0.1)
        assert.ErrorContains(t, err, "недопустимая скидка")
    })
}
```

---

### **8.6 Полезные команды**
- Запуск конкретного теста:
  ```bash
  go test -run TestAdd
  ```
- Пропустить кеширование:
  ```bash
  go test -count=1
  ``` 

## **9. Задание: Написание тестов для функции проверки палиндрома**

## **Цель задания**
1. Написать функцию `IsPalindrome`, которая проверяет, является ли строка палиндромом.
2. Написать тесты для этой функции, используя стандартный пакет `testing`.
3. Научиться запускать тесты и анализировать их результаты.

---

## **Что такое палиндром?**
Палиндром — это строка, которая читается одинаково слева направо и справа налево.  
**Примеры**:  
- "level"  
- "madam"  
- "A man a plan a canal Panama" (если игнорировать пробелы и регистр).

---

## **Шаг 1: Создайте проект**
1. Создайте папку для проекта, например: `palindrome`.
2. Перейдите в нее:
   ```bash
   mkdir palindrome && cd palindrome
   ```

---

## **Шаг 2: Напишите функцию `IsPalindrome`**
Создайте файл `palindrome.go`:
```go
package palindrome

// IsPalindrome проверяет, является ли строка палиндромом.
// Пока игнорируем пробелы и регистр (это можно добавить позже).
func IsPalindrome(s string) bool {
    for i := 0; i < len(s)/2; i++ {
        if s[i] != s[len(s)-1-i] {
            return false
        }
    }
    return true
}
```

---

## **Шаг 3: Напишите тесты**
Создайте файл `palindrome_test.go`:
```go
package palindrome

import "testing"

// Тест для палиндромов
func TestIsPalindrome(t *testing.T) {
    // Таблица тестовых случаев
    tests := []struct {
        name string
        input string
        want bool
    }{
        {"Пустая строка", "", true},
        {"Один символ", "a", true},
        {"Палиндром (нечетная длина)", "level", true},
        {"Палиндром (четная длина)", "abba", true},
        {"Не палиндром", "hello", false},
        {"Строка с разным регистром", "Level", false}, // Пока не обрабатываем регистр
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := IsPalindrome(tt.input); got != tt.want {
                t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, got, tt.want)
            }
        })
    }
}
```

---

## **Шаг 4: Запустите тесты**
1. Откройте терминал в папке проекта.
2. Выполните команду:
   ```bash
   go test -v
   ```
   - `-v` выводит подробную информацию о каждом тесте.

**Пример вывода:**
```
=== RUN   TestIsPalindrome
=== RUN   TestIsPalindrome/Пустая_строка
=== RUN   TestIsPalindrome/Один_символ
=== RUN   TestIsPalindrome/Палиндром_(нечетная_длина)
=== RUN   TestIsPalindrome/Палиндром_(четная_длина)
=== RUN   TestIsPalindrome/Не_палиндром
=== RUN   TestIsPalindrome/Строка_с_разным_регистром
--- FAIL: TestIsPalindrome (0.00s)
    --- PASS: TestIsPalindrome/Пустая_строка (0.00s)
    --- PASS: TestIsPalindrome/Один_символ (0.00s)
    --- PASS: TestIsPalindrome/Палиндром_(нечетная_длина) (0.00s)
    --- PASS: TestIsPalindrome/Палиндром_(четная_длина) (0.00s)
    --- PASS: TestIsPalindrome/Не_палиндром (0.00s)
    --- FAIL: TestIsPalindrome/Строка_с_разным_регистром (0.00s)
        palindrome_test.go:24: IsPalindrome("Level") = false, want true
FAIL
exit status 1
FAIL    palindrome    0.002s
```

---

## **Шаг 5: Анализ результатов**
- Тест `Строка с разным регистром` провален, потому что текущая реализация не игнорирует регистр.
- Это нормально! Тесты помогают находить недоработки.

---

## **Шаг 6: Улучшите функцию (доп. задание)**
Измените `IsPalindrome`, чтобы она:
1. Игнорировала регистр символов.
2. Удаляла пробелы из строки перед проверкой.

**Подсказка:**
- Используйте `strings.ToLower()` для приведения к нижнему регистру.
- Используйте `strings.ReplaceAll()` для удаления пробелов.

---

## **Шаг 7: Обновите тесты**
Добавьте новые тестовые случаи в таблицу:
```go
{"Палиндром с пробелами", "A man a plan a canal Panama", true},
{"Палиндром с регистром и пробелами", "Madam Im Adam", true},
```

---

## **Инструкции для новичка**
1. **Как называть тесты**:  
   - Функции тестов должны начинаться с `Test` (например, `TestIsPalindrome`).  
   - Используйте `t.Run()` для подтестов с понятными именами.

2. **Что установить**:  
   - Только Go. Дополнительные библиотеки не нужны (мы используем стандартный `testing`).

3. **Как запускать**:  
   - `go test -v` — запуск всех тестов.  
   - `go test -run TestIsPalindrome` — запуск конкретного теста.  
   - `go test -cover` — проверка покрытия кода тестами.

---

## **Пример улучшенной функции**
```go
package palindrome

import "strings"

func IsPalindrome(s string) bool {
    // Приводим к нижнему регистру и удаляем пробелы
    s = strings.ToLower(s)
    s = strings.ReplaceAll(s, " ", "")

    // Проверяем палиндром
    for i := 0; i < len(s)/2; i++ {
        if s[i] != s[len(s)-1-i] {
            return false
        }
    }
    return true
}
``` 