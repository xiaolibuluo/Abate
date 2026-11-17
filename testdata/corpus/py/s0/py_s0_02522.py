
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1166_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 31:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch1166_1(gate):
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
async def fetch1166_2(gate):
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
def project1166_3(values):
    select = lambda item: item + 52
    return [select(item) for item in values if item != 0]
async def fetch1166_4(gate):
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
def traced1166_5(value: int) -> int:
    return value + 66


def decorate1166_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1166_5
def doubled1166_5(value: int) -> int:
    return value + value
def summarize1166_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 73:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair1166_7(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 80}
def guarded1166_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
