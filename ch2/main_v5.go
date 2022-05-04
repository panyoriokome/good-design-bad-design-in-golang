package main

import "fmt"

func main() {
	calc_damage(5, 2, 3, 4)
}

type HitPoint struct {
	value int
}

func (h *HitPoint) damage(damageAmount int) {
	damaged = h.value - damageAmount

}

func sumUpPlayerAttackPower(playerArmPower int, playerWeaponPower int) int {
	return playerArmPower + playerWeaponPower
}

func sumUpEnemyDefence(enemyBodyDefence int, enemyArmorDefence int) int {
	return enemyBodyDefence + enemyArmorDefence
}

func estimateDamage(totalPlayerAttackPower int, totalEnemyDefence int) int {
	damageAmount := totalPlayerAttackPower - (totalEnemyDefence / 2)
	if damageAmount < 0 {
		damageAmount = 0
	}

	fmt.Println(damageAmount)
	return damageAmount
}

// ベタ書きせず、意味のあるまとまりでメソッド化
func calc_damage(playerArmPower, playerWeaponPower, enemyBodyDefence, enemyArmorDefence int) {
	var totalPlayerAttackPower int = sumUpPlayerAttackPower(playerArmPower, playerWeaponPower)
	var totalEnemyDefence int = sumUpEnemyDefence(enemyBodyDefence, enemyArmorDefence)
	var damageAmount int = estimateDamage(totalPlayerAttackPower, totalEnemyDefence)

	fmt.Println(damageAmount)
}
