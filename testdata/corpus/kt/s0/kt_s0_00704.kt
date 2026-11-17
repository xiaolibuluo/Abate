package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome5790 {
    data class Success(val payload: String) : Outcome5790()
    data object Failure : Outcome5790()
}

fun describe5790(outcome: Outcome5790): String = when (outcome) {
    is Outcome5790.Success -> outcome.payload
    Outcome5790.Failure -> "failed"
}
data class Pair5791<A, B>(val first: A, val second: B)

fun combine5791(left: Int, right: Int): Pair5791<Int, Int> {
    return Pair5791(left + 70, right + 47)
}
class Counter5792 {
    var current: Int by Delegates.observable(77) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked5793(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome5794 {
    data class Success(val payload: String) : Outcome5794()
    data object Failure : Outcome5794()
}

fun describe5794(outcome: Outcome5794): String = when (outcome) {
    is Outcome5794.Success -> outcome.payload
    Outcome5794.Failure -> "failed"
}
fun <T> List<T>.secondOrNull5795(): T? = if (size > 1) this[1] else null

fun render5795(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter5796 {
    var current: Int by Delegates.observable(8) { _, _, new -> println(new) }

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
