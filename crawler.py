"""从国务院官网抓取放假信息"""

import os
import re
import json
from datetime import datetime, timezone, timedelta
from typing import Iterator, Tuple

import requests


def main():
    """更新节假日信息"""
    comments: list[str] = [
        "// automatically generated on",
        "//    and manually checked on DATA NOT VERIFIED",
    ]

    for year, link, holidays in data():
        print(year, link, holidays, sep='\n')
        file = f"./data/{year}.txt"

        if os.path.isfile(file):
            with open(file, encoding='utf-8') as f_obj:
                existing = f_obj.read()
                if "DATA NOT VERIFIED" not in existing:
                    continue  # 数据已人工确认

        with open(file, 'w', encoding='utf-8') as f_obj:
            f_obj.write(
                f"{comments[0]} {beijing_time().strftime('%-m/%-d/%Y')}\n"
                f"{comments[1]}\n// source: {link}\n\n{holidays}"
            )

    with open('./README.md', 'r', encoding='utf-8') as f_obj:
        content = f_obj.read().split('\n')

    update_info = "> Calendar data updated "
    for i, line in enumerate(content):
        if line.startswith(update_info):
            content[i] = update_info + beijing_time().strftime("at %-H:%M on %B %-d, %Y")

    with open('./README.md', 'w', encoding='utf-8') as f_obj:
        f_obj.write('\n'.join(content))


def data() -> Iterator[Tuple[str, str, str]]:
    """爬取国务院网站数据"""
    for year, link in source():
        print(f"\n\n{year}: {link}")
        results: list[str] = []

        response = requests.get(link, timeout=(5, 10))
        response.encoding = response.apparent_encoding

        line_regex = r"(?P<id>.)、(?P<name>.*)：(</.*?>)?(?P<detail>.*放假.*。)"
        for line in response.text.replace('<br/>', '\n').replace('</p>', '\n').split('\n'):
            if match := re.search(line_regex, line):
                work, rest, *_ = match.group("detail").split('。')
                dates = ';'.join((match.group("name"), parse(work, year), parse(rest, year)))
                print(dates)  # 已知需要人工干预如下情况: 1.与周末连休, 2.补休
                results.append(f"{dates:50} // {match.group('detail')}")

        yield year, link, '\n'.join(results)


def parse(text: str, year: str) -> str:
    """解析节假日安排数据"""
    results: list[str] = []
    range_type_a = r"(?P<m1>\d?\d)月(?P<d1>\d?\d)日(（[^）]+）)?至(?P<m2>\d?\d)月(?P<d2>\d?\d)日"
    range_type_b = r"(?P<m1>\d?\d)月(?P<d1>\d?\d)日(（[^）]+）)?至(?P<d2>\d?\d)日"
    single_date = r"(?P<m1>\d?\d)月(?P<d1>\d?\d)日"

    for item in text.split('、'):
        if match := re.search(range_type_a, item):
            results.append(f"{year}.{match.group('m1')}.{match.group('d1')}-"
                           f"{year}.{match.group('m2')}.{match.group('d2')}")
            print(f"\tA: {results[-1]:25} {item}")

        elif match := re.search(range_type_b, item):
            results.append(f"{year}.{match.group('m1')}.{match.group('d1')}-"
                           f"{year}.{match.group('m1')}.{match.group('d2')}")
            print(f"\tB: {results[-1]:25} {item}")

        elif match := re.search(single_date, item):
            results.append(f"{year}.{match.group('m1')}.{match.group('d1')}")
            print(f"\tS: {results[-1]:25} {item}")

        else:
            print(f"\tX: {'':15} {item}")

    return ','.join(results)


def source() -> Iterator[Tuple[str, str]]:
    """获取官网发布通知列表"""
    search_url = ("https://sousuoht.www.gov.cn/athena/forward/"
                  "2B22E8E39E850E17F95A016A74FCB6B673336FA8B6FEC0E2955907EF9AEE06BE")
    search_resp = requests.post(
        search_url,
        timeout=(5, 10),
        json={
            "code": "17da70961a7",
            "dataTypeId": "107",
            "orderBy": "time",
            "searchBy": "title",
            "pageNo": 1,
            "pageSize": 10,
            "searchWord": "节假日安排"
        },
        headers={
            "Content-Type": "application/json;charset=UTF-8",
            "Athenaappname": "%E5%9B%BD%E7%BD%91%E6%90%9C%E7%B4%A2",  # 国网搜索
            "Athenaappkey": "R6kU1sKz%2BSHMWevvn8zRHAGJOwe6o53KEf7v0BJPVWY0Vxb4gIMCWDshQDW"
                            "xw2Ua5CtSpEalXVlVzlj4mulMw6lnQzuqTEDDRC833wtXJyC%2F1kPbME1oCi"
                            "jkQOluRhbXZj3iRK1nXCzq5Sw%2F%2B2XbJbm%2BFEPLZRwoNNeYoOLcueg%3D"
        },
    )

    search_resp.encoding = search_resp.apparent_encoding
    search_data = json.loads(search_resp.text)
    # print(search_data)

    for result in search_data["result"]["data"]["middle"]["list"]:
        if match := re.search(r"^国务院办公厅关于(?P<year>20\d\d)年.*通知", result["title"]):
            print(match.group("year"), result["url"])
            print(result, end="\n\n")
            yield match.group("year"), result["url"]


def beijing_time() -> datetime:
    """获取当前北京时间"""
    utc_time = datetime.utcnow().replace(tzinfo=timezone.utc)
    return utc_time.astimezone(timezone(timedelta(hours=8)))


if __name__ == '__main__':
    main()
