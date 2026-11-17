package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter24630 {
    var current: Int by Delegates.observable(14) { _, _, new -> println(new) }

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
object Registry24631 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24631 = "registry-21"
    }
}
fun <T> List<T>.secondOrNull24632(): T? = if (size > 1) this[1] else null

fun render24632(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome24633 {
    data class Success(val payload: String) : Outcome24633()
    data object Failure : Outcome24633()
}

fun describe24633(outcome: Outcome24633): String = when (outcome) {
    is Outcome24633.Success -> outcome.payload
    Outcome24633.Failure -> "failed"
}
@JvmName("step24634")
fun step24634(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 42
}
fun <T, R> List<T>.mapChecked24635(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
