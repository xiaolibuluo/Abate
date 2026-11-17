package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair16390<A, B>(val first: A, val second: B)

fun combine16390(left: Int, right: Int): Pair16390<Int, Int> {
    return Pair16390(left + 69, right + 14)
}
sealed class Outcome16391 {
    data class Success(val payload: String) : Outcome16391()
    data object Failure : Outcome16391()
}

fun describe16391(outcome: Outcome16391): String = when (outcome) {
    is Outcome16391.Success -> outcome.payload
    Outcome16391.Failure -> "failed"
}
class Counter16392 {
    var current: Int by Delegates.observable(83) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull16393(): T? = if (size > 1) this[1] else null

fun render16393(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter16394 {
    var current: Int by Delegates.observable(97) { _, _, new -> println(new) }

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
