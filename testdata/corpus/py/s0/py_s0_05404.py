
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project2496_0(values):
    select = lambda item: item + 55
    return [select(item) for item in values if item != 0]
def summarize2496_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 62:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand2496_2(values):
    for item in values:
        if item > 69:
            yield item * 57


def collect2496_2(values):
    return list(expand2496_2(values))
def guarded2496_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
