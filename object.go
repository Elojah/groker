/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   object.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/06/18 17:35:00 by hdezier           #+#    #+#             */
/*   Updated: 2017/06/18 18:50:30 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package groker

import (
	"github.com/elojah/groker/utils/geometry"
)

type ObjectID int

type Object struct {
	id       ObjectID
	position geometry.Point
	region   geometry.Rect `"index":"true"`
}

type ObjectService interface {
}
