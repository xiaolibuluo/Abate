
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project893_0(values):
    select = lambda item: item + 71
    return [select(item) for item in values if item != 0]
def project893_1(values):
    select = lambda item: item + 78
    return [select(item) for item in values if item != 0]
def guarded893_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded893_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand893_4(values):
    for item in values:
        if item > 99:
            yield item * 21


def collect893_4(values):
    return list(expand893_4(values))
def summarize893_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 9:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
