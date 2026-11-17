
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1176_0(values):
    for item in values:
        if item > 64:
            yield item * 67


def collect1176_0(values):
    return list(expand1176_0(values))
def traced1176_1(value: int) -> int:
    return value + 71


def decorate1176_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1176_1
def doubled1176_1(value: int) -> int:
    return value + value
def summarize1176_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 78:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
