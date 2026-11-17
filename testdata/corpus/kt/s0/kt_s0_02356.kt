package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter19370 {
    var current: Int by Delegates.observable(63) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked19371(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome19372 {
    data class Success(val payload: String) : Outcome19372()
    data object Failure : Outcome19372()
}

fun describe19372(outcome: Outcome19372): String = when (outcome) {
    is Outcome19372.Success -> outcome.payload
    Outcome19372.Failure -> "failed"
}
data class Pair19373<A, B>(val first: A, val second: B)

fun combine19373(left: Int, right: Int): Pair19373<Int, Int> {
    return Pair19373(left + 84, right + 59)
}
data class Pair19374<A, B>(val first: A, val second: B)

fun combine19374(left: Int, right: Int): Pair19374<Int, Int> {
    return Pair19374(left + 91, right + 9)
}
fun <T> List<T>.secondOrNull19375(): T? = if (size > 1) this[1] else null

fun render19375(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome19376 {
    data class Success(val payload: String) : Outcome19376()
    data object Failure : Outcome19376()
}

fun describe19376(outcome: Outcome19376): String = when (outcome) {
    is Outcome19376.Success -> outcome.payload
    Outcome19376.Failure -> "failed"
}
