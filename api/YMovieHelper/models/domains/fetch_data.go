package domains

import (
	"encoding/json"
	"fmt"

	"github.com/itkmaingit/YMovieHelper/utils"
)

func (item *DynamicItem) FetchData() error {

	body, err := utils.ReadFile(item.ItemUrl)
	if err != nil {
		return fmt.Errorf("DynamicItem.FetchData: %w", err)
	}
	err = json.Unmarshal(body, &item.Item)
	if err != nil {
		return fmt.Errorf("DynamicItem.FetchData: %w", err)
	}

	return err
}

func (item *SingleItem) FetchData() error {

	body, err := utils.ReadFile(item.ItemUrl)
	if err != nil {
		return fmt.Errorf("SingleItem.FetchData: %w", err)
	}
	err = json.Unmarshal(body, &item.Item)
	if err != nil {
		return fmt.Errorf("SingleItem.FetchData: %w", err)
	}

	return err
}

func (item *MultipleItem) FetchData() error {
	var items []*ItemOnYMMP

	body, err := utils.ReadFile(item.ItemUrl)
	if err != nil {
		return fmt.Errorf("MultipleItem.FetchData: %w", err)
	}
	err = json.Unmarshal(body, &items)
	if err != nil {
		return fmt.Errorf("MultipleItem.FetchData: %w", err)
	}

	item.Items = items

	return err
}

func FetchEmotionData(data map[string]map[string]string) (map[string]map[string]*TachieFaceParameter, error) {
	emotionMapData := make(map[string]map[string]*TachieFaceParameter)
	for characterName, emotionMap := range data {
		_, ok := emotionMapData[characterName]
		if !ok {
			innerMap := make(map[string]*TachieFaceParameter)
			emotionMapData[characterName] = innerMap
		}

		for emotionKey, emotionUrl := range emotionMap {
			var emotion *TachieFaceParameter

			body, err := utils.ReadFile(emotionUrl)
			if err != nil {
				return nil, fmt.Errorf("FetchEmotionData: %w", err)
			}

			err = json.Unmarshal(body, &emotion)
			if err != nil {
				return nil, fmt.Errorf("FetchEmotionData: %w", err)
			}

			emotionMapData[characterName][emotionKey] = emotion

		}
	}

	return emotionMapData, nil
}
