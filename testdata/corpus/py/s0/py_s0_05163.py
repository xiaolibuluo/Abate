
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2406_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize2406_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 56:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch2406_2(gate):
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
def guarded2406_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced2406_4(value: int) -> int:
    return value + 77


def decorate2406_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2406_4
def doubled2406_4(value: int) -> int:
    return value + value
