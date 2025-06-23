# vikunja-mcp User Guide: AI-Powered Task Management for ADHD

## Table of Contents
1. [Quick Start](#quick-start)
2. [Understanding the System](#understanding-the-system)
3. [Setup & Configuration](#setup--configuration)
4. [Core Features & Tools](#core-features--tools)
5. [Daily Workflow Examples](#daily-workflow-examples)
6. [Advanced Usage Patterns](#advanced-usage-patterns)
7. [Troubleshooting](#troubleshooting)
8. [Tips for ADHD Users](#tips-for-adhd-users)

---

## Quick Start

**vikunja-mcp** transforms your existing Vikunja task manager into an AI-powered productivity system. It learns your energy patterns, suggests optimal tasks for your current state, and helps you maintain focus through intelligent recommendations.

### What Makes This Different?
- **Energy-Aware**: Matches tasks to your current mental capacity
- **Context-Smart**: Reduces decision paralysis with AI recommendations  
- **ADHD-Optimized**: Built specifically for neurodivergent productivity patterns
- **Non-Disruptive**: Works with your existing Vikunja setup without migration

### Prerequisites
- Active Vikunja instance (self-hosted or cloud)
- OpenAI API key for intelligent recommendations
- Claude AI access (via this interface)

---

## Understanding the System

### Architecture Overview
```
You ↔ Claude AI ↔ vikunja-mcp ↔ OpenAI GPT-4 ↔ Your Vikunja Tasks
```

**What happens when you ask for recommendations:**
1. You describe your current state ("I have 2 hours, medium energy, need to do deep work")
2. Claude passes this to vikunja-mcp
3. vikunja-mcp fetches your tasks from Vikunja
4. OpenAI analyzes tasks and suggests the best matches
5. You get a curated list optimized for your current capacity

### Key Concepts

**Energy Levels:**
- **Low**: Simple, routine tasks that don't require deep thinking
- **Medium**: Moderate complexity work, administrative tasks
- **High**: Complex problem-solving, creative work, challenging projects
- **Social**: Communication, meetings, collaborative work

**Work Modes:**
- **Deep**: Sustained focus on complex, important work
- **Quick**: Short bursts of productive action (5-25 minutes)
- **Admin**: Organizational tasks, email, planning, maintenance

**Hyperfocus Compatibility (1-5 scale):**
- **1-2**: Light tasks, good for warm-up or cool-down
- **3**: Balanced tasks suitable for most sessions
- **4-5**: Prime hyperfocus material, save for peak energy

---

## Setup & Configuration

### Option 1: Docker Deployment (Recommended)

```bash
# Pull and run the container
docker run -d \
  --name vikunja-mcp \
  -e VIKUNJA_URL="https://your-vikunja-instance.com" \
  -e VIKUNJA_TOKEN="your-vikunja-api-token" \
  -e OPENAI_API_KEY="your-openai-api-key" \
  vikunja-mcp:latest

# Check if container is running
docker ps | grep vikunja-mcp
docker logs vikunja-mcp
```

### Option 2: Local Development

```bash
# Clone and build
git clone https://github.com/BelKirill/vikunja-mcp.git
cd vikunja-mcp
make build

# Set environment variables
export VIKUNJA_URL="https://your-vikunja-instance.com"
export VIKUNJA_TOKEN="your-vikunja-api-token" 
export OPENAI_API_KEY="your-openai-api-key"

# Run the MCP server
./bin/mcp
```

### Getting Your Vikunja API Token

1. Log into your Vikunja instance
2. Go to Settings → API Tokens
3. Create a new token with full access
4. Copy the token for your environment variables

### Connecting to Claude

Once your vikunja-mcp server is running, it automatically becomes available as MCP tools in this Claude interface. The MCP tools will appear and you can start using commands like:

```
I have 3 hours, high energy, want to do deep work. What should I focus on?
```

**Note**: The MCP server runs as a background process that Claude connects to via the MCP protocol, not as a web service.

---

## Core Features & Tools

### 1. `daily-focus` - Your AI Productivity Assistant

**Purpose**: Get personalized task recommendations based on your current state

**Usage:**
```
I have 3 hours, high energy, want to do deep work. What should I focus on?
```

**Parameters:**
- **energy**: `low|medium|high|social` - Your current mental capacity
- **mode**: `deep|quick|admin` - Type of work you want to do
- **hours**: Available time (0.1 to 8 hours)
- **max_items**: Number of tasks to recommend (default: 5)
- **instructions**: Special preferences or constraints

**Example Output:**
```
🎯 AI Recommendations for Deep Work (High Energy, 3 hours)

1. **Refactor authentication system** (Hyperfocus: 5/5)
   - Perfect for high-energy deep work
   - Estimated: 2.5 hours
   - Why: Complex problem-solving matches your peak capacity

2. **Write technical architecture doc** (Hyperfocus: 4/5)
   - High-value documentation work
   - Estimated: 2 hours
   - Why: Creative writing task suited for focused sessions
```

### 2. `upsert-task` - Intelligent Task Creation

**Purpose**: Create or update tasks with rich metadata for better AI recommendations

**Usage:**
```
Create a task: "Review Q3 budget analysis" - this is admin work, medium priority, should take about 1 hour
```

**Parameters:**
- **title**: Task name
- **description**: Detailed description (metadata will be embedded here)
- **priority**: 1-5 (1=highest, 5=lowest)
- **project_id**: Vikunja project ID
- **metadata**: Energy level, work mode, duration, hyperfocus rating

**The AI automatically adds metadata like:**
```json
{
  "energy": "medium",
  "mode": "admin", 
  "minutes": 60,
  "hyper_focus_comp": 3
}
```

### 3. `get-filtered-tasks` - Smart Task Discovery

**Purpose**: Find tasks using natural language or specific filters

**Natural Language Examples:**
```
Show me quick tasks I can do in 15 minutes
Find high-priority tasks for deep work
What admin tasks are overdue?
Show me creative projects I've been avoiding
List tasks with priority 1-2 in my current project
Find all tasks tagged "urgent" that aren't done
Show me tasks created this week that need review
```

**Manual Filter Examples (Working Vikunja Syntax):**
```
# High-priority incomplete tasks (priorities 1-2, with 1 being highest)
priority <= 2 && done = false

# All incomplete tasks (basic productivity filter)
done = false

# Recent tasks (this week)
created >= "2025-06-17" && done = false

# Time-based filtering (specific date ranges)
updated >= "2025-06-01" && updated <= "2025-06-23" && done = false

# Urgent incomplete tasks by priority
priority = 1 && done = false

# Medium to high priority tasks
priority <= 3 && done = false

# Recently created high-priority tasks
created >= "2025-06-20" && priority <= 2

# Complex filtering (high priority OR recent updates)
(priority <= 2 || updated >= "2025-06-20") && done = false
```

**Filter Tips:**
- **Priority**: Lower numbers = higher priority (1 = highest, 5 = lowest)
- **Dates**: Use YYYY-MM-DD format in quotes
- **Status**: `done = false` for incomplete, `done = true` for completed
- **Operators**: `>=` (greater/equal), `<=` (less/equal), `=` (equals), `!=` (not equals)
- **Logic**: `&&` (and), `||` (or), parentheses for grouping

### 4. `get-full-task` - Detailed Task Analysis

**Purpose**: Get complete information about a specific task including metadata

**Usage:**
```
Show me details for task ID 1234
```

**Returns:**
- Full task description
- Embedded metadata (energy, mode, duration)
- Creation/update timestamps
- Project context
- Priority and status

### 5. `add-comment` - Task Progress Tracking

**Purpose**: Add timestamped comments to track progress or insights

**Usage:**
```
Add comment to task 1234: "Completed first draft, needs peer review before finalizing"
```

---

## Daily Workflow Examples

### Morning Planning Session

```
🌅 MORNING ROUTINE

You: "I have 4 hours today, feeling medium energy, mix of deep work and admin. What should I tackle?"

Claude: [Uses daily-focus to analyze your tasks]

Response: "Here's your optimized day:
• 9-11am: Deep work on customer dashboard (High focus needed)
• 11-11:30am: Quick admin - expense reports (Low energy OK) 
• 11:30am-12:30pm: Review and respond to client emails (Medium energy)
• 1:30-2:30pm: Admin time - project planning for next week"
```

### Mid-Day Energy Check

```
⚡ ENERGY SHIFT

You: "My energy crashed after lunch. Show me low-energy tasks for the next hour."

Claude: [Filters for low-energy, quick tasks]

Response: "Perfect for a post-lunch dip:
• Organize project files (15 min)
• Update task priorities (10 min)
• Review calendar for tomorrow (5 min)
• Clean up email inbox (20 min)"
```

### End-of-Day Capture

```
🌙 DAILY WRAP-UP

You: "Create tasks for tomorrow: 'Prepare presentation slides' (2 hours, deep work) and 'Call vendor about pricing' (15 min admin)"

Claude: [Creates tasks with appropriate metadata]

Response: "Tasks created with optimal scheduling data:
• Presentation slides → Tagged for high-energy deep work sessions
• Vendor call → Flagged as quick admin task for low-energy periods"
```

### Manual Weekly Review (Current Approach)

```
📊 PATTERN ANALYSIS

You: "Show me all tasks I completed this week"
Claude: [Uses get-filtered-tasks with date range and done=true]

You: "What high-priority tasks did I avoid this week?"
Claude: [Uses get-filtered-tasks with priority filters and done=false]

You: "Show me tasks I created but haven't started"
Claude: [Uses get-filtered-tasks with creation date filters]

// Note: This is manual analysis. A dedicated weekly-review tool 
// would automate this pattern analysis and provide insights.
```

---

## Advanced Usage Patterns

### Energy-Based Batching

**Concept**: Group similar tasks by energy requirement for better flow

```
# High Energy Blocks (90-120 minutes)
- Complex problem solving
- Creative work
- Learning new concepts
- Difficult conversations

# Medium Energy Blocks (25-60 minutes)  
- Writing and documentation
- Code reviews
- Planning and strategy
- Routine development work

# Low Energy Blocks (5-25 minutes)
- Email and communication
- File organization
- Simple data entry
- Meeting preparation
```

### Hyperfocus Session Planning

**When you feel hyperfocus coming on:**

```
You: "I'm feeling hyperfocused and have 3 hours. What's the highest-value work I can do?"

Claude: [Identifies tasks with hyperfocus rating 4-5]

Response: "Hyperfocus session optimized:
1. **Architect new feature** (5/5 rating) - Perfect complex challenge
2. **Write technical deep-dive** (4/5 rating) - Creative writing flow
3. Backup: **Debug performance issue** (4/5 rating) - If you finish early"
```

### Context Switching Minimization

**Stay within project boundaries:**

```
You: "I'm working on the mobile app project. Keep me focused there for the next 2 hours."

Claude: [Filters to mobile app project only]

Response: "Mobile app focus session:
• UI component implementation (Deep work, 90 min)
• Update mobile app documentation (Medium work, 30 min)
• Test on different devices (Quick tasks, 20 min total)"
```

### Procrastination Intervention

**When avoiding important work:**

```
You: "I keep avoiding the quarterly report. It's due Friday and I'm stressed about it."

Claude: [Suggests task breakdown and energy matching]

Response: "Let's break this down:
• Start with outline (15 min, low energy) - Less intimidating entry point
• Draft financial section (45 min, medium energy) - Use existing data
• Write executive summary (30 min, high energy) - Save creative work for peak
• Review and polish (20 min, medium energy) - Final details

Start with the outline right now - just 15 minutes to reduce anxiety."
```

---

## Tips for ADHD Users

### Working WITH Your Brain, Not Against It

**Embrace Natural Rhythms:**
- Track your energy patterns for 1-2 weeks
- Schedule demanding work during your peak hours
- Use low-energy periods for routine tasks
- Don't fight your natural flow

**Reduce Decision Fatigue:**
- Let the AI choose your next task
- Batch similar decisions together
- Pre-plan your week during high-energy periods
- Create "default" routines for common scenarios

**Manage Hyperfocus:**
- Always have backup tasks ready (rating 4-5)
- Set gentle alerts for breaks (but don't force them)
- Keep water and snacks nearby
- Note what triggers hyperfocus for future planning

**Handle Overwhelm:**
- Start with quick wins (5-15 minute tasks)
- Break large projects into smaller pieces
- Use the "minimum viable progress" approach
- Celebrate small accomplishments

### Pomodoro Optimization

**Traditional Pomodoro**: 25 min work, 5 min break
**ADHD-Optimized Pomodoro**: 
- Start with 15-25 minutes (based on task complexity)
- Extend naturally if in flow (don't force breaks)
- Take longer breaks when energy is low
- Use breaks for movement, not more mental work

### Emergency Protocols

**When Everything Feels Urgent:**
```
1. Brain dump all tasks (don't organize yet)
2. Ask AI: "Help me prioritize this chaos"
3. Pick ONE task to do right now
4. Set a 25-minute timer
5. Ignore everything else until timer goes off
```

**When Motivation is Zero:**
```
1. Ask for 5-minute tasks only
2. Do just one thing
3. Use completion dopamine to build momentum
4. Gradually increase task difficulty
5. Celebrate every small win
```

---

## Troubleshooting

### Common Issues

**"AI recommendations don't match my preferences"**
- Update task metadata to better reflect actual energy/time requirements
- Use the `instructions` parameter to specify preferences
- Track which recommendations work and adjust task descriptions

**"Tasks aren't showing up in recommendations"**
- Check that tasks have proper metadata embedded
- Verify task status (completed tasks are filtered out)
- Ensure project_id matches if filtering by project

**"System feels overwhelming instead of helpful"**
- Start with fewer recommendations (max_items: 2-3)
- Focus on one work mode at a time
- Use shorter time blocks (1-2 hours maximum)

**"Energy levels don't seem accurate"**
- Energy is subjective and changes throughout the day
- Start with your best guess and adjust based on results
- Track patterns: what types of work feel energizing vs. draining?

### Technical Issues

**Connection Problems:**
```bash
# Check if MCP server is running
docker ps | grep vikunja-mcp
# or for local development:
ps aux | grep mcp

# Check MCP server logs
docker logs vikunja-mcp

# Verify environment variables are set
docker exec vikunja-mcp env | grep -E "(VIKUNJA|OPENAI)"
```

**Slow AI Responses:**
- OpenAI API calls typically take 2-3 seconds
- Reduce max_items if recommendations are too slow
- Check your OpenAI API rate limits
- MCP protocol adds minimal overhead to AI processing

---

## Future Enhancements

Based on your project roadmap, exciting features are coming:

### **🔄 Planned Tools (In Development)**
- **Weekly Review Tool**: Automated pattern analysis and productivity insights
- **Session Management**: Persistent AI context across work sessions
- **Advanced Bulk Operations**: Intelligent mass task updates
- **Label Management**: Smart categorization and filtering
- **Project Context Awareness**: Minimize context switching between work streams

### **💡 Potential Future Features**
- **Energy Pattern Learning**: Automatic detection of your optimal work windows
- **Procrastination Intervention**: Early detection and mitigation strategies
- **Team Collaboration**: Shared productivity features
- **Advanced Analytics**: Deep insights into work patterns and efficiency

The system is actively evolving based on real-world usage patterns and ADHD productivity research.

## Getting the Most Value

### Week 1: Foundation
- Set up the system and create 10-15 tasks with good metadata
- Use daily-focus every morning for one week
- Track which recommendations actually work for you

### Week 2: Patterns
- Notice your natural energy rhythms
- Identify which task types you consistently avoid
- Start optimizing work modes (deep vs. quick vs. admin)

### Week 3: Optimization
- Refine task metadata based on actual experience
- Experiment with different time blocks and energy combinations
- Begin using the system for weekly planning

### Week 4: Mastery
- Trust the AI recommendations and reduce manual task selection
- Use the system to identify and address procrastination patterns
- Share insights with others or contribute improvements

### Long-term Success
- The system gets better as your task database grows
- Regular weekly reviews help identify productivity patterns
- Consider contributing features that solve your specific challenges

---

## Support & Community

**Project Repository**: https://github.com/BelKirill/vikunja-mcp
**Documentation**: Comprehensive README with technical details
**Issues**: Report bugs or request features on GitHub

**This is an active project built by someone who uses it daily.** Your feedback and real-world usage patterns help improve the system for the entire ADHD community.

---

*Remember: This system is designed to work with your brain, not fix it. The goal is to reduce friction and cognitive load so you can focus on what matters most to you.*