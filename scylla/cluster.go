/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   cluster.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/06/17 18:53:31 by hdezier           #+#    #+#             */
/*   Updated: 2017/06/18 16:36:38 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package scylla

import (
	"github.com/gocql/gocql"
)

func InitCluster(keyspace string) (*gocql.Session, error) {
	// HARDCODED cluster adress (docker scylla default)
	cluster = gocql.NewCluster("172.17.0.2")
	cluster.Keyspace = keyspace
	// HARDCODED cluster consistency
	cluster.Consistency = gocql.One
	return cluster.CreateSession()
}

func Insert(session *gocql.Session, i interface{}, table string) (err error) {
	query := Serialize(i, INSERT)
	return session.Query(query).Consistency(gocql.One).Exec()
}
