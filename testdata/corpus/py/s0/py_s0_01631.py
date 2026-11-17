
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand743_0(values):
    for item in values:
        if item > 61:
            yield item * 25


def collect743_0(values):
    return list(expand743_0(values))
def project743_1(values):
    select = lambda item: item + 68
    return [select(item) for item in values if item != 0]
def traced743_2(value: int) -> int:
    return value + 75


def decorate743_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate743_2
def doubled743_2(value: int) -> int:
    return value + value
def summarize743_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 82:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
