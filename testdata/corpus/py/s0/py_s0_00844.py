
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch385_0(gate):
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
def traced385_1(value: int) -> int:
    return value + 70


def decorate385_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate385_1
def doubled385_1(value: int) -> int:
    return value + value
def guarded385_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch385_3(gate):
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
def summarize385_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 91:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced385_5(value: int) -> int:
    return value + 98


def decorate385_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate385_5
def doubled385_5(value: int) -> int:
    return value + value
async def fetch385_6(gate):
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
def summarize385_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 15:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
