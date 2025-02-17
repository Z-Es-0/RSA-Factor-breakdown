/*
 * @Author: Z-Es-0 zes18642300628@qq.com
 * @Date: 2025-02-16 13:31:31
 * @LastEditors: Z-Es-0 zes18642300628@qq.com
 * @LastEditTime: 2025-02-17 19:23:19
 * @FilePath: \Cryptography_tools\RSA\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"math/big"
	"os"
	"rsa/factorbreakdown"
)

func ErrorInput() {
	fmt.Println("请提供一个整数参数")
	os.Exit(1)
}

func main() {

	if len(os.Args) < 2 {
		ErrorInput()
	}

	arg0 := os.Args[1]
	// arg1 := os.Args[2]

	n, _ := new(big.Int).SetString(arg0, 10)

	worksize := 100000

	if n == nil {
		ErrorInput()
	}

	fmt.Println("n is", n)

	factorbreakdown.BuildFactory(worksize, *n)

}
