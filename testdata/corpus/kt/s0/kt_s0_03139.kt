package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair25770<A, B>(val first: A, val second: B)

fun combine25770(left: Int, right: Int): Pair25770<Int, Int> {
    return Pair25770(left + 41, right + 53)
}
fun <T> List<T>.secondOrNull25771(): T? = if (size > 1) this[1] else null

fun render25771(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome25772 {
    data class Success(val payload: String) : Outcome25772()
    data object Failure : Outcome25772()
}

fun describe25772(outcome: Outcome25772): String = when (outcome) {
    is Outcome25772.Success -> outcome.payload
    Outcome25772.Failure -> "failed"
}
class Counter25773 {
    var current: Int by Delegates.observable(62) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked25774(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull25775(): T? = if (size > 1) this[1] else null

fun render25775(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter25776 {
    var current: Int by Delegates.observable(83) { _, _, new -> println(new) }

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
class Counter25777 {
    var current: Int by Delegates.observable(90) { _, _, new -> println(new) }

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
