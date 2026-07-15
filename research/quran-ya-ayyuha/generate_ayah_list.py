#!/usr/bin/env python3
"""Generate list of Quran ayahs beginning with يَا أَيُّهَا الَّذِينَ آمَنُوا."""

import csv
import json
import re
import urllib.request
from pathlib import Path

API_URL = "https://api.alquran.cloud/v1/quran/quran-uthmani"
OUTPUT_DIR = Path(__file__).parent

# Standard scholarly list of 89 ayahs (Quran order).
STANDARD_89 = [
    (2, 104), (2, 153), (2, 172), (2, 178), (2, 183), (2, 208), (2, 254),
    (2, 264), (2, 267), (2, 278), (2, 282),
    (3, 100), (3, 102), (3, 118), (3, 130), (3, 149), (3, 156), (3, 200),
    (4, 19), (4, 29), (4, 43), (4, 59), (4, 71), (4, 94), (4, 135), (4, 136), (4, 144),
    (5, 1), (5, 2), (5, 6), (5, 8), (5, 11), (5, 35), (5, 51), (5, 54), (5, 57),
    (5, 87), (5, 90), (5, 94), (5, 95), (5, 101), (5, 105), (5, 106),
    (8, 15), (8, 20), (8, 24), (8, 27), (8, 29), (8, 45),
    (9, 23), (9, 28), (9, 34), (9, 38), (9, 119), (9, 123),
    (22, 77),
    (24, 21), (24, 27), (24, 58),
    (33, 9), (33, 41), (33, 49), (33, 53), (33, 56), (33, 69), (33, 70),
    (47, 7), (47, 33),
    (49, 1), (49, 2), (49, 6), (49, 11), (49, 12),
    (57, 28),
    (58, 9), (58, 11), (58, 12),
    (59, 18),
    (60, 1), (60, 10), (60, 13),
    (61, 2), (61, 10), (61, 14),
    (62, 9),
    (63, 9),
    (64, 14),
    (66, 6), (66, 8),
]

SPECIAL_NOTES = {
    (5, 1): "bismillah_prefix",
    (49, 1): "bismillah_prefix",
    (60, 1): "bismillah_prefix",
    (33, 56): "mid_ayah_clause",
}

SURAH_NAMES = {
    2: "Al-Baqarah", 3: "Aal-Imran", 4: "An-Nisa", 5: "Al-Ma'idah",
    8: "Al-Anfal", 9: "At-Tawbah", 22: "Al-Hajj", 24: "An-Noor",
    33: "Al-Ahzab", 47: "Muhammad", 49: "Al-Hujurat", 57: "Al-Hadeed",
    58: "Al-Mujadalah", 59: "Al-Hashr", 60: "Al-Mumtahinah", 61: "As-Saff",
    62: "Al-Jumuah", 63: "Al-Munafiqoon", 64: "At-Taghabun", 66: "At-Tahreem",
}


def normalize(text: str) -> str:
    text = re.sub(r"[\u064B-\u065F\u0670\u06D6-\u06ED\u06DF\u06E2\u06ED]", "", text)
    text = re.sub(r"[\u0622\u0623\u0625\u0671\u0626\u0624]", "\u0627", text)
    text = text.replace("\u0621", "")
    text = text.replace("\u0649", "\u064A")
    text = text.replace("\u0640", "")
    text = re.sub(r"[\u200F\u200E\u0610-\u061A]", "", text)
    text = re.sub(r"\s+", "", text)
    return text.strip()


PHRASE_PATTERN = re.compile(r"^يايهاالذينامنوا")
BISMILLAH_PATTERN = re.compile(r"^بسماللهالرحمنالرحيميايهاالذينامنوا")


def fetch_quran() -> dict:
    with urllib.request.urlopen(API_URL) as response:
        return json.load(response)


def build_ayah_index(data: dict) -> dict[tuple[int, int], str]:
    index = {}
    for surah in data["data"]["surahs"]:
        sn = surah["number"]
        for ayah in surah["ayahs"]:
            index[(sn, ayah["numberInSurah"])] = ayah["text"]
    return index


def classify_ayah(text: str) -> str:
    norm = normalize(text)
    if PHRASE_PATTERN.match(norm):
        return "direct_start"
    if BISMILLAH_PATTERN.match(norm):
        return "bismillah_prefix"
    if "يايهاالذينامنوا" in norm:
        return "mid_ayah_clause"
    return "no_match"


def note_description(note_key: str | None) -> str:
    if note_key == "bismillah_prefix":
        return "Ayah begins with Bismillah, then يَا أَيُّهَا الَّذِينَ آمَنُوا"
    if note_key == "mid_ayah_clause":
        return "Phrase appears as second clause in the ayah (after salawat on the Prophet)"
    return ""


def main() -> None:
    data = fetch_quran()
    ayah_index = build_ayah_index(data)

    entries = []
    for surah, ayah in STANDARD_89:
        text = ayah_index[(surah, ayah)]
        classification = classify_ayah(text)
        note_key = SPECIAL_NOTES.get((surah, ayah))
        entries.append({
            "reference": f"{surah}:{ayah}",
            "surah": surah,
            "ayah": ayah,
            "surah_name": SURAH_NAMES.get(surah, ""),
            "arabic_text": text,
            "classification": classification,
            "note": note_key or "",
            "note_description": note_description(note_key),
        })

    # Verify API scan matches standard list
    scanned_direct = []
    for (surah, ayah), text in sorted(ayah_index.items()):
        if classify_ayah(text) == "direct_start":
            scanned_direct.append(f"{surah}:{ayah}")

    standard_refs = {f"{s}:{a}" for s, a in STANDARD_89}
    direct_refs = set(scanned_direct)
    special_refs = {f"{s}:{a}" for s, a in SPECIAL_NOTES}
    assert len(entries) == 89
    assert direct_refs | special_refs == standard_refs, (
        f"Mismatch: missing={standard_refs - direct_refs - special_refs}, "
        f"extra={direct_refs | special_refs - standard_refs}"
    )

    output = {
        "phrase": "يَا أَيُّهَا الَّذِينَ آمَنُوا",
        "meaning": "O you who have believed",
        "total_count": 89,
        "direct_start_count": 85,
        "special_cases_count": 4,
        "ayahs": entries,
    }

    json_path = OUTPUT_DIR / "ayah-ya-ayyuha.json"
    csv_path = OUTPUT_DIR / "ayah-ya-ayyuha.csv"

    with open(json_path, "w", encoding="utf-8") as f:
        json.dump(output, f, ensure_ascii=False, indent=2)

    with open(csv_path, "w", encoding="utf-8", newline="") as f:
        writer = csv.DictWriter(
            f,
            fieldnames=[
                "reference", "surah", "ayah", "surah_name",
                "classification", "note", "note_description", "arabic_text",
            ],
        )
        writer.writeheader()
        writer.writerows(entries)

    print(f"Generated {len(entries)} ayahs")
    print(f"JSON: {json_path}")
    print(f"CSV:  {csv_path}")


if __name__ == "__main__":
    main()
