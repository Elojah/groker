/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   actor.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/06/18 17:34:50 by hdezier           #+#    #+#             */
/*   Updated: 2017/06/29 15:40:56 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package groker

import (
	"github.com/elojah/groker/utils/geometry"
)

type ActorID int

type Actor struct {
	Id       ActorID `"create":"PRIMARY KEY"`
	Position geometry.Point
	Stats    Stats
	Skills   []Skill `"id_only":"true"`
}

func (a Actor) GetSurroundRect() geometry.Rect {
	return geometry.Rect{
		O: geometry.Point{
			X: a.Position.X - SkillZoneWidth/2,
			Y: a.Position.Y - SkillZoneHeight/2,
		},
		W: SkillZoneWidth,
		H: SkillZoneHeight,
	}
}

func (a Actor) IsEnemy(enemy Actor) bool {
	/*
		TODO Implement some complex shit with guilds and all
	*/
	return true
}

func (a Actor) Cast(sID SkillID, as ActorService, ss SkillService) error {
	skill, err := ss.Get(sID)
	// First filter on surround targets
	targets := as.Get(skill.GetFilters(a))
	if err != nil {
		return err
	}
	// Double loop, be careful here...
	for subSkill := range skill.subs {
		for target := range ActorFilter[subSkill.Targets](a, targets) {
			s := subSkill
			victim := &a
			// We suppose here counterskill can be on caster ONLY
			// Skillback send a MonoSkill to apply on caster and not a Skill
			for ; ; s != nil {
				if victim == &a {
					victim = &target
				} else {
					victim = &a
				}
				s, err = *victim.Take(s)
				if err != nil {
					return err
				}
			}
		}
	}
}

func (a Actor) Take(ms MonoSkill) (skillback MonoSkill, err error) {
	a.Stats.Apply(ms)
}

type ActorService interface {
	New() (ActorID, error)
	Delete(ActorID) error

	UpdatePosition(ActorID, geometry.Vec2) error
	UpdateStats(ActorID, Stats) error

	AddSkill(ActorID, SkillID, Skill) error
	DeleteSkill(ActorID, SkillID) error

	Cast(ActorID, SkillID) error
	Receive(ActorID, MonoSkill) error

	Get([]ZoneFilter) chan Actor
}
