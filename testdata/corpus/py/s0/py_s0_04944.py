
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project2294_0(values):
    select = lambda item: item + 48
    return [select(item) for item in values if item != 0]
def pair2294_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 55}
def summarize2294_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 62:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
