/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   skill.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/06/18 15:55:28 by hdezier           #+#    #+#             */
/*   Updated: 2017/06/29 15:16:46 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package groker

import (
	"github.com/elojah/groker/utils/geometry"
)

type Effect int

const (
	NONE Effect = iota
	DMG
	HEAL
	DRAIN
	CONTROL
	MOVE
)

type Shape int

const (
	RECT Shape = iota
	CIRCLE
	CONE
)

/*
Area must cover all skill shapes possible. Need Shape + some attributes, add if necessary
*/
type Area struct {
	shape Shape
	o     geometry.Point
	w     geometry.Val
	h     geometry.Val
}

/*
	Value of dmg, heal, etc.
*/
type SkillValue int

type MonoSkill struct {
	Value  SkillValue
	Effect Effect
}

type SubSkill struct {
	Mono    MonoSkill
	Targets Targets
	Area    Area
}

const (
	/*
		All skills must have a range < (min(SkillZoneWidth, SkillZoneHeight) / 2 - 1)
	*/
	SkillZoneWidth  = 21
	SkillZoneHeight = 21
)

type SkillID int

type Skill struct {
	Id   SkillID `"create":"PRIMARY KEY"`
	Subs []SubSkill
}

/*
	TODO !!! Use this type to filter ASAP potential targets only (enemies/allies/etc.)
*/
type ZoneFilter interface {
}

/*
	TODO !!!
*/
func (s Skill) GetFilters(a Actor) []ZoneFilter {
	return []ZoneFilter{}
}

type SkillService interface {
	New([]SubSkill) (Skill, error)
	Update(SkillID, []SubSkill) (SkillID, error)
	Delete(SkillID) error
	Get(SkillID) (Skill, error)
}
