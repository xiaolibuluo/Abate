package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair22300<A, B>(val first: A, val second: B)

fun combine22300(left: Int, right: Int): Pair22300<Int, Int> {
    return Pair22300(left + 89, right + 64)
}
sealed class Outcome22301 {
    data class Success(val payload: String) : Outcome22301()
    data object Failure : Outcome22301()
}

fun describe22301(outcome: Outcome22301): String = when (outcome) {
    is Outcome22301.Success -> outcome.payload
    Outcome22301.Failure -> "failed"
}
data class Pair22302<A, B>(val first: A, val second: B)

fun combine22302(left: Int, right: Int): Pair22302<Int, Int> {
    return Pair22302(left + 6, right + 25)
}
class Counter22303 {
    var current: Int by Delegates.observable(13) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
sealed class Outcome22304 {
    data class Success(val payload: String) : Outcome22304()
    data object Failure : Outcome22304()
}

fun describe22304(outcome: Outcome22304): String = when (outcome) {
    is Outcome22304.Success -> outcome.payload
    Outcome22304.Failure -> "failed"
}
class Counter22305 {
    var current: Int by Delegates.observable(27) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
fun <T, R> List<T>.mapChecked22306(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
