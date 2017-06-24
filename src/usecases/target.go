/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   target.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/06/18 19:41:14 by hdezier           #+#    #+#             */
/*   Updated: 2017/06/24 21:58:07 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package groker

import (
	d "github.com/elojah/groker/src/domain"
	"github.com/elojah/groker/src/utils/geometry"
)

type Targets int

const (
	SELF Targets = iota
	CLOSEST_ENEMY
	LOWEST_HP_ENEMY
	ALL_ENEMIES
)

/*
	Used to implement more complex actor filtering
*/

type filterSkillOptions func(Actor, Actor) bool
type filterActorFn func(Actor, chan Actor, ...filterSkillOptions) chan Actor

func filterOptions(a Actor, target Actor, options ...filterSkillOptions) bool {
	for _, filterOpt := range options {
		if !filterOpt(a, target) {
			return false
		}
	}
	return true
}

func isInZone(a Actor, target Actor) bool {

}

const (
	ActorFilter = map[Targets]filterActorFn{

		SELF: func(a Actor, targets chan Actor, options ...filterSkillOptions) chan Actor {
			<-a
		},

		CLOSEST_ENEMY: func(a Actor, targets chan Actor, options ...filterSkillOptions) chan Actor {
			// We compare with QuickRelativeDistance so distance must be at least this
			closestdistance := SkillZoneWidth + SkillZoneHeight + 1
			closestTarget := nil
			for target := range targets {
				if !filterOptions(a, target, options...) {
					continue
				}
				if !a.IsEnemy(target) {
					continue
				}
				distance := geometry.QuickRelativeDistance(a.Position, target.position)
				if distance < closestdistance {
					closestTarget = &target
					closestdistance = distance
				}
			}
			if closestTarget != nil {
				<-*closestTarget
			}
		},

		LOWEST_HP_ENEMY: func(a Actor, targets chan Actor, options ...filterSkillOptions) chan Actor {
			lowestHP = MaxHp
			lowestTarget := nil
			for target := range targets {
				if !filterOptions(a, target, options...) {
					continue
				}
				if !a.IsEnemy(target) {
					continue
				}
				if target.stats.Hp < lowestHP {
					lowestTarget = &target
					lowestHp = target.stats.Hp
				}
			}
			if lowestTarget != nil {
				<-*lowestTarget
			}
		},

		ALL_ENEMIES: func(a Actor, targets chan Actor, options ...filterSkillOptions) chan Actor {
			for target := range targets {
				if !filterOptions(a, target, options...) {
					continue
				}
				if a.IsEnemy(target) {
					<-target
				}
			}
		},
	}
)
