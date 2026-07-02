# Commerce AI Platform Architecture

## Overview

Commerce AI Platform is an enterprise-grade AI application that demonstrates modern AI engineering practices using Go and Ollama.

The platform allows users to interact with an AI shopping assistant capable of:

- Searching products
- Checking prices and availability
- Creating orders
- Returning orders
- Answering FAQ questions using RAG

The project emphasizes production-quality architecture including:

- AI Agents
- Tool Calling
- Retrieval-Augmented Generation (RAG)
- OpenTelemetry
- AI Observability
- LLM Evaluations

---

# High Level Architecture

```
                          User
                            │
                            ▼
                    REST API (Gin)
                            │
                            ▼
                     AI Agent Runtime
                            │
        ┌───────────────────┼────────────────────┐
        │                   │                    │
        ▼                   ▼                    ▼
  Product Tool        Order Tool          FAQ Tool (RAG)
        │                   │                    │
        └──────────────┬────┴────────────────────┘
                       ▼
                    SQLite
             (Products / Orders)

                       ▲
                       │
                sqlite-vec
                FAQ Embeddings

                       ▲
                       │
              Ollama Embeddings

                       │
                       ▼
                    Ollama LLM

                       │
                       ▼
              OpenTelemetry SDK

                       │
                       ▼
                 Arize Phoenix
```

---

# Core Components

## API Layer

Responsibilities

- REST endpoints
- Request validation
- Authentication (future)
- Streaming responses

Technology

- Gin

---

## AI Agent

Responsibilities

- Conversation management
- Tool selection
- Tool execution
- Response generation

---

## Tool Framework

Tools expose business capabilities to the LLM.

Examples

- Product Search
- Price Lookup
- Availability
- Order Creation
- Return Processing
- FAQ Search

---

## Database

SQLite stores transactional data.

Tables

- Products
- Orders
- Order Items

---

## Retrieval Augmented Generation

Knowledge stored outside the LLM.

Sources

- FAQ
- Return Policy
- Shipping Policy
- Store Information

Workflow

```
FAQ

↓

Chunking

↓

Embeddings

↓

sqlite-vec

↓

Similarity Search

↓

LLM
```

---

## Observability

OpenTelemetry captures

- HTTP requests
- LLM calls
- Tool execution
- SQL queries
- Errors

Telemetry is exported to

- Arize Phoenix

---

## Evaluation

The platform includes automated evaluation.

Metrics include

- Retrieval Quality
- Tool Selection
- Business Rule Compliance
- Hallucination Detection
- Response Quality

---

# Project Structure

```
commerce-ai-platform/

cmd/
    api/
    evals/
    ingest/

config/

internal/

    agent/
    tools/
    rag/
    evaluation/

    platform/
        config/
        logging/
        telemetry/
        database/

docs/

data/

web/
```

---

# Design Principles

- Keep business logic outside the LLM.
- Use tools instead of prompting for transactional operations.
- Use RAG for unstructured knowledge.
- Instrument everything with OpenTelemetry.
- Continuously evaluate agent quality.
- Prefer modular, testable components.

---

# Future Architecture

The platform will evolve into a multi-agent system.

```
                    Gateway

                        │

       ┌────────────────┼────────────────┐

       ▼                ▼                ▼

Shopping Agent    Returns Agent    Support Agent

       │                │                │

       └───────────────Shared Tools──────┘
```