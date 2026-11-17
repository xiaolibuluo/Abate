package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome11480 {
    data class Success(val payload: String) : Outcome11480()
    data object Failure : Outcome11480()
}

fun describe11480(outcome: Outcome11480): String = when (outcome) {
    is Outcome11480.Success -> outcome.payload
    Outcome11480.Failure -> "failed"
}
fun <T> List<T>.secondOrNull11481(): T? = if (size > 1) this[1] else null

fun render11481(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step11482")
fun step11482(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 5
}
class Counter11483 {
    var current: Int by Delegates.observable(12) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked11484(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
