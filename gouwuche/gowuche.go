package main

import "fmt"

var (
	firstNumber     string
	Commodity       string
	Commodity_money int
	YES             string
)

func main() {
LABEL1:
	firstNumber := map[string]int{
		"iphone":  6288,
		"sanxing": 5288,
		"xiaomi":  1999,
		"huawei":  3999}
	fmt.Print("商店有下列商品:\n", firstNumber, "\n", "请输入你要购买的物品:")
	fmt.Scanln(&Commodity)
	fmt.Print("你选择购买的商品是:", Commodity)
	Commodity_price := firstNumber[Commodity]
	fmt.Print("它的价格是:", Commodity_price)

	if Commodity_money == 0 {
		fmt.Print("你的账户的金钱是:0.00￥,", "是否要充值[y/n]", "(提示:最少充值金额大于或等于你选择商品的价格):")
		fmt.Scanln(&YES)
		fmt.Scanln(&Commodity_money)
		if YES != "y" {
			fmt.Println("您没有充值,返回商品列表")
			goto LABEL1
		} else {
			fmt.Println("你已经成功充值:", Commodity_money)
		}
	}

	if Commodity_price > Commodity_money {
		fmt.Print("你充值的金钱少于商品的价格", "请你重新充值,最少充值金额大于或等于你选择商品的价格:")
		fmt.Scanln(&Commodity_money)
	}

	for Commodity_price <= Commodity_money {
		fmt.Print("是否要将", Commodity, "添加到购物车[y/n]:")
		fmt.Scanln(&YES)
		if YES != "y" {
			fmt.Println("你成功取消该商品的购买！")
			break
		} else {
			fmt.Println("已将", Commodity, "成功添加到购物车中")
			break
		}
	}
	fmt.Println("你已经成功购买", Commodity, "商品")
}
