# Debug Logging Guide

This document describes all the debug logging that has been added to the SkillDar Client application to help trace performance issues and verify the fixes applied.

## Overview

Comprehensive debug logging has been added to trace three main flows:

1. **Category Button Interactions** - When user taps a category button
2. **Worker Data Loading** - When workers are loaded from API
3. **Map Loading Lifecycle** - When map loads, gets cancelled, or completes

All debug logs include timestamps in format `HH:MM:SS.mmm` for precise timing analysis.

---

## 1. Category Button Logging (categories.go)

### Location

File: `pkg/ui/categories.go` - Function: `createCategoryButtonWithFilter()`

### Logs Produced

When a user taps a category button, two debug messages are logged:

```
[DEBUG] [15:04:05.123] Category button tapped: Plumbing (API: plumber)
[DEBUG] Triggering worker reload with new category...
```

### What It Tells You

- **Exact time** the user tapped the category button
- **Display name** shown to user (e.g., "Plumbing")
- **API name** sent to backend (e.g., "plumber")
- Confirmation that worker reload was triggered

### Example Output Sequence

```
[DEBUG] [15:04:10.450] Category button tapped: Plumbing (API: plumber)
[DEBUG] Triggering worker reload with new category...
[DEBUG] [15:04:10.451] Map loading cancelled due to category filter selection
[DEBUG] [15:04:10.452] Starting worker load - Page: 1, Limit: 20, Category: plumber
```

---

## 2. Worker Loading Logging (workers_list.go)

### Location

File: `pkg/ui/workers_list.go` - Function: `loadWorkers()`

### Logs Produced

#### 2.1 Load Start

```
[DEBUG] [15:04:10.452] Starting worker load - Page: 1, Limit: 20, Category: plumber
```

#### 2.2 Load Completion

```
[DEBUG] [15:04:11.234] Worker load completed: 8 workers received, total in category: 247
```

#### 2.3 Rendering Start

```
[DEBUG] [15:04:11.235] Rendering 5 visible workers immediately, deferring 3 workers
```

#### 2.4 Lazy Rendering Completion (One entry per deferred worker)

```
[DEBUG] [15:04:11.255] Lazy rendered worker #6
[DEBUG] [15:04:11.275] Lazy rendered worker #7
[DEBUG] [15:04:11.295] Lazy rendered worker #8
```

### What It Tells You

- **API call start time** and parameters (page, limit, category filter)
- **API response** received with worker count and total available
- **Rendering strategy** (5 immediately, rest deferred)
- **Lazy rendering completion** with 20ms stagger between each worker
- **Total load-to-display time** from start log to last lazy render log

### Example Full Flow

```
[DEBUG] [15:04:10.452] Starting worker load - Page: 1, Limit: 20, Category: plumber
[DEBUG] [15:04:11.234] Worker load completed: 8 workers received, total in category: 247
[DEBUG] [15:04:11.235] Rendering 5 visible workers immediately, deferring 3 workers
[DEBUG] [15:04:11.255] Lazy rendered worker #6
[DEBUG] [15:04:11.275] Lazy rendered worker #7
[DEBUG] [15:04:11.295] Lazy rendered worker #8
```

**Total time: ~843ms** (from 15:04:10.452 to 15:04:11.295)

---

## 3. Map Loading Logging (map_widget.go)

### Location

File: `pkg/ui/map_widget.go` - Function: `createMapSection()`

### Logs Produced

#### 3.1 Map Load Scheduled (on screen load)

```
[DEBUG] [15:04:05.100] Map load scheduled: waiting 5 seconds before starting...
```

#### 3.2 Delay Completion Check (5 seconds later)

```
[DEBUG] [15:04:10.101] Map delay completed, checking if should load...
```

#### 3.3a Map Load Starts (if not cancelled)

```
[DEBUG] [15:04:10.101] Starting map load (not cancelled, not already loaded)
[DEBUG] [15:04:10.102] Starting map widget creation...
[DEBUG] [15:04:10.103] Creating map widget asynchronously (10s timeout)...
[DEBUG] [15:04:10.104] Map widget created, configuring zoom and location...
[DEBUG] [15:04:10.105] Map configured - Zoom: 13, Location: Tunis (36.8065, 10.1815)
[DEBUG] [15:04:11.456] Map successfully loaded and displayed
```

#### 3.3b Map Load Timeout (if creation takes >10 seconds)

