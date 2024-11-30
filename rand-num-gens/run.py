import math

# Заданные значения

for Z in range(0, 100):
    Z /= 100

    for U in range(1, 100):
        U /= 100

        # print(Z, U)

        # Постоянная e
        e = math.e

        # Шаг 1: Вычисляем V
        V = math.sqrt(8 / e) * (Z - 0.5)

        # Шаг 2: Вычисляем X
        X = V / U

        # Шаг 3: Проверяем условия
        # Верхняя грань
        upper_bound = 5 - 4 * U * e**(1 / 4)
        upper_check = X**2 <= upper_bound

        # Нижняя грань
        lower_bound = (4 * e**(-1.35)) / U + 1.4
        lower_check = X**2 <= lower_bound

        # Основная проверка
        log_check = X**2 <= -4 * math.log(U)
        #print(upper_check, lower_check)
        if upper_check and lower_check:
            print(U, V, upper_check, lower_check, log_check)
