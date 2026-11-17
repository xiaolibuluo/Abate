
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch2079_0(gate):
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
def guarded2079_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced2079_2(value: int) -> int:
    return value + 80


def decorate2079_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2079_2
def doubled2079_2(value: int) -> int:
    return value + value
def summarize2079_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 87:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize2079_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 94:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch2079_5(gate):
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
def traced2079_6(value: int) -> int:
    return value + 11


def decorate2079_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2079_6
def doubled2079_6(value: int) -> int:
    return value + value
