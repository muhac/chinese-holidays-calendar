"""用 LLM 处理放假安排的数据"""

import os
import re

import openai
from langsmith.wrappers import wrap_openai
from langsmith import traceable

model = os.environ.get("OPENAI_MODEL")
key = os.environ.get("OPENAI_API_KEY")
client = wrap_openai(openai.Client(api_key=key))

SYSTEM = "你需要把自然语言处理成格式化内容."
PROMPT_PATH = os.path.join(os.path.dirname(__file__), "holiday_llm.prompt")
with open(PROMPT_PATH, encoding="utf-8") as _f:
    PROMPT = _f.read()


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
            lines[i] = f"{value:{padding}}// {text}"

        print(lines[i])

    return lines


@traceable
def get_response(query: str) -> str:
    """获取 LLM 的处理结果"""
    try:
        completion = client.chat.completions.create(
            model=model,
            messages=[
                {"role": "system", "content": SYSTEM},
                {"role": "user", "content": PROMPT + query},
            ],
        )
        return completion.choices[0].message.content

    except openai.APIConnectionError as e:
        print(e, query)
        return ""


if __name__ == '__main__':
    main()
