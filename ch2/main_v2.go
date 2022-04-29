package main

func main() {
	calc_damage(1, 2, 3, 4)
}

// 再代入ではなく、目的ごとに意味のある変数を用意
func calc_damage(playerArmPower, playerWeaponPower, enemyBodyDefence, enemyArmorDefence int) {
	var totalPlayerAttackPower int = playerArmPower + playerWeaponPower
	var totalEnemyDefence int = enemyBodyDefence + enemyArmorDefence
	var damageAmount = totalPlayerAttackPower - (totalEnemyDefence / 2)
	if damageAmount < 0 {
		damageAmount = 0
	}
}
