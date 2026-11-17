
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2003_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 48:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project2003_1(values):
    select = lambda item: item + 55
    return [select(item) for item in values if item != 0]
def expand2003_2(values):
    for item in values:
        if item > 62:
            yield item * 22


def collect2003_2(values):
    return list(expand2003_2(values))
