---
sidebar_position: 0
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

# Article

## Definition
<Tabs>
  <TabItem value="json" label="JSON" default>

```json title="article.json"
{
  "id": 0,
  "categoryID": 0,
  "articleName": "",
  "status": 0,
  "createdAt": "",
  "updatedAt": "",
}
```

  </TabItem>
  <TabItem value="typescript" label="Typescript">

```tsx title="article.ts"
export interface IArticle {
    id: number;
    categoryID: number;
    articleName: string;
    status: number;
    createdAt: Date | null;
    updatedAt: Date | null;
}
```

  </TabItem>
  <TabItem value="kotlin" label="Kotlin">
  
  </TabItem>
</Tabs>

## Get Articles

<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X GET "${baseUrl}/article/?filter=${filter}&orderBy=${orderBy}&page=${page}&pageSize=${pageSize}" \
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

## Get Article
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X GET "${baseUrl}/article/${id}" \
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

## Apply Article.ID
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X POST "${baseUrl}/apply/article/id/" \
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

## Create Article
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X POST "${baseUrl}/article/" \
   -H "Authorization: Bearer ${accessToken}" \
   -H "Content-Type: application/json" \
   -H "Accept: application/json" \
   -d @article.json
```
:::info
```json title="article.json"
{
  "categoryID": 0,
  "articleName": "",
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

## Update Article
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
#!/bin/bash
curl -X PUT "${baseUrl}/article/${id}" \
   -H "Authorization: Bearer ${accessToken}" \
   -H "Content-Type: application/json" \
   -H "Accept: application/json" \
   -d @article.json
```
:::info
```json title="article.json"
{
  "categoryID": 0,
  "articleName": "",
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

## Patch Article
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X PATCH "${baseUrl}/article/${id}" \
   -H "Authorization: Bearer ${accessToken}" \
   -H "Content-Type: application/json" \
   -H "Accept: application/json" \
   -H "Attrs: categoryID,articleName,status" \
   -d @article.json
```
:::info
```json title="article.json"
{
  "categoryID": 0,
  "articleName": "",
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

## Update Article Status
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X PATCH "${baseUrl}/article/${id}/status/" \
   -H "Authorization: Bearer ${accessToken}" \
   -H "Content-Type: application/json" \
   -H "Accept: application/json" \
   -d @article.json
```
:::info
```json title="article.json"
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

## Destroy Article
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X DELETE "${baseUrl}/article/${id}" \
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

## Get Tags
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X GET "${baseUrl}/article/${id}/tag/?filter=${filter}&orderBy=${orderBy}&page=${page}&pageSize=${pageSize}" \
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

## Link Tags
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X POST "${baseUrl}/article/${id}/tag/" \
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

## UnLink Tags
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X DELETE "${baseUrl}/article/${id}/tag/" \
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
