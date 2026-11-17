
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand785_0(values):
    for item in values:
        if item > 25:
            yield item * 14


def collect785_0(values):
    return list(expand785_0(values))
def guarded785_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize785_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 39:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
