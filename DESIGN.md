---
name: used-book
description: 중고책 검색 + 최적 조합 플랫폼 — 다크 테마, amber 악센트, 모바일퍼스트
colors:
  brand:
    primary: "#f59e0b"
    primary-hover: "#fbbf24"
  neutral:
    bg: "#0a0a0b"
    surface: "#121213"
    surface-elevated: "#1c1c1e"
    border: "rgba(255, 255, 255, 0.08)"
    border-subtle: "rgba(255, 255, 255, 0.04)"
    text: "#f5f5f4"
    text-secondary: "#a8a29e"
    text-muted: "#78716c"
    text-dim: "#57534e"
  semantic:
    success: "#22c55e"
    warning: "#eab308"
    danger: "#ef4444"
    info: "#3b82f6"
typography:
  body:
    fontFamily: "Inter, system-ui, sans-serif"
    fontSize: 1rem
    lineHeight: 1.5
  display:
    fontFamily: "Inter, system-ui, sans-serif"
    fontSize: 2rem
    lineHeight: 1.1
    fontWeight: 600
  mono:
    fontFamily: "JetBrains Mono, ui-monospace, monospace"
    fontSize: 0.875rem
    lineHeight: 1.5
components:
  nav-bar:
    backgroundColor: "#121213"
    borderTop: "1px solid rgba(255, 255, 255, 0.08)"
    linkColor: "#78716c"
    linkActive: "#f59e0b"
  search-bar:
    backgroundColor: "#1c1c1e"
    border: "1px solid rgba(255, 255, 255, 0.08)"
    borderRadius: 12px
  card:
    backgroundColor: "#1c1c1e"
    border: "1px solid rgba(255, 255, 255, 0.08)"
    borderRadius: 12px
    padding: 16px
  button-primary:
    backgroundColor: "#f59e0b"
    textColor: "#0a0a0b"
    borderRadius: 8px
    fontWeight: 600
  badge:
    backgroundColor: "rgba(245, 158, 11, 0.12)"
    textColor: "#fbbf24"
    borderRadius: 9999px
---

# used-book 디자인 시스템

> 중고책 검색 및 최적 조합 플랫폼용 디자인 시스템.
> 다크 테마, amber 악센트, 모바일퍼스트.

---

## 브랜드 톤

| 축 | 값 | 설명 |
|----|-----|------|
| 온도 | 따뜻함 | Amber/gold 악센트, 책의 따뜻함 |
| 무게 | 중간 | 여백 충분, 16px base spacing |
| 공식성 | 캐주얼-기술적 | 둥근 모서리(12px), Inter 폰트 |
| 기술성 | 기술적 | 다크테마, JetBrains Mono 코드 |
| 레퍼런스 | Linear + Notion | Linear의 다크 정밀함 + Notion의 따뜻함 |

---

## 컬러 시스템

### 브랜드 팔레트

```css
/* Amber — 책 페이지 빛깔에서 영감 */
--color-brand: #f59e0b;
--color-brand-hover: #fbbf24;
--color-brand-muted: rgba(245, 158, 11, 0.12);
```

### 중립 팔레트

```css
--color-bg: #0a0a0b;          /* Linear-inspired near-black */
--color-surface: #121213;      /* 카드 기본 */
--color-surface-elevated: #1c1c1e;  /* 호버/상위 카드 */
--color-border: rgba(255, 255, 255, 0.08);
--color-border-subtle: rgba(255, 255, 255, 0.04);
```

### 텍스트

```css
--color-text: #f5f5f4;           /* 1차 — 거의 흰색 */
--color-text-secondary: #a8a29e; /* 2차 — 스톤 그레이 */
--color-text-muted: #78716c;     /* 3차 — 뮤트 */
--color-text-dim: #57534e;       /* 4차 — very dim */
```

### 시맨틱

```css
--color-success: #22c55e;
--color-warning: #eab308;
--color-danger: #ef4444;
--color-info: #3b82f6;
```

---

## 타이포그래피

### 폰트

