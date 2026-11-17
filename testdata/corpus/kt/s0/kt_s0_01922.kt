package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair15860<A, B>(val first: A, val second: B)

fun combine15860(left: Int, right: Int): Pair15860<Int, Int> {
    return Pair15860(left + 59, right + 9)
}
fun <T> List<T>.secondOrNull15861(): T? = if (size > 1) this[1] else null

fun render15861(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull15862(): T? = if (size > 1) this[1] else null

fun render15862(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter15863 {
    var current: Int by Delegates.observable(80) { _, _, new -> println(new) }

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
sealed class Outcome15864 {
    data class Success(val payload: String) : Outcome15864()
    data object Failure : Outcome15864()
}

fun describe15864(outcome: Outcome15864): String = when (outcome) {
    is Outcome15864.Success -> outcome.payload
    Outcome15864.Failure -> "failed"
}
