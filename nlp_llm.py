"""用 LLM 处理放假安排的数据"""

import os
import re

import openai

client = openai.OpenAI(
    base_url="http://localhost:8000/v1",
    api_key="NA",
)

PROMPT = """你需要把自然语言处理成格式化内容.
你将收到一个中国节假日的放假安排, 它包含了节日名, 放假日期和调休日期.

数据格式为: 节日名;放假日期;补班日期
如果没有补班, 可以省略为: 节日名;放假日期

其中, 日期格式形如 Y.M.D,Y.M.D-Y.M.D
用“,”拼接多个日期, 并使用“-”表示日期区间

例如:

输入
2013年春节: 2月9日至15日放假调休，共7天。2月16日(星期六)、2月17日(星期日)上班。
输出
春节;2013.2.9-2013.2.15;2013.2.16,2013.2.17

输入
2023年元旦: 2022年12月31日至2023年1月2日放假调休，共3天。
输出
元旦;2022.12.31-2023.1.2

输入
2025年国庆节、中秋节: 10月1日（周三）至8日（周三）放假调休，共8天。9月28日（周日）、10月11日（周六）上班。
输出
国庆节、中秋节;2025.10.1-2025.10.8;2025.9.28,2025.10.11

根据上面的例子完成下面这个数据的处理. 你应该仅输出最终结果, 必须只有一行, 格式和上面一致.

输入
"""


def main():
    """更新节假日信息文件"""
    for file in os.listdir('./data'):
        if file.endswith('.txt'):
            with open(f'./data/{file}', encoding='utf-8') as f_obj:
                content = f_obj.read()

            if "DATA NOT VERIFIED" not in content:
                continue

            data = parse(file[:4], content.split('\n'))
            data = '\n'.join(data).replace("DATA NOT VERIFIED", "// BY AI")

            with open(f'./data/{file}', 'w', encoding='utf-8') as f_obj:
                f_obj.write(data)


def parse(year: str, lines: list[str]) -> list[str]:
    """按行处理节假日信息"""
    print(year)

    padding = max(len(line.split('//')[0]) for line in lines)

    for i, line in enumerate(lines):
        if len(line.split(';')) < 2:
            continue

        name = line.split(';')[0]
        text = line.split('//')[-1]

        query = f"{year}年{name}: {text}"
        value = get_response(query).strip()

        if value and '\n' not in value and re.search(r";\d{4}\.\d{1,2}\.\d{1,2}", value):
            lines[i] = f"{value:{padding}}// {line}"

        print(lines[i])

    return lines


def get_response(query: str) -> str:
    """获取 LLM 的处理结果"""
    try:
        completion = client.chat.completions.create(
            model="Qwen/Qwen2.5-1.5B-Instruct",
            messages=[
                {"role": "user", "content": PROMPT + query},
            ]
        )
        return completion.choices[0].message.content

    except openai.APIConnectionError as e:
        print(e, query)
        return ""


if __name__ == '__main__':
    main()
