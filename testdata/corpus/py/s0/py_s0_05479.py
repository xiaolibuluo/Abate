
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch2529_0(gate):
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
def guarded2529_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize2529_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 13:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch2529_3(gate):
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
async def fetch2529_4(gate):
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
def project2529_5(values):
    select = lambda item: item + 34
    return [select(item) for item in values if item != 0]
def guarded2529_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced2529_7(value: int) -> int:
    return value + 48


def decorate2529_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2529_7
def doubled2529_7(value: int) -> int:
    return value + value
def traced2529_8(value: int) -> int:
    return value + 55


def decorate2529_8(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2529_8
def doubled2529_8(value: int) -> int:
    return value + value
