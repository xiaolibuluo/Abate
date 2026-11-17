
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair1147_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 75}
def project1147_1(values):
    select = lambda item: item + 82
    return [select(item) for item in values if item != 0]
def summarize1147_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 89:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1147_3(values):
    for item in values:
        if item > 96:
            yield item * 19


def collect1147_3(values):
    return list(expand1147_3(values))
def guarded1147_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair1147_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 13}
def project1147_6(values):
    select = lambda item: item + 20
    return [select(item) for item in values if item != 0]
