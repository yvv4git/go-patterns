# TIMEOUT / ЗАДЕРЖКА

В 1991 году работая в Sun Microsystems Питер Дойч сформулировал заблуждения распределенных систем:
1. Сеть надежна.
2. Задержка равна нулю.
3. Пропускная способность бесконечна.
4. Сеть безопасна.
5. Топология не меняется.
6. Существует один администратор.
7. Транспортные расходы равны нулю.
8. Сеть однородна.

В действительности так и происходит, любое звено распределенной системы может выйти из строя, следовательно, целесообразнее
завершить работу некоторой функции по истечению заданного интервала времени, вместо того, чтобы бесконечно ждать. 