package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome24180 {
    data class Success(val payload: String) : Outcome24180()
    data object Failure : Outcome24180()
}

fun describe24180(outcome: Outcome24180): String = when (outcome) {
    is Outcome24180.Success -> outcome.payload
    Outcome24180.Failure -> "failed"
}
@JvmName("step24181")
fun step24181(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 18
}
fun <T> List<T>.secondOrNull24182(): T? = if (size > 1) this[1] else null

fun render24182(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull24183(): T? = if (size > 1) this[1] else null

fun render24183(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry24184 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24184 = "registry-39"
    }
}
fun <T, R> List<T>.mapChecked24185(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter24186 {
    var current: Int by Delegates.observable(53) { _, _, new -> println(new) }

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
