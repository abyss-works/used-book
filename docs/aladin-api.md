# 알라딘 Open API 필드 참조

> 기준: API Version `20131101`, Output `JS` (JSON)

---

## 1. ItemSearch.aspx — 상품 검색

**요청 파라미터**

| 필드 | 타입 | 필수 | 설명 |
|------|------|------|------|
| `ttbkey` | string | Y | 알라딘 API 인증키 |
| `Query` | string | Y | 검색어 |
| `QueryType` | string | N | `Title`(기본), `Author`, `Publisher`, `Keyword` |
| `MaxResults` | int | N | 최대 결과 수 (1~100, 기본 10) |
| `start` | int | N | 시작 위치 (기본 1) |
| `SearchTarget` | string | N | `Book`(기본), `Foreign`, `Music`, `DVD`, `Used` |
| `output` | string | N | `xml`(기본), `js` |
| `Version` | string | N | `20131101` |
| `Cover` | string | N | 표지 크기: `Big`, `MidBig`, `Mid`(기본), `Small`, `Mini`, `None` |
| `CategoryId` | int | N | 특정 카테고리로 제한 |
| `Sort` | string | N | 정렬: `Accuracy`(기본, 정확도), `SalesPoint`(판매량), `PublishDate`(출간일), `CustomerRating`(고객평점), `MyReview`(마이리뷰순) |

**응답 (루트)**

| 필드 | 타입 | 설명 |
|------|------|------|
| `version` | string | API 버전 |
| `logo` | string | 알라딘 로고 URL |
| `title` | string | 검색 결과 타이틀 |
| `link` | string | 알라딘 페이지 URL |
| `pubDate` | string | 응답 생성 시간 |
| `totalResults` | int | 총 검색 결과 수 |
| `startIndex` | int | 시작 인덱스 |
| `itemsPerPage` | int | 페이지당 항목 수 |
| `query` | string | 검색어 |
| `searchCategoryId` | int | 검색 카테고리 ID |
| `searchCategoryName` | string | 검색 카테고리명 |
| `item[]` | array | 상품 목록 |

**응답 `item[]` (상품)**

| 필드 | 타입 | 설명 |
|------|------|------|
| `itemId` | int | 알라딘 고유 상품 ID |
| `title` | string | 제목 |
| `link` | string | 알라딘 상품 페이지 URL |
| `author` | string | 저자 |
| `pubDate` | string | 출간일 (`YYYY-MM-DD`) |
| `description` | string | 책 소개 |
| `isbn` | string | ISBN10 |
| `isbn13` | string | ISBN13 |
| `priceSales` | int | 판매가 |
| `priceStandard` | int | 정가 |
| `mallType` | string | `BOOK`, `MUSIC`, `DVD`, `USED` |
| `stockStatus` | string | 재고 상태 (빈 문자열=정상) |
| `mileage` | int | 마일리지 |
| `cover` | string | 표지 이미지 URL |
| `categoryId` | int | 카테고리 ID |
| `categoryName` | string | 카테고리명 (예: `국내도서>컴퓨터/모바일>컴퓨터 공학>소프트웨어 공학`) |
| `publisher` | string | 출판사 |
| `salesPoint` | int | 판매지수 |
| `adult` | bool | 성인 여부 |
| `fixedPrice` | bool | 정가제 적용 여부 |
| `customerReviewRank` | int | 고객 평점 (1~10) |
| `seriesInfo` | object | 시리즈 정보 (옵션) |
| `subInfo` | object | 부가 정보 (옵션, itemLookUp에서 전체 제공) |

**`seriesInfo`**

| 필드 | 타입 | 설명 |
|------|------|------|
| `seriesId` | int | 시리즈 ID |
| `seriesLink` | string | 시리즈 URL |
| `seriesName` | string | 시리즈명 |

---

## 2. ItemLookUp.aspx — 상품 조회

**요청 파라미터**

| 필드 | 타입 | 필수 | 설명 |
|------|------|------|------|
| `ttbkey` | string | Y | 알라딘 API 인증키 |
| `ItemId` | string | Y | ISBN 또는 알라딘 상품 ID (`itemIdType`에 따라 다름) |
| `itemIdType` | string | N | `ISBN`(기본, 10자리), `ISBN13`(13자리), `ItemId`(알라딘ID) |
| `output` | string | N | `xml`(기본), `js` |
| `Version` | string | N | `20131101` |
| `OptResult` | string | N | 쉼표 구분 부가정보 옵션 (아래 참조) |
| `Cover` | string | N | 표지 크기 |
| `offCode` | string | N | 중고매장 코드 |

**`OptResult` 옵션값**

| 값 | 설명 |
|----|------|
| `ebookList` | 전자책(eBook) 정보 |
| `usedList` | 중고상품 요약 정보 |
| `reviewList` | 리뷰 정보 |
| `fileFormatList` | 전자책 파일 포맷/용량 |
| `c2binfo` | C2B(알라딘 매입) 정보 |
| `packing` | 판형/포장 정보 |
| `b2bSupply` | 전자책 B2B 납품 가능 여부 |
| `subbarcode` | 부가기호 |
| `cardReviewImgList` | 카드리뷰 이미지 |

**응답 (루트)**

ItemSearch와 동일한 구조 + 추가 부가정보.

| 필드 | 타입 | 설명 |
|------|------|------|
| `version` | string | API 버전 |
| `title` | string | 상품 정보 타이틀 |
| `link` | string | 알라딘 페이지 URL |
| `pubDate` | string | 응답 생성 시간 |
| `totalResults` | int | 1 |
| `startIndex` | int | 1 |
| `itemsPerPage` | int | 1 |
| `item[]` | array | 상품 배열 (단일 항목) |

