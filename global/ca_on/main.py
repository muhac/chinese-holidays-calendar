"""
Generate the public holidays in Ontario for the next year
Thanks to ChatGPT and Copilot for the code
"""

import datetime
from dateutil.easter import easter


class Holiday:
    """节假日"""

    def __init__(self, name_cn, name_en, date):
        self.name_cn = name_cn
        self.name_en = name_en
        self.date = date
        self.long_weekend_dates = None

    def __str__(self):
        """格式化 debug 输出"""
        return f"{self.name_cn}: {self.format_date(self.date)}"

    @staticmethod
    def format_date(date):
        """格式化日期为 data.txt 的格式"""
        return date.strftime('%Y.%-m.%-d')


def get_holidays(year: int) -> list[Holiday]:
    """获取给定年份的节假日"""
    holidays = [
        # 固定假日
        # 元旦 (1月1日)
        Holiday("元旦", "New Year's Day", datetime.date(year, 1, 1)),
        # 劳动节 (5月1日)
        Holiday("国庆节", "Canada Day", datetime.date(year, 7, 1)),
        # 圣诞节 (12月25日)
        Holiday("圣诞节", "Christmas Day", datetime.date(year, 12, 25)),
        # 节礼日 (12月26日)
        Holiday("节礼日", "Boxing Day", datetime.date(year, 12, 26)),

        # 变动假日
        # 耶稣受难日（复活节前2天）
        Holiday("耶稣受难日", "Good Friday", easter(year) - datetime.timedelta(days=2)),
        # 维多利亚日（5月最后一个周一）
        Holiday("维多利亚日", "Victoria Day", get_last_monday(datetime.date(year, 5, 25))),
        # 劳动节（9月的第一个周一）
        Holiday("劳动节", "Labour Day", get_first_monday(datetime.date(year, 9, 1))),
        # 感恩节（10月的第二个周一）
        Holiday("感恩节", "Thanksgiving", get_second_monday(datetime.date(year, 10, 1))),
        # 家庭日（2月的第三个周一）
        Holiday("家庭日", "Family Day", get_third_monday(datetime.date(year, 2, 1))),
    ]

    # 处理假日如果落在周末的情况
    adjust_for_weekends(holidays)
    connect_long_weekends(holidays)

    return holidays


def get_first_monday(date):
    """获取给定日期所在月的第一个星期一"""
    while date.weekday() != 0:
        date += datetime.timedelta(days=1)
    return date


def get_second_monday(date):
    """获取给定日期所在月的第二个星期一"""
    first_monday = get_first_monday(date)
    return first_monday + datetime.timedelta(days=7)


def get_third_monday(date):
    """获取给定日期所在月的第三个星期一"""
    first_monday = get_first_monday(date)
    return first_monday + datetime.timedelta(days=14)


def get_last_monday(date):
    """获取给定日期所在月的最后一个星期一"""
    while date.weekday() != 0:
        date -= datetime.timedelta(days=1)
    return date


def adjust_for_weekends(holidays: list[Holiday]):
    """如果假日落在周末，将假日顺延到下一个工作日，特殊处理圣诞节和节礼日"""
    boxing_day_delta = 0
    for holiday in holidays:
        # 特殊处理圣诞节和节礼日
        if holiday.name_cn == "圣诞节" and holiday.date.weekday() == 5:
            # 圣诞节在周六, 节礼日在周日
            holiday.date += datetime.timedelta(days=2)
            boxing_day_delta = 2
        elif holiday.name_cn == "圣诞节" and holiday.date.weekday() == 6:
            # 圣诞节在周日, 节礼日在周一
            holiday.date += datetime.timedelta(days=1)
            boxing_day_delta = 1
        elif holiday.name_cn == "节礼日" and holiday.date.weekday() == 5:
            # 节礼日在周六
            holiday.date += datetime.timedelta(days=2)
        elif holiday.name_cn == "节礼日":
            holiday.date += datetime.timedelta(days=boxing_day_delta)
        else:
            # 其他假期处理
            if holiday.date.weekday() == 5:  # 如果是假日落在周六
                holiday.date += datetime.timedelta(days=2)
            elif holiday.date.weekday() == 6:  # 如果是假日落在周日
                holiday.date += datetime.timedelta(days=1)


def connect_long_weekends(holidays: list[Holiday]):
    """将连续的假期连接起来"""
    boxing_day = next(h.date for h in holidays if h.name_cn == "节礼日")

    for holiday in holidays:
        dates = [holiday.date]
        if holiday.name_cn == "节礼日":
            continue
        if holiday.name_cn == "圣诞节":
            dates.append(boxing_day)

        long_weekend_start = holiday.date
        while is_weekend(prev_day := long_weekend_start - datetime.timedelta(days=1), dates):
            long_weekend_start = prev_day

        long_weekend_end = holiday.date
        while is_weekend(next_day := long_weekend_end + datetime.timedelta(days=1), dates):
            long_weekend_end = next_day

        holiday.long_weekend_dates = (
            Holiday.format_date(long_weekend_start),
            Holiday.format_date(long_weekend_end),
        )


def is_weekend(date, extra_holidays=None):
    """判断给定日期是否是周末或者公共假日"""
    return date.weekday() >= 5 or (extra_holidays and date in extra_holidays)


def main(year=None):
    """生成下一年的节假日数据"""
    next_year = year if year else datetime.datetime.now().year + 1
    print(f"Generating Ontario holidays for {next_year}")

    holidays = get_holidays(next_year)
    holidays.sort(key=lambda x: x.date)

    data = (f"// Ontario Public Holidays in {next_year}\n"
            f"// I hope ChatGPT gets the dates correct\n\n")

    for holiday in holidays:
        print(holiday)
        title = f"{holiday.name_cn} {holiday.name_en}".replace(' ', '_')
        data += (f"{title};{Holiday.format_date(holiday.date)};"
                 f"{f'{d[0]}-{d[1]}' if (d := holiday.long_weekend_dates) else ''}\n")  # pylint: disable=unsubscriptable-object

    with open(f"./data/{next_year}.txt", "w", encoding='utf-8') as f:
        f.write(data)


if __name__ == '__main__':
    main()
