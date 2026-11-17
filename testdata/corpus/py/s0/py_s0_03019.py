
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1385_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 65:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair1385_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 72}
def project1385_2(values):
    select = lambda item: item + 79
    return [select(item) for item in values if item != 0]
def guarded1385_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded1385_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1385_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 100:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
