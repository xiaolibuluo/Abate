package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step26090")
fun step26090(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 69
}
class Counter26091 {
    var current: Int by Delegates.observable(76) { _, _, new -> println(new) }

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
data class Pair26092<A, B>(val first: A, val second: B)

fun combine26092(left: Int, right: Int): Pair26092<Int, Int> {
    return Pair26092(left + 83, right + 55)
}
fun <T, R> List<T>.mapChecked26093(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome26094 {
    data class Success(val payload: String) : Outcome26094()
    data object Failure : Outcome26094()
}

fun describe26094(outcome: Outcome26094): String = when (outcome) {
    is Outcome26094.Success -> outcome.payload
    Outcome26094.Failure -> "failed"
}
data class Pair26095<A, B>(val first: A, val second: B)

fun combine26095(left: Int, right: Int): Pair26095<Int, Int> {
    return Pair26095(left + 7, right + 27)
}
