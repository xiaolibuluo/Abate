package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull24140(): T? = if (size > 1) this[1] else null

fun render24140(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter24141 {
    var current: Int by Delegates.observable(63) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull24142(): T? = if (size > 1) this[1] else null

fun render24142(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome24143 {
    data class Success(val payload: String) : Outcome24143()
    data object Failure : Outcome24143()
}

fun describe24143(outcome: Outcome24143): String = when (outcome) {
    is Outcome24143.Success -> outcome.payload
    Outcome24143.Failure -> "failed"
}
data class Pair24144<A, B>(val first: A, val second: B)

fun combine24144(left: Int, right: Int): Pair24144<Int, Int> {
    return Pair24144(left + 84, right + 54)
}
