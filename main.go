package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Human - родительская структура
type Human struct {
	Name      string
	Age       int
	BirthDate time.Time
}

// Вспомогательные функции для склонений
func getAgeSuffix(age int) string {
	lastDigit := age % 10
	lastTwoDigits := age % 100

	if lastTwoDigits >= 11 && lastTwoDigits <= 14 {
		return "лет"
	}

	switch lastDigit {
	case 1:
		return "год"
	case 2, 3, 4:
		return "года"
	default:
		return "лет"
	}
}

func getDaysSuffix(days int) string {
	lastDigit := days % 10
	lastTwoDigits := days % 100

	if lastTwoDigits >= 11 && lastTwoDigits <= 14 {
		return "дней"
	}

	switch lastDigit {
	case 1:
		return "день"
	case 2, 3, 4:
		return "дня"
	default:
		return "дней"
	}
}

// Методы Human
func (h *Human) SayHello() string {
	return fmt.Sprintf("Привет, меня зовут %s!", h.Name)
}

func (h *Human) Introduce() string {
	ageSuffix := getAgeSuffix(h.Age)
	return fmt.Sprintf("Я %s, мне %d %s. Родился %s.",
		h.Name, h.Age, ageSuffix, h.BirthDate.Format("02.01.2006"))
}

func (h *Human) IsAdult() bool {
	return h.Age >= 18
}

// Расчет дней до дня рождения
func (h *Human) DaysUntilBirthday() int {
	now := time.Now()

	// Дата рождения в текущем году
	birthdayThisYear := time.Date(now.Year(), h.BirthDate.Month(), h.BirthDate.Day(), 0, 0, 0, 0, time.Local)

	// Если день рождения уже прошел в этом году
	if now.After(birthdayThisYear) {
		birthdayNextYear := time.Date(now.Year()+1, h.BirthDate.Month(), h.BirthDate.Day(), 0, 0, 0, 0, time.Local)
		days := int(birthdayNextYear.Sub(now).Hours() / 24)
		return days
	}

	// День рождения еще не наступил
	days := int(birthdayThisYear.Sub(now).Hours() / 24)
	return days
}

// Информация о дне рождения
func (h *Human) GetBirthdayInfo() string {
	daysUntil := h.DaysUntilBirthday()
	daysSuffix := getDaysSuffix(daysUntil)

	if daysUntil == 0 {
		return "🎂 Сегодня ваш день рождения! Поздравляем!"
	}

	nextAge := h.Age + 1
	ageSuffix := getAgeSuffix(nextAge)

	return fmt.Sprintf("До вашего %d-%s осталось %d %s", nextAge, ageSuffix, daysUntil, daysSuffix)
}

// Определение знака зодиака
func (h *Human) GetZodiacSign() string {
	day := h.BirthDate.Day()
	month := h.BirthDate.Month()

	switch month {
	case time.January:
		if day < 20 {
			return "Козерог ♑"
		}
		return "Водолей ♒"
	case time.February:
		if day < 19 {
			return "Водолей ♒"
		}
		return "Рыбы ♓"
	case time.March:
		if day < 21 {
			return "Рыбы ♓"
		}
		return "Овен ♈"
	case time.April:
		if day < 20 {
			return "Овен ♈"
		}
		return "Телец ♉"
	case time.May:
		if day < 21 {
			return "Телец ♉"
		}
		return "Близнецы ♊"
	case time.June:
		if day < 21 {
			return "Близнецы ♊"
		}
		return "Рак ♋"
	case time.July:
		if day < 23 {
			return "Рак ♋"
		}
		return "Лев ♌"
	case time.August:
		if day < 23 {
			return "Лев ♌"
		}
		return "Дева ♍"
	case time.September:
		if day < 23 {
			return "Дева ♍"
		}
		return "Весы ♎"
	case time.October:
		if day < 23 {
			return "Весы ♎"
		}
		return "Скорпион ♏"
	case time.November:
		if day < 22 {
			return "Скорпион ♏"
		}
		return "Стрелец ♐"
	case time.December:
		if day < 22 {
			return "Стрелец ♐"
		}
		return "Козерог ♑"
	default:
		return "Неизвестно"
	}
}

