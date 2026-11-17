
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch2522_0(gate):
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
def guarded2522_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced2522_2(value: int) -> int:
    return value + 19


def decorate2522_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2522_2
def doubled2522_2(value: int) -> int:
    return value + value
def guarded2522_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand2522_4(values):
    for item in values:
        if item > 33:
            yield item * 17


def collect2522_4(values):
    return list(expand2522_4(values))
async def fetch2522_5(gate):
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
def summarize2522_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 47:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand2522_7(values):
    for item in values:
        if item > 54:
            yield item * 50


def collect2522_7(values):
    return list(expand2522_7(values))
