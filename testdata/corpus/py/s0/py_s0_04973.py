
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair2307_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 23}
def expand2307_1(values):
    for item in values:
        if item > 30:
            yield item * 65


def collect2307_1(values):
    return list(expand2307_1(values))
def summarize2307_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 37:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced2307_3(value: int) -> int:
    return value + 44


def decorate2307_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2307_3
def doubled2307_3(value: int) -> int:
    return value + value
def summarize2307_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 51:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
