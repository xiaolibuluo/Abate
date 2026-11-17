
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair2542_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 71}
def summarize2542_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 78:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair2542_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 85}
def traced2542_3(value: int) -> int:
    return value + 92


def decorate2542_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2542_3
def doubled2542_3(value: int) -> int:
    return value + value
