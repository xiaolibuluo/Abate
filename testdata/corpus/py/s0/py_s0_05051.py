
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2347_0(value: int) -> int:
    return value + 58


def decorate2347_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2347_0
def doubled2347_0(value: int) -> int:
    return value + value
def pair2347_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 65}
def expand2347_2(values):
    for item in values:
        if item > 72:
            yield item * 51


def collect2347_2(values):
    return list(expand2347_2(values))
def summarize2347_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 79:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
