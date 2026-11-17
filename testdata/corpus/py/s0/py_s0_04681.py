
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2169_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 72:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand2169_1(values):
    for item in values:
        if item > 79:
            yield item * 14


def collect2169_1(values):
    return list(expand2169_1(values))
def project2169_2(values):
    select = lambda item: item + 86
    return [select(item) for item in values if item != 0]
def pair2169_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 93}
def summarize2169_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 100:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced2169_5(value: int) -> int:
    return value + 10


def decorate2169_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2169_5
def doubled2169_5(value: int) -> int:
    return value + value
