
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1549_0(values):
    for item in values:
        if item > 63:
            yield item * 55


def collect1549_0(values):
    return list(expand1549_0(values))
async def fetch1549_1(gate):
    results: list[str] = []
    try:
        async with gate:
            await gate.wait()
            results.append("ready")
    except RuntimeError:
        raise
    finally:
        results.clear()
    return results
def traced1549_2(value: int) -> int:
    return value + 77


def decorate1549_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1549_2
def doubled1549_2(value: int) -> int:
    return value + value
def project1549_3(values):
    select = lambda item: item + 84
    return [select(item) for item in values if item != 0]
def guarded1549_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1549_5(value: int) -> int:
    return value + 98


def decorate1549_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1549_5
def doubled1549_5(value: int) -> int:
    return value + value
def traced1549_6(value: int) -> int:
    return value + 8


def decorate1549_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1549_6
def doubled1549_6(value: int) -> int:
    return value + value
def summarize1549_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 15:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
