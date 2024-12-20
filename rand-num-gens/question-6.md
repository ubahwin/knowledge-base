# Вопрос №6. Критерий сериальной корреляции для проверки генератора случайных чисел

Критерий сериальной корреляции применяется для проверки независимости
генерируемых значений от предыдущих с некоторым лагом $q$. При этом проверяется
гипотеза о независимости двух последовательностей
$\left\{Y_i\right\}=\left\{U_q, U_{q+1}, \ldots, U_n, U_1, U_2, \ldots, U_{q-1}\right\}$
и $\left\{X_i\right\} = \left\{U_1, U_2, \ldots, U_n\right\}$.

Для проверки независимости находят коэффициент корреляции

$$
r_{XY} = \dfrac
{
n \cdot \sum\limits_{i = 1}^{n} \left(X_i \cdot Y_i\right) - \left(\sum\limits_{i=1}^{n} X_i\right) \cdot \left(\sum\limits_{i=1}^{n} Y_i\right)
}
{
\sqrt{
n \cdot \sum\limits_{i = 1}^{n} X_i^2 - \left(\sum\limits_{i = 1}^{n} X_i\right)^2
}
\cdot
\sqrt{
n \cdot \sum\limits_{i = 1}^{n} Y_i^2 - \left(\sum\limits_{i = 1}^{n} Y_i\right)^2
}
}
$$

Чтобы указанные последовательности можно было считать независимыми, коэффициент
корреляции должен быть близок к нулю.

Из книги Д. Кнута «Искусство программирования, том 2. Получисленные алгоритмы» и
<...> известно, что с вероятностью $0.95$ значение коэффициента корреляции,
приемлемое для независимости, находится в интервале
$\left(\upmu_n - 2 \sigma_n, \, \upmu_n + 2  \sigma_n\right)$, где
$\upmu_n = \dfrac{-1}{n-1}$,
$\sigma_n^2 = \dfrac{n^2}{\left(n-1\right)^2 \left(n - 2\right)}$.

<!-- TODO: откуда ещё? ... — [8] -->
