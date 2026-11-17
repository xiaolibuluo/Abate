package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome12410 {
    data class Success(val payload: String) : Outcome12410()
    data object Failure : Outcome12410()
}

fun describe12410(outcome: Outcome12410): String = when (outcome) {
    is Outcome12410.Success -> outcome.payload
    Outcome12410.Failure -> "failed"
}
class Counter12411 {
    var current: Int by Delegates.observable(43) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked12412(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step12413")
fun step12413(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 57
}
fun <T> List<T>.secondOrNull12414(): T? = if (size > 1) this[1] else null

fun render12414(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
