package models

import (
	"crypto/md5"
	"fmt"
)

//Attribute ...
type Attribute struct {
	number int
	object *XMLTree
}

//Item ...
type Item struct {
	number        int
	object        *XMLTree
	attributesMap map[string]*Attribute
}

//Channel ...
type Channel struct {
	object   *XMLTree
	itemsMap map[string]*Item
}

//RSSFeed ...
type RSSFeed struct {
	RSSFeedXMLTree *XMLTree
	RSSFeedChannel *Channel
}

//RssFeedFromXML ...
func RssFeedFromXML(rootXML *XMLTree) *RSSFeed {
	rssFeed := new(RSSFeed)
	rssFeed.RSSFeedXMLTree = rootXML
	rssFeed.RSSFeedChannel = new(Channel)
	channel := rssFeed.RSSFeedChannel
	channel.itemsMap = make(map[string]*Item)
	channel.object = rootXML.Branches[0]
	channelTree := channel.object
	for i, channelBranch := range channelTree.Branches {
		item := new(Item)
		item.number = i
		item.object = channelBranch
		item.attributesMap = make(map[string]*Attribute)
		for j, itemBranch := range channelBranch.Branches {
			attribute := new(Attribute)
			attribute.number = j
			attribute.object = itemBranch
			item.attributesMap[HashOfString(itemBranch.Name)] = attribute
		}
		var key string
		if channelBranch.Name == "item" {
			key = item.attributesMap[HashOfString("guid")].object.Value
		} else {
			key = channelBranch.Name
		}
		channel.itemsMap[HashOfString(key)] = item
	}
	return rssFeed
}

//HashOfString ...
func HashOfString(s string) string {
	contentBytes := []byte(s)
	return fmt.Sprintf("%x", md5.Sum(contentBytes))
}

//RssFeedCompareString ...
func RssFeedCompareString(oldRSSFeed string, newRSSFeed string) *XMLTree {
	xmlTree1 := ReadXMLFromString(oldRSSFeed)
	xmlTree2 := ReadXMLFromString(newRSSFeed)
	rss1 := RssFeedFromXML(xmlTree1)
	rss2 := RssFeedFromXML(xmlTree2)
	RssFeedCompare(*rss1, rss2)
	return rss2.RSSFeedXMLTree
}

//RssFeedCompare ...
func RssFeedCompare(oldRSSFeed RSSFeed, newRSSFeed *RSSFeed) {
	newRSSFeedChannel := newRSSFeed.RSSFeedChannel
	oldRSSFeedChannel := oldRSSFeed.RSSFeedChannel

	if newRSSFeed.RSSFeedXMLTree.ContentHash == oldRSSFeed.RSSFeedXMLTree.ContentHash {
		newRSSFeed.RSSFeedXMLTree.Status = OLD
		newRSSFeedChannel.object.Status = OLD
	} else {
		newRSSFeed.RSSFeedXMLTree.Status = NEW
		newRSSFeedChannel.object.Status = NEW
	}
	for id, newRSSFeedItem := range newRSSFeedChannel.itemsMap {
		oldRSSFeedItem, oldItemStatus := oldRSSFeedChannel.itemsMap[id]
		// true если данные есть в двух rssFeed
		if oldItemStatus {
			// проверка на несовпадение содержания в двух item-x
			if oldRSSFeedItem.object.ContentHash != newRSSFeedItem.object.ContentHash {
				if newRSSFeedItem.number != oldRSSFeedItem.number {
					newRSSFeedItem.object.Status = CHANGEDVALUEORDER
				} else {
					newRSSFeedItem.object.Status = CHANGEDVALUE
				}
				for jd, newRSSFeedAttr := range newRSSFeedItem.attributesMap {
					oldRSSFeedAttr, oldAttrStatus := oldRSSFeedItem.attributesMap[jd]
					// true если данные есть в двух rssFeed
					if oldAttrStatus {
						// проверка на совпадение содержания в двух атрибутах
						if newRSSFeedAttr.object.ContentHash != oldRSSFeedAttr.object.ContentHash {
							newRSSFeedAttr.object.Status = NEW
						} else {
							if newRSSFeedAttr.number != oldRSSFeedAttr.number {
								newRSSFeedAttr.object.Status = CHANGEDORDER
							} else {
								newRSSFeedAttr.object.Status = OLD
							}
						}
					} else {
						newRSSFeedAttr.object.Status = NEW
					}
				}
				for jd, oldRSSFeedAttr := range oldRSSFeedItem.attributesMap {
					_, newAttrStatus := newRSSFeedItem.attributesMap[jd]
					// если данные отсутствуют
					if !newAttrStatus {
						currentItem := newRSSFeedItem.object
						deletedAttr := new(XMLTree)
						*deletedAttr = *oldRSSFeedAttr.object
						deletedAttr.Status = DELETED
						currentItem.Branches = append(currentItem.Branches, deletedAttr)
					}
				}
			} else {
				if newRSSFeedItem.number != oldRSSFeedItem.number {
					newRSSFeedItem.object.Status = CHANGEDORDER
				} else {
					newRSSFeedItem.object.Status = OLD
				}
			}
		} else {
			newRSSFeedItem.object.Status = NEW
		}
	}
	//добавляем удаленные item-ы c статусом DELETED
	for id, oldRSSFeedItem := range oldRSSFeedChannel.itemsMap {
		_, newItemStatus := newRSSFeedChannel.itemsMap[id]
		// если данные отсутствуют
		if !newItemStatus {
			channel := newRSSFeed.RSSFeedXMLTree.Branches[0]
			deletedItem := new(XMLTree)
			*deletedItem = *oldRSSFeedItem.object
			deletedItem.Status = DELETED
			channel.Branches = append(channel.Branches, deletedItem)
		}

	}
}
