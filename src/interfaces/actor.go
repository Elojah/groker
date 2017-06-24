/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   actor.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/06/24 21:16:49 by hdezier           #+#    #+#             */
/*   Updated: 2017/06/24 21:57:31 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package interfaces

type ActorService interface {
	New() (ActorID, error)
	Delete(ActorID) error

	UpdatePosition(ActorID, geometry.Vec2) error
	UpdateStats(ActorID, Stats) error

	AddSkill(ActorID, SkillID, Skill) error
	DeleteSkill(ActorID, SkillID) error
}
