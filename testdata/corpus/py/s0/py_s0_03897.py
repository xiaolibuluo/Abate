
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1803_0(values):
    select = lambda item: item + 67
    return [select(item) for item in values if item != 0]
def project1803_1(values):
    select = lambda item: item + 74
    return [select(item) for item in values if item != 0]
def summarize1803_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 81:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
