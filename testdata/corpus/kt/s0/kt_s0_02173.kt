package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked17880(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome17881 {
    data class Success(val payload: String) : Outcome17881()
    data object Failure : Outcome17881()
}

fun describe17881(outcome: Outcome17881): String = when (outcome) {
    is Outcome17881.Success -> outcome.payload
    Outcome17881.Failure -> "failed"
}
class Counter17882 {
    var current: Int by Delegates.observable(80) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull17883(): T? = if (size > 1) this[1] else null

fun render17883(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair17884<A, B>(val first: A, val second: B)

fun combine17884(left: Int, right: Int): Pair17884<Int, Int> {
    return Pair17884(left + 94, right + 64)
}
fun <T> List<T>.secondOrNull17885(): T? = if (size > 1) this[1] else null

fun render17885(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull17886(): T? = if (size > 1) this[1] else null

fun render17886(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
