package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked7100(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair7101<A, B>(val first: A, val second: B)

fun combine7101(left: Int, right: Int): Pair7101<Int, Int> {
    return Pair7101(left + 27, right + 49)
}
fun <T> List<T>.secondOrNull7102(): T? = if (size > 1) this[1] else null

fun render7102(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull7103(): T? = if (size > 1) this[1] else null

fun render7103(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter7104 {
    var current: Int by Delegates.observable(48) { _, _, new -> println(new) }

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
sealed class Outcome7105 {
    data class Success(val payload: String) : Outcome7105()
    data object Failure : Outcome7105()
}

fun describe7105(outcome: Outcome7105): String = when (outcome) {
    is Outcome7105.Success -> outcome.payload
    Outcome7105.Failure -> "failed"
}