```
[DEBUG] [15:04:10.101] Starting map load (not cancelled, not already loaded)
[DEBUG] [15:04:10.102] Starting map widget creation...
[DEBUG] [15:04:10.103] Creating map widget asynchronously (10s timeout)...
[DEBUG] [15:04:20.105] Map widget creation timed out
```

#### 3.3c Map Load Cancelled (if category filter active)

```
[DEBUG] [15:04:10.101] Map delay completed, checking if should load...
[DEBUG] [15:04:10.101] Map load skipped - was cancelled during delay
```

#### 3.3c Map Already Loaded

```
[DEBUG] [15:04:10.101] Map load skipped - already loaded
```

#### 3.4 Cancellation Event (when user taps category filter)

```
[DEBUG] [15:04:10.451] Map loading cancelled due to category filter selection
```

### What It Tells You

- **Exact time** map load begins (deferred 5 seconds)
- **5-second delay** progress (logs at start and completion)
- **Map creation steps** (widget creation, configuration, display)
- **Load status** (whether map was cancelled, already loaded, or completed)
- **Total map load time** from start to completion

### Example Scenarios

#### Scenario 1: Map loads normally (no category filter)

```
[DEBUG] [15:04:05.100] Map load scheduled: waiting 5 seconds before starting...
[DEBUG] [15:04:10.101] Map delay completed, checking if should load...
[DEBUG] [15:04:10.101] Starting map load (not cancelled, not already loaded)
[DEBUG] [15:04:10.102] Starting map widget creation...
[DEBUG] [15:04:10.103] Creating map widget asynchronously...
[DEBUG] [15:04:10.104] Map widget created, configuring zoom and location...
[DEBUG] [15:04:10.105] Map configured - Zoom: 13, Location: Tunis (36.8065, 10.1815)
[DEBUG] [15:04:11.456] Map successfully loaded and displayed
```

**Total time: ~6.3 seconds** (5s delay + ~1.3s creation)

#### Scenario 2: User filters category before map loads

```
[DEBUG] [15:04:05.100] Map load scheduled: waiting 5 seconds before starting...
[DEBUG] [15:04:10.450] Category button tapped: Plumbing (API: plumber)
[DEBUG] [15:04:10.451] Map loading cancelled due to category filter selection
[DEBUG] [15:04:10.452] Starting worker load - Page: 1, Limit: 20, Category: plumber
[DEBUG] [15:04:10.101] Map delay completed, checking if should load...
[DEBUG] [15:04:10.101] Map load skipped - was cancelled during delay
```

**Map never loads** - workers load immediately instead

---

## 4. Complete User Flow Example

Here's what a typical user session looks like in the logs:

```
=== App Started ===
[DEBUG] [15:04:05.100] Map load scheduled: waiting 5 seconds before starting...

=== Workers Load Immediately ===
[DEBUG] [15:04:05.150] Starting worker load - Page: 1, Limit: 20, Category:
[DEBUG] [15:04:06.234] Worker load completed: 20 workers received, total in category: 247
[DEBUG] [15:04:06.235] Rendering 5 visible workers immediately, deferring 15 workers
[DEBUG] [15:04:06.255] Lazy rendered worker #6
[DEBUG] [15:04:06.275] Lazy rendered worker #7
... (15 more lazy render logs with 20ms stagger)
[DEBUG] [15:04:06.555] Lazy rendered worker #20

=== Map Delay Progresses (users already see workers) ===
[DEBUG] [15:04:10.101] Map delay completed, checking if should load...
[DEBUG] [15:04:10.101] Starting map load (not cancelled, not already loaded)
[DEBUG] [15:04:10.102] Starting map widget creation...
[DEBUG] [15:04:10.103] Creating map widget asynchronously...
[DEBUG] [15:04:10.104] Map widget created, configuring zoom and location...
[DEBUG] [15:04:10.105] Map configured - Zoom: 13, Location: Tunis (36.8065, 10.1815)
[DEBUG] [15:04:11.456] Map successfully loaded and displayed

=== User Taps Category Filter ===
[DEBUG] [15:04:15.450] Category button tapped: Electricity (API: electrician)
[DEBUG] [15:04:15.451] Map loading cancelled due to category filter selection
[DEBUG] [15:04:15.452] Starting worker load - Page: 1, Limit: 20, Category: electrician
[DEBUG] [15:04:16.234] Worker load completed: 12 workers received, total in category: 89
[DEBUG] [15:04:16.235] Rendering 5 visible workers immediately, deferring 7 workers
... (lazy render logs)

=== No Map Loading During Filter (was cancelled) ===
```

