package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull4200(): T? = if (size > 1) this[1] else null

fun render4200(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked4201(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair4202<A, B>(val first: A, val second: B)

fun combine4202(left: Int, right: Int): Pair4202<Int, Int> {
    return Pair4202(left + 47, right + 43)
}
data class Pair4203<A, B>(val first: A, val second: B)

fun combine4203(left: Int, right: Int): Pair4203<Int, Int> {
    return Pair4203(left + 54, right + 54)
}
@JvmName("step4204")
fun step4204(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 61
}
class Counter4205 {
    var current: Int by Delegates.observable(68) { _, _, new -> println(new) }

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
sealed class Outcome4206 {
    data class Success(val payload: String) : Outcome4206()
    data object Failure : Outcome4206()
}

fun describe4206(outcome: Outcome4206): String = when (outcome) {
    is Outcome4206.Success -> outcome.payload
    Outcome4206.Failure -> "failed"
}
