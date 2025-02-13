package handler

import (
	"context"
	"fmt"
	"hash/fnv"
	"strings"
	"time"

	"github.com/micro/micro/v3/service/errors"
	log "github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/service/model"
	"github.com/micro/services/pkg/tenant"
	pb "github.com/micro/services/rss/proto"
)

type Rss struct {
	feeds            model.Model
	entries          model.Model
	feedsIdIndex     model.Index
	feedsNameIndex   model.Index
	entriesDateIndex model.Index
	entriesURLIndex  model.Index
}

func idFromName(name string) string {
	hash := fnv.New64a()
	hash.Write([]byte(name))
	return fmt.Sprintf("%d", hash.Sum64())
}

func NewRss() *Rss {
	idIndex := model.ByEquality("id")
	idIndex.Order.Type = model.OrderTypeUnordered

	nameIndex := model.ByEquality("name")
	nameIndex.Order.Type = model.OrderTypeUnordered

	dateIndex := model.ByEquality("date")
	dateIndex.Order.Type = model.OrderTypeDesc

	entriesURLIndex := model.ByEquality("feed")
	entriesURLIndex.Order.Type = model.OrderTypeDesc
	entriesURLIndex.Order.FieldName = "date"

	f := &Rss{
		feeds: model.NewModel(
			model.WithNamespace("feeds"),
			model.WithIndexes(idIndex, nameIndex),
		),
		entries: model.NewModel(
			model.WithNamespace("entries"),
			model.WithIndexes(dateIndex, entriesURLIndex),
		),
		feedsIdIndex:     idIndex,
		feedsNameIndex:   nameIndex,
		entriesDateIndex: dateIndex,
		entriesURLIndex:  entriesURLIndex,
	}

	// register model instances
	f.feeds.Register(new(pb.Feed))
	f.entries.Register(new(pb.Entry))

	go f.crawl()
	return f
}

func (e *Rss) crawl() {
	e.fetchAll()
	tick := time.NewTicker(1 * time.Minute)
	for _ = range tick.C {
		e.fetchAll()
	}
}

func (e *Rss) Add(ctx context.Context, req *pb.AddRequest, rsp *pb.AddResponse) error {
	log.Info("Received Rss.Add request")

	if len(req.Name) == 0 {
		return errors.BadRequest("rss.add", "require name")
	}

	// get the tenantID
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		tenantID = "micro"
	}

	f := pb.Feed{
		Id:       tenantID + "/" + idFromName(req.Name),
		Name:     req.Name,
		Url:      req.Url,
		Category: req.Category,
	}

	// create the feed
	e.feeds.Create(f)

	// schedule immediate fetch
	go e.fetch(&f)

	return nil
}

func (e *Rss) Feed(ctx context.Context, req *pb.FeedRequest, rsp *pb.FeedResponse) error {
	log.Info("Received Rss.Entries request")

	// get the tenantID
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		tenantID = "micro"
	}

	// feeds by url
	feedUrls := map[string]bool{}
	var feed *pb.Feed

	if len(req.Name) > 0 {
		id := tenantID + "/" + idFromName(req.Name)
		q := model.QueryEquals("ID", id)

		// get the feed
		if err := e.feeds.Read(q, &feed); err != nil {
			return errors.InternalServerError("rss.feeds", "could not read feed")
		}

		feedUrls[feed.Url] = true
	} else {
		// get all the feeds for a user
		var feeds []*pb.Feed
		q := model.QueryAll()
		if err := e.feeds.Read(q, &feeds); err != nil {
			return errors.InternalServerError("rss.feeds", "could not read feed")
		}
		for _, feed := range feeds {
			if !strings.HasPrefix(feed.Id, tenantID+"/") {
				continue
			}
			feedUrls[feed.Url] = true
		}
	}

	if req.Limit == 0 {
		req.Limit = int64(25)
	}

	// default query all
	q := e.entriesDateIndex.ToQuery(nil)

	// if the need is not nil, then use one url
	if feed != nil {
		q = e.entriesURLIndex.ToQuery(feed.Url)
	}

	q.Limit = req.Limit
	q.Offset = req.Offset

	// iterate until entries hits the limit
	for len(rsp.Entries) < int(req.Limit) {
		var entries []*pb.Entry
		// get the entries for each
		err := e.entries.Read(q, &entries)
		if err != nil {
			return errors.InternalServerError("rss.feeds", "could not read feed")
		}

		// find the relevant entries
		for _, entry := range entries {
			// check its a url we care about
			if _, ok := feedUrls[entry.Feed]; !ok {
				continue
			}

			// add the entry
			rsp.Entries = append(rsp.Entries, entry)

			// once you hit the limit return
			if len(rsp.Entries) == int(req.Limit) {
				return nil
			}
		}

		// no more entries or less than the limit
		if len(entries) == 0 || len(entries) < int(req.Limit) {
			return nil
		}

		// increase the offset
		q.Offset += q.Limit
	}

	return nil
}

func (e *Rss) List(ctx context.Context, req *pb.ListRequest, rsp *pb.ListResponse) error {
	var feeds []*pb.Feed
	q := model.QueryAll()

	// TODO: find a way to query only by tenant
	err := e.feeds.Read(q, &feeds)
	if err != nil {
		return errors.InternalServerError("rss.list", "failed to read list of feeds: %v", err)
	}

	// get the tenantID
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		tenantID = "micro"
	}

	for _, feed := range feeds {
		// filter for the tenant
		if !strings.HasPrefix(feed.Id, tenantID+"/") {
			continue
		}

		rsp.Feeds = append(rsp.Feeds, feed)
	}

	return nil
}

func (e *Rss) Remove(ctx context.Context, req *pb.RemoveRequest, rsp *pb.RemoveResponse) error {
	if len(req.Name) == 0 {
		return errors.BadRequest("rss.remove", "blank name provided")
	}

	// get the tenantID
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		tenantID = "micro"
	}

	id := tenantID + "/" + idFromName(req.Name)

	e.feeds.Delete(model.QueryEquals("ID", id))
	return nil
}
