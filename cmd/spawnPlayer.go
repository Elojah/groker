/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   spawnPlayer.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/06/18 16:11:48 by hdezier           #+#    #+#             */
/*   Updated: 2017/06/18 18:45:02 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package cmd

import (
	"github.com/elojah/groker"
	"github.com/elojah/groker/scylla"
)

func spawnPlayer() {
	scylla.InitCluster(`groker`)
}