**응답 `item[]` — Search와 동일한 필드 + 다음 추가 필드:**

**`subInfo` (OptResult 포함 시)**

| 필드 | 타입 | 설명 |
|------|------|------|
| `subTitle` | string | 부제목 |
| `originalTitle` | string | 원제 |
| `itemPage` | int | 쪽수 |
| `usedList` | object | 중고상품 요약 (아래 참조) |
| `ebookList` | object | eBook 정보 (OptResult=ebookList) |

**`subInfo.usedList` (OptResult=usedList)**

| 필드 | 타입 | 설명 |
|------|------|------|
| `aladinUsed` | object | 알라딘 직접 판매 중고 |
| `userUsed` | object | 개인 간 중고 거래 |
| `spaceUsed` | object | 오프라인 매장 중고 |

각 카테고리 객체:

| 필드 | 타입 | 설명 |
|------|------|------|
| `itemCount` | int | 매물 수 |
| `minPrice` | int | 최저 가격 |
| `link` | string | 상세 매물 페이지 URL |

> **중요**: usedList는 개별 매물 배열이 아니라 **3개 카테고리 요약 정보**만 제공.
> 개별 중고 매물 조회는 `ItemOffStore.aspx` API 또는 각 `link` URL 접근 필요.

---

## 3. ItemList.aspx — 상품 리스트

**요청 파라미터**

| 필드 | 타입 | 필수 | 설명 |
|------|------|------|------|
| `ttbkey` | string | Y | 알라딘 API 인증키 |
| `QueryType` | string | Y | 리스트 타입 (아래 참조) |
| `MaxResults` | int | N | 최대 결과 수 |
| `start` | int | N | 시작 위치 |
| `SearchTarget` | string | N | `Book`(기본), `Foreign`, `Music`, `DVD`, `Used` |
| `output` | string | N | `xml`, `js` |
| `Version` | string | N | `20131101` |
| `CategoryId` | int | N | 카테고리 필터 |
| `Cover` | string | N | 표지 크기 |

**`QueryType` 값**

| 값 | 설명 |
|----|------|
| `ItemNewAll` | 신간 전체 |
| `ItemNewSpecial` | 주목할 만한 신간 |
| `ItemEditorChoice` | 편집자 추천 |
| `Bestseller` | 베스트셀러 |
| `BlogBest` | 북플 베스트 (국내도서만) |

**응답**: ItemSearch와 동일 구조 (item[] 배열)

---

## 4. ItemOffStore.aspx — 중고상품 보유 매장 검색

> 현재 used-book 프로젝트에서 미구현.

**요청 파라미터**

| 필드 | 타입 | 필수 | 설명 |
|------|------|------|------|
| `ttbkey` | string | Y | 알라딘 API 인증키 |
| `itemIdType` | string | N | `ISBN`(기본), `ISBN13`, `ItemId` |
| `ItemId` | string | Y | 상품 ID |
| `output` | string | N | `xml`, `js` |
| `Version` | string | N | `20131101` |
| `offCode` | string | N | 특정 매장 코드 (생략 시 전체) |

---

## 5. Go 모델 — `model/book.go`

```go
type Book struct {
    ID       string `json:"id"`
    Title    string `json:"title"`
    Author   string `json:"author"`
    Isbn     string `json:"isbn"`
    Cover    string `json:"cover"`
    Category string `json:"category"`
}

type UsedItem struct {
    SellerID    string `json:"seller_id"`
    SellerName  string `json:"seller_name"`
    Price       int    `json:"price"`
    Condition   string `json:"condition"`
    DeliveryFee int    `json:"delivery_fee"`
    Stock       int    `json:"stock"`
    Link        string `json:"link"`
}

type AladinSearchResult struct {
    Books []Book `json:"books"`
    Total int    `json:"total"`
    Start int    `json:"start"`
    Max   int    `json:"max"`
}

type AladinUsedResult struct {
    Book  Book       `json:"book"`
    Items []UsedItem `json:"items"`
}
```

**검색 API → 우리 API 매핑**

| Aladin `item` 필드 | 우리 `book` 필드 |
|-------------------|-----------------|
| `itemId` (int) | `id` (string, fmt.Sprintf) |
| `title` | `title` |
| `author` | `author` |
| `isbn13` | `isbn` |
| `cover` | `cover` |
| — | `category` (미사용, 항상 빈 문자열) |

**조회 API → our API 매핑 (`usedList` 카테고리 → `items[]`)**

| 카테고리 키 | `seller_name` | `condition` | 비고 |
|-----------|--------------|------------|------|
| `aladinUsed` | 알라딘 중고 | 상 | `price`=minPrice, `stock`=itemCount |
| `userUsed` | 개인 판매자 | 중 | 상동 |
| `spaceUsed` | 알라딘 매장 | 상 | 상동 |

---

## 6. 참고사항

- 기본 응답 형식: XML (`output` 생략 시)
- JSON 요청 시 `output=JS` 지정
- 하루 API 호출 제한: **5,000회** (프리미엄 신청 시 증액 가능)
- 인코딩: UTF-8
- `ItemId`와 `ISBN`은 서로 다른 식별자 — `itemIdType`을 올바르게 설정 필요
- `ItemLookUp.aspx`에서 `item` 필드는 **배열** `[]`로 응답 (단일 결과도 배열)
