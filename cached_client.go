package ogetit

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

type CachedClient struct {
	client Client
	server *cache.Cache
	report *cache.Cache
}

func NewCachedClient(apikey string, reportTTL, serverTTL time.Duration) (Client, error) {
	client, err := NewClient(apikey)
	if err != nil {
		return nil, err
	}

	return &CachedClient{
		client: client,
		server: cache.New(serverTTL, reportTTL*2),
		report: cache.New(reportTTL, reportTTL*2),
	}, nil
}

func (cc *CachedClient) GetCombatReport(reportKey string) (*CombatReport, error) {
	var report *CombatReport
	var err error

	if key, ok := cc.server.Get(reportKey); !ok {
		logrus.Debugf("cache miss for %s", reportKey)
		report, err = cc.client.GetCombatReport(reportKey)
		if err != nil {
			return nil, err
		}
		cc.server.Add(reportKey, report, cache.DefaultExpiration)
	} else {
		logrus.Infof("cache hit for %s", reportKey)
		report = key.(*CombatReport)
	}

	return report, nil
}

func (cc *CachedClient) GetSpyReport(reportKey string) (*SpyReport, error) {
	var report *SpyReport
	var err error

	if key, ok := cc.server.Get(reportKey); !ok {
		logrus.Debugf("cache miss for %s", reportKey)
		report, err = cc.client.GetSpyReport(reportKey)
		if err != nil {
			return nil, err
		}
		cc.server.Add(reportKey, report, cache.DefaultExpiration)
	} else {
		logrus.Infof("cache hit for %s", reportKey)
		report = key.(*SpyReport)
	}

	return report, nil
}

func (cc *CachedClient) GetRecycleReport(reportKey string) (*RecycleReport, error) {
	var report *RecycleReport
	var err error

	if key, ok := cc.server.Get(reportKey); !ok {
		logrus.Debugf("cache miss for %s", reportKey)
		report, err = cc.client.GetRecycleReport(reportKey)
		if err != nil {
			return nil, err
		}
		cc.server.Add(reportKey, report, cache.DefaultExpiration)
	} else {
		logrus.Infof("cache hit for %s", reportKey)
		report = key.(*RecycleReport)
	}

	return report, nil
}

func (cc *CachedClient) GetServerData(universeID int, community string) (*ServerData, error) {
	cacheKey := fmt.Sprintf("%s_%d", community, universeID)
	var serverData *ServerData
	var err error

	if key, ok := cc.server.Get(cacheKey); !ok {
		logrus.Debugf("cache miss for %s", cacheKey)
		serverData, err = cc.client.GetServerData(universeID, community)
		if err != nil {
			return nil, err
		}
		cc.server.Add(cacheKey, serverData, cache.DefaultExpiration)
	} else {
		logrus.Infof("cache hit for %s", cacheKey)
		serverData = key.(*ServerData)
	}

	return serverData, nil
}

func (cc *CachedClient) Close() error {
	return cc.client.Close()
}
