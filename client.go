package ogetit

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	TYPE_SPY_REPORT string = "sr"
)

type Client interface {
	GetCombatReport(string) (*CombatReport, error)
	GetSpyReport(string) (*SpyReport, error)
	GetRecycleReport(string) (*RecycleReport, error)
	GetServerData(int, string) (*ServerData, error)
	Close() error
}

type OGetIt struct {
	ApiKey  string
	Version string
}

func NewClient(apiKey string) (Client, error) {
	return &OGetIt{
		ApiKey:  apiKey,
		Version: "v1",
	}, nil
}

func ParseKey(reportId string) (string, string, int, string, error) {
	parts := strings.SplitN(reportId, "-", 4)

	if len(parts) != 4 {
		return "", "", 0, "", errors.New("key format error")
	}

	universeID, err := strconv.Atoi(parts[2])
	if err != nil {
		return "", "", 0, "", errors.New("universe is not a number")
	}

	return parts[0], parts[1], universeID, parts[3], nil
}

func (o *OGetIt) GetCombatReport(reportKey string) (*CombatReport, error) {
	keytype, community, universeID, reportID, err := ParseKey(reportKey)
	if err != nil {
		return nil, err
	}
	if keytype != "cr" {
		return nil, fmt.Errorf("invalid key type: %s", keytype)
	}

	result := &CombatReport{}
	return result, o.get(community, universeID, "combat", keytype, reportID, result)
}

func (o *OGetIt) GetRecycleReport(reportKey string) (*RecycleReport, error) {
	keytype, community, universeID, reportID, err := ParseKey(reportKey)
	if err != nil {
		return nil, err
	}
	if keytype != "rr" {
		return nil, fmt.Errorf("invalid key type: %s", keytype)
	}

	result := &RecycleReport{}
	return result, o.get(community, universeID, "recycle", keytype, reportID, result)
}

func (o *OGetIt) GetSpyReport(reportKey string) (*SpyReport, error) {
	keytype, community, universeID, reportID, err := ParseKey(reportKey)
	if err != nil {
		return nil, err
	}
	if keytype != TYPE_SPY_REPORT {
		return nil, fmt.Errorf("invalid key type: %s", keytype)
	}

	result := &SpyReport{}
	return result, o.get(community, universeID, "spy", keytype, reportID, result)
}

func (o *OGetIt) GetSpyReportFromJSON(file string) (*SpyReport, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	result := &SpyReport{}
	return result, json.Unmarshal(data, result)
}

func (o *OGetIt) GetMissleReport() {

}

func (o *OGetIt) GetServerData(universeID int, community string) (*ServerData, error) {
	url := fmt.Sprintf("https://s%d-%s.ogame.gameforge.com/api/serverData.xml", universeID, community)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	serverData := &ServerData{}
	return serverData, xml.Unmarshal(data, serverData)
}

func (o *OGetIt) Close() error {
	return nil
}

func (o *OGetIt) get(community string, universeID int, reportType, keyType, reportID string, result interface{}) error {
	url := o.getUrl(universeID, community, reportType, keyType, reportID)
	logrus.Debugf("getting url: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, result)
}

func (o *OGetIt) getUrl(universeId int, community, reportType, keyType, reportID string) string {
	return fmt.Sprintf(
		"https://s%d-%s.ogame.gameforge.com/api/v1/%s/report?api_key=%s&%s_id=%s",
		universeId,
		community,
		reportType,
		o.ApiKey,
		keyType,
		reportID,
	)
}
