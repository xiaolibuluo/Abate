
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize441_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 15:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
