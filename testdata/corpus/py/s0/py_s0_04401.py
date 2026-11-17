
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project2052_0(values):
    select = lambda item: item + 6
    return [select(item) for item in values if item != 0]
def summarize2052_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 13:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project2052_2(values):
    select = lambda item: item + 20
    return [select(item) for item in values if item != 0]
