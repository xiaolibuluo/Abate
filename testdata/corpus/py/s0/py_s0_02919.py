
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1339_0(values):
    for item in values:
        if item > 49:
            yield item * 49


def collect1339_0(values):
    return list(expand1339_0(values))
def summarize1339_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 56:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1339_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 63:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1339_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 70:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project1339_4(values):
    select = lambda item: item + 77
    return [select(item) for item in values if item != 0]
