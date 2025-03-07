# Вопрос №10. Генерация случайных величин, имеющих распределение Бернулли

Случайная величина $X$ имеет распределение Бернулли, ели она принимает всего два
значения: $1$ и $0$ с вероятностями $p$ и $q = 1 - p$ соответственно.

Таким образом, $\mathrm{P}\left(X = 1\right) = p$,
$\mathrm{P}\left(X = 0\right) = q$.

> Принято говорить, что событие $\left\{X = 1\right\}$ соответствует «успеху», а
> событие $\left\{X = 0\right\}$ — «неудаче». Эти названия условные, и в
> зависимости от конкретной задачи могут быть заменены на противоположные.

Алгоритм генерации случайной величины, имеющей распределение Бернулли:

1. Выбирается вероятность успеха $p$ ($0 \leqslant p \leqslant 1$)
2. [Генерируется](./question-1.md) равномерно распределённое в
   $\left[0, \, 1\right)$ число $U$.
3. Если $U < p$, то искомая случайная величина $X = 1$, иначе $X = 0$.
