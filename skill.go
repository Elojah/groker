/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   skill.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/06/18 15:55:28 by hdezier           #+#    #+#             */
/*   Updated: 2017/06/18 20:23:22 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package scylla

import (
	"github.com/elojah/groker/geometry"
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

type SkillService interface {
	New([]SubSkill) (Skill, error)
	Update(SkillID, []SubSkill) (SkillID, error)
	Delete(SkillID) error
	Get(SkillID) (Skill, error)
}
