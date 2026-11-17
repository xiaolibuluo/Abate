package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked17430(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair17431<A, B>(val first: A, val second: B)

fun combine17431(left: Int, right: Int): Pair17431<Int, Int> {
    return Pair17431(left + 70, right + 21)
}
@JvmName("step17432")
fun step17432(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 77
}
fun <T> List<T>.secondOrNull17433(): T? = if (size > 1) this[1] else null

fun render17433(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull17434(): T? = if (size > 1) this[1] else null

fun render17434(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome17435 {
    data class Success(val payload: String) : Outcome17435()
    data object Failure : Outcome17435()
}

fun describe17435(outcome: Outcome17435): String = when (outcome) {
    is Outcome17435.Success -> outcome.payload
    Outcome17435.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked17436(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter17437 {
    var current: Int by Delegates.observable(15) { _, _, new -> println(new) }

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
