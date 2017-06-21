/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   types.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/06/17 16:25:46 by hdezier           #+#    #+#             */
/*   Updated: 2017/06/18 20:04:48 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package geometry

type Val int

type Point struct {
	x Val
	y Val
}

type Rect struct {
	o Point
	h Val
	w Val
}

type Vec2 struct {
	x Val
	y Val
}
