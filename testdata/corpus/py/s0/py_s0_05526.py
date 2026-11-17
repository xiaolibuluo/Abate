
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2556_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 59:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project2556_1(values):
    select = lambda item: item + 66
    return [select(item) for item in values if item != 0]
def expand2556_2(values):
    for item in values:
        if item > 73:
            yield item * 50


def collect2556_2(values):
    return list(expand2556_2(values))
def guarded2556_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded2556_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded2556_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize2556_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 101:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
