/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   entity.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/06/18 14:02:37 by hdezier           #+#    #+#             */
/*   Updated: 2017/06/24 21:26:49 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package scylla

import (
	"github.com/elojah/groker/src/utils/geometry"
	"github.com/gocql/gocql"
)

type UUID int

/*
	Implement: Actor, Entity
*/
type Actor struct {
	id       UUID `"create":"PRIMARY KEY"`
	position geometry.Point
	stats    Stats
	skills   []Skill
}

/*
	Implement: Object, Entity
*/
type Object struct {
	id       UUID `"create":"primary key"`
	position geometry.Point
	region   geometry.Rect `"index":"true"`
}
