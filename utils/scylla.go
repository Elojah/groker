/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   scylla.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/06/16 18:46:43 by hdezier           #+#    #+#             */
/*   Updated: 2017/06/18 15:46:09 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
)

func main() {
	// connect to the cluster
	cluster := gocql.NewCluster("172.17.0.2")
	cluster.Keyspace = "gotest"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// if err := session.Query(`CONSISTENCY ONE`).Exec(); err != nil {
	// 	fmt.Println(`[error] during consistency edition`)
	// 	log.Fatal(err)
	// }

	// insert a tweet
	if err := session.Query(`INSERT INTO tweet (id, text) VALUES (?, ?)`,
		gocql.TimeUUID(),
		"hello world",
	).Consistency(gocql.One).Exec(); err != nil {
		fmt.Println(`[error] during insert`)
		log.Fatal(err)
	}

	var id gocql.UUID
	var text string
	/*
	 Search for a specific set of records whose 'timeline' column matches
	 * the value 'me'. The secondary index that we created earlier will be
	 * used for optimizing the search
	*/
	if err := session.Query(`SELECT id, text FROM tweet LIMIT 1`).Consistency(gocql.One).Scan(&id, &text); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tweet:", id, text)
}
