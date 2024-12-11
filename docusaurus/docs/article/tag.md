---
sidebar_position: 0
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

# Tag

## Definition
<Tabs>
  <TabItem value="json" label="JSON" default>

```json title="tag.json"
{
  "id": 0,
  "tagName": "",
  "status": 0,
  "createdAt": "",
  "updatedAt": "",
}
```

  </TabItem>
  <TabItem value="typescript" label="Typescript">

```tsx title="tag.ts"
export interface ITag {
    id: number;
    tagName: string;
    status: number;
    createdAt: Date | null;
    updatedAt: Date | null;
}
```

  </TabItem>
  <TabItem value="kotlin" label="Kotlin">
  
  </TabItem>
</Tabs>

## Get Tags

<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X GET "${baseUrl}/tag/?filter=${filter}&orderBy=${orderBy}&page=${page}&pageSize=${pageSize}" \
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

## Get Tag
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X GET "${baseUrl}/tag/${id}" \
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

## Apply Tag.ID
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X POST "${baseUrl}/apply/tag/id/" \
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

## Create Tag
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X POST "${baseUrl}/tag/" \
   -H "Authorization: Bearer ${accessToken}" \
   -H "Content-Type: application/json" \
   -H "Accept: application/json" \
   -d @tag.json
```
:::info
```json title="tag.json"
{
  "tagName": "",
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

## Update Tag
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
#!/bin/bash
curl -X PUT "${baseUrl}/tag/${id}" \
   -H "Authorization: Bearer ${accessToken}" \
   -H "Content-Type: application/json" \
   -H "Accept: application/json" \
   -d @tag.json
```
:::info
```json title="tag.json"
{
  "tagName": "",
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

## Patch Tag
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X PATCH "${baseUrl}/tag/${id}" \
   -H "Authorization: Bearer ${accessToken}" \
   -H "Content-Type: application/json" \
   -H "Accept: application/json" \
   -H "Attrs: tagName,status" \
   -d @tag.json
```
:::info
```json title="tag.json"
{
  "tagName": "",
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

## Update Tag Status
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X PATCH "${baseUrl}/tag/${id}/status/" \
   -H "Authorization: Bearer ${accessToken}" \
   -H "Content-Type: application/json" \
   -H "Accept: application/json" \
   -d @tag.json
```
:::info
```json title="tag.json"
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

## Destroy Tag
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X DELETE "${baseUrl}/tag/${id}" \
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

## Get Articles
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X GET "${baseUrl}/tag/${id}/article/?filter=${filter}&orderBy=${orderBy}&page=${page}&pageSize=${pageSize}" \
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

## Link Articles
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X POST "${baseUrl}/tag/${id}/article/" \
   -H "Authorization: Bearer ${accessToken}" \
   -H "Content-Type: application/json" \
   -H "Accept: application/json" \
   -d "[1,2,3,4]"
```
  </TabItem>
  <TabItem value="typescript" label="Typescript">

  </TabItem>
  <TabItem value="kotlin" label="Kotlin">
	
  </TabItem>
</Tabs>

## UnLink Articles
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X DELETE "${baseUrl}/tag/${id}/article/" \
   -H "Authorization: Bearer ${accessToken}" \
   -H "Content-Type: application/json" \
   -H "Accept: application/json" \
   -d "[1,2,3,4]"
```
  </TabItem>
  <TabItem value="typescript" label="Typescript">

  </TabItem>
  <TabItem value="kotlin" label="Kotlin">
	
  </TabItem>
</Tabs>
