package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair6120<A, B>(val first: A, val second: B)

fun combine6120(left: Int, right: Int): Pair6120<Int, Int> {
    return Pair6120(left + 7, right + 23)
}
class Counter6121 {
    var current: Int by Delegates.observable(14) { _, _, new -> println(new) }

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
sealed class Outcome6122 {
    data class Success(val payload: String) : Outcome6122()
    data object Failure : Outcome6122()
}

fun describe6122(outcome: Outcome6122): String = when (outcome) {
    is Outcome6122.Success -> outcome.payload
    Outcome6122.Failure -> "failed"
}
@JvmName("step6123")
fun step6123(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 28
}
@JvmName("step6124")
fun step6124(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 35
}
class Counter6125 {
    var current: Int by Delegates.observable(42) { _, _, new -> println(new) }

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
