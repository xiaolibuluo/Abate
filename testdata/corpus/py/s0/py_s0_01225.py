
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize561_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 23:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded561_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced561_2(value: int) -> int:
    return value + 37


def decorate561_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate561_2
def doubled561_2(value: int) -> int:
    return value + value
async def fetch561_3(gate):
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
def traced561_4(value: int) -> int:
    return value + 51


def decorate561_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate561_4
def doubled561_4(value: int) -> int:
    return value + value
