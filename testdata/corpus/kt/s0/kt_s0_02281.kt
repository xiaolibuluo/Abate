package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome18720 {
    data class Success(val payload: String) : Outcome18720()
    data object Failure : Outcome18720()
}

fun describe18720(outcome: Outcome18720): String = when (outcome) {
    is Outcome18720.Success -> outcome.payload
    Outcome18720.Failure -> "failed"
}
class Counter18721 {
    var current: Int by Delegates.observable(98) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked18722(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull18723(): T? = if (size > 1) this[1] else null

fun render18723(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
