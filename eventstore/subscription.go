package eventstore

import (
  "log"
  "encoding/json"
  "fmt"
  "errors"
)

type Subscription struct{
  StreamName                  string
  SubscriptionName            string
  MinCheckPointCount          int
  StartFrom                   int
  ResolveLinkTos              bool
  ReadBatchSize               int
  NamedConsumerStrategy       string
  ExtraStatistics             bool
  MaxRetryCount               int
  LiveBufferSize              int
  MessageTimeoutMilliseconds  int
  MaxCheckPointCount          int
  MaxSubscriberCount          int
  CheckPointAfterMilliseconds int
  BufferSize                  int
}

func (client *Client) GetSubscription(streamName string, subscriptionName string) (*Subscription, error) {
  var data map[string]interface{}
  err := client.makeRequest("GET", fmt.Sprintf("/subscriptions/%s/%s/info", streamName, subscriptionName), nil, &data)

  if err != nil {
    log.Print(err)
    return nil, err
  }

  return &Subscription{
    StreamName:                  data["eventStreamId"].(string),
    SubscriptionName:            data["groupName"].(string),
    MinCheckPointCount:          int(data["config"].(map[string]interface{})["minCheckPointCount"].(float64)),
    StartFrom:                   int(data["config"].(map[string]interface{})["startFrom"].(float64)),
    ResolveLinkTos:              data["config"].(map[string]interface{})["resolveLinktos"].(bool),
    ReadBatchSize:               int(data["config"].(map[string]interface{})["readBatchSize"].(float64)),
    NamedConsumerStrategy:       data["config"].(map[string]interface{})["namedConsumerStrategy"].(string),
    ExtraStatistics:             data["config"].(map[string]interface{})["extraStatistics"].(bool),
    MaxRetryCount:               int(data["config"].(map[string]interface{})["maxRetryCount"].(float64)),
    LiveBufferSize:              int(data["config"].(map[string]interface{})["liveBufferSize"].(float64)),
    MessageTimeoutMilliseconds:  int(data["config"].(map[string]interface{})["messageTimeoutMilliseconds"].(float64)),
    MaxCheckPointCount:          int(data["config"].(map[string]interface{})["maxCheckPointCount"].(float64)),
    MaxSubscriberCount:          int(data["config"].(map[string]interface{})["maxSubscriberCount"].(float64)),
    CheckPointAfterMilliseconds: int(data["config"].(map[string]interface{})["checkPointAfterMilliseconds"].(float64)),
    BufferSize:                  int(data["config"].(map[string]interface{})["bufferSize"].(float64)),
  }, nil
}  

func (client *Client) CreateSubscription(streamName                  string,
                                         subscriptionName            string,
                                         minCheckPointCount          int,
                                         startFrom                   int,
                                         resolveLinkTos              bool,
                                         readBatchSize               int,
                                         namedConsumerStrategy       string,
                                         extraStatistics             bool,
                                         maxRetryCount               int,
                                         liveBufferSize              int,
                                         messageTimeoutMilliseconds  int,
                                         maxCheckPointCount          int,
                                         maxSubscriberCount          int,
                                         checkPointAfterMilliseconds int,
                                         bufferSize                  int) (*Subscription, error) {

  userData, err := json.Marshal(map[string]interface{}{
    "minCheckPointCount" :         minCheckPointCount,
    "startFrom":                   startFrom,
    "ResolveLinkTos":              resolveLinkTos,
    "readBatchSize":               readBatchSize,
    "namedConsumerStrategy":       namedConsumerStrategy,
    "extraStatistics":             extraStatistics,
    "maxRetryCount":               maxRetryCount,
    "liveBufferSize":              liveBufferSize,
    "messageTimeoutMilliseconds":  messageTimeoutMilliseconds,
    "maxCheckPointCount":          maxCheckPointCount,
    "maxSubscriberCount":          maxSubscriberCount,
    "checkPointAfterMilliseconds": checkPointAfterMilliseconds,
    "bufferSize":                  bufferSize,
  })

  if err != nil {
    log.Print(err)
    return nil, err
  }

  var data map[string]interface{}
  err = client.makeRequest("PUT", fmt.Sprintf("/subscriptions/%s/%s", streamName, subscriptionName), userData, &data)

  if err != nil {
    log.Print(err)
    return nil, err
  }

  if data["result"].(string) != "Success" {
    log.Print(data["reason"].(string))
    return nil, errors.New(data["reason"].(string))
  }

  return client.GetSubscription(streamName, subscriptionName)
}

func (client *Client) UpdateSubscription(streamName                  string,
                                         subscriptionName            string,
                                         minCheckPointCount          int,
                                         startFrom                   int,
                                         resolveLinkTos              bool,
                                         readBatchSize               int,
                                         namedConsumerStrategy       string,
                                         extraStatistics             bool,
                                         maxRetryCount               int,
                                         liveBufferSize              int,
                                         messageTimeoutMilliseconds  int,
                                         maxCheckPointCount          int,
                                         maxSubscriberCount          int,
                                         checkPointAfterMilliseconds int,
                                         bufferSize                  int) (*Subscription, error) {

  userData, err := json.Marshal(map[string]interface{}{
    "minCheckPointCount" :         minCheckPointCount,
    "startFrom":                   startFrom,
    "ResolveLinkTos":              resolveLinkTos,
    "readBatchSize":               readBatchSize,
    "namedConsumerStrategy":       namedConsumerStrategy,
    "extraStatistics":             extraStatistics,
    "maxRetryCount":               maxRetryCount,
    "liveBufferSize":              liveBufferSize,
    "messageTimeoutMilliseconds":  messageTimeoutMilliseconds,
    "maxCheckPointCount":          maxCheckPointCount,
    "maxSubscriberCount":          maxSubscriberCount,
    "checkPointAfterMilliseconds": checkPointAfterMilliseconds,
    "bufferSize":                  bufferSize,
  })

  if err != nil {
    log.Print(err)
    return nil, err
  }

  var data map[string]interface{}
  err = client.makeRequest("POST", fmt.Sprintf("/subscriptions/%s/%s", streamName, subscriptionName), userData, &data)

  if err != nil {
    log.Print(err)
    return nil, err
  }

  if data["result"].(string) != "Success" {
    log.Print(data["reason"].(string))
    return nil, errors.New(data["reason"].(string))
  }

  return client.GetSubscription(streamName, subscriptionName)
}

func (client *Client) DeleteSubscription(streamName string, subscriptionName string) (bool, error) {
  var data map[string]interface{}
  err := client.makeRequest("DELETE", fmt.Sprintf("/subscriptions/%s/%s", streamName, subscriptionName), nil, &data)

  if err != nil {
   log.Print(err)
   return false, err
  }

  if data["result"].(string) != "Success" {
    return false, nil
  }

  return true, nil
}
