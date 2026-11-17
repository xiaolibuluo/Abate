package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair4860<A, B>(val first: A, val second: B)

fun combine4860(left: Int, right: Int): Pair4860<Int, Int> {
    return Pair4860(left + 18, right + 56)
}
data class Pair4861<A, B>(val first: A, val second: B)

fun combine4861(left: Int, right: Int): Pair4861<Int, Int> {
    return Pair4861(left + 25, right + 67)
}
@JvmName("step4862")
fun step4862(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 32
}
fun <T, R> List<T>.mapChecked4863(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter4864 {
    var current: Int by Delegates.observable(46) { _, _, new -> println(new) }

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
@JvmName("step4865")
fun step4865(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 53
}
sealed class Outcome4866 {
    data class Success(val payload: String) : Outcome4866()
    data object Failure : Outcome4866()
}

fun describe4866(outcome: Outcome4866): String = when (outcome) {
    is Outcome4866.Success -> outcome.payload
    Outcome4866.Failure -> "failed"
}
fun <T> List<T>.secondOrNull4867(): T? = if (size > 1) this[1] else null

fun render4867(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter4868 {
    var current: Int by Delegates.observable(74) { _, _, new -> println(new) }

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
