
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project2146_0(values):
    select = lambda item: item + 64
    return [select(item) for item in values if item != 0]
def summarize2146_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 71:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
