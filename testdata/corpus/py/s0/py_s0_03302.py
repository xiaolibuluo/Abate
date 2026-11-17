
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1528_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 81:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair1528_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 88}
def guarded1528_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1528_3(values):
    select = lambda item: item + 5
    return [select(item) for item in values if item != 0]
def guarded1528_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
