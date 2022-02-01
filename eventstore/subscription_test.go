package eventstore

import (
  "testing"
)

func TestGetSubscription(t *testing.T) {

  client, _ := NewClient("http://eventstore.hostname:2113", "myuser", "mypass")

  httpClient := NewMockHttpClient()
  client.httpClient = httpClient

  httpClient.addHttpClientResponse(`
    {
      "config": {
        "resolveLinktos": true,
        "startFrom": 0,
        "messageTimeoutMilliseconds": 30000,
        "extraStatistics": false,
        "maxRetryCount": 10,
        "liveBufferSize": 500,
        "bufferSize": 1000,
        "readBatchSize": 500,
        "preferRoundRobin": true,
        "checkPointAfterMilliseconds": 2000,
        "minCheckPointCount": 10,
        "maxCheckPointCount": 1000,
        "maxSubscriberCount": 10,
        "namedConsumerStrategy": "RoundRobin"
      },
      "eventStreamId": "stream-1",
      "groupName":     "subscription-1"
    }`, 200, "200 OK", nil)

  client.GetSubscription("stream-1", "subscription-1")
}

func TestCreateSubscription(t *testing.T) {

  client, _ := NewClient("http://eventstore.hostname:2113", "myuser", "mypass")

  httpClient := NewMockHttpClient()
  client.httpClient = httpClient

  httpClient.addHttpClientResponse(`
    {
      "result": "Success",
      "reason": ""
    }`, 200, "200 OK", nil)

  httpClient.addHttpClientResponse(`
    {
      "config": {
        "resolveLinktos": true,
        "startFrom": 0,
        "messageTimeoutMilliseconds": 30000,
        "extraStatistics": false,
        "maxRetryCount": 10,
        "liveBufferSize": 500,
        "bufferSize": 1000,
        "readBatchSize": 500,
        "preferRoundRobin": true,
        "checkPointAfterMilliseconds": 2000,
        "minCheckPointCount": 10,
        "maxCheckPointCount": 1000,
        "maxSubscriberCount": 10,
        "namedConsumerStrategy": "RoundRobin"
      },
      "eventStreamId": "stream-1",
      "groupName":     "subscription-1"
    }`, 200, "200 OK", nil)

  client.CreateSubscription(
    "stream-1",
    "subscription-1",
    2,
    0,
    true,
    5,
    "RoundRobin",
    true,
    7,
    1,
    3,
    2,
    9,
    6,
    5)
}

func TestUpdateSubscription(t *testing.T) {

  client, _ := NewClient("http://eventstore.hostname:2113", "myuser", "mypass")

  httpClient := NewMockHttpClient()
  client.httpClient = httpClient

  httpClient.addHttpClientResponse(`
    {
      "result": "Success",
      "reason": ""
    }`, 200, "200 OK", nil)

  httpClient.addHttpClientResponse(`
    {
      "config": {
        "resolveLinktos": true,
        "startFrom": 0,
        "messageTimeoutMilliseconds": 30000,
        "extraStatistics": false,
        "maxRetryCount": 10,
        "liveBufferSize": 500,
        "bufferSize": 1000,
        "readBatchSize": 500,
        "preferRoundRobin": true,
        "checkPointAfterMilliseconds": 2000,
        "minCheckPointCount": 10,
        "maxCheckPointCount": 1000,
        "maxSubscriberCount": 10,
        "namedConsumerStrategy": "RoundRobin"
      },
      "eventStreamId": "stream-1",
      "groupName":     "subscription-1"
    }`, 200, "200 OK", nil)

  client.UpdateSubscription(
    "stream-1",
    "subscription-1",
    2,
    0,
    true,
    5,
    "RoundRobin",
    true,
    7,
    1,
    3,
    2,
    9,
    6,
    5)
}


func TestDeleteSubscription(t *testing.T) {

  client, _ := NewClient("http://eventstore.hostname:2113", "myuser", "mypass")

  httpClient := NewMockHttpClient()
  client.httpClient = httpClient

  httpClient.addHttpClientResponse(`
    {
      "result": "Success",
      "reason": ""
    }`, 200, "200 OK", nil)

  client.DeleteSubscription("stream-1","subscription-1")
}
