# 🧾 Transaction Event Processor

> **Difficulty:** Easy–Medium  
> **Target Time:** 30 minutes  
> **Language:** Go  
> **Topics:** Deduplication · Validation · Aggregation · Sorting

---

## 📌 1. Scenario

You receive a slice of **transaction events** from a payment service. Due to network retries, some transactions are **duplicated**.

Your task is to process these events and produce a **clean, validated, aggregated, and sorted summary per user**.

---

## 🎯 2. Requirements

### ✅ Deduplication

Remove duplicate events by `ID`, keeping only the **first occurrence**.

### ✅ Validation

An event is considered **valid** only if:

- `ID` is not empty
- `UserID` is not empty
- `Amount` is not negative (`Amount >= 0`)

### ✅ Aggregation

For each unique user, compute:

- **`TotalAmount`** → sum of all unique valid event amounts
- **`EventCount`** → number of unique valid events

### ✅ Sorting

Sort the final summary list using **two conditions**:

| Priority    | Field         | Order                      |
| ----------- | ------------- | -------------------------- |
| Primary     | `TotalAmount` | Descending (highest first) |
| Tie-breaker | `UserID`      | Ascending (A → Z)          |

---

## 🛠️ 3. Solution Architecture & Strategy

The solution combines several idiomatic Go + DSA patterns:

### 🔹 A. The "Set" Pattern for Deduplication

**Data Structure:** `map[string]struct{}`

Go has no native `Set` type, so we emulate one using a map with **empty struct values** (`struct{}{}`).

> 💡 **Why `struct{}{}`?**  
> Empty structs occupy **zero bytes** of memory. This makes them the most memory-efficient way to track "seen" items in Go — a well-known idiomatic trick.

````go
seen := make(map[string]struct{})

# B. Pointer-Based Aggregation

**Data Structure:** `map[string]*UserSummary`

To group data by `UserID` and accumulate running totals, we store **pointers** to `UserSummary` structs in a map.

> 💡 **Why pointers?**
> Mutating `TotalAmount` and `EventCount` **in place** avoids re-inserting structures into the map on every update, keeping the code concise and efficient.

```go
type UserSummary struct {
    UserID      string
    TotalAmount float64
    EventCount  int
}

C. Custom Multi-Condition Sorting
Function: sort.Slice

After aggregation, we extract the map values into a flat slice and sort it using a custom comparator that handles:

Descending monetary totals

Ascending UserID tie-breaking

go
sort.Slice(summaries, func(i, j int) bool {
    if summaries[i].TotalAmount != summaries[j].TotalAmount {
        return summaries[i].TotalAmount > summaries[j].TotalAmount
    }
    return summaries[i].UserID < summaries[j].UserID
})
📊 4. Complexity Analysis
⏱️ Time Complexity: O(N log N)
Step	Cost
Iterate over events (map lookups/insertions)	O(N)
Extract map items into result slice	O(K)
Sort the user summary slice	O(K log K)
Where N = number of events, K = number of unique users.
The dominant term is the sorting step → O(N log N).

💾 Space Complexity: O(N)
Structure	Usage
seen map	Up to N unique event IDs
userMap	Up to K unique users
Overall auxiliary space required → O(N).

🔄 5. High-Level Flow
text
┌────────────────────┐
│  Input: Events[]   │
└─────────┬──────────┘
          ▼
┌────────────────────┐
│   Deduplicate by   │
│        ID          │
└─────────┬──────────┘
          ▼
┌────────────────────┐
│     Validate       │
│ (ID, UserID, Amt)  │
└─────────┬──────────┘
          ▼
┌────────────────────┐
│     Aggregate      │
│  per UserID        │
└─────────┬──────────┘
          ▼
┌────────────────────┐
│      Sort          │
│ Total ↓ / UserID ↑ │
└─────────┬──────────┘
          ▼
┌────────────────────┐
│ Output: Summaries  │
└────────────────────┘
🧠 6. Key Takeaways
✅ Idiomatic Go Set → map[string]struct{} for zero-memory-cost tracking.

✅ Pointer maps → mutate aggregates in place without re-insertion overhead.

✅ Custom comparators → sort.Slice handles multi-key sorting cleanly.

✅ Single-pass dedup + validate + aggregate → keeps the solution efficient and readable.

🏁 7. Summary
A clean, production-style Go implementation that:

🧹 Removes duplicate retry events

🔍 Filters out invalid records

📦 Aggregates totals and counts per user

📉 Sorts by spending (descending) with deterministic tie-breaking

Built with three simple patterns and running in O(N log N) time and O(N) space.
````
