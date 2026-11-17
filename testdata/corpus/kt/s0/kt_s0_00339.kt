package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull2820(): T? = if (size > 1) this[1] else null

fun render2820(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step2821")
fun step2821(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 89
}
fun <T> List<T>.secondOrNull2822(): T? = if (size > 1) this[1] else null

fun render2822(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair2823<A, B>(val first: A, val second: B)

fun combine2823(left: Int, right: Int): Pair2823<Int, Int> {
    return Pair2823(left + 6, right + 64)
}
@JvmName("step2824")
fun step2824(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 13
}
class Counter2825 {
    var current: Int by Delegates.observable(20) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked2826(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome2827 {
    data class Success(val payload: String) : Outcome2827()
    data object Failure : Outcome2827()
}

fun describe2827(outcome: Outcome2827): String = when (outcome) {
    is Outcome2827.Success -> outcome.payload
    Outcome2827.Failure -> "failed"
}
