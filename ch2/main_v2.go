package main

func main() {
	calc_damage(1, 2, 3, 4)
}

// 変数名を全て意味のある単語にすることで、コードを見ただけでロジックの意味がわかるように変更
func calc_damage(playerArmPower, playerWeaponPower, enemyBodyDefence, enemyArmorDefence int) {
	var damageAmount int = 0
	damageAmount = playerArmPower + playerWeaponPower
	damageAmount = damageAmount - ((enemyBodyDefence + enemyArmorDefence) / 2)
	if damageAmount < 0 {
		damageAmount = 0
	}
}
