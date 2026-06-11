# Infrastructure — used-book

> 인프라, 배포, Git 워크플로우.

---

## 프로젝트 실행

### 로컬 개발

1. **DB 실행**: `docker run -d --name usedbook-pg -e POSTGRES_DB=usedbook -e POSTGRES_PASSWORD=postgres -p 5432:5432 postgres:16`
2. **백엔드 실행**: `ALADIN_TTB_KEY=... DB_HOST=localhost go run .`
3. **프론트엔드 실행**: `cd frontend && npm run dev`

### Kubernetes 배포

```bash
kubectl apply -f k8s/postgres.yaml
kubectl apply -f k8s/app.yaml
kubectl apply -f k8s/ingress.yaml
```

---

## Git 브랜치 전략

| 구분 | 설명 |
|------|------|
| 메인 브랜치 | `main` |
| 작업 브랜치 | `feat/*`, `fix/*` |
| 커밋 컨벤션 | [commit-convention.md](../commit-convention.md) 참조 |
| 커밋 형식 | `feat:`, `fix:`, `chore:`, `docs:`, `init:` |

---

## 환경 변수

| 변수 | 기본값 | 설명 |
|------|--------|------|
| `DB_HOST` | `postgres` | PostgreSQL 호스트 |
| `DB_PORT` | `5432` | PostgreSQL 포트 |
| `DB_USER` | `postgres` | DB 사용자 |
| `DB_PASSWORD` | `postgres` | DB 비밀번호 |
| `DB_NAME` | `usedbook` | DB 이름 |
| `PORT` | `8080` | 서버 포트 |
| `ALADIN_TTB_KEY` | — | 알라딘 Open API 인증키 |

---

## k8s 매니페스트

| 파일 | 내용 |
|------|------|
| `k8s/postgres.yaml` | PostgreSQL Deployment + PVC + Service |
| `k8s/app.yaml` | used-book Deployment + Service |
| `k8s/ingress.yaml` | Ingress 라우팅 (used-book.abyssworks.dev) |
| `Dockerfile` | 멀티스테이지 Go 빌드 → 경량 실행 이미지 |

네임스페이스: `usedbook-prod`

---

## 빌드

- **백엔드**: `go build -o used-book .` → 단일 바이너리
- **프론트엔드**: `npm run build` → `frontend/dist/`
- **Docker**: `docker build -t used-book .`
