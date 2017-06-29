/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   stats.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/06/18 16:50:43 by hdezier           #+#    #+#             */
/*   Updated: 2017/06/29 11:46:44 by hdezier          ###   ########.fr       */
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

const (
	nullSkill = MonoSkill{
		Effect: NONE,
		Value:  0,
	}
)

func (s *Stats) Apply(ms MonoSkill) (MonoSkill, error) {
	// TODO, use a table instead
	switch ms.Effect {
	case (NONE):
		return nullSkill, nil
	case (DMG):
		s.Hp = s.Hp - ms.Value
		return nullSkill, nil
	case (HEAL):
		s.Hp = s.Hp + ms.Value
		return nullSkill, nil
	case (DRAIN):
		s.Hp = s.Hp - ms.Value
		ms := MonoSkill{
			Effect: HEAL,
			Value:  ms.Value,
		}
		return ms, nil
	case (CONTROL):
		/*TODO*/
		return nullSkill, nil
	case (MOVE):
		/*TODO*/
		return nullSkill, nil
	}
}
