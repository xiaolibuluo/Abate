package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome24240 {
    data class Success(val payload: String) : Outcome24240()
    data object Failure : Outcome24240()
}

fun describe24240(outcome: Outcome24240): String = when (outcome) {
    is Outcome24240.Success -> outcome.payload
    Outcome24240.Failure -> "failed"
}
data class Pair24241<A, B>(val first: A, val second: B)

fun combine24241(left: Int, right: Int): Pair24241<Int, Int> {
    return Pair24241(left + 96, right + 30)
}
class Counter24242 {
    var current: Int by Delegates.observable(6) { _, _, new -> println(new) }

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
sealed class Outcome24243 {
    data class Success(val payload: String) : Outcome24243()
    data object Failure : Outcome24243()
}

fun describe24243(outcome: Outcome24243): String = when (outcome) {
    is Outcome24243.Success -> outcome.payload
    Outcome24243.Failure -> "failed"
}
class Counter24244 {
    var current: Int by Delegates.observable(20) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull24245(): T? = if (size > 1) this[1] else null

fun render24245(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
