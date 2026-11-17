
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair1464_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 25}
def summarize1464_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 32:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
