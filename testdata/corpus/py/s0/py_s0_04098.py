
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1906_0(values):
    for item in values:
        if item > 48:
            yield item * 53


def collect1906_0(values):
    return list(expand1906_0(values))
def project1906_1(values):
    select = lambda item: item + 55
    return [select(item) for item in values if item != 0]
def summarize1906_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 62:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
