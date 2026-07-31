package main

import (
	"fmt"
	"math/big"
)

func factorialRecursionCalk(currentArg, targetArg int, acc *big.Int, resultsChan chan *big.Int) *big.Int {
	acc.Mul(acc, big.NewInt(int64(currentArg)))
	if currentArg == targetArg {
		resultsChan <- acc
		// Прерываем рекурсию.
		// Иначе на следующем шаге условие if currentArg == targetArg уже никогда не выполнится (так как аргумент стал больше),
		// и рекурсия будет уходить в бесконечность, пока не закончится память (Stack Overflow), а главный поток main бесконечно ждёт ответа.
		return acc
	}
	return factorialRecursionCalk(currentArg+1, targetArg, acc, resultsChan)
}

func main() {
	argFactorial := 45

	// Вычисление argFactorial! используя горутины.
	// Вычисляем факториал до середны диапазона:
	middle := argFactorial / 2

	// Создаем канал для передачи больших чисел (*big.Int)
	resultsChan := make(chan *big.Int)

	// 1. Вычисление факториала от 1 до 22 в рекурсии.
	// Можем начать вычисление с 2! т.к. третий параметр (аккумулятор) уже несет факториал 1 (1)
	go factorialRecursionCalk(2, middle, big.NewInt(1), resultsChan)

	// 2. Вычисление промежуточное произведением диапазона от 23 до 45 в цикле.
	// В англоязычной литературе это называется partial product или range product.
	// Сам по себе факториал всегда начинается с единицы.
	// Но из-за свойства ассоциативности умножения мы можем считать кусками: 45!=res1*res2
	go factorialCalk(middle, argFactorial, resultsChan)

	// Главный поток (main) ждет и забирает оба значения из канала
	// Первая прочитанная переменная получит то число, которое посчиталось БЫСТРЕЕ
	res1 := <-resultsChan
	res2 := <-resultsChan

	// 5. Перемножаем две половины между собой для получения финального 45!
	finalResult := new(big.Int).Mul(res1, res2)

	fmt.Printf("Factorial %d is equal to:\n%s\n", argFactorial, finalResult.String())

}

func factorialCalk(border int, argFactorial int, resultsChan chan *big.Int) {
	acc := big.NewInt(1)                          // Начинаем с 1, чтобы первым значением было 23
	for i := border + 1; i <= argFactorial; i++ { // от 23 до 45
		acc.Mul(acc, big.NewInt(int64(i))) // Умножаем строго на i (23, 24... 45)
	}
	resultsChan <- acc
}
