# HieroglyphsScraper

## Project

Scraper written in Golang for fetching Hieroglyphs from Wikipedia

2 JSON files are generated:
- `categories.json` containing the Gardiner's sign categories (26) ([Source](https://en.wikipedia.org/wiki/Gardiner%27s_sign_list)):

```json
[
  {
    "id": "A",
    "name": "A. Man and his occupations"
  },
  {
    "id": "B",
    "name": "B. Woman and her occupations"
  },
  {
    "id": "C",
    "name": "C. Anthropomorphic deities"
  },
  ...
]
```

- `hieroglyphs.json` containing the hieroglyphs (1071) ([Source](https://en.wikipedia.org/wiki/List_of_Egyptian_hieroglyphs)):

```json
[
  {
    "id": "A1",
    "category_id": "A",
    "unicode": "U+13000",
    "description": "seated man",
    "transliteration": "I (Masculine) (paeu)\nmale, man, typical masculine role, son, courtier (Masculine)",
    "phonetic": "",
    "note": "Commonly placed behind a name to indicate masculine sex of named person."
  },
  {
    "id": "A2",
    "category_id": "A",
    "unicode": "U+13001",
    "description": "man with hand to mouth",
    "transliteration": "eat (wnm)\ndrink (swr)\nspeak, think, feel, tell (sdjd)\nrefrain from speech (gr) advise/counsel (kAj)\nlove (mrj)",
    "phonetic": "",
    "note": "Activities involving the mouth, head, or ideas."
  },
  ...
]
```

## Usage

```bash
go build -o hieroglyphs-scraper hieroglyphs-scraper.go
./hieroglyphs-scraper
```
