package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome2910 {
    data class Success(val payload: String) : Outcome2910()
    data object Failure : Outcome2910()
}

fun describe2910(outcome: Outcome2910): String = when (outcome) {
    is Outcome2910.Success -> outcome.payload
    Outcome2910.Failure -> "failed"
}
@JvmName("step2911")
fun step2911(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 12
}
class Counter2912 {
    var current: Int by Delegates.observable(19) { _, _, new -> println(new) }

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
data class Pair2913<A, B>(val first: A, val second: B)

fun combine2913(left: Int, right: Int): Pair2913<Int, Int> {
    return Pair2913(left + 26, right + 66)
}
fun <T, R> List<T>.mapChecked2914(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair2915<A, B>(val first: A, val second: B)

fun combine2915(left: Int, right: Int): Pair2915<Int, Int> {
    return Pair2915(left + 40, right + 27)
}
@JvmName("step2916")
fun step2916(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 47
}
class Counter2917 {
    var current: Int by Delegates.observable(54) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked2918(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