---

## 5. Performance Metrics to Analyze

### Key Timings

1. **Worker Load Time**

   - From "Starting worker load" to "Worker load completed"
   - Expected: 800ms - 2000ms depending on network
   - Example: `15:04:10.452` to `15:04:11.234` = 782ms ✅

2. **Worker Rendering Time**

   - From "Rendering X visible workers" to last "Lazy rendered worker"
   - Expected: 200ms - 300ms for 20 workers
   - Immediate workers render instantly, deferred workers stagger at 20ms

3. **Map Load Time**

   - From "Starting map widget creation" to "Map successfully loaded"
   - Expected: 1000ms - 2000ms depending on system
   - Example: `15:04:10.102` to `15:04:11.456` = 1354ms

4. **Total Initial Load Time**
   - From first worker load start to map completion
   - Expected: ~6-7 seconds total
   - 5 seconds map delay + ~1 second worker load + ~1 second map creation

### Optimization Insights

- **Workers visible time** = Worker load time (~1s) - This is what users see
- **Map doesn't block** = 5 second delay allows categories to be responsive
- **Lazy rendering** = Users see 5 workers instantly, rest fill in over 300ms
- **Category filter response** = Map cancels immediately, workers reload

---

## 6. Troubleshooting Guide

### Problem: Category filter seems to hang

Look for:

- No "Starting worker load" log after "Category button tapped"
- No "Map loading cancelled" log
  **Solution**: Check if API is responding, verify token hasn't expired

### Problem: Workers load slowly

Look for:

- Long gap between "Starting worker load" and "Worker load completed"
- Log the exact timestamps and calculate duration
  **Solution**: Network issue, server slow, or API pagination issue

### Problem: Map keeps loading and interfering

Look for:

- "Map loading cancelled" doesn't appear
- Map creates after category filter completes
- **New**: "Map widget creation timed out" (means map took >10 seconds to load)

**Solution**: Check if cancelMap flag is working, may need to increase delay. **Timeout protection**: Map widget creation has a 10-second timeout. If creation takes too long, it will abort and log "Map widget creation timed out" to prevent UI freezing. This keeps the app responsive even if map loading stalls.

### Problem: Lazy rendering not working

Look for:

- No "Lazy rendered worker" logs after visible workers
- All workers logged at same time
  **Solution**: Check if stagger duration (20ms) is working, verify time.Sleep in code

---

## 7. Log Output Format

All debug logs follow this format:

```
[DEBUG] [HH:MM:SS.mmm] Message text...
```

- `[DEBUG]` - Prefix for all debug messages
- `[HH:MM:SS.mmm]` - Timestamp in 24-hour format with milliseconds
- Rest of message is action-specific

This format makes it easy to:

- Filter logs: `grep "\[DEBUG\]"`
- Correlate with other system logs
- Measure timing between events
- Track exact sequence of operations

---

## 8. Running with Debug Logs

The app automatically outputs debug logs to console when running:

### Android

```bash
make android
adb logcat | grep DEBUG
```

### Desktop

```bash
./skillDarClient 2>&1 | grep DEBUG
```

### View all logs

Remove grep filter to see everything:

```bash
adb logcat
./skillDarClient
```

---

## 9. Summary of Changes

| File              | Function                           | Changes                                                                 |
| ----------------- | ---------------------------------- | ----------------------------------------------------------------------- |
| `categories.go`   | `createCategoryButtonWithFilter()` | Added `time` import, button tap logging with timestamp                  |
| `workers_list.go` | `loadWorkers()`                    | Added load start/completion logging with timestamp, lazy render logging |
| `map_widget.go`   | `createMapSection()`               | Added 5s delay logging, load progress logging, cancellation logging     |

All changes are additive (debug logging only) - no logic changes or behavior modifications.

---

## 10. Next Steps

After deploying with these debug logs:

1. **Run the app** on Android device
2. **Use typical user flow** (tap categories, scroll, wait for map)
3. **Capture logs** using adb logcat
4. **Analyze timings** to verify fixes are working
5. **Fine-tune sleep times** if needed based on real network speeds

Expected behavior with logs:

- Workers load in ~1 second
- Map cancels immediately when category tapped
- Category filter responsive (no lag)
- Map loads after 5 seconds if no filter
- Total initial load ~6-7 seconds
