package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter21280 {
    var current: Int by Delegates.observable(24) { _, _, new -> println(new) }

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
sealed class Outcome21281 {
    data class Success(val payload: String) : Outcome21281()
    data object Failure : Outcome21281()
}

fun describe21281(outcome: Outcome21281): String = when (outcome) {
    is Outcome21281.Success -> outcome.payload
    Outcome21281.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked21282(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked21283(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome21284 {
    data class Success(val payload: String) : Outcome21284()
    data object Failure : Outcome21284()
}

fun describe21284(outcome: Outcome21284): String = when (outcome) {
    is Outcome21284.Success -> outcome.payload
    Outcome21284.Failure -> "failed"
}
data class Pair21285<A, B>(val first: A, val second: B)

fun combine21285(left: Int, right: Int): Pair21285<Int, Int> {
    return Pair21285(left + 59, right + 15)
}
