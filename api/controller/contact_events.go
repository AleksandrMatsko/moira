package controller

import (
	"fmt"

	"github.com/moira-alert/moira"
	"github.com/moira-alert/moira/api"
	"github.com/moira-alert/moira/api/dto"
)

func GetContactEventsHistoryById(database moira.Database, contactID string, from int64, to int64, page int64, size int64,
) (*dto.ContactEventItemList, *api.ErrorResponse) {
	events, err := database.GetNotificationsHistoryByContactId(contactID, from, to, page, size)
	if err != nil {
		return nil, api.ErrorInternalServer(fmt.Errorf("GetContactEventsHistoryById: can't get notifications for contact with id %v", contactID))
	}

	eventsList := dto.ContactEventItemList{
		List: make([]dto.ContactEventItem, 0, len(events)),
	}
	for _, i := range events {
		contactEventItem := &dto.ContactEventItem{
			TimeStamp: i.TimeStamp,
			Metric:    i.Metric,
			State:     string(i.State),
			OldState:  string(i.OldState),
			TriggerID: i.TriggerID,
		}
		eventsList.List = append(eventsList.List, *contactEventItem)
	}

	return &eventsList, nil
}
