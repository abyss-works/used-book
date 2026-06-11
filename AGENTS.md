# AGENTS.md — used-book 프로젝트

> 이 문서는 used-book 프로젝트의 기술 스택, 구조, 컨벤션을 정의한다.

---

## 프로젝트 개요

| 항목 | 내용 |
|------|------|
| 저장소 | `abyss-works/used-book` |
| 언어 | Go 1.25 |
| DB | PostgreSQL (lib/pq) |
| API | REST (net/http, no framework) |
| 프론트엔드 | Vue 3 + TypeScript + Vite (임베디드 정적 파일, `embed.FS`) |
| 외부 API | 알라딘 Open API (TTB Key, ItemLookUp + usedList) |
| 빌드 | `go build -o used-book .` |
| 배포 | Docker → kind k8s 클러스터 |

## 로컬 저장소 경로 규칙

모든 Git 저장소는 `~/<org-name>/<repo-name>/` 경로에 clone한다.

| 저장소 | 로컬 경로 |
|--------|----------|
| `abyss-works/board` | `~/abyss-works/board/` |
| `abyss-works/used-book` | `~/abyss-works/used-book/` |

## 디렉토리 구조

```
used-book/
├── main.go              # 진입점, 라우팅, 핸들러
├── aladin/              # 알라딘 Open API 클라이언트
│   └── client.go
├── optimizer/           # 중고책 최적 조합 엔진
│   └── optimizer.go
├── model/               # 도메인 모델
│   └── book.go
├── handler/             # HTTP 핸들러
│   └── api.go
├── frontend/            # Vue 3 + TypeScript + Vite
│   └── src/
│       ├── api/         # Axios HTTP calls
│       ├── stores/      # Pinia state
│       ├── views/       # Pages
│       └── components/  # Reusable
├── Dockerfile            # 멀티스테이지 빌드
├── k8s/                  # k8s manifests
│   ├── app.yaml          # Deployment + Service
│   └── postgres.yaml     # PostgreSQL StatefulSet
├── AGENTS.md             # (이 파일)
└── CLAUDE.md
```

## 빌드 & 테스트

```bash
go build -o used-book .   # 빌드
go vet ./...               # 정적 분석
```

## API 패턴

- 표준 `net/http` 핸들러
- JSON 요청/응답
- 라우팅: `http.HandleFunc` with path prefix matching
- DB: `database/sql` + `lib/pq`
- 포트: 환경변수 `PORT` (기본 8080)
- DB 연결: `DATABASE_URL` 환경변수
- 알라딘 API Key: 환경변수 `ALADIN_TTB_KEY`

## 배포 (k8s)

- kind 클러스터 (`abyssworks`)
- 네임스페이스: `default`
- Ingress: `*.abyssworks.dev`
- PostgreSQL: 동일 클러스터 내 StatefulSet

## 커밋 컨벤션

### 브랜치 전략

- `main` — 릴리즈 브랜치
- `dev` — 개발/다듬기 브랜치
- feature 브랜치 — `feat/*` 또는 `fix/*`

### 커밋 메시지

- 타입: `feat:`, `fix:`, `refactor:`, `chore:`, `docs:`, `test:`, `init:`
- 제목: 한글 30자 이내, 명사형 종결
- 본문: 불릿 리스트, 파일 단위 변경 추적
- 커밋 단위: 원자적, 계층별 분리 (Data → API → UI → Docs)
