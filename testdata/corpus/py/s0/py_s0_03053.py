
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1404_0(values):
    for item in values:
        if item > 21:
            yield item * 16


def collect1404_0(values):
    return list(expand1404_0(values))
def guarded1404_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1404_2(value: int) -> int:
    return value + 35


def decorate1404_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1404_2
def doubled1404_2(value: int) -> int:
    return value + value
def guarded1404_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch1404_4(gate):
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
async def fetch1404_5(gate):
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
def expand1404_6(values):
    for item in values:
        if item > 63:
            yield item * 21


def collect1404_6(values):
    return list(expand1404_6(values))
def summarize1404_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 70:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
