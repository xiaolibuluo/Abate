package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair19450<A, B>(val first: A, val second: B)

fun combine19450(left: Int, right: Int): Pair19450<Int, Int> {
    return Pair19450(left + 70, right + 21)
}
data class Pair19451<A, B>(val first: A, val second: B)

fun combine19451(left: Int, right: Int): Pair19451<Int, Int> {
    return Pair19451(left + 77, right + 32)
}
@JvmName("step19452")
fun step19452(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 84
}
class Counter19453 {
    var current: Int by Delegates.observable(91) { _, _, new -> println(new) }

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
@JvmName("step19454")
fun step19454(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 98
}
class Counter19455 {
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
data class Pair19456<A, B>(val first: A, val second: B)

fun combine19456(left: Int, right: Int): Pair19456<Int, Int> {
    return Pair19456(left + 15, right + 26)
}
sealed class Outcome19457 {
    data class Success(val payload: String) : Outcome19457()
    data object Failure : Outcome19457()
}

fun describe19457(outcome: Outcome19457): String = when (outcome) {
    is Outcome19457.Success -> outcome.payload
    Outcome19457.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked19458(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
