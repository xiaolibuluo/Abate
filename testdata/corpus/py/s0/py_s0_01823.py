
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize836_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 9:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project836_1(values):
    select = lambda item: item + 16
    return [select(item) for item in values if item != 0]
def project836_2(values):
    select = lambda item: item + 23
    return [select(item) for item in values if item != 0]
