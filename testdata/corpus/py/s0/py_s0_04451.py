
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2074_0(value: int) -> int:
    return value + 98


def decorate2074_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2074_0
def doubled2074_0(value: int) -> int:
    return value + value
def guarded2074_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded2074_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced2074_3(value: int) -> int:
    return value + 22


def decorate2074_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2074_3
def doubled2074_3(value: int) -> int:
    return value + value
async def fetch2074_4(gate):
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
def summarize2074_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 36:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
