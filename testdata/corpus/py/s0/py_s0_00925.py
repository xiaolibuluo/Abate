
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize425_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 98:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand425_1(values):
    for item in values:
        if item > 8:
            yield item * 67


def collect425_1(values):
    return list(expand425_1(values))
def summarize425_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 15:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced425_3(value: int) -> int:
    return value + 22


def decorate425_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate425_3
def doubled425_3(value: int) -> int:
    return value + value
def summarize425_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 29:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project425_5(values):
    select = lambda item: item + 36
    return [select(item) for item in values if item != 0]
def summarize425_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 43:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
