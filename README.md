# Unit Converter README

Этот проект представляет собой консольное приложение на языке Go для конвертации различных единиц измерений. Пользователь может выбрать тип конверсии (масса, длина, площадь, объём, скорость, давление, углы и температура), затем ввести исходную и целевую единицу измерения и значение для конверсии. Программа выполнит вычисления и выведет результат на экран.

# Особенности
Поддерживает несколько типов единиц: масса, длина, площадь, объём, скорость, давление, углы, температура.
Возможность выбора типа единицы измерения и конвертация между различными единицами.
Программа поддерживает конверсию между метрами, километрами, литрами, килограммами, градусами и многими другими единицами.

# Установка и запуск
1. Установите Go
Убедитесь, что у вас установлен Go (версия 1.18 или выше). Для этого следуйте официальной документации Go.

2. Скачайте или клонируйте проект
Вы можете скачать исходный код проекта, либо клонировать репозиторий с помощью git:

   git clone https://github.com/AlexDillz/project_for_practice_2.git

   cd project_for_practice_2

4. Запустите программу
Перейдите в директорию с проектом и запустите программу:

   go run convertor.go
   
5. Использование программы
После запуска программы в терминале будет предложено выбрать тип конверсии. Выберите нужную опцию, затем выберите единицы измерения для конверсии (например, килограммы в граммы) и введите значение для конвертации.

Пример вывода:

Выберите тип конверсии:
1. Масса
2. Длина
3. Площадь
4. Объём
5. Скорость
6. Давление
7. Углы
8. Температура
0. Выход

Введите номер опции: 1
Выберите единицу (введите номер):
1. kilogram
2. gram
3. pound
4. tonne
5. milligram

Введите номер: 1
Выберите единицу (введите номер):
1. kilogram
2. gram
3. pound
4. tonne
5. milligram

Введите номер: 2
Введите значение: 5
5.000000000000000 kilogram = 5000.000000000000000 gram

# Поддерживаемые единицы:

Масса:
kilogram, gram, pound, tonne, milligram

Длина:
meter, kilometer, centimeter, decimeter ,millimeter

Площадь:
square_meter, square_kilometer, square_centimeter, square_decimeter, square_millimeter

Объём:
volume_meter, volume_kilometer, volume_centimeter, volume_decimeter, volume_millimeter, liter

Скорость:
meter_per_seco, kilometer_per_second, kilometer_per_hour, sea_speed, Mach_number

Давление:
pascal, atmosphere, bar, metres_of_the_water_column, millimetres_of_the_mercury_column

Углы:
degrees, radians, grads, turnovers, sec, min, hour

Температура:
Celsius, Fahrenheit, Kelvin

# Важные замечания:
Программа повторяет запросы на ввод, если введены неверные данные (например, некорректное значение или неверный выбор).
В случае ошибок в конверсии будет выведено сообщение о том, что конвертация не поддерживается для выбранных единиц.
