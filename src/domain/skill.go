/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   skill.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/06/18 15:55:28 by hdezier           #+#    #+#             */
/*   Updated: 2017/06/24 21:42:14 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package scylla

import (
	"github.com/elojah/groker/src/utils/geometry"
)

type SkillID int
type SkillValue int
type Effect int

const (
	DMG Effect = iota
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

const (
	/*
		All skills must have a range < (min(MaxSkillZoneWidth, MaxSkillZoneHeight) / 2 - 1)
	*/
	MaxSkillZoneWidth  = 21
	MaxSkillZoneHeight = 21
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

type MonoSkill struct {
	value  SkillValue
	effect Effect
}

func (MonoSkill) Affect(a Actor) err {
}

type SubSkill struct {
	mono    MonoSkill
	targets Targets
	area    Area
}

type Skill struct {
	id   SkillID `"create":"PRIMARY KEY"`
	subs []SubSkill
}

/*
	TODO !!!
*/
func (s Skill) GetFilters(a Actor) ZoneFilter {
}