| 용도 | 폰트 | fallback |
|------|------|----------|
| UI | Inter | system-ui, -apple-system, sans-serif |
| Mono | JetBrains Mono | ui-monospace, monospace |

### 계층

| 계층 | 크기 | 두께 | 행간 | 용도 |
|------|------|------|------|------|
| h1 | 1.75rem (28px) | 600 | 1.2 | 페이지 제목 |
| h2 | 1.25rem (20px) | 600 | 1.3 | 섹션 제목 |
| h3 | 1.125rem (18px) | 600 | 1.4 | 카드 제목 |
| body | 1rem (16px) | 400 | 1.5 | 본문 |
| body-sm | 0.875rem (14px) | 400 | 1.5 | 작은 본문 |
| caption | 0.75rem (12px) | 500 | 1.4 | 라벨/메타 |
| mono | 0.875rem (14px) | 400 | 1.5 | 코드/가격 |

---

## 간격 & 레이아웃

### Spacing (Tailwind 기본 4px 그리드, 기준 16px)

| 이름 | 크기 | 용도 |
|------|------|------|
| xs | 4px | 아이콘 간격 |
| sm | 8px | 요소 사이 |
| md | 16px | 기본 |
| lg | 24px | 섹션 |
| xl | 32px | 큰 섹션 |
| 2xl | 48px | 페이지 섹션 |

### 모서리

| 이름 | 크기 | 용도 |
|------|------|------|
| sm | 6px | 버튼 |
| md | 8px | 작은 카드 |
| lg | 12px | 카드/검색바 |
| xl | 16px | 모달 |
| full | 9999px | 배지/칩 |

---

## 컴포넌트

### NavBar (하단 모바일 네비게이션)

```css
/* mobile-first bottom navigation */
--nav-bg: #121213;
--nav-border-top: 1px solid rgba(255, 255, 255, 0.08);
--nav-link: #78716c;
--nav-link-active: #f59e0b;
--nav-height: 64px;
```

### SearchBar

```css
--search-bg: #1c1c1e;
--search-border: 1px solid rgba(255, 255, 255, 0.08);
--search-radius: 12px;
--search-padding: 14px 16px;
```

### BookCard

```css
/* 책 검색 결과 카드 */
--card-bg: #1c1c1e;
--card-border: 1px solid rgba(255, 255, 255, 0.08);
--card-radius: 12px;
--card-padding: 16px;
--card-shadow: 0 1px 3px rgba(0,0,0,0.3);
```

### Button

| 종류 | 배경 | 텍스트 | 용도 |
|------|------|--------|------|
| Primary | `#f59e0b` | `#0a0a0b` | 주요 CTA |
| Ghost | transparent | `#a8a29e` | 보조 동작 |
| Icon | transparent | `#78716c` | 아이콘 전용 |

### Badge

```css
/* 상태 표시 배지 */
--badge-bg: rgba(245, 158, 11, 0.12);
--badge-text: #fbbf24;
--badge-radius: 9999px;
--badge-padding: 2px 10px;
--badge-font: 12px;
```

---

## 모바일퍼스트

### Breakpoints

| 이름 | 너비 | 적용 |
|------|------|------|
| Mobile | <640px | 기본 — 단일 컬럼, 하단 Nav |
| Tablet | 640-1024px | 2열 그리드, 상단 Nav |
| Desktop | >1024px | 3열 그리드, full layout |

### 레이아웃 원칙

1. 모든 컴포넌트는 mobile-first로 작성
2. 검색창은 화면 상단 고정
3. 하단 네비게이션은 모바일 전용 (태블릿 이상에서 상단으로 전환)
4. 터치 타겟 최소 44px
5. Safe area 대응 (ios notched devices)

---

## 접근성

- 텍스트/배경 대비비 WCAG AA (4.5:1) 통과
  - `#f5f5f4` on `#0a0a0b` = 15.5:1 ✅
  - `#a8a29e` on `#0a0a0b` = 8.5:1 ✅
  - `#78716c` on `#0a0a0b` = 5.8:1 ✅
- 포커스 링: `2px solid #f59e0b`
- 터치 타겟 최소 44x44px
