
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair2557_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 72}
def summarize2557_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 79:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize2557_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 86:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair2557_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 93}
def guarded2557_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded2557_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
