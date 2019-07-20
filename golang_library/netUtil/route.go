package main

import "sync"
import log "github.com/golang/glog"


type Route struct {
	appid  int64
	mutex   sync.Mutex
	clients map[uint64]ClientSet
	room_clients map[int64]ClientSet
}

func NewRoute(appid int64) *Route {
	route := new(Route)
	route.appid = appid
	route.clients = make(map[uint64]ClientSet)
	route.room_clients = make(map[int64]ClientSet)
	return route
}

func (route *Route) AddRoomClient(room_id int64, client *Client) {
	route.mutex.Lock()
	defer route.mutex.Unlock()
	set, ok := route.room_clients[room_id];
	if !ok {
		set = NewClientSet()
		route.room_clients[room_id] = set
	}
	set.Add(client)
}

//todo optimise client set clone
func (route *Route) FindRoomClientSet(room_id int64) ClientSet {
	route.mutex.Lock()
	defer route.mutex.Unlock()

	set, ok := route.room_clients[room_id]
	if ok {
		return set.Clone()
	} else {
		return nil
	}
}

func (route *Route) RemoveRoomClient(room_id int64, client *Client) bool {
	route.mutex.Lock()
	defer route.mutex.Unlock()
	if set, ok := route.room_clients[room_id]; ok {
		set.Remove(client)
		if set.Count() == 0 {
			delete(route.room_clients, room_id)
		}
		return true
	}
	log.Info("room client non exists")
	return false
}

func (route *Route) AddClient(client *Client) {
	route.mutex.Lock()
	defer route.mutex.Unlock()
	set, ok := route.clients[client.connID];
	if !ok {
		set = NewClientSet()
		route.clients[client.connID] = set
	}
	set.Add(client)
}

func (route *Route) RemoveClient(client *Client) bool {
	route.mutex.Lock()
	defer route.mutex.Unlock()
	if set, ok := route.clients[client.connID]; ok {
		set.Remove(client)
		if set.Count() == 0 {
			delete(route.clients, client.connID)
		}
		return true
	}
	log.Info("client non exists")
	return false
}

func (route *Route) FindClientSet(connID uint64) ClientSet {
	route.mutex.Lock()
	defer route.mutex.Unlock()

	set, ok := route.clients[connID]
	if ok {
		return set.Clone()
	} else {
		return nil
	}
}

func (route *Route) IsOnline(connID uint64) bool {
	route.mutex.Lock()
	defer route.mutex.Unlock()

	set, ok := route.clients[connID]
	if ok {
		return len(set) > 0
	}
	return false
}


func (route *Route) GetUserIDs() IntSet {
	return NewIntSet()
}
