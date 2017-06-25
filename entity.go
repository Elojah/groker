/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   entity.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/06/17 16:40:01 by hdezier           #+#    #+#             */
/*   Updated: 2017/06/18 18:57:56 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package actor

import (
	"github.com/elojah/groker/geometry"
)

const (
	/*
		All skills must have a range < (min(SkillZoneWidth, SkillZoneHeight) / 2 - 1)
	*/
	SkillZoneWidth  = 21
	SkillZoneHeight = 21
)

type Entity interface {
	GetPosition()
}

/*
	TODO !!! Use this type to filter ASAP potential targets only (enemies/allies/etc.)
*/
type ZoneFilter interface {
}

type EntitiesService interface {
	GetActors(ZoneFilter) (chan Actor, error)
	GetObjects(ZoneFilter) (chan Object, error)
}
