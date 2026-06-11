# used-book

> 중고책 최적 조합 — 여러 판매자의 중고책을 배송비까지 고려한 최적 구매 조합 탐색

## 개요

알라딘 Open API를 통해 중고책 정보를 수집하고, 판매자별 가격·상태·배송비를 종합적으로 고려하여
**총 비용이 가장 낮은 구매 조합**을 찾아주는 웹 애플리케이션.

### 문제 상황

- 같은 책이라도 판매자마다 가격·상태가 다름
- 여러 판매자에게서 구매할 경우 배송비가 중복됨
- 단순 최저가가 아닌, **배송비를 고려한 최적 조합** 필요

### 기술 스택

| 영역 | 기술 |
|------|------|
| 백엔드 | Go 1.25 + `net/http` |
| 프론트엔드 | Vue 3 + TypeScript + Vite + Tailwind CSS |
| 데이터베이스 | PostgreSQL |
| 인프라 | Docker + Kubernetes (Kind) |
| CI/CD | GitHub Actions (Self-hosted Runner) |
| 외부 API | 알라딘 Open API (TTB Key) |

## 프로젝트 구조

```
used-book/
├── main.go              # 서버 진입점
├── aladin/              # 알라딘 API 클라이언트
├── optimizer/           # 최적 조합 엔진
├── model/               # 도메인 모델
├── handler/             # HTTP 핸들러
├── frontend/            # Vue 3 SPA
├── k8s/                 # Kubernetes 매니페스트
└── .github/workflows/  # CI/CD 파이프라인
```

## 라이선스

MIT
