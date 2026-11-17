package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome22180 {
    data class Success(val payload: String) : Outcome22180()
    data object Failure : Outcome22180()
}

fun describe22180(outcome: Outcome22180): String = when (outcome) {
    is Outcome22180.Success -> outcome.payload
    Outcome22180.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked22181(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter22182 {
    var current: Int by Delegates.observable(44) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull22183(): T? = if (size > 1) this[1] else null

fun render22183(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull22184(): T? = if (size > 1) this[1] else null

fun render22184(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
