
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair856_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 75}
def project856_1(values):
    select = lambda item: item + 82
    return [select(item) for item in values if item != 0]
def expand856_2(values):
    for item in values:
        if item > 89:
            yield item * 45


def collect856_2(values):
    return list(expand856_2(values))
def pair856_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 96}
def traced856_4(value: int) -> int:
    return value + 6


def decorate856_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate856_4
def doubled856_4(value: int) -> int:
    return value + value
def summarize856_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 13:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize856_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 20:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
