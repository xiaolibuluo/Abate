
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1255_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 24:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1255_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1255_2(values):
    for item in values:
        if item > 38:
            yield item * 32


def collect1255_2(values):
    return list(expand1255_2(values))
def project1255_3(values):
    select = lambda item: item + 45
    return [select(item) for item in values if item != 0]
def traced1255_4(value: int) -> int:
    return value + 52


def decorate1255_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1255_4
def doubled1255_4(value: int) -> int:
    return value + value
def summarize1255_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 59:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1255_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1255_7(values):
    select = lambda item: item + 73
    return [select(item) for item in values if item != 0]
async def fetch1255_8(gate):
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
