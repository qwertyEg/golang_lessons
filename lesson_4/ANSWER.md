# Решение задания: Проверка палиндромов

## **Шаг 1: Создание проекта**
1. Создайте папку для проекта:
   ```bash
   mkdir palindrome && cd palindrome
   ```

2. Инициализируйте Go-модуль:
   ```bash
   go mod init github.com/ваш-никнейм/palindrome
   ```
   Это создаст файл `go.mod` (для управления зависимостями).

---

## **Шаг 2: Написание функции**
1. Создайте файл `palindrome.go`:
   ```bash
   touch palindrome.go
   ```

2. Вставьте код функции:
   ```go
   package palindrome

   import "strings"

   // IsPalindrome проверяет, является ли строка палиндромом (игнорирует регистр и пробелы).
   func IsPalindrome(s string) bool {
       s = strings.ToLower(s)
       s = strings.ReplaceAll(s, " ", "")

       for i := 0; i < len(s)/2; i++ {
           if s[i] != s[len(s)-1-i] {
               return false
           }
       }
       return true
   }
   ```

---

## **Шаг 3: Написание тестов**
1. Создайте файл `palindrome_test.go`:
   ```bash
   touch palindrome_test.go
   ```

2. Вставьте код тестов:
   ```go
   package palindrome

   import "testing"

   func TestIsPalindrome(t *testing.T) {
       tests := []struct {
           name  string
           input string
           want  bool
       }{
           {"Пустая строка", "", true},
           {"Один символ", "a", true},
           {"Палиндром (нечетная длина)", "level", true},
           {"Палиндром (четная длина)", "abba", true},
           {"Не палиндром", "hello", false},
           {"Строка с регистром", "Level", true},
           {"Палиндром с пробелами", "A man a plan a canal Panama", true},
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

## **Шаг 4: Запуск тестов**
1. Запустите все тесты:
   ```bash
   go test -v
   ```

2. **Пример успешного вывода:**
   ```
   === RUN   TestIsPalindrome
   === RUN   TestIsPalindrome/Пустая_строка
   === RUN   TestIsPalindrome/Один_символ
   === RUN   TestIsPalindrome/Палиндром_(нечетная_длина)
   === RUN   TestIsPalindrome/Палиндром_(четная_длина)
   === RUN   TestIsPalindrome/Не_палиндром
   === RUN   TestIsPalindrome/Строка_с_регистром
   === RUN   TestIsPalindrome/Палиндром_с_пробелами
   --- PASS: TestIsPalindrome (0.00s)
       --- PASS: TestIsPalindrome/Пустая_строка (0.00s)
       --- PASS: TestIsPalindrome/Один_символ (0.00s)
       --- PASS: TestIsPalindrome/Палиндром_(нечетная_длина) (0.00s)
       --- PASS: TestIsPalindrome/Палиндром_(четная_длина) (0.00s)
       --- PASS: TestIsPalindrome/Не_палиндром (0.00s)
       --- PASS: TestIsPalindrome/Строка_с_регистром (0.00s)
       --- PASS: TestIsPalindrome/Палиндром_с_пробелами (0.00s)
   PASS
   ok      github.com/ваш-никнейм/palindrome    0.002s
   ```

---

## **Шаг 5: Проверка покрытия кода**
Узнайте, какая часть кода покрыта тестами:
```bash
go test -cover
```
**Пример вывода:**
```
PASS
coverage: 100.0% of statements
ok      github.com/ваш-никнейм/palindrome    0.002s
```

---

## **Шаг 6: Дополнительные команды**
- **Запуск одного теста**:
  ```bash
  go test -v -run TestIsPalindrome/Палиндром_с_пробелами
  ```

- **Просмотр покрытия в HTML**:
  ```bash
  go test -coverprofile=coverage.out
  go tool cover -html=coverage.out
  ```
  Откроется браузер с визуализацией покрытия.

---

## **Шаг 8: Возможные ошибки и исправления**
### **Ошибка 1: Файлы не найдены**
Убедитесь, что файлы `palindrome.go` и `palindrome_test.go` находятся в одной папке.

### **Ошибка 2: Неправильный импорт пакета**
Если вы видите ошибку `undefined: IsPalindrome`, проверьте:
- Оба файла объявлены в пакете `palindrome`.
- Функция `IsPalindrome` начинается с заглавной буквы (экспортирована). 