
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project703_0(values):
    select = lambda item: item + 26
    return [select(item) for item in values if item != 0]
def expand703_1(values):
    for item in values:
        if item > 33:
            yield item * 61


def collect703_1(values):
    return list(expand703_1(values))
def guarded703_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize703_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 47:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded703_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand703_5(values):
    for item in values:
        if item > 61:
            yield item * 44


def collect703_5(values):
    return list(expand703_5(values))
