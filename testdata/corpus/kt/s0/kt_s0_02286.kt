package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome18750 {
    data class Success(val payload: String) : Outcome18750()
    data object Failure : Outcome18750()
}

fun describe18750(outcome: Outcome18750): String = when (outcome) {
    is Outcome18750.Success -> outcome.payload
    Outcome18750.Failure -> "failed"
}
@JvmName("step18751")
fun step18751(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 40
}
data class Pair18752<A, B>(val first: A, val second: B)

fun combine18752(left: Int, right: Int): Pair18752<Int, Int> {
    return Pair18752(left + 47, right + 41)
}
fun <T, R> List<T>.mapChecked18753(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair18754<A, B>(val first: A, val second: B)

fun combine18754(left: Int, right: Int): Pair18754<Int, Int> {
    return Pair18754(left + 61, right + 63)
}
class Counter18755 {
    var current: Int by Delegates.observable(68) { _, _, new -> println(new) }

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
