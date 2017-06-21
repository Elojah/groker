/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   serializer.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/06/18 16:19:33 by hdezier           #+#    #+#             */
/*   Updated: 2017/06/18 16:51:39 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package scylla

import ()

type fields int

const (
	CREATE fields = iota
	INSERT
	SELECT
)

func Serialize(i interface{}, f fields) string {
}
