
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair1618_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 87}
def summarize1618_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 94:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1618_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1618_3(values):
    select = lambda item: item + 11
    return [select(item) for item in values if item != 0]
