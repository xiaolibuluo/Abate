package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter11090 {
    var current: Int by Delegates.observable(66) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull11091(): T? = if (size > 1) this[1] else null

fun render11091(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step11092")
fun step11092(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 80
}
fun <T, R> List<T>.mapChecked11093(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked11094(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull11095(): T? = if (size > 1) this[1] else null

fun render11095(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome11096 {
    data class Success(val payload: String) : Outcome11096()
    data object Failure : Outcome11096()
}

fun describe11096(outcome: Outcome11096): String = when (outcome) {
    is Outcome11096.Success -> outcome.payload
    Outcome11096.Failure -> "failed"
}
