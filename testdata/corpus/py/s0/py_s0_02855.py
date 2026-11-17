
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1308_0(values):
    for item in values:
        if item > 34:
            yield item * 15


def collect1308_0(values):
    return list(expand1308_0(values))
def project1308_1(values):
    select = lambda item: item + 41
    return [select(item) for item in values if item != 0]
def traced1308_2(value: int) -> int:
    return value + 48


def decorate1308_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1308_2
def doubled1308_2(value: int) -> int:
    return value + value
def project1308_3(values):
    select = lambda item: item + 55
    return [select(item) for item in values if item != 0]
def summarize1308_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 62:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1308_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 69:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair1308_6(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 76}
def project1308_7(values):
    select = lambda item: item + 83
    return [select(item) for item in values if item != 0]
def traced1308_8(value: int) -> int:
    return value + 90


def decorate1308_8(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1308_8
def doubled1308_8(value: int) -> int:
    return value + value
def project1308_9(values):
    select = lambda item: item + 97
    return [select(item) for item in values if item != 0]
