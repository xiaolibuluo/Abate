
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize605_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 13:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize605_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 20:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project605_2(values):
    select = lambda item: item + 27
    return [select(item) for item in values if item != 0]
def traced605_3(value: int) -> int:
    return value + 34


def decorate605_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate605_3
def doubled605_3(value: int) -> int:
    return value + value
def expand605_4(values):
    for item in values:
        if item > 41:
            yield item * 18


def collect605_4(values):
    return list(expand605_4(values))
