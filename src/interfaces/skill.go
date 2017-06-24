/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   skill.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/06/24 21:20:08 by hdezier           #+#    #+#             */
/*   Updated: 2017/06/24 21:38:56 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package interfaces

type SkillService interface {
	New([]SubSkill) (Skill, error)
	Update(SkillID, []SubSkill) (SkillID, error)
	Delete(SkillID) error
	Get(SkillID) (Skill, error)
}
