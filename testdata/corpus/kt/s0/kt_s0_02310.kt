package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair18970<A, B>(val first: A, val second: B)

fun combine18970(left: Int, right: Int): Pair18970<Int, Int> {
    return Pair18970(left + 28, right + 51)
}
sealed class Outcome18971 {
    data class Success(val payload: String) : Outcome18971()
    data object Failure : Outcome18971()
}

fun describe18971(outcome: Outcome18971): String = when (outcome) {
    is Outcome18971.Success -> outcome.payload
    Outcome18971.Failure -> "failed"
}
data class Pair18972<A, B>(val first: A, val second: B)

fun combine18972(left: Int, right: Int): Pair18972<Int, Int> {
    return Pair18972(left + 42, right + 12)
}
fun <T, R> List<T>.mapChecked18973(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome18974 {
    data class Success(val payload: String) : Outcome18974()
    data object Failure : Outcome18974()
}

fun describe18974(outcome: Outcome18974): String = when (outcome) {
    is Outcome18974.Success -> outcome.payload
    Outcome18974.Failure -> "failed"
}
class Counter18975 {
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
fun <T> List<T>.secondOrNull18976(): T? = if (size > 1) this[1] else null

fun render18976(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair18977<A, B>(val first: A, val second: B)

fun combine18977(left: Int, right: Int): Pair18977<Int, Int> {
    return Pair18977(left + 77, right + 67)
}
class Counter18978 {
    var current: Int by Delegates.observable(84) { _, _, new -> println(new) }

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
