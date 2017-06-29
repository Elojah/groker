/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   types.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/06/17 16:25:46 by hdezier           #+#    #+#             */
/*   Updated: 2017/06/29 15:07:27 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package geometry

type Val int

type Point struct {
	X Val
	Y Val
}

type Rect struct {
	O Point
	H Val
	W Val
}

type Vec2 struct {
	X Val
	Y Val
}