// Структура с встраиванием Human
type Action struct {
	Human

	Occupation string
	Skills     []string
	IsActive   bool
}

// Методы Action
func (a *Action) Work() string {
	if a.IsActive {
		return fmt.Sprintf("%s работает %s", a.Name, a.Occupation)
	}
	return fmt.Sprintf("%s сейчас не работает", a.Name)
}

func (a *Action) DisplaySkills() string {
	if len(a.Skills) == 0 {
		return fmt.Sprintf("%s пока не имеет навыков", a.Name)
	}
	return fmt.Sprintf("Навыки %s: %s", a.Name, strings.Join(a.Skills, ", "))
}

func (a *Action) SpecialIntroduce() string {
	baseIntro := a.Introduce()
	return fmt.Sprintf("%s Я работаю %s. %s",
		baseIntro, a.Occupation, a.DisplaySkills())
}

// Функции для ввода данных
func readString(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func readInt(prompt string) int {
	for {
		input := readString(prompt)
		value, err := strconv.Atoi(input)
		if err == nil {
			return value
		}
		fmt.Println("❌ Ошибка! Введите целое число.")
	}
}

func readDate(prompt string) time.Time {
	for {
		input := readString(prompt + " (формат: ДД.ММ.ГГГГ): ")
		date, err := time.Parse("02.01.2006", input)
		if err == nil {
			return date
		}
		fmt.Println("❌ Ошибка! Используйте формат ДД.ММ.ГГГГ (например: 15.03.1990)")
	}
}

func readSkills() []string {
	var skills []string
	fmt.Println("\n💡 Введите навыки (для завершения введите 'готово'):")

	for {
		skill := readString("Навык: ")
		if strings.ToLower(skill) == "готово" {
			break
		}
		if skill != "" {
			skills = append(skills, skill)
		}
	}
	return skills
}

func readYesNo(prompt string) bool {
	for {
		input := strings.ToLower(readString(prompt + " (да/нет): "))
		if input == "да" || input == "д" {
			return true
		}
		if input == "нет" || input == "н" {
			return false
		}
		fmt.Println("❌ Ошибка! Введите 'да' или 'нет'.")
	}
}

func main() {
	fmt.Println("👤 Введите ваши данные:")
	fmt.Println("========================")

	// Ввод основных данных
	name := readString("Имя: ")
	age := readInt("Возраст: ")
	birthDate := readDate("Дата рождения")

	// Ввод профессиональных данных
	fmt.Println("\n💼 Введите профессиональные данные:")
	fmt.Println("========================")

	occupation := readString("Кем Вы работаете? ")
	skills := readSkills()
	isActive := readYesNo("Сейчас работаете?")

	// Создаем экземпляр Action
	person := Action{
		Human: Human{
			Name:      name,
			Age:       age,
			BirthDate: birthDate,
		},
		Occupation: occupation,
		Skills:     skills,
		IsActive:   isActive,
	}

	// Демонстрация возможностей
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("🎯 РЕЗУЛЬТАТЫ:")
	fmt.Println(strings.Repeat("=", 50))

	// Методы Human
	fmt.Println("1.", person.SayHello())
	fmt.Println("2.", person.Introduce())

	// Информация о дне рождения
	fmt.Println("3.", person.GetBirthdayInfo())
	fmt.Printf("4. Знак зодиака: %s\n", person.GetZodiacSign())

	if person.IsAdult() {
		fmt.Println("5. ✅ Совершеннолетний")
	} else {
		fmt.Println("5. ❌ Несовершеннолетний")
	}

	// Методы Action
	fmt.Println("6.", person.Work())
	fmt.Println("7.", person.DisplaySkills())
	fmt.Println("8.", person.SpecialIntroduce())

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("🎉 ПРОГРАММА ЗАВЕРШЕНА!")
	fmt.Println("📊 Финальные данные:")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println(person.SpecialIntroduce())

	ageSuffix := getAgeSuffix(person.Age)
	fmt.Printf("🎂 Возраст: %d %s\n", person.Age, ageSuffix)

	fmt.Printf("📅 %s\n", person.GetBirthdayInfo())
	fmt.Printf("♈ Знак зодиака: %s\n", person.GetZodiacSign())
}
