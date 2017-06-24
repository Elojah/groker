/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   stats.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/06/18 16:50:43 by hdezier           #+#    #+#             */
/*   Updated: 2017/06/24 21:42:22 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package groker

type StatValue int

const (
	MaxHp = 10000
)

type Stats struct {
	Atk         StatValue
	Def         StatValue
	Hp          StatValue
	Mana        StatValue
	AtkSpeed    StatValue
	MovSpeed    StatValue
	Penetration StatValue
}
