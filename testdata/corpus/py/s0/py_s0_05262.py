
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2441_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 19:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded2441_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair2441_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 33}
