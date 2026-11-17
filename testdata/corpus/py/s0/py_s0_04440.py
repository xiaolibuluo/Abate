
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2068_0(values):
    for item in values:
        if item > 20:
            yield item * 28


def collect2068_0(values):
    return list(expand2068_0(values))
def summarize2068_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 27:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize2068_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 34:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand2068_3(values):
    for item in values:
        if item > 41:
            yield item * 61


def collect2068_3(values):
    return list(expand2068_3(values))
def summarize2068_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 48:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
