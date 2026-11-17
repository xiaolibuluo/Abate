
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair220_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 52}
def pair220_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 59}
def guarded220_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize220_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 73:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize220_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 80:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded220_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
