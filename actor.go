/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   actor.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/06/18 17:34:50 by hdezier           #+#    #+#             */
/*   Updated: 2017/06/18 20:25:13 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package groker

import (
	"github.com/elojah/groker/geometry"
)

type ActorID int

type Actor struct {
	id       ActorID `"create":"PRIMARY KEY"`
	position geometry.Point
	stats    Stats
	skills   []Skill `"id_only":"true"`
}

func (a Actor) GetSurroundRect() geometry.Rect {
	return geometry.Rect{
		o: Point{
			x: a.position.x - SkillZoneWidth/2,
			y: a.position.y - SkillZoneHeight/2,
		},
		w: SkillZoneWidth,
		h: SkillZoneHeight,
	}
}

func (a Actor) IsEnemy(enemy Actor) bool {
	/*
		TODO Implement some complex shit with guilds and all
	*/
	return true
}

func (a Actor) Exec(sID SkillID, es EntitiesService, ss SkillService) error {
	skill, err := ss.Get(sID)
	targets := e.GetActors(skill.GetFilters(a))
	if err != nil {
		return err
	}
	// Double loop, be careful here...
	for subSkill := range skill.subs {
		for target := range ActorFilter[subSkill.targets](a, targets, IsInZone(subSkill.area)) {

		}
	}
}

type ActorService interface {
	New() (ActorID, error)
	Delete(ActorID) error

	UpdatePosition(ActorID, geometry.Vec2) error
	UpdateStats(ActorID, Stats) error

	AddSkill(ActorID, SkillID, Skill) error
	DeleteSkill(ActorID, SkillID) error

	Exec(ActorID, SkillID) error
	Receive(ActorID, MonoSkill) error
}
