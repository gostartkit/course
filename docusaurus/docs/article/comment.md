---
sidebar_position: 0
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

# Comment

## Definition
<Tabs>
  <TabItem value="json" label="JSON" default>

```json title="comment.json"
{
  "id": 0,
  "articleID": 0,
  "commentName": "",
  "status": 0,
  "createdAt": "",
  "updatedAt": "",
}
```

  </TabItem>
  <TabItem value="typescript" label="Typescript">

```tsx title="comment.ts"
export interface IComment {
    id: number;
    articleID: number;
    commentName: string;
    status: number;
    createdAt: Date | null;
    updatedAt: Date | null;
}
```

  </TabItem>
  <TabItem value="kotlin" label="Kotlin">
  
  </TabItem>
</Tabs>

## Get Comments

<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X GET "${baseUrl}/comment/?filter=${filter}&orderBy=${orderBy}&page=${page}&pageSize=${pageSize}" \
   -H "Authorization: Bearer ${accessToken}" \
   -H "Content-Type: application/json" \
   -H "Accept: application/json"
```
  </TabItem>
  <TabItem value="typescript" label="Typescript">

  </TabItem>
  <TabItem value="kotlin" label="Kotlin">

  </TabItem>
</Tabs>

## Get Comment
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X GET "${baseUrl}/comment/${id}" \
   -H "Authorization: Bearer ${accessToken}" \
   -H "Content-Type: application/json" \
   -H "Accept: application/json"
```
  </TabItem>
  <TabItem value="typescript" label="Typescript">

  </TabItem>
  <TabItem value="kotlin" label="Kotlin">
  
  </TabItem>
</Tabs>

## Apply Comment.ID
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X POST "${baseUrl}/apply/comment/id/" \
   -H "Authorization: Bearer ${accessToken}" \
   -H "Content-Type: application/json" \
   -H "Accept: application/json"
```
  </TabItem>
  <TabItem value="typescript" label="Typescript">

  </TabItem>
  <TabItem value="kotlin" label="Kotlin">
  
  </TabItem>
</Tabs>

## Create Comment
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X POST "${baseUrl}/comment/" \
   -H "Authorization: Bearer ${accessToken}" \
   -H "Content-Type: application/json" \
   -H "Accept: application/json" \
   -d @comment.json
```
:::info
```json title="comment.json"
{
  "articleID": 0,
  "commentName": "",
  "status": 0,
}
```
:::

:::tip
success status: 201
:::
  </TabItem>
  <TabItem value="typescript" label="Typescript">

  </TabItem>
  <TabItem value="kotlin" label="Kotlin">
  
  </TabItem>
</Tabs>

## Update Comment
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
#!/bin/bash
curl -X PUT "${baseUrl}/comment/${id}" \
   -H "Authorization: Bearer ${accessToken}" \
   -H "Content-Type: application/json" \
   -H "Accept: application/json" \
   -d @comment.json
```
:::info
```json title="comment.json"
{
  "articleID": 0,
  "commentName": "",
  "status": 0,
}
```
:::
  </TabItem>
  <TabItem value="typescript" label="Typescript">

  </TabItem>
  <TabItem value="kotlin" label="Kotlin">
  
  </TabItem>
</Tabs>

## Patch Comment
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X PATCH "${baseUrl}/comment/${id}" \
   -H "Authorization: Bearer ${accessToken}" \
   -H "Content-Type: application/json" \
   -H "Accept: application/json" \
   -H "Attrs: articleID,commentName,status" \
   -d @comment.json
```
:::info
```json title="comment.json"
{
  "articleID": 0,
  "commentName": "",
  "status": 0,
}
```
:::
  </TabItem>
  <TabItem value="typescript" label="Typescript">

  </TabItem>
  <TabItem value="kotlin" label="Kotlin">
  
  </TabItem>
</Tabs>

## Update Comment Status
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X PATCH "${baseUrl}/comment/${id}/status/" \
   -H "Authorization: Bearer ${accessToken}" \
   -H "Content-Type: application/json" \
   -H "Accept: application/json" \
   -d @comment.json
```
:::info
```json title="comment.json"
{
  "status": 0
}
```
:::
  </TabItem>
  <TabItem value="typescript" label="Typescript">

  </TabItem>
  <TabItem value="kotlin" label="Kotlin">
  
  </TabItem>
</Tabs>

## Destroy Comment
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X DELETE "${baseUrl}/comment/${id}" \
   -H "Authorization: Bearer ${accessToken}" \
   -H "Content-Type: application/json" \
   -H "Accept: application/json"
```
  </TabItem>
  <TabItem value="typescript" label="Typescript">

  </TabItem>
  <TabItem value="kotlin" label="Kotlin">
  
  </TabItem>
</Tabs>
