package main

import "fmt"

func main() {
	h := &HitPoint{100}
	fmt.Println("ダメージ計算前のヒットポイント: ", h.value)

	damage_amount := calc_damage(5, 2, 3, 4)
	h.damage(damage_amount)
	fmt.Println("ダメージ計算後のヒットポイント: ", h.value)
}

type HitPoint struct {
	value int
}

func (h *HitPoint) damage(damageAmount int) {
	damaged := h.value - damageAmount
	h.value = damaged
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
func calc_damage(playerArmPower, playerWeaponPower, enemyBodyDefence, enemyArmorDefence int) int {
	var totalPlayerAttackPower int = sumUpPlayerAttackPower(playerArmPower, playerWeaponPower)
	var totalEnemyDefence int = sumUpEnemyDefence(enemyBodyDefence, enemyArmorDefence)
	var damageAmount int = estimateDamage(totalPlayerAttackPower, totalEnemyDefence)

	return damageAmount
}
