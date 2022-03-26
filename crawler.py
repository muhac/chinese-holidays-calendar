import os
import re
from datetime import datetime
from typing import Iterator, Tuple

import requests


def main():
    comments: list[str] = [
        "// automatically generated by crawler.py",
        "// manually checked by DATA NOT VERIFIED",
    ]

    for year, link, holidays in data():
        print(year, link, holidays, sep='\n')
        file = f'./data/{year}.txt'

        if os.path.isfile(file):
            with open(file) as f:
                existing = f.read()
                if comments[0] in existing and comments[1] not in existing:
                    continue

        with open(file, 'w') as f:
            f.write(
                f"{comments[0]} ({datetime.now().strftime('%-m/%-d/%Y')})\n"
                f"{comments[1]}\n\n// source: {link}\n\n{holidays}"
            )

    update_info = "> data updated at: "
    with open('./README.md', 'r') as f:
        content = f.read().split('\n')
    for i in range(len(content)):
        if content[i].startswith(update_info):
            content[i] = update_info + datetime.now().strftime("%B %-d, %Y")
    with open('./README.md', 'w') as f:
        f.write('\n'.join(content))


def data() -> Iterator[Tuple[str, str, str]]:
    for year, link in source():
        print(f"\n\n{year}: {link}")
        results: list[str] = []

        r = requests.get(link)
        r.encoding = r.apparent_encoding

        line_regex = r"(?P<id>.)、(?P<name>.*)：(</.*?>)?(?P<detail>.*放假.*。)"
        for line in r.text.replace('<br/>', '\n').split('\n'):
            match = re.search(line_regex, line)
            if match is None:
                continue

            work, rest, *_ = match.group('detail').split('。')
            dates = ';'.join((match.group('name'), parse(work), parse(rest)))
            print(dates)  # todo: 需要人工干预如下情况: 1.与周末连休, 2.补休
            results.append(f"{dates:30} // {match.group('detail')}")

        yield year, link, '\n'.join(results)


def parse(text: str) -> str:
    results: list[str] = []
    range_type_a = r"(?P<m1>\d?\d)月(?P<d1>\d?\d)日至(?P<m2>\d?\d)月(?P<d2>\d?\d)日"
    range_type_b = r"(?P<m1>\d?\d)月(?P<d1>\d?\d)日至(?P<d2>\d?\d)日"
    single_date = r"(?P<m1>\d?\d)月(?P<d1>\d?\d)日"

    for item in text.split('、'):
        match = re.search(range_type_a, item)
        if match is not None:
            results.append(f"{match.group('m1')}.{match.group('d1')}-{match.group('m2')}.{match.group('d2')}")
            print(f"\tA: {results[-1]:15} {item}")
            continue

        match = re.search(range_type_b, item)
        if match is not None:
            results.append(f"{match.group('m1')}.{match.group('d1')}-{match.group('m1')}.{match.group('d2')}")
            print(f"\tB: {results[-1]:15} {item}")
            continue

        match = re.search(single_date, item)
        if match is not None:
            results.append(f"{match.group('m1')}.{match.group('d1')}")
            print(f"\tS: {results[-1]:15} {item}")
            continue

        print(f"\tX: {'':15} {item}")

    return ','.join(results)


def source() -> Iterator[Tuple[str, str]]:
    search_url = "http://sousuo.gov.cn/s.htm?t=paper&advance=false&n=&codeYear=&codeCode=" \
                 "&searchfield=title&sort=&q=%E8%8A%82%E5%81%87%E6%97%A5%E5%AE%89%E6%8E%92"
    link_regex = r"href=['\"](?P<link>.*?)['\"].*国务院办公厅关于(?P<year>20\d\d)年.*通知"

    for line in requests.get(search_url).text.split('\n'):
        match = re.search(link_regex, line)
        if match is None:
            continue
        yield match.group('year'), match.group('link')


if __name__ == '__main__':
    main()
