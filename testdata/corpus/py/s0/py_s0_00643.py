
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize296_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 70:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize296_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 77:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
