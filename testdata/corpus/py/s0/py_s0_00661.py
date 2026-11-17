
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize304_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 77:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project304_1(values):
    select = lambda item: item + 84
    return [select(item) for item in values if item != 0]
def traced304_2(value: int) -> int:
    return value + 91


def decorate304_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate304_2
def doubled304_2(value: int) -> int:
    return value + value
