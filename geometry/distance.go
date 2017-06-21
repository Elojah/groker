/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   distance.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/06/18 19:12:56 by hdezier           #+#    #+#             */
/*   Updated: 2017/06/18 20:05:45 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package geometry

/*
	QuickRelativeDistance should be use to compare distances only
*/
func abs(a Val) Val {
	if a > 0 {
		return a
	} else {
		return (-a)
	}
}
func QuickRelativeDistance(a Point, b Point) Val {
	return (abs(a.x-b.x) + abs(a.y-b.y))
}
