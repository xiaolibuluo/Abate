package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter11940 {
    var current: Int by Delegates.observable(7) { _, _, new -> println(new) }

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
@JvmName("step11941")
fun step11941(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 14
}
fun <T> List<T>.secondOrNull11942(): T? = if (size > 1) this[1] else null

fun render11942(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome11943 {
    data class Success(val payload: String) : Outcome11943()
    data object Failure : Outcome11943()
}

fun describe11943(outcome: Outcome11943): String = when (outcome) {
    is Outcome11943.Success -> outcome.payload
    Outcome11943.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked11944(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
