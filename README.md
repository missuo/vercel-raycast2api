# Raycast to OpenAI API for Vercel

Access Raycast AI models through an OpenAI-compatible API, deployed effortlessly with Vercel.

[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/clone?repository-url=https%3A%2F%2Fgithub.com%2Fmissuo%2Fvercel-raycast2api)

Unfortunatly, Vercel has not implemented `http.Flusher` yet, so the streaming is not supported.

## 🌟 Features

- 🧠 OpenAI-compatible interface for Raycast AI
- 🚀 Instant Vercel deployment
- 🛡️ Optional API key protection
- ⚙️ Simple configuration via environment variables
- ⚡️ Fully compatible with tools like [Cursor](https://cursor.sh)

## 🔧 Setup

### 1. Deploy to Vercel

Click the button above to instantly deploy on Vercel. During deployment, you'll be prompted to configure the following environment variables:

| Variable | Required | Description |
|:--------|:--------:|:------------|
| `RAYCAST_BEARER_TOKEN` | ✅ | Your Raycast Bearer Token |
| `API_KEY` | ⭕ | Optional – set to require bearer token in requests |
| `PORT` | ❌ | Ignored on Vercel (automatically managed) |

> 📌 You can get your `RAYCAST_BEARER_TOKEN` from your Raycast account.

### 2. After Deployment

Your deployed endpoint will be available at:

```
https://<your-vercel-subdomain>.vercel.app/v1
```

## 📡 API Usage

All endpoints follow the OpenAI API pattern:

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/v1/models` | GET | List available models |
| `/v1/chat/completions` | POST | Create chat completions |
| `/v1/refresh-models` | GET | Refresh model list |
| `/health` | GET | Health check |

### Example `curl`:

```bash
curl https://<your-endpoint>.vercel.app/v1/chat/completions \
  -H "Authorization: Bearer your-api-key" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "claude-3-opus-20240229",
    "messages": [{"role": "user", "content": "Hello!"}]
  }'
```

## ✅ Supported Models

Over 30 models from providers like Anthropic, OpenAI, Google, Groq, Mistral, Perplexity, and Raycast.

See full model list via:

```
GET /v1/models
```

## 🖥️ Local Development (Optional)

If you prefer to run locally:

```bash
git clone https://github.com/missuo/raycast2api
cd raycast2api
go mod tidy
go build
./raycast2api
```

Or with Docker:

```bash
docker compose up -d --build
```

## 🙏 Credits

- [szcharlesji/raycast-relay](https://github.com/szcharlesji/raycast-relay)
- [xjasonlyu/vercel-deeplx](https://github.com/xjasonlyu/vercel-deeplx)
- [missuo/raycast2api](https://github.com/missuo/raycast2api)

## 📄 License

MIT License

