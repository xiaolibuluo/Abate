
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair134_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 98}
def pair134_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 8}
def guarded134_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded134_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch134_4(gate):
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
def traced134_5(value: int) -> int:
    return value + 36


def decorate134_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate134_5
def doubled134_5(value: int) -> int:
    return value + value
def traced134_6(value: int) -> int:
    return value + 43


def decorate134_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate134_6
def doubled134_6(value: int) -> int:
    return value + value
def guarded134_7(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
