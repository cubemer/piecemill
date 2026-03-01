package main

import (
	"math"
	"net/http"
	"sort"
	"strconv"

	"github.com/pocketbase/pocketbase/core"
)

type machineScore struct {
	MachineID int
	Score     float64
}

func handleRecommend(e *core.RequestEvent) error {
	templateID := e.Request.URL.Query().Get("template_id")
	quantityStr := e.Request.URL.Query().Get("quantity")
	rankMode := e.Request.URL.Query().Get("rank")

	if templateID == "" || quantityStr == "" {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"error": "missing required parameters: template_id and quantity",
		})
	}

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil || quantity < 1 {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"error": "quantity must be a positive integer",
		})
	}

	// Look up the template.
	tmpl, err := e.App.FindFirstRecordByData("templates", "template_id", templateID)
	if err != nil {
		return e.JSON(http.StatusNotFound, map[string]string{
			"error": "template_id not found",
		})
	}

	cutTimeAt25 := tmpl.GetFloat("cut_time_at_25")

	// Fetch all machines.
	machines, err := e.App.FindAllRecords("machines")
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to fetch machines",
		})
	}

	// Build a map of machine_id -> speed.
	speedMap := make(map[int]float64)
	for _, m := range machines {
		mid := int(m.GetFloat("machine_id"))
		spd := m.GetFloat("speed")
		speedMap[mid] = spd
	}

	// Fetch all queue entries.
	queueEntries, err := e.App.FindAllRecords("queue_entries")
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to fetch queue entries",
		})
	}

	// Calculate committed time per machine.
	committedTime := make(map[int]float64)
	for _, entry := range queueEntries {
		entryMID := int(entry.GetFloat("machine_id"))
		entryTID := entry.GetString("template_id")
		entryQty := entry.GetFloat("quantity")

		// Look up this entry's template cut_time_at_25.
		entryTmpl, err := e.App.FindFirstRecordByData("templates", "template_id", entryTID)
		if err != nil {
			continue // skip entries with unknown templates
		}

		entryCutTime := entryTmpl.GetFloat("cut_time_at_25")
		speed := speedMap[entryMID]
		if speed == 0 {
			continue
		}

		actualTime := entryCutTime / (speed / 25.0)
		committedTime[entryMID] += actualTime * entryQty
	}

	// Score each machine: committed + incoming.
	scores := make([]machineScore, 0, 12)
	for mid := 1; mid <= 12; mid++ {
		speed := speedMap[mid]
		if speed == 0 {
			speed = 25 // fallback
		}

		incomingTime := (cutTimeAt25 / (speed / 25.0)) * float64(quantity)
		total := committedTime[mid] + incomingTime

		scores = append(scores, machineScore{
			MachineID: mid,
			Score:     total,
		})
	}

	// Sort by score ascending, then by machine_id ascending for tie-breaking.
	sort.SliceStable(scores, func(i, j int) bool {
		if math.Abs(scores[i].Score-scores[j].Score) < 1e-9 {
			return scores[i].MachineID < scores[j].MachineID
		}
		return scores[i].Score < scores[j].Score
	})

	// Build result.
	var result []int
	if rankMode == "full" {
		for _, s := range scores {
			result = append(result, s.MachineID)
		}
	} else {
		// Return top result(s) — all machines tied for the lowest score.
		bestScore := scores[0].Score
		for _, s := range scores {
			if math.Abs(s.Score-bestScore) < 1e-9 {
				result = append(result, s.MachineID)
			} else {
				break
			}
		}
	}

	return e.JSON(http.StatusOK, result)
}
