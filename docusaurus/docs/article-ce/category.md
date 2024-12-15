---
sidebar_position: 1
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

# Category

## Definition
<Tabs>
  <TabItem value="json" label="JSON" default>

```json title="category.json"
{
  "id": 0,
  "ref": 0,
  "categoryName": "",
  "status": 0,
  "createdAt": "",
  "updatedAt": "",
}
```

  </TabItem>
  <TabItem value="typescript" label="Typescript">

```tsx title="category.ts"
export interface ICategory {
    id: number;
    ref: number;
    categoryName: string;
    status: number;
    createdAt: Date | null;
    updatedAt: Date | null;
}
```

  </TabItem>
  <TabItem value="kotlin" label="Kotlin">
  
  </TabItem>
</Tabs>

## Get Categories

<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X GET "${baseUrl}/category/?filter=${filter}&orderBy=${orderBy}&page=${page}&pageSize=${pageSize}" \
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

## Get Category
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X GET "${baseUrl}/category/${id}" \
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

## Apply Category.ID
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X POST "${baseUrl}/apply/category/id/" \
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

## Create Category
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X POST "${baseUrl}/category/" \
   -H "Authorization: Bearer ${accessToken}" \
   -H "Content-Type: application/json" \
   -H "Accept: application/json" \
   -d @category.json
```
:::info
```json title="category.json"
{
  "ref": 0,
  "categoryName": "",
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

## Update Category
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
#!/bin/bash
curl -X PUT "${baseUrl}/category/${id}" \
   -H "Authorization: Bearer ${accessToken}" \
   -H "Content-Type: application/json" \
   -H "Accept: application/json" \
   -d @category.json
```
:::info
```json title="category.json"
{
  "ref": 0,
  "categoryName": "",
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

## Patch Category
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X PATCH "${baseUrl}/category/${id}" \
   -H "Authorization: Bearer ${accessToken}" \
   -H "Content-Type: application/json" \
   -H "Accept: application/json" \
   -H "Attrs: ref,categoryName,status" \
   -d @category.json
```
:::info
```json title="category.json"
{
  "ref": 0,
  "categoryName": "",
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

## Update Category Status
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X PATCH "${baseUrl}/category/${id}/status/" \
   -H "Authorization: Bearer ${accessToken}" \
   -H "Content-Type: application/json" \
   -H "Accept: application/json" \
   -d @category.json
```
:::info
```json title="category.json"
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

## Destroy Category
<Tabs>
  <TabItem value="curl" label="Curl" default>

```bash
curl -X DELETE "${baseUrl}/category/${id}" \
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
