## **1. Статический анализ (Static Analysis)**

### **Что это?**  
Это автоматическая проверка кода **без его выполнения**.  
Анализатор ищет потенциальные ошибки, недочеты, нарушение стандартов и «запахи» кода.

**Примеры проблем, которые ловит статический анализ:**
- Неправильное использование `fmt.Printf` (например, `%d` для строки).
- Неиспользуемые переменные или импорты.
- Подозрительные конструкции: вечный цикл `for {}`, dead code.

---

### **Инструменты статического анализа в Go**
#### **1.1 `go vet`**  
Встроенный в Go инструмент. Проверяет код на распространенные ошибки.

**Установка**: Уже есть в Go.  
**Запуск**:
```bash
go vet ./...
```

**Пример**:
```go
func main() {
    fmt.Printf("%d", "hello") // Ошибка: %d для строки
}
```
**Вывод**:
```
printf: %d format arg "hello" is a string value
```

---

#### **1.2 `staticcheck`**  
Продвинутый статический анализатор. Находит больше ошибок, чем `go vet`.

**Установка**:
```bash
go install honnef.co/go/tools/cmd/staticcheck@latest
```

**Запуск**:
```bash
staticcheck ./...
```

**Пример**:
```go
func calc() error {
    return nil
}

func main() {
    calc() // Ошибка: результат функции calc() игнорируется
}
```
**Вывод**:
```
main.go:7:2: result of call to calc is unused (SA4017)
```

---

## **2. Линтеры (Linters)**

### **Что это?**  
Линтеры проверяют код на соответствие **стилистическим правилам** и **best practices**.  
Они не ищут ошибки выполнения, но следят за единообразием кода.

**Примеры правил линтеров**:
- Имена функций должны быть в `CamelCase`.
- Комментарии для экспортированных функций.
- Максимальная длина строки (например, 80 символов).

---

### **Инструменты-линтеры в Go**
#### **2.1 `golint`** (устаревший)  
Проверяет стиль кода по рекомендациям Go.

**Установка**:
```bash
go install golang.org/x/lint/golint@latest
```

**Пример**:
```go
// Плохо: функция экспортирована (начинается с заглавной), но нет комментария
func GetData() {}
```
**Запуск**:
```bash
golint ./...
```
**Вывод**:
```
main.go:3:1: exported function GetData should have comment or be unexported
```

---

#### **2.2 `revive`**  
Современная замена `golint` с настройкой правил.

**Установка**:
```bash
go install github.com/mgechev/revive@latest
```

**Запуск**:
```bash
revive ./...
```

**Пример настройки правил** (файл `revive.toml`):
```toml
[rule.exported]
  # Требовать комментарии для экспортированных функций
  arguments = [["allowUncommentedExportedStruct"]]
```

---

## **3. Различия: Статический анализ vs Линтеры**

| **Критерий**          | **Статический анализ**               | **Линтеры**                     |
|-----------------------|--------------------------------------|----------------------------------|
| **Цель**              | Найти ошибки и баги                  | Проверить стиль и соглашения    |
| **Примеры проблем**   | Неправильный формат строки           | Отсутствие комментария к функции|
| **Инструменты**       | `go vet`, `staticcheck`              | `golint`, `revive`              |
| **Важность**          | Критично для работы                  | Важно для читаемости            |

---

## **4. Как использовать всё вместе?**

### **4.1 `golangci-lint`**  
Универсальный инструмент, который объединяет множество линтеров и статических анализаторов (включая `staticcheck`, `revive`, `govet`).

**Установка**:
```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

**Запуск**:
```bash
golangci-lint run
```

**Пример вывода**:
```
main.go:7:2: SA4017: result of call to calc is unused (staticcheck)
main.go:3:1: exported function GetData should have comment or be unexported (revive)
```

---

### **4.2 Настройка `.golangci.yml`**  
Создайте файл конфигурации, чтобы выбрать нужные линтеры:

```yaml
linters:
  enable:
    - revive
    - staticcheck
    - govet
```

---

## **5. Примеры ошибок**

### **Статический анализ (`staticcheck`)**  
**Код**:
```go
func main() {
    x := 10
    if x > 5 {
        return // Ошибка: return вне функции
    }
}
```
**Вывод**:
```
main.go:4:9: unreachable code (SA4018)
```

---

### **Линтер (`revive`)**  
**Код**:
```go
func get_user() {} // Стиль: должно быть getUser
```
**Вывод**:
```
main.go:3:1: function name should be in CamelCase (revive)
```

---


## **6. Итог**
- **Статический анализ** ищет **ошибки** в коде.
- **Линтеры** следят за **стилем** и **соглашениями**.
- Используйте `golangci-lint`, чтобы не выбирать между ними — он объединяет оба подхода. 