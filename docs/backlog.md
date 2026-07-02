# Commerce AI Platform Backlog

## Overview

This document tracks the implementation backlog for the Commerce AI Platform.

Each backlog item has:

- Unique ID
- Priority
- Estimated effort
- Status
- GitHub Issue
- Dependencies

---

# Status Legend

| Status | Meaning |
|---------|---------|
| ⬜ | Not Started |
| 🚧 | In Progress |
| ✅ | Completed |
| ⏸️ | Blocked |

---

# Milestone 1 — Foundation

| ID | Title | Priority | Estimate | Status | GitHub Issue |
|----|-------|----------|----------|--------|--------------|
| CAP-001 | Initialize Go project and repository structure | High | S | 🚧 | |
| CAP-002 | Configuration management | High | S | ⬜ | |
| CAP-003 | Logging framework | High | S | ⬜ | |
| CAP-004 | HTTP server and routing | High | M | ⬜ | |
| CAP-005 | SQLite schema | High | M | ⬜ | |
| CAP-006 | Database migrations | High | M | ⬜ | |
| CAP-007 | Seed sample data | Medium | S | ⬜ | |
| CAP-008 | Repository layer | High | M | ⬜ | |

---

# Milestone 2 — AI Agent

| ID | Title | Priority | Estimate | Status | GitHub Issue |
|----|-------|----------|----------|--------|--------------|
| CAP-009 | Integrate Ollama | High | M | ⬜ | |
| CAP-010 | Agent runtime | High | L | ⬜ | |
| CAP-011 | Conversation memory | Medium | M | ⬜ | |
| CAP-012 | Prompt management | Medium | M | ⬜ | |
| CAP-013 | Tool registry | High | M | ⬜ | |
| CAP-014 | Tool execution engine | High | M | ⬜ | |

---

# Milestone 3 — Commerce Tools

| ID | Title | Priority | Estimate | Status | GitHub Issue |
|----|-------|----------|----------|--------|--------------|
| CAP-015 | Product search tool | High | M | ⬜ | |
| CAP-016 | Product pricing tool | High | S | ⬜ | |
| CAP-017 | Product availability tool | High | S | ⬜ | |
| CAP-018 | Order creation tool | High | L | ⬜ | |
| CAP-019 | Order retrieval tool | Medium | M | ⬜ | |
| CAP-020 | Return eligibility tool | High | M | ⬜ | |
| CAP-021 | Return processing tool | High | M | ⬜ | |

---

# Milestone 4 — RAG

| ID | Title | Priority | Estimate | Status | GitHub Issue |
|----|-------|----------|----------|--------|--------------|
| CAP-022 | FAQ ingestion pipeline | High | M | ⬜ | |
| CAP-023 | Chunking strategy | Medium | S | ⬜ | |
| CAP-024 | Generate embeddings | High | M | ⬜ | |
| CAP-025 | sqlite-vec integration | High | M | ⬜ | |
| CAP-026 | Semantic retrieval | High | M | ⬜ | |
| CAP-027 | FAQ tool | High | M | ⬜ | |

---

# Milestone 5 — API

| ID | Title | Priority | Estimate | Status | GitHub Issue |
|----|-------|----------|----------|--------|--------------|
| CAP-028 | Chat endpoint | High | M | ⬜ | |
| CAP-029 | Streaming responses | Medium | M | ⬜ | |
| CAP-030 | Request validation | Medium | S | ⬜ | |
| CAP-031 | Error handling | Medium | S | ⬜ | |

---

# Milestone 6 — Observability

| ID | Title | Priority | Estimate | Status | GitHub Issue |
|----|-------|----------|----------|--------|--------------|
| CAP-032 | OpenTelemetry setup | High | M | ⬜ | |
| CAP-033 | HTTP tracing | High | S | ⬜ | |
| CAP-034 | Tool tracing | High | M | ⬜ | |
| CAP-035 | SQLite tracing | High | M | ⬜ | |
| CAP-036 | Ollama tracing | High | M | ⬜ | |
| CAP-037 | Phoenix integration | High | M | ⬜ | |

---

# Milestone 7 — Evaluation

| ID | Title | Priority | Estimate | Status | GitHub Issue |
|----|-------|----------|----------|--------|--------------|
| CAP-038 | Evaluation dataset | High | M | ⬜ | |
| CAP-039 | Evaluation runner | High | L | ⬜ | |
| CAP-040 | Retrieval evaluation | Medium | M | ⬜ | |
| CAP-041 | Tool evaluation | Medium | M | ⬜ | |
| CAP-042 | Business rule evaluation | Medium | M | ⬜ | |
| CAP-043 | LLM-as-a-Judge | Medium | L | ⬜ | |
| CAP-044 | Evaluation reports | Medium | M | ⬜ | |

---

# Milestone 8 — Testing

| ID | Title | Priority | Estimate | Status | GitHub Issue |
|----|-------|----------|----------|--------|--------------|
| CAP-045 | Unit tests | High | M | ⬜ | |
| CAP-046 | Integration tests | High | L | ⬜ | |
| CAP-047 | End-to-end tests | Medium | L | ⬜ | |

---

# Milestone 9 — Frontend

| ID | Title | Priority | Estimate | Status | GitHub Issue |
|----|-------|----------|----------|--------|--------------|
| CAP-048 | React chat UI | Medium | L | ⬜ | |
| CAP-049 | Streaming UI | Medium | M | ⬜ | |
| CAP-050 | Conversation history | Medium | M | ⬜ | |

---

# Milestone 10 — Production Readiness

| ID | Title | Priority | Estimate | Status | GitHub Issue |
|----|-------|----------|----------|--------|--------------|
| CAP-051 | Docker support | Medium | M | ⬜ | |
| CAP-052 | GitHub Actions | Medium | M | ⬜ | |
| CAP-053 | Run evaluations in CI | Medium | M | ⬜ | |
| CAP-054 | Documentation | Medium | M | ⬜ | |
| CAP-055 | Performance testing | Low | M | ⬜ | |

---

# Future Enhancements

- Multi-agent architecture
- Agent Gateway
- MCP Server
- Recommendation engine
- Authentication
- Authorization
- Shopping cart agent
- Inventory agent
- Customer support agent
- Human approval workflows
- Multiple LLM providers
- Kubernetes deployment

---

# Definition of Done

A backlog item is considered complete when:

- Implementation completed
- Unit tests added
- Documentation updated
- Code reviewed
- CI passes
- Merged into `main`