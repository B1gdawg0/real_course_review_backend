import requests
import json
import sys

URL1 = "https://cs.sci.ku.ac.th/api/program-year/undergraduate"
URL2 = "https://cs.sci.ku.ac.th/api/course/undergraduate"


def fetch_courses(year: int) -> dict:
    """GET request to URL1/{year} and return parsed JSON."""
    url = f"{URL1}/{year}"
    print(f"  Fetching: {url}")
    response = requests.get(url)
    response.encoding = "utf-8"
    response.raise_for_status()
    return response.json()


def fetch_course_detail(year: int, code: str) -> tuple[str, str]:
    """
    GET request to URL2/{year}/{code}.
    Returns (description_thai, response_code).
    On any error, returns ("NEED_HUMAN_FIX", "NEED_HUMAN_FIX").
    """
    try:
        url = f"{URL2}/{year}/{code}"
        response = requests.get(url)
        response.encoding = "utf-8"
        response.raise_for_status()
        data = response.json()

        description = data.get("description", {}).get("thai", "NEED_HUMAN_FIX")
        response_code = data.get("information", {}).get("code", "NEED_HUMAN_FIX")
        return description, response_code

    except Exception:
        return "NEED_HUMAN_FIX", "NEED_HUMAN_FIX"


def extract_course(input_year: int, raw: dict) -> dict:
    """Extract the 6 required keys from a single course object."""
    return {
        "name":        raw.get("name_th", ""),
        "description": "",
        "semester":    f"{input_year}",
        "code":        raw.get("subject_id", ""),
        "credit":      raw.get("credit", ""),
        "course_type": "",
    }


def main():
    raw_input = input("Enter Thai year (e.g. 2560): ").strip()
    if not raw_input.isdigit():
        print("Error: please enter a numeric year.")
        sys.exit(1)
    input_year = int(raw_input)

    print(f"\n[1] Fetching course list for year {input_year}...")
    data = fetch_courses(input_year)

    print("[2] Processing courses...")
    raw_courses = data if isinstance(data, list) else data.get("courses", [])

    courses = []
    flagged = 0

    for raw in raw_courses:
        course = extract_course(input_year, raw)
        code = course["code"]

        if code:
            description, response_code = fetch_course_detail(input_year, code)

            if response_code == code:
                course["description"] = description
            else:
                course["description"] = "NEED_HUMAN_FIX"
                course["credit"]      = "NEED_HUMAN_FIX"
                course["semester"]    = "NEED_HUMAN_FIX"
                flagged += 1

        courses.append(course)

    print(f"      → {len(courses)} courses processed. ({flagged} flagged as NEED_HUMAN_FIX)")

    output_filename = f"{input_year}_courses.json"
    output = {"courses": courses}

    with open(output_filename, "w", encoding="utf-8") as f:
        json.dump(output, f, ensure_ascii=False, indent=2)

    print(f"\n✓ Saved → {output_filename}")


if __name__ == "__main__":
    main()
