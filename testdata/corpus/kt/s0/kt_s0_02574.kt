package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked21210(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked21211(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair21212<A, B>(val first: A, val second: B)

fun combine21212(left: Int, right: Int): Pair21212<Int, Int> {
    return Pair21212(left + 44, right + 55)
}
sealed class Outcome21213 {
    data class Success(val payload: String) : Outcome21213()
    data object Failure : Outcome21213()
}

fun describe21213(outcome: Outcome21213): String = when (outcome) {
    is Outcome21213.Success -> outcome.payload
    Outcome21213.Failure -> "failed"
}
data class Pair21214<A, B>(val first: A, val second: B)

fun combine21214(left: Int, right: Int): Pair21214<Int, Int> {
    return Pair21214(left + 58, right + 16)
}
class Counter21215 {
    var current: Int by Delegates.observable(65) { _, _, new -> println(new) }

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
