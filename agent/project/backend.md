# Backend — used-book

> 백엔드 기술 스택, 모듈 구조, 계층 아키텍처, API.

---

## 기술 스택

| 영역 | 기술 |
|------|------|
| 런타임 | Go 1.25 |
| 프레임워크 | `net/http` (표준 라이브러리) |
| 빌드 | `go build` |
| DB 드라이버 | `github.com/lib/pq` (PostgreSQL) |
| DB | PostgreSQL (운영/개발 공통) |
| 외부 API | 알라딘 Open API (TTB Key 인증) |

---

## 모듈 구조

```
used-book/
├── main.go              # 진입점: DB 초기화, 라우트 등록, 서버 시작
├── aladin/              # 알라딘 Open API 클라이언트
│   └── client.go        #   ItemSearch, ItemLookUp(+usedList) 호출
├── optimizer/           # 중고책 최적 조합 엔진
│   └── optimizer.go     #   배송비·가격·상태 종합 최적화
├── model/               # 도메인 모델
│   └── book.go          #   Book, UsedItem, Seller, WishlistEntry
├── handler/             # HTTP 핸들러
│   └── api.go           #   /api/search, /api/optimize, /api/wishlist
├── frontend/            # Vue 3 + TypeScript + Vite (임베디드)
```

---

## 계층 아키텍처

```
Handler (HTTP 요청/응답)
  → Aladin Client (외부 API 호출)
  → Optimizer (비즈니스 로직 — 최적 조합 계산)
    → Model (도메인 객체)
```

- Handler는 입력 검증 + JSON 직렬화만 담당
- aladin.Client는 외부 API 호출과 응답 파싱
- optimizer는 순수 계산 로직 (DB 없이 in-memory)

---

## API 엔드포인트

### 위시리스트 (`/api/wishlist`)

| 메서드 | 경로 | 설명 |
|--------|------|------|
| GET | `/api/wishlist` | 위시리스트 목록 조회 |
| POST | `/api/wishlist` | 위시리스트에 책 추가 |
| DELETE | `/api/wishlist/{id}` | 위시리스트에서 제거 |

### 검색 및 최적화

| 메서드 | 경로 | 설명 |
|--------|------|------|
| GET | `/api/search?q={query}` | 알라딘 도서 검색 |
| GET | `/api/optimize` | 위시리스트 기반 최적 조합 계산 |
| POST | `/api/optimize` | 특정 리스트 기반 최적 조합 계산 |

---

## 엔티티

```go
type Book struct {
    ID          string  `json:"id"`          // 알라딘 ID (itemId)
    Title       string  `json:"title"`
    Author      string  `json:"author"`
    Isbn        string  `json:"isbn"`
    Cover       string  `json:"cover"`
}

type UsedItem struct {
    Seller     string `json:"seller"`
    Price      int    `json:"price"`       // 원
    Condition  string `json:"condition"`   // 최상/상/중
    DeliveryFee int   `json:"delivery_fee"` // 배송비
    Stock      int    `json:"stock"`
}

type OptimizationResult struct {
    TotalCost int               `json:"total_cost"`
    ShipCount int               `json:"ship_count"`
    Items     []PurchasedItem   `json:"items"`
}
```

---

## 구조 원칙

| 원칙 | 설명 |
|------|------|
| **PK** | 컬럼명 `id`, `SERIAL` (auto-increment) |
| **JSON 태그** | snake_case 사용 |
| **시간** | `time.Time`, DB는 `TIMESTAMP` |
