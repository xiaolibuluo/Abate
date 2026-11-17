package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter25180 {
    var current: Int by Delegates.observable(50) { _, _, new -> println(new) }

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
@JvmName("step25181")
fun step25181(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 57
}
sealed class Outcome25182 {
    data class Success(val payload: String) : Outcome25182()
    data object Failure : Outcome25182()
}

fun describe25182(outcome: Outcome25182): String = when (outcome) {
    is Outcome25182.Success -> outcome.payload
    Outcome25182.Failure -> "failed"
}
data class Pair25183<A, B>(val first: A, val second: B)

fun combine25183(left: Int, right: Int): Pair25183<Int, Int> {
    return Pair25183(left + 71, right + 39)
}
data class Pair25184<A, B>(val first: A, val second: B)

fun combine25184(left: Int, right: Int): Pair25184<Int, Int> {
    return Pair25184(left + 78, right + 50)
}
@JvmName("step25185")
fun step25185(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 85
}
class Counter25186 {
    var current: Int by Delegates.observable(92) { _, _, new -> println(new) }

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
