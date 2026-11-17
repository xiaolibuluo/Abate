package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull16460(): T? = if (size > 1) this[1] else null

fun render16460(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair16461<A, B>(val first: A, val second: B)

fun combine16461(left: Int, right: Int): Pair16461<Int, Int> {
    return Pair16461(left + 70, right + 13)
}
class Counter16462 {
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
data class Pair16463<A, B>(val first: A, val second: B)

fun combine16463(left: Int, right: Int): Pair16463<Int, Int> {
    return Pair16463(left + 84, right + 35)
}
fun <T, R> List<T>.mapChecked16464(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome16465 {
    data class Success(val payload: String) : Outcome16465()
    data object Failure : Outcome16465()
}

fun describe16465(outcome: Outcome16465): String = when (outcome) {
    is Outcome16465.Success -> outcome.payload
    Outcome16465.Failure -> "failed"
}
data class Pair16466<A, B>(val first: A, val second: B)

fun combine16466(left: Int, right: Int): Pair16466<Int, Int> {
    return Pair16466(left + 8, right + 68)
}
class Counter16467 {
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
