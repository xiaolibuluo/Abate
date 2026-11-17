
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1071_0(values):
    select = lambda item: item + 57
    return [select(item) for item in values if item != 0]
def summarize1071_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 64:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
