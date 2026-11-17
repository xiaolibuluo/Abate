package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter7230 {
    var current: Int by Delegates.observable(92) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull7231(): T? = if (size > 1) this[1] else null

fun render7231(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair7232<A, B>(val first: A, val second: B)

fun combine7232(left: Int, right: Int): Pair7232<Int, Int> {
    return Pair7232(left + 9, right + 29)
}
fun <T, R> List<T>.mapChecked7233(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry7234 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7234 = "registry-23"
    }
}
class Counter7235 {
    var current: Int by Delegates.observable(30) { _, _, new -> println(new) }

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
sealed class Outcome7236 {
    data class Success(val payload: String) : Outcome7236()
    data object Failure : Outcome7236()
}

fun describe7236(outcome: Outcome7236): String = when (outcome) {
    is Outcome7236.Success -> outcome.payload
    Outcome7236.Failure -> "failed"
}
