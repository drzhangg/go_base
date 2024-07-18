package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type PrometheusQueryResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric map[string]string `json:"metric"`
			Value  []interface{}     `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

func formatSize(size float64) string {
	const (
		_  = 1 << (10 * iota)
		KiB
		MiB
		GiB
		TiB
	)

	switch {
	case size >= TiB:
		return fmt.Sprintf("%.2fTi", size/TiB)
	case size >= GiB:
		return fmt.Sprintf("%.2fGi", size/GiB)
	case size >= MiB:
		return fmt.Sprintf("%.2fMi", size/MiB)
	default:
		return fmt.Sprintf("%.2fKi", size/KiB)
	}
}

func prometheusQuery(query string) ([]byte, error) {
	promUrl := ""

	queryURL := fmt.Sprintf("%s?query=%s", promUrl, url.QueryEscape(query))
	fmt.Println("queryUrl::", queryURL)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", queryURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Cookie", "Hm_lvt_9a20840d3cf815b336728226955921de=1713523334; Hm_lvt_f099bd07c3f6962ce977f4884adce005=1713523338; _ga=GA1.2.524612032.1713523338; Hm_lpvt_9a20840d3cf815b336728226955921de=1713523442; pt_s_77ebaa58=vt=1713523442358&cad=; Hm_lpvt_f099bd07c3f6962ce977f4884adce005=1713523442; _ga_EK35F0NKZ2=GS1.2.1713523340.1.1.1713523447.55.0.0; pt_77ebaa58=uid=FUzdP8zDfdS1HaqpNmWAZQ&nid=0&vid=0D0fbt4dRZoQFmfxnGaFfQ&vn=1&pvn=4&sact=1713523447164&to_flag=1&pl=lHBHCKnZ033KtDLvKHulQw*pt*1713523442358; tfstk=fKBoMbwpxEgsQcmLrtv73pvkczPAV49BSwHpJpLUgE8XpgIKpsvhyws-pbIe8vbcAUKpeM8h-ZLXpUBRw6m5nNDKezN5YvvpLPUTBRISVp9UW03MsgYWAi8UfKg5Ng9BLre0hlJAVNYgqZNkLov2vhiyL38zmE-HuY-EUHu4mHT28p8eLmy2jhTEaHRefAn2yw753ldylzbvisBDqIVA4Eo2dtxkSFSyivkEb3Ak7g80WJ_4CBbkZteo53fVjLtlRRMWUBfP_LW3S8vVEaXyiwonipBl0TKf-clR0B5AFC68uYvN1gpRTZrELef2YgomgfSDtvt4vtlIOQ-XmF6gATL6gZc-Woqm_1Oyc3yTmocI2Q-XmFE0mfW2antzB; token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYmYiOjE3MTk5OTEyMzcsInJvbGUiOiJhZG1pbiIsInVzZXJfaWQiOiIxMTA1MiIsInVzZXJuYW1lIjoiMTEwNTIifQ.Rt3qvgn2gUqweCzkfjtHwQG1fz2xcA3bo81tdW1-GUo")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func getDiskMetrics(query string) (map[string]float64, error) {
	body, err := prometheusQuery(query)
	if err != nil {
		return nil, err
	}

	var queryResponse PrometheusQueryResponse
	err = json.Unmarshal(body, &queryResponse)
	if err != nil {
		return nil, err
	}

	metrics := make(map[string]float64)
	for _, result := range queryResponse.Data.Result {
		node := result.Metric["instance"]
		value, _ := result.Value[1].(string)
		val, _ := strconv.ParseFloat(value, 64)
		metrics[node] = val
	}

	return metrics, nil
}

func main() {
	totalDiskQuery := `sum(node_filesystem_size_bytes) by (instance)`
	usedDiskQuery := `sum(node_filesystem_size_bytes - node_filesystem_free_bytes) by (instance)`

	totalDisk, err := getDiskMetrics(totalDiskQuery)
	if err != nil {
		fmt.Println("Error getting total disk metrics:", err)
		return
	}

	usedDisk, err := getDiskMetrics(usedDiskQuery)
	if err != nil {
		fmt.Println("Error getting used disk metrics:", err)
		return
	}

	for node, total := range totalDisk {
		used := usedDisk[node]
		usagePercent := (used / total) * 100

		fmt.Printf("Node: %s\n", node)
		fmt.Printf("  Total Disk: %s\n", formatSize(total))
		fmt.Printf("  Used Disk: %s (%.2f%%)\n", formatSize(used), usagePercent)
	}
}